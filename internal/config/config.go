package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string        `yaml:"env"`
	Server *ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func MustNewConfig(filename string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(filename, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
