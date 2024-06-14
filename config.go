package notify

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PushPlus PushPlusConfig `yaml:"pushplus"`
}

type PushPlusConfig struct {
	Token string `json:"token" yaml:"token"`
}

// LoadConfig loads the configuration from the specified YAML file.
// It reads the file, unmarshals the YAML data, and returns the loaded configuration.
func LoadConfig(configYamlFile string) (config *Config, err error) {
	// read file
	data, err := os.ReadFile(configYamlFile)
	if err != nil {
		return nil, fmt.Errorf("Load:os.ReadFile-> %w", err)
	}

	// unmarshal yaml
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Load:yaml.Unmarshal-> %w", err)
	}

	return
}
