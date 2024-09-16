package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("environment", "development")
	viper.SetDefault("service_name", "Go Microservice")
	viper.SetDefault("server_host", "localhost")
	viper.SetDefault("server_port", 4280)
	viper.SetDefault("sentry_url", "")
}
