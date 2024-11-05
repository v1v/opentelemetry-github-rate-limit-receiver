package githubratelimit

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/v1v/opentelemetry-github-rate-limit-receiver/internal/metadata"
)

// Config defines the configuration for the TCP stats receiver.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"` // ScraperControllerSettings to configure scraping interval (default: 10s)
	metadata.MetricsBuilderConfig           `mapstructure:",squash"` // MetricsBuilderConfig to enable/disable specific metrics (default: all enabled)
}

func createDefaultConfig() component.Config {
	return &Config{
		ScraperControllerSettings: scraperhelper.NewDefaultScraperControllerSettings(metadata.Type),
		MetricsBuilderConfig:      metadata.DefaultMetricsBuilderConfig(),
	}
}

func (c Config) Validate() error {

	return nil
}
