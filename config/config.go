package config

import (
	_ "embed"
	"fmt"
	"github.com/goccy/go-yaml"
	"os"
	"sync"
)

type Config struct {
	Registry struct {
		RegisterAddress []string `yaml:"register_address"`
		UserName        string   `yaml:"user_name"`
		Password        string   `yaml:"password"`
	} `yaml:"Registry"`
	Hertz struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"Hertz"`
}

var (
	configInstance *Config
	once           sync.Once
)

//go:embed config.yaml
var data []byte

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		configInstance, err = loadConfig()
	})
	return configInstance, err
}

func loadConfig() (*Config, error) {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Registry
	registryAddrEnv := os.Getenv("REGISTRY_ADDR")
	if registryAddrEnv != "" {
		conf.Registry.RegisterAddress = []string{registryAddrEnv}
	}
	registryUserNameEnv := os.Getenv("REGISTRY_USERNAME")
	if registryUserNameEnv != "" {
		conf.Registry.UserName = ""
	}
	registryPasswordEnv := os.Getenv("REGISTRY_PASSWORD")
	if registryPasswordEnv != "" {
		conf.Registry.Password = registryPasswordEnv
	}
	return &conf, nil
}
