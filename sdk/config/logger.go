package config

import (
	"github.com/geniussheep/ymir/sdk/pkg/logger"
	"go.uber.org/zap"
)

type Logger struct {
	Type   string     `yaml:"type" mapstructure:"type"`
	Path   string     `yaml:"path" mapstructure:"path"`
	Level  string     `yaml:"level" mapstructure:"level"`
	Stdout string     `yaml:"stdout" mapstructure:"stdout"`
	Cap    uint       `yaml:"cap" mapstructure:"cap"`
	ZapCfg zap.Config `yaml:"zapCfg" mapstructure:"zapCfg"`
}

func getPath(path string) string {
	if len(path) <= 0 {
		return "logs"
	}
	return path
}

func getType(t string) string {
	if len(t) <= 0 {
		return "zap"
	}
	return t
}

func getLevel(level string) string {
	if len(level) <= 0 {
		return "info"
	}
	return level
}

func getStdout(stdout string) string {
	if len(stdout) <= 0 {
		return "default"
	}
	return stdout
}

func getZapCfg(cfg zap.Config) zap.Config {
	return cfg
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(getType(e.Type)),
		logger.WithPath(getPath(e.Path)),
		logger.WithLevel(getLevel(e.Level)),
		logger.WithStdout(getStdout(e.Stdout)),
		logger.WithCap(e.Cap),
		logger.WithZapCfg(getZapCfg(e.ZapCfg)),
	)
}

var LoggerConfig = new(Logger)
