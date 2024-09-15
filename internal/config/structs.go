package config

type Config struct {
	ServiceName string `mapstructure:"service_name"`
	ServerHost  string `mapstructure:"server_host"`
	ServerPort  int    `mapstructure:"server_port"`
}
