// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	gokitlog "github.com/go-kit/log"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/prometheus/discovery"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/allocation"
	"github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/collector"
	"github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/config"
	"github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/prehook"
	"github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/server"
	"github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/target"
	allocatorWatcher "github.com/open-telemetry/opentelemetry-operator/cmd/otel-allocator/watcher"
)

var (
	setupLog     = ctrl.Log.WithName("setup")
	eventsMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "opentelemetry_allocator_events",
		Help: "Number of events in the channel.",
	}, []string{"source"})
)

func main() {
	var (
		// allocatorPrehook will be nil if filterStrategy is not set or
		// unrecognized. No filtering will be used in this case.
		allocatorPrehook prehook.Hook
		allocator        allocation.Allocator
		discoveryManager *discovery.Manager
		collectorWatcher *collector.Client
		fileWatcher      allocatorWatcher.Watcher
		promWatcher      allocatorWatcher.Watcher
		targetDiscoverer *target.Discoverer

		discoveryCancel context.CancelFunc
		runGroup        run.Group
		eventChan       = make(chan allocatorWatcher.Event)
		eventCloser     = make(chan bool, 1)
		interrupts      = make(chan os.Signal, 1)
		errChan         = make(chan error)
	)

	// EULA statement is required for Arc extension
	clusterResourceId := os.Getenv("CLUSTER")
	if strings.EqualFold(clusterResourceId, "connectedclusters") {
		setupLog.Info("MICROSOFT SOFTWARE LICENSE TERMS\n\nMICROSOFT Azure Arc-enabled Kubernetes\n\nThis software is licensed to you as part of your or your company's subscription license for Microsoft Azure Services. You may only use the software with Microsoft Azure Services and subject to the terms and conditions of the agreement under which you obtained Microsoft Azure Services. If you do not have an active subscription license for Microsoft Azure Services, you may not use the software. Microsoft Azure Legal Information: https://azure.microsoft.com/en-us/support/legal/")
	}

	cliConf, err := config.ParseCLI()
	if err != nil {
		setupLog.Error(err, "Failed to parse parameters")
		os.Exit(1)
	}

	// Defaulting to consistent hashing
	allocationStrategy := "consistent-hashing"
	// Config file will not exist at startup, so not attempting to load the file which results in an error and just using defaults here.
	cfg := config.Config{
		AllocationStrategy: &allocationStrategy,
		LabelSelector: map[string]string{
			"rsName":                         "ama-metrics",
			"kubernetes.azure.com/managedby": "aks",
		},
	}

	if validationErr := config.ValidateConfig(&cfg, &cliConf); validationErr != nil {
		setupLog.Error(validationErr, "Invalid configuration")
	}

	cliConf.RootLogger.Info("Starting the Target Allocator")
	ctx := context.Background()
	log := ctrl.Log.WithName("allocator")

	allocatorPrehook = prehook.New(cfg.GetTargetsFilterStrategy(), log)
	allocator, err = allocation.New(cfg.GetAllocationStrategy(), log, allocation.WithFilter(allocatorPrehook))
	if err != nil {
		setupLog.Error(err, "Unable to initialize allocation strategy")
		os.Exit(1)
	}
	srv := server.NewServer(log, allocator, cliConf.ListenAddr)

	discoveryCtx, discoveryCancel := context.WithCancel(ctx)
	discoveryManager = discovery.NewManager(discoveryCtx, gokitlog.NewNopLogger())
	targetDiscoverer = target.NewDiscoverer(log, discoveryManager, allocatorPrehook, srv)
	collectorWatcher, collectorWatcherErr := collector.NewClient(log, cliConf.ClusterConfig)
	if collectorWatcherErr != nil {
		setupLog.Error(collectorWatcherErr, "Unable to initialize collector watcher")
		os.Exit(1)
	}
	fileWatcher, err = allocatorWatcher.NewFileWatcher(setupLog.WithName("file-watcher"), cliConf)
	if err != nil {
		setupLog.Error(err, "Can't start the file watcher")
		os.Exit(1)
	}
	signal.Notify(interrupts, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer close(interrupts)

	if *cliConf.PromCRWatcherConf.Enabled {
		promWatcher, err = allocatorWatcher.NewPrometheusCRWatcher(setupLog.WithName("prometheus-cr-watcher"), cfg, cliConf)
		if err != nil {
			setupLog.Error(err, "Can't start the prometheus watcher")
			os.Exit(1)
		}
		runGroup.Add(
			func() error {
				promWatcherErr := promWatcher.Watch(eventChan, errChan)
				setupLog.Info("Prometheus watcher exited")
				return promWatcherErr
			},
			func(_ error) {
				setupLog.Info("Closing prometheus watcher")
				promWatcherErr := promWatcher.Close()
				if promWatcherErr != nil {
					setupLog.Error(promWatcherErr, "prometheus watcher failed to close")
				}
			})
	}
	runGroup.Add(
		func() error {
			fileWatcherErr := fileWatcher.Watch(eventChan, errChan)
			setupLog.Info("File watcher exited")
			return fileWatcherErr
		},
		func(_ error) {
			setupLog.Info("Closing file watcher")
			fileWatcherErr := fileWatcher.Close()
			if fileWatcherErr != nil {
				setupLog.Error(fileWatcherErr, "file watcher failed to close")
			}
		})
	runGroup.Add(
		func() error {
			discoveryManagerErr := discoveryManager.Run()
			setupLog.Info("Discovery manager exited")
			return discoveryManagerErr
		},
		func(_ error) {
			setupLog.Info("Closing discovery manager")
			discoveryCancel()
		})
	runGroup.Add(
		func() error {
			// Initial loading of the config file's scrape config
			err = targetDiscoverer.ApplyConfig(allocatorWatcher.EventSourceConfigMap, cfg.Config)
			if err != nil {
				setupLog.Error(err, "Unable to apply initial configuration")
				return err
			}
			err := targetDiscoverer.Watch(allocator.SetTargets)
			setupLog.Info("Target discoverer exited")
			return err
		},
		func(_ error) {
			setupLog.Info("Closing target discoverer")
			targetDiscoverer.Close()
		})
	runGroup.Add(
		func() error {
			err := collectorWatcher.Watch(ctx, cfg.LabelSelector, allocator.SetCollectors)
			setupLog.Info("Collector watcher exited")
			return err
		},
		func(_ error) {
			setupLog.Info("Closing collector watcher")
			collectorWatcher.Close()
		})
	runGroup.Add(
		func() error {
			err := srv.Start()
			setupLog.Info("Server failed to start")
			return err
		},
		func(_ error) {
			setupLog.Info("Closing server")
			if shutdownErr := srv.Shutdown(ctx); shutdownErr != nil {
				setupLog.Error(shutdownErr, "Error on server shutdown")
			}
		})
	runGroup.Add(
		func() error {
			for {
				select {
				case event := <-eventChan:
					eventsMetric.WithLabelValues(event.Source.String()).Inc()
					loadConfig, err := event.Watcher.LoadConfig(ctx)
					if err != nil {
						setupLog.Error(err, "Unable to load configuration")
						continue
					}
					err = targetDiscoverer.ApplyConfig(event.Source, loadConfig)
					if err != nil {
						setupLog.Error(err, "Unable to apply configuration")
						continue
					}
				case err := <-errChan:
					setupLog.Error(err, "Watcher error")
				case <-eventCloser:
					return nil
				}
			}
		},
		func(_ error) {
			setupLog.Info("Closing watcher loop")
			close(eventCloser)
		})
	runGroup.Add(
		func() error {
			for {
				select {
				case <-interrupts:
					setupLog.Info("Received interrupt")
					return nil
				case <-eventCloser:
					return nil
				}
			}
		},
		func(_ error) {
			setupLog.Info("Closing interrupt loop")
		})
	if runErr := runGroup.Run(); runErr != nil {
		setupLog.Error(runErr, "run group exited")
	}
	setupLog.Info("Target allocator exited.")
}