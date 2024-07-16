package config

import "C"
import (
	"github.com/geniussheep/ymir/config"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/sdk/common"
	"os"
	"path/filepath"
	"sync"
)

var (
	ExtendConfig any
	setting      *Setting
	mux          sync.RWMutex
	isInit       bool = false
)

// Setting 配置集合
type Setting struct {
	Application *Application           `yaml:"application"`
	Logger      *Logger                `yaml:"logger"`
	Databases   *map[string]*Database  `yaml:"databases"`
	Redis       *map[string]*Redis     `yaml:"redis"`
	Zookeeper   *map[string]*Zookeeper `yaml:"zookeeper"`
	K8S         *map[string]*K8S       `yaml:"k8s"`
	Http        *Http                  `yaml:"http"`
	Extend      any                    `yaml:"extend"`
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
	defaultConfigPath := common.DefaultConfigFilePath
	logger.DefaultLogger.Logf(logger.InfoLevel, "load default config:%s", defaultConfigPath)
	Setup(defaultConfigPath)
}

//// 设置Extend的默认值
//func setExtDV(c *config.Config, ext any) {
//	if ext == nil {
//		return
//	}
//	extType := reflect.ValueOf(ext)
//	if extType.Kind() != reflect.Struct {
//		return
//	}
//
//	for i := 0; i < extType.NumField(); i++ {
//		switch extType.Type().Field(i).Type().Kind() {
//
//		case reflect.Struct:
//			break;
//		case reflect.Bool:
//		case reflect.Int:
//		case reflect.Int8:
//		case reflect.Int16:
//		case reflect.Int32:
//		case reflect.Int64:
//		case reflect.Uint:
//		case reflect.Uint8:
//		case reflect.Uint16:
//		case reflect.Uint32:
//		case reflect.Uint64:
//		case reflect.Uintptr:
//		case reflect.Float32:
//		case reflect.Float64:
//		case reflect.String:
//			break
//		case reflect.Complex64:
//		case reflect.Complex128:
//		case reflect.Array:
//		case reflect.Chan:
//		case reflect.Func:
//		case reflect.Interface:
//		case reflect.Map:
//		case reflect.Pointer:
//		case reflect.Slice:
//		case reflect.Invalid:
//			break;
//		}
//	}
//}

func Setup(configPath string, cbs ...func()) {
	if isInit {
		return
	}
	mux.Lock()
	defer mux.Unlock()
	currentPath, _ := os.Getwd()
	configPath = filepath.Join(currentPath, configPath)
	logger.DefaultLogger.Logf(logger.InfoLevel, "will load config: %s", configPath)
	c, err := config.New(configPath)
	if err != nil {
		logger.DefaultLogger.Logf(logger.ErrorLevel, "load config failed, error: %s", err)
		os.Exit(1)
	}
	setting = &Setting{
		Application: ApplicationConfig,
		Logger:      LoggerConfig,
		Databases:   &DatabaseConfig,
		Redis:       &RedisConfig,
		Zookeeper:   &ZookeeperConfig,
		K8S:         &K8SConfig,
		Http:        HttpConfig,
		Extend:      ExtendConfig,
		callbacks:   cbs,
	}
	err = c.Unmarshal(&setting)
	if err != nil {
		logger.DefaultLogger.Logf(logger.ErrorLevel, "load config failed, error: %s", err)
		os.Exit(1)
	}
	setting.Http.SetDefault()
	isInit = true
	setting.init()
	logger.DefaultLogger.Logf(logger.InfoLevel, "load config success")
}
