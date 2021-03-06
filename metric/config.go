package metric

import (
	"github.com/nirnanaaa/asparagus/metric/adapters/cloudwatch"
	"github.com/nirnanaaa/asparagus/metric/adapters/console"
	"github.com/nirnanaaa/asparagus/metric/adapters/slack"
)

const (
	//DefaultMetricsEnable is default state for metrics
	DefaultMetricsEnable = false
)

// Config represents the meta configuration.
type Config struct {
	Enabled    bool              `toml:"enable-metrics"`
	Slack      slack.Config      `toml:"slack"`
	Console    console.Config    `toml:"console"`
	Cloudwatch cloudwatch.Config `toml:"cloudwatch"`
}

// NewConfig builds a new configuration with default values.
func NewConfig() *Config {
	return &Config{
		Enabled:    DefaultMetricsEnable,
		Cloudwatch: cloudwatch.NewConfig(),
		Slack:      slack.NewConfig(),
		Console:    console.NewConfig(),
	}
}

// Validate the config
func (c *Config) Validate() error {
	return nil
}
