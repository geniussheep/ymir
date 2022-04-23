package config

type Tls struct {
	Cert string `yaml:"cert" mapstructure:"cert"`
	Key  string `yaml:"key" mapstructure:"key"`
	Ca   string `yaml:"ca" mapstructure:"ca"`
}

type Redis struct {
	Addr       string `yaml:"addr" mapstructure:"addr"`
	Username   string `yaml:"username" mapstructure:"username"`
	Password   string `yaml:"password" mapstructure:"password"`
	DB         int    `yaml:"db" mapstructure:"db"`
	PoolSize   int    `yaml:"poolSize" mapstructure:"poolSize"`
	MaxRetries int    `yaml:"maxRetries" mapstructure:"maxRetries"`
	//Tls        *Tls   `yaml:"tls" mapstructure:"tls"`
}

var RedisConfig = make(map[string]*Redis)
