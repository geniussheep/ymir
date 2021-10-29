package config

import "gitlab.benlai.work/go/ymir/sdk/pkg/logger"

type Logger struct {
	Type      string `yaml:"type"`
	Path      string `yaml:"path"`
	Level     string `yaml:"level"`
	Stdout    string `yaml:"stdout"`
	EnabledDB bool   `yaml:"enabledDB"`
	Cap       uint   `yaml:"cap"`
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithCap(e.Cap),
	)
}

var LoggerConfig = new(Logger)