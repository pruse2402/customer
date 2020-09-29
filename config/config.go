package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Name     string `yaml:"name", envconfig:"DB_NAME"`
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"password", envconfig:"DB_PASSWORD"`
		Address  string `yaml:"address", envconfig:"DB_ADDRESS"`
	} `yaml:"database"`
}

var Cfg Config

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config) {
	f, err := os.Open("config/dev.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
