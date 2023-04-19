package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type panelConfig struct {
	Name        string `mapstructure:"name"`
	URL         string `mapstructure:"url"`
	Icon        string `mapstructure:"icon"`
	ImageWidth  string
	ImageHeight string
	Image       bool
	Description string
}

type DashboardConfig struct {
	Panels        []panelConfig `mapstructure:"panels"`
	ChatGPTAPIKey string        `mapstructure:"chatgpt_key"`
}

func ReadConfig() *DashboardConfig {

	var dbconfig DashboardConfig

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&dbconfig)

	return &dbconfig

}
