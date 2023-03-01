package config

type Zookeeper struct {
	Servers        string `yaml:"servers"  mapstructure:"servers"`
	SessionTimeout int64  `yaml:"sessionTimeout"  mapstructure:"sessionTimeout"`
}

var ZookeeperConfig = make(map[string]*Zookeeper)
