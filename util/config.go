package util

import (
	"fmt"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/spf13/viper"
)

type Config struct {
	PrimeURL      string
	RegionURLs    []string
	ZoneURLs      [][]string
	Location      common.Location
	Auto          bool
	Optimize      bool
	OptimizeTimer int
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err = viper.ReadInConfig()    // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	err = viper.Unmarshal(&config)
	return
}
