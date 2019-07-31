package settings

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// DefaultDataTimeout is the default timeout for all operations performed in the MongoDB database
	DefaultDataTimeout = "5s"
	// DefaultDataURL is the default URL for the MongoDB database
	DefaultDataURL = "mongodb://localhost:27017"
	// DefaultPort is the default Port to use for the MongoDB database
	DefaultPort = "3080"
)

// Settings contain the configuration for the runtime operation of this API
type Settings struct {
	Data struct {
		Mongo struct {
			Database string
			Timeout  string
			URL      string
		}
	}
	Server struct {
		Port string
	}
}

// NewSettings provides a new settings object loaded with pertinent configuration
func NewSettings() (*Settings, error) {
	settings, err := initViper()
	if err != nil {
		return &settings, err
	}

	return &settings, err
}

func initViper() (Settings, error) {
	viper.SetConfigName("settings") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")        // Search the root directory for the configuration file
	viper.WatchConfig()             // Watch for changes to the configuration file and recompile

	// reload when there are configuration changes
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("settings file changes detected: %s", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("unable to re-read configuration changes found in %s", e.Name)
		}
	})

	// set configuration defaults
	viper.SetDefault("Data.Mongo.Timeout", DefaultDataTimeout)
	viper.SetDefault("Data.Mongo.URL", DefaultDataURL)
	viper.SetDefault("Server.Port", DefaultPort)

	// read configuration file in for the first time
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("unable to read settings file")
		return Settings{}, err
	}

	var settings Settings
	err := viper.Unmarshal(&settings)

	return settings, err
}
