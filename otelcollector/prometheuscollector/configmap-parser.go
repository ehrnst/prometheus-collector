// package configmapparserpackage

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

// func configmapparser() {
// 	// Set agent config schema version
// 	setConfigSchemaVersion()

// 	// Set agent config file version
// 	setConfigFileVersion()

// 	// Parse the settings for pod annotations
// 	parsePodAnnotations()

// 	// Parse the configmap to set the right environment variables for Prometheus collector settings
// 	parsePrometheusCollectorSettings()

// 	// Parse the settings for default scrape configs
// 	parseDefaultScrapeSettings()

// 	// Parse the settings for debug mode
// 	parseDebugMode()

// 	// Parse the settings for default targets metrics keep list config
// 	parseDefaultTargetsMetricsKeepList()

// 	// Parse the settings for default-targets-scrape-interval-settings config
// 	parseScrapeInterval()

// 	// Merge default and custom Prometheus config
// 	mergePrometheusConfig()

// 	// Set environment variables from the prom-config-validator
// 	setPromConfigValidatorEnvVars()

// 	fmt.Println("prom-config-validator::Use default Prometheus config:", os.Getenv("AZMON_USE_DEFAULT_PROMETHEUS_CONFIG"))
// }

// func setConfigSchemaVersion() {
// 	if existsAndNotEmpty("/etc/config/settings/schema-version") {
// 		configSchemaVersion := readFileTrim("/etc/config/settings/schema-version")
// 		configSchemaVersion = strings.ReplaceAll(configSchemaVersion, " ", "")
// 		configSchemaVersion = configSchemaVersion[:10]
// 		os.Setenv("AZMON_AGENT_CFG_SCHEMA_VERSION", configSchemaVersion)
// 		appendAndSourceBashrc("AZMON_AGENT_CFG_SCHEMA_VERSION", configSchemaVersion)
// 	}
// }

// func setConfigFileVersion() {
// 	if existsAndNotEmpty("/etc/config/settings/config-version") {
// 		configFileVersion := readFileTrim("/etc/config/settings/config-version")
// 		configFileVersion = strings.ReplaceAll(configFileVersion, " ", "")
// 		configFileVersion = configFileVersion[:10]
// 		os.Setenv("AZMON_AGENT_CFG_FILE_VERSION", configFileVersion)
// 		appendAndSourceBashrc("AZMON_AGENT_CFG_FILE_VERSION", configFileVersion)
// 	}
// }

// func parsePodAnnotations() {
// 	runRubyScript("/opt/microsoft/configmapparser/tomlparser-pod-annotation-based-scraping.rb")

// 	if exists("/opt/microsoft/configmapparser/config_def_pod_annotation_based_scraping") {
// 		appendFileToBashrc("/opt/microsoft/configmapparser/config_def_pod_annotation_based_scraping")
// 		sourceBashrc()
// 	}
// }

// func parsePrometheusCollectorSettings() {
// 	runRubyScript("/opt/microsoft/configmapparser/tomlparser-prometheus-collector-settings.rb")
// 	appendFileToBashrc("/opt/microsoft/configmapparser/config_prometheus_collector_settings_env_var")
// 	sourceBashrc()
// }

// func parseDefaultScrapeSettings() {
// 	runRubyScript("/opt/microsoft/configmapparser/tomlparser-default-scrape-settings.rb")

// 	if exists("/opt/microsoft/configmapparser/config_default_scrape_settings_env_var") {
// 		appendFileToBashrc("/opt/microsoft/configmapparser/config_default_scrape_settings_env_var")
// 		sourceBashrc()
// 	}
// }

// func parseDebugMode() {
// 	runRubyScript("/opt/microsoft/configmapparser/tomlparser-debug-mode.rb")

// 	if exists("/opt/microsoft/configmapparser/config_debug_mode_env_var") {
// 		appendFileToBashrc("/opt/microsoft/configmapparser/config_debug_mode_env_var")
// 		sourceBashrc()
// 	}
// }

// func parseDefaultTargetsMetricsKeepList() {
// 	runRubyScript("/opt/microsoft/configmapparser/tomlparser-default-targets-metrics-keep-list.rb")
// }

// func parseScrapeInterval() {
// 	runRubyScript("/opt/microsoft/configmapparser/tomlparser-scrape-interval.rb")
// }

// func mergePrometheusConfig() {
// 	var rubyScriptPath string

// 	if os.Getenv("AZMON_OPERATOR_ENABLED") == "true" || os.Getenv("CONTAINER_TYPE") == "ConfigReaderSidecar" {
// 		rubyScriptPath = "/opt/microsoft/configmapparser/prometheus-config-merger-with-operator.rb"
// 	} else {
// 		rubyScriptPath = "/opt/microsoft/configmapparser/prometheus-config-merger.rb"
// 	}

// 	runRubyScript(rubyScriptPath)

// 	appendAndSourceBashrc("AZMON_INVALID_CUSTOM_PROMETHEUS_CONFIG", "false")
// 	os.Setenv("AZMON_INVALID_CUSTOM_PROMETHEUS_CONFIG", "false")
// 	appendAndSourceBashrc("CONFIG_VALIDATOR_RUNNING_IN_AGENT", "true")
// 	os.Setenv("CONFIG_VALIDATOR_RUNNING_IN_AGENT", "true")

// 	if exists("/opt/promMergedConfig.yml") {
// 		err := runPromConfigValidator("/opt/promMergedConfig.yml", "/opt/microsoft/otelcollector/collector-config.yml", "/opt/microsoft/otelcollector/collector-config-template.yml")
// 		if err != nil || !exists("/opt/microsoft/otelcollector/collector-config.yml") {
// 			fmt.Println("prom-config-validator::Prometheus custom config validation failed. The custom config will not be used")
// 			appendAndSourceBashrc("AZMON_INVALID_CUSTOM_PROMETHEUS_CONFIG", "true")
// 			os.Setenv("AZMON_INVALID_CUSTOM_PROMETHEUS_CONFIG", "true")

// 			if exists("/opt/defaultsMergedConfig.yml") {
// 				fmt.Println("prom-config-validator::Running validator on just default scrape configs")
// 				err = runPromConfigValidator("/opt/defaultsMergedConfig.yml", "/opt/collector-config-with-defaults.yml", "/opt/microsoft/otelcollector/collector-config-template.yml")
// 				if err != nil || !exists("/opt/collector-config-with-defaults.yml") {
// 					fmt.Println("prom-config-validator::Prometheus default scrape config validation failed. No scrape configs will be used")
// 				} else {
// 					copyFile("/opt/collector-config-with-defaults.yml", "/opt/microsoft/otelcollector/collector-config-default.yml")
// 				}
// 			}

// 			appendAndSourceBashrc("AZMON_USE_DEFAULT_PROMETHEUS_CONFIG", "true")
// 			os.Setenv("AZMON_USE_DEFAULT_PROMETHEUS_CONFIG", "true")
// 		}
// 	} else if exists("/opt/defaultsMergedConfig.yml") {
// 		fmt.Println("prom-config-validator::No custom Prometheus config found. Only using default scrape configs")
// 		err := runPromConfigValidator("/opt/defaultsMergedConfig.yml", "/opt/collector-config-with-defaults.yml", "/opt/microsoft/otelcollector/collector-config-template.yml")
// 		if err != nil || !exists("/opt/collector-config-with-defaults.yml") {
// 			fmt.Println("prom-config-validator::Prometheus default scrape config validation failed. No scrape configs will be used")
// 		} else {
// 			fmt.Println("prom-config-validator::Prometheus default scrape config validation succeeded, using this as collector config")
// 			copyFile("/opt/collector-config-with-defaults.yml", "/opt/microsoft/otelcollector/collector-config-default.yml")
// 		}
// 		appendAndSourceBashrc("AZMON_USE_DEFAULT_PROMETHEUS_CONFIG", "true")
// 		os.Setenv("AZMON_USE_DEFAULT_PROMETHEUS_CONFIG", "true")
// 	} else {
// 		fmt.Println("prom-config-validator::No custom config via configmap or default scrape configs enabled.")
// 		appendAndSourceBashrc("AZMON_USE_DEFAULT_PROMETHEUS_CONFIG", "true")
// 		os.Setenv("AZMON_USE_DEFAULT_PROM
