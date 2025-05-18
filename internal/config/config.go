package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Env    string           `yaml:"env" env-default:"local"`
	Server HttpServerConfig `yaml:"http_server"`
	DB     DBConfig         `yaml:"database"`
}

type HttpServerConfig struct {
	Addr        string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"admin123"`
	Name     string `yaml:"name" env-default:"url"`
}

func LoadConfig(configPath string) *Config {
	var cfg Config
	cleanenv.ReadConfig(configPath, &cfg)
	return &cfg
}
