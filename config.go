package imgi

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	General   GeneralConfig
	Log       LogConfig
	Locations map[string]string
}

type GeneralConfig struct {
	Port int
}

type LogConfig struct {
	File  string
	Level string
}

func LoadConfig(configFile string) (Config, error) {
	config := Config{}
	cfg, err := ioutil.ReadFile(configFile)
	if err != nil {
		return config, err
	}

	if err := toml.Unmarshal(cfg, &config); err != nil {
		return config, err
	}

	if err := normalizeLocations(&config); err != nil {
		return config, err
	}

	return config, nil
}

func normalizeLocations(config *Config) error {
	if len(config.Locations) < 1 {
		return fmt.Errorf("locations must be configurated")
	}
	for k, v := range config.Locations {
		path, err := filepath.Abs(v)
		if err != nil {
			return err
		}
		config.Locations[k] = path
	}
	return nil
}
