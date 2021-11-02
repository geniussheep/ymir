package config

import (
	"gitlab.benlai.work/go/ymir/config"
	"gitlab.benlai.work/go/ymir/logger"
	"os"
	"path/filepath"
	"sync"
)

var (
	setting *Setting
	mux     sync.RWMutex
	isInit  bool = false
)

// Setting 配置集合
type Setting struct {
	Application *Application `yaml:"application"`
	Logger      *Logger      `yaml:"logger"`
	Databases   *Database    `yaml:"databases"`
}

func Instance() *Setting {
	if isInit {
		return setting
	}
	Default()
	return setting
}

func Default() {
	defaultConfigPath := "conf/config.yaml"
	logger.DefaultLogger.Logf(logger.InfoLevel, "load default config:%s", defaultConfigPath)
	Setup(defaultConfigPath)
}

func Setup(configPath string) {
	if isInit {
		return
	}
	mux.Lock()
	defer mux.Unlock()
	currentPath, _ := os.Getwd()
	configPath = filepath.Join(currentPath, configPath)
	logger.DefaultLogger.Logf(logger.InfoLevel, "will load config: %s", configPath)
	c := config.New(configPath)
	err := c.Unmarshal(&setting)
	if err != nil {
		logger.DefaultLogger.Logf(logger.ErrorLevel, "load config failed, error: %s", err)
		os.Exit(1)
	}
	isInit = true
	logger.DefaultLogger.Logf(logger.InfoLevel, "load config success")
}
