package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server struct {
		Bucket    string
		AccessKey string
		SecretKey string
		Region    string
		SourceDir string
	}
}

func (c Config) IsNotSet() bool {
	return c.Server.Bucket == "" || c.Server.AccessKey == "" || c.Server.SecretKey == ""
}

func GetConfig() (Config, error) {
	if config.IsNotSet() {

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")

		if err := viper.ReadInConfig(); err != nil {
			return Config{}, fmt.Errorf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}

		// Read AccessKey and SecretKey from environment variables
		accessKey := os.Getenv("ACCESS_KEY")
		secretKey := os.Getenv("SECRET_KEY")

		if accessKey != "" {
			config.Server.AccessKey = accessKey
		}
		if secretKey != "" {
			config.Server.SecretKey = secretKey
		}

		if config.IsNotSet() {
			return Config{}, fmt.Errorf("Configuration file is missing values")
		}
	}

	return config, nil
}

func (c *Config) ShowConfig() {
	fmt.Printf("Server Bucket: %s\n", c.Server.Bucket)
	fmt.Printf("Server AccessKey: %s\n", c.Server.AccessKey)
	fmt.Printf("Database SecretKey: %s\n", c.Server.SecretKey)
}
