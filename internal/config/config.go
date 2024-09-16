package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const envPrefix = "app"

var C Config

func Load() {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found (skipping): %v", err)
	}

	viper.SetEnvPrefix(envPrefix)

	setDefaults()

	viper.AutomaticEnv()

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode config variables into struct: %v", err)
	}
}
