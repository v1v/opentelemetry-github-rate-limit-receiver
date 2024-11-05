// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for githubratelimit metrics.
type MetricsConfig struct {
	GithubAPIRateLimitCoreRemaining MetricConfig `mapstructure:"github.api.rate-limit.core.remaining"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		GithubAPIRateLimitCoreRemaining: MetricConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for githubratelimit metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}