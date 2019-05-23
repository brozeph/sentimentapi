package settings

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	DefaultPort = "3080"
)

type Settings struct {
	Data struct {
		Mongo struct {
			URL      string
			Database string
		}
	}
	Server struct {
		Port string
	}
}

func New() (*Settings, error) {
	settings, err := initViper()
	if err != nil {
		return &settings, err
	}

	return &settings, err
}

func initViper() (Settings, error) {
	viper.SetConfigName("settings") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")        // Search the root directory for the configuration file

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return Settings{}, err
	}

	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("settings file changes detected: %s", e.Name)
	})

	viper.SetDefault("Server.Port", DefaultPort)
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("error encountered while reading settings file: %s", err)
	}

	var settings Settings
	err = viper.Unmarshal(&settings)

	return settings, err
}
