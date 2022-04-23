package config

import "C"
import (
	"gitlab.benlai.work/go/ymir/config"
	"gitlab.benlai.work/go/ymir/logger"
	"os"
	"path/filepath"
	"sync"
)

var (
	ExtendConfig interface{}
	setting      *Setting
	mux          sync.RWMutex
	isInit       bool = false
)

// Setting 配置集合
type Setting struct {
	Application *Application          `yaml:"application"`
	Logger      *Logger               `yaml:"logger"`
	Databases   *map[string]*Database `yaml:"databases"`
	Redis       *map[string]*Redis    `yaml:"redis"`
	Extend      interface{}           `yaml:"extend"`
	callbacks   []func()
}

func Instance() *Setting {
	if isInit {
		return setting
	}
	Default()
	return setting
}

func (cfg *Setting) runCallback() {
	for i := range cfg.callbacks {
		cfg.callbacks[i]()
	}
}

func (cfg *Setting) init() {
	cfg.Logger.Setup()
	cfg.runCallback()
}

func Default() {
	defaultConfigPath := "conf/config.yaml"
	logger.DefaultLogger.Logf(logger.InfoLevel, "load default config:%s", defaultConfigPath)
	Setup(defaultConfigPath)
}

func Setup(configPath string, cbs ...func()) {
	if isInit {
		return
	}
	mux.Lock()
	defer mux.Unlock()
	currentPath, _ := os.Getwd()
	configPath = filepath.Join(currentPath, configPath)
	logger.DefaultLogger.Logf(logger.InfoLevel, "will load config: %s", configPath)
	c := config.New(configPath)
	setting = &Setting{
		Application: ApplicationConfig,
		Logger:      LoggerConfig,
		Databases:   &DatabaseConfig,
		Redis:       &RedisConfig,
		Extend:      ExtendConfig,
		callbacks:   cbs,
	}
	err := c.Unmarshal(&setting)
	if err != nil {
		logger.DefaultLogger.Logf(logger.ErrorLevel, "load config failed, error: %s", err)
		os.Exit(1)
	}
	isInit = true
	setting.init()
	logger.DefaultLogger.Logf(logger.InfoLevel, "load config success")
}
