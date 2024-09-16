package config

type Config struct {
	Environment string `mapstructure:"environment"`
	ServiceName string `mapstructure:"service_name"`
	ServerHost  string `mapstructure:"server_host"`
	ServerPort  int    `mapstructure:"server_port"`
	SentryUrl   string `mapstructure:"sentry_url"`
}
