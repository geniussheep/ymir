package config

import (
	"gitlab.benlai.work/go/ymir/config"
	"sync"
)

var (
	cfg *Config
	once   sync.Once
)

// Config 配置集合
type Config struct {
	Application *Application `yaml:"application"`
	Logger      *Logger      `yaml:"logger"`
	Databases   *Database    `yaml:"databases"`
}

func Current() Config {
	once.Do(func() {
		cfg = Default()
	})

	return *cfg
}

func Default() *Config {
	return Setup("./conf/config.yaml")
}

func Setup(configPath string) *Config {
	c := config.New(configPath)
	c.Unmarshal(cfg)
	return cfg
}