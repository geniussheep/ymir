package config

type Database struct {
	Dsn       string         `yaml:"dsn"  mapstructure:"dsn"`
	Driver    string         `yaml:"driver"  mapstructure:"driver"`
	UseDbms   bool           `yaml:"useDbms" mapstructure:"useDbms"`
	LogConfig DatabaseLogger `yaml:"logConfig" mapstructure:"logConfig"`
}

type DatabaseLogger struct {
	SlowThreshold             int    `yaml:"slowThresholdMS" mapstructure:"slowThresholdMS"`                      // Slow SQL threshold
	LogLevel                  string `yaml:"level"  mapstructure:"level"`                                         // Log level
	IgnoreRecordNotFoundError bool   `yaml:"ignoreRecordNotFoundError"  mapstructure:"ignoreRecordNotFoundError"` // Ignore ErrRecordNotFound error for logger
	Colorful                  bool   `yaml:"colorful"  mapstructure:"colorful"`                                   // Disable color
}

var DatabaseConfig = make(map[string]*Database)
