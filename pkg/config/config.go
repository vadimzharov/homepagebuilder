package config

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type panelConfig struct {
	Name        string `mapstructure:"name"`
	URL         string `mapstructure:"url"`
	ImageURL    string `mapstructure:"image"`
	ImageWidth  string `mapstructure:"imagewidth"`
	ImageHeight string `mapstructure:"imageheight"`
	Image       bool
	Description string
}

type DashboardConfig struct {
	Panels        []panelConfig `mapstructure:"panels"`
	ChatGPTAPIKey string        `mapstructure:"chatgpt_key"`
}

type CustomQueries struct {
	Mainpagequery string `mapstructure:"custommainpagequery"`
	Panelquery    string `mapstructure:"custompanelquery"`
}

func ReadConfig() (*DashboardConfig, *CustomQueries) {

	var dbconfig DashboardConfig
	var customQueries CustomQueries

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")     // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&dbconfig)

	tuneConfig(&dbconfig)

	var customQueriesConfigFile string

	if flag.Lookup("custom-queries") == nil {
		flag.StringVar(&customQueriesConfigFile, "custom-queries", "", "Use this flag to set configuration file with custom ChatGPT queries")
		flag.Parse()

		if customQueriesConfigFile != "" {
			log.Println("Using configuration file with custom queries", customQueriesConfigFile)

			customQueriesConfigFileName := filepath.Base(customQueriesConfigFile[:len(customQueriesConfigFile)-len(filepath.Ext(customQueriesConfigFile))])

			viper.SetConfigName(customQueriesConfigFileName) // name of config file (without extension)
			viper.SetConfigType("yaml")                      // REQUIRED if the config file does not have the extension in the name
			viper.AddConfigPath("./")                        // optionally look for config in the working directory
			err := viper.ReadInConfig()                      // Find and read the config file
			if err != nil {                                  // Handle errors reading the config file
				panic(fmt.Errorf("Fatal error config file: %w", err))
			}

			viper.Unmarshal(&customQueries)

		}

	}

	return &dbconfig, &customQueries

}

func tuneConfig(dbconfig *DashboardConfig) {

	for i := range dbconfig.Panels {
		if dbconfig.Panels[i].ImageURL != "" {
			dbconfig.Panels[i].Image = true
			if dbconfig.Panels[i].ImageHeight == "" {
				dbconfig.Panels[i].ImageHeight = "100"
			}
			if dbconfig.Panels[i].ImageWidth == "" {
				dbconfig.Panels[i].ImageWidth = "100"
			}
		}
	}

}
