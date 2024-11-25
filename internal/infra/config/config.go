package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type (
	Config struct {
		Env     string     `yaml:"env" env-default:"dev"`
		Postgre PostgreSQL `yaml:"postgres"`
	}
	PostgreSQL struct {
		Host     string `yaml:"host" env-default:"localhost"`
		Port     int    `yaml:"port" env-default:"5121"`
		User     string `yaml:"user" env-default:"postgres"`
		Password string `yaml:"password" env-default:"root"`
		Database string `yaml:"database" env-default:"students"`
	}
)

var instance *Config

func MustLoad() *Config {
	if instance != nil {
		return nil
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	instance = &cfg

	return instance
}
