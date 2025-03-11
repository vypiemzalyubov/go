package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func readConfig(configPath string, config interface{}) error {
	if configPath == `` {
		return fmt.Errorf(`no config path`)
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", configPath, err)
	}

	if err = yaml.Unmarshal(configBytes, config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

func NewConfig() (Config, error) {
	var config Config

	err := readConfig("../../internal/infrastructure/config-local-grpc.yml", &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
