package config

import (
	"gitlab.benlai.work/go/ymir/config"
	"gitlab.benlai.work/go/ymir/logger"
	"os"
	"path/filepath"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
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
	_cfg, err := Setup("conf/config.yaml")
	if err != nil {
		logger.DefaultLogger.Logf(logger.ErrorLevel, "load config failed, error: %s", err)
		os.Exit(1)
	}
	return _cfg
}

func Setup(configPath string) (*Config, error) {
	currentPath, _ := os.Getwd()
	configPath = filepath.Join(currentPath, configPath)
	logger.DefaultLogger.Logf(logger.InfoLevel, "will load config: %s", configPath)
	c := config.New(configPath)
	err := c.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
