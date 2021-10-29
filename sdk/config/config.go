package config

// Config 配置集合
type Config struct {
	Application *Application `yaml:"application"`
	Logger      *Logger      `yaml:"logger"`
	Databases   *Database    `yaml:"databases"`
	Extend      interface{}  `yaml:"extend"`
}

