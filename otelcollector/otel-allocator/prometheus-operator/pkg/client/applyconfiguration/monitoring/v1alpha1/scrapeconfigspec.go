// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
)

// ScrapeConfigSpecApplyConfiguration represents an declarative configuration of the ScrapeConfigSpec type for use
// with apply.
type ScrapeConfigSpecApplyConfiguration struct {
	StaticConfigs         []StaticConfigApplyConfiguration                  `json:"staticConfigs,omitempty"`
	FileSDConfigs         []FileSDConfigApplyConfiguration                  `json:"fileSDConfigs,omitempty"`
	HTTPSDConfigs         []HTTPSDConfigApplyConfiguration                  `json:"httpSDConfigs,omitempty"`
	KubernetesSDConfigs   []KubernetesSDConfigApplyConfiguration            `json:"kubernetesSDConfigs,omitempty"`
	ConsulSDConfigs       []ConsulSDConfigApplyConfiguration                `json:"consulSDConfigs,omitempty"`
	DNSSDConfigs          []DNSSDConfigApplyConfiguration                   `json:"dnsSDConfigs,omitempty"`
	EC2SDConfigs          []EC2SDConfigApplyConfiguration                   `json:"ec2SDConfigs,omitempty"`
	RelabelConfigs        []*v1.RelabelConfig                               `json:"relabelings,omitempty"`
	MetricsPath           *string                                           `json:"metricsPath,omitempty"`
	ScrapeInterval        *v1.Duration                                      `json:"scrapeInterval,omitempty"`
	ScrapeTimeout         *v1.Duration                                      `json:"scrapeTimeout,omitempty"`
	HonorTimestamps       *bool                                             `json:"honorTimestamps,omitempty"`
	HonorLabels           *bool                                             `json:"honorLabels,omitempty"`
	Params                map[string][]string                               `json:"params,omitempty"`
	Scheme                *string                                           `json:"scheme,omitempty"`
	BasicAuth             *monitoringv1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	Authorization         *monitoringv1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	TLSConfig             *monitoringv1.SafeTLSConfigApplyConfiguration     `json:"tlsConfig,omitempty"`
	SampleLimit           *uint64                                           `json:"sampleLimit,omitempty"`
	TargetLimit           *uint64                                           `json:"targetLimit,omitempty"`
	LabelLimit            *uint64                                           `json:"labelLimit,omitempty"`
	LabelNameLengthLimit  *uint64                                           `json:"labelNameLengthLimit,omitempty"`
	LabelValueLengthLimit *uint64                                           `json:"labelValueLengthLimit,omitempty"`
	KeepDroppedTargets    *uint64                                           `json:"keepDroppedTargets,omitempty"`
	MetricRelabelConfigs  []*v1.RelabelConfig                               `json:"metricRelabelings,omitempty"`
}

// ScrapeConfigSpecApplyConfiguration constructs an declarative configuration of the ScrapeConfigSpec type for use with
// apply.
func ScrapeConfigSpec() *ScrapeConfigSpecApplyConfiguration {
	return &ScrapeConfigSpecApplyConfiguration{}
}

// WithStaticConfigs adds the given value to the StaticConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the StaticConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithStaticConfigs(values ...*StaticConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithStaticConfigs")
		}
		b.StaticConfigs = append(b.StaticConfigs, *values[i])
	}
	return b
}

// WithFileSDConfigs adds the given value to the FileSDConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FileSDConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithFileSDConfigs(values ...*FileSDConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFileSDConfigs")
		}
		b.FileSDConfigs = append(b.FileSDConfigs, *values[i])
	}
	return b
}

// WithHTTPSDConfigs adds the given value to the HTTPSDConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HTTPSDConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithHTTPSDConfigs(values ...*HTTPSDConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHTTPSDConfigs")
		}
		b.HTTPSDConfigs = append(b.HTTPSDConfigs, *values[i])
	}
	return b
}

// WithKubernetesSDConfigs adds the given value to the KubernetesSDConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the KubernetesSDConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithKubernetesSDConfigs(values ...*KubernetesSDConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithKubernetesSDConfigs")
		}
		b.KubernetesSDConfigs = append(b.KubernetesSDConfigs, *values[i])
	}
	return b
}

// WithConsulSDConfigs adds the given value to the ConsulSDConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConsulSDConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithConsulSDConfigs(values ...*ConsulSDConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConsulSDConfigs")
		}
		b.ConsulSDConfigs = append(b.ConsulSDConfigs, *values[i])
	}
	return b
}

// WithDNSSDConfigs adds the given value to the DNSSDConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DNSSDConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithDNSSDConfigs(values ...*DNSSDConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDNSSDConfigs")
		}
		b.DNSSDConfigs = append(b.DNSSDConfigs, *values[i])
	}
	return b
}

// WithEC2SDConfigs adds the given value to the EC2SDConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EC2SDConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithEC2SDConfigs(values ...*EC2SDConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEC2SDConfigs")
		}
		b.EC2SDConfigs = append(b.EC2SDConfigs, *values[i])
	}
	return b
}

// WithRelabelConfigs adds the given value to the RelabelConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RelabelConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithRelabelConfigs(values ...**v1.RelabelConfig) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRelabelConfigs")
		}
		b.RelabelConfigs = append(b.RelabelConfigs, *values[i])
	}
	return b
}

// WithMetricsPath sets the MetricsPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MetricsPath field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithMetricsPath(value string) *ScrapeConfigSpecApplyConfiguration {
	b.MetricsPath = &value
	return b
}

// WithScrapeInterval sets the ScrapeInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeInterval field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithScrapeInterval(value v1.Duration) *ScrapeConfigSpecApplyConfiguration {
	b.ScrapeInterval = &value
	return b
}

// WithScrapeTimeout sets the ScrapeTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeTimeout field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithScrapeTimeout(value v1.Duration) *ScrapeConfigSpecApplyConfiguration {
	b.ScrapeTimeout = &value
	return b
}

// WithHonorTimestamps sets the HonorTimestamps field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HonorTimestamps field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithHonorTimestamps(value bool) *ScrapeConfigSpecApplyConfiguration {
	b.HonorTimestamps = &value
	return b
}

// WithHonorLabels sets the HonorLabels field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HonorLabels field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithHonorLabels(value bool) *ScrapeConfigSpecApplyConfiguration {
	b.HonorLabels = &value
	return b
}

// WithParams puts the entries into the Params field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Params field,
// overwriting an existing map entries in Params field with the same key.
func (b *ScrapeConfigSpecApplyConfiguration) WithParams(entries map[string][]string) *ScrapeConfigSpecApplyConfiguration {
	if b.Params == nil && len(entries) > 0 {
		b.Params = make(map[string][]string, len(entries))
	}
	for k, v := range entries {
		b.Params[k] = v
	}
	return b
}

// WithScheme sets the Scheme field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scheme field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithScheme(value string) *ScrapeConfigSpecApplyConfiguration {
	b.Scheme = &value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithBasicAuth(value *monitoringv1.BasicAuthApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithAuthorization(value *monitoringv1.SafeAuthorizationApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	b.Authorization = value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithTLSConfig(value *monitoringv1.SafeTLSConfigApplyConfiguration) *ScrapeConfigSpecApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithSampleLimit sets the SampleLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampleLimit field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithSampleLimit(value uint64) *ScrapeConfigSpecApplyConfiguration {
	b.SampleLimit = &value
	return b
}

// WithTargetLimit sets the TargetLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetLimit field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithTargetLimit(value uint64) *ScrapeConfigSpecApplyConfiguration {
	b.TargetLimit = &value
	return b
}

// WithLabelLimit sets the LabelLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelLimit field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithLabelLimit(value uint64) *ScrapeConfigSpecApplyConfiguration {
	b.LabelLimit = &value
	return b
}

// WithLabelNameLengthLimit sets the LabelNameLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelNameLengthLimit field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithLabelNameLengthLimit(value uint64) *ScrapeConfigSpecApplyConfiguration {
	b.LabelNameLengthLimit = &value
	return b
}

// WithLabelValueLengthLimit sets the LabelValueLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelValueLengthLimit field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithLabelValueLengthLimit(value uint64) *ScrapeConfigSpecApplyConfiguration {
	b.LabelValueLengthLimit = &value
	return b
}

// WithKeepDroppedTargets sets the KeepDroppedTargets field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepDroppedTargets field is set to the value of the last call.
func (b *ScrapeConfigSpecApplyConfiguration) WithKeepDroppedTargets(value uint64) *ScrapeConfigSpecApplyConfiguration {
	b.KeepDroppedTargets = &value
	return b
}

// WithMetricRelabelConfigs adds the given value to the MetricRelabelConfigs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MetricRelabelConfigs field.
func (b *ScrapeConfigSpecApplyConfiguration) WithMetricRelabelConfigs(values ...**v1.RelabelConfig) *ScrapeConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMetricRelabelConfigs")
		}
		b.MetricRelabelConfigs = append(b.MetricRelabelConfigs, *values[i])
	}
	return b
}
