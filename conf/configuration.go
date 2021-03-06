package conf

import (
	"github.com/netlify/netlify-commons/messaging"
	"github.com/netlify/netlify-commons/nconf"
)

type Config struct {
	LogConf     nconf.LoggingConfig  `mapstructure:"log_conf"`
	NatsConf    *NatsConfig          `mapstructure:"nats_conf"`
	MetricsConf *nconf.MetricsConfig `mapstructure:"metrics_conf"`
	DataCenter  string               `mapstructure:"data_center"`
	NumWorkers  int                  `mapstructure:"num_workers"`
}

type NatsConfig struct {
	messaging.NatsConfig `mapstructure:",squash"`
	Subject              string `mapstructure:"command_subject"`
	Group                string `mapstructure:"command_group"`
	DurableName          string `mapstructure:"durable_name"`
}
