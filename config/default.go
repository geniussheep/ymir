package config

import (
	"github.com/spf13/viper"
	"time"
)

type config struct {
	instance *viper.Viper
}

//func (c *config) Init(path string, configName string, configType string) {
//	c.instance = viper.New()
//	c.instance.AddConfigPath(path)
//	c.instance.SetConfigName(configName)
//	c.instance.SetConfigType(configType)
//}
//
//func (c *config) AddConfigPath(path ...string) {
//	for _, p := range path {
//		c.instance.AddConfigPath(p)
//	}
//}

func initConfig(configPath string) (Config, error) {
	c := config{instance: viper.New()}
	c.instance.SetConfigFile(configPath)
	err := c.instance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func New(configPath string) (Config, error) {
	return initConfig(configPath)
}

func (c *config) Get(key string) interface{} {
	return c.instance.Get(key)
}
func (c *config) GetBool(key string) bool {
	return c.instance.GetBool(key)
}
func (c *config) GetFloat64(key string) float64 {
	return c.instance.GetFloat64(key)
}
func (c *config) GetInt(key string) int {
	return c.instance.GetInt(key)
}
func (c *config) GetString(key string) string {
	return c.instance.GetString(key)
}
func (c *config) GetStringMap(key string) map[string]interface{} {
	return c.instance.GetStringMap(key)
}
func (c *config) GetStringMapString(key string) map[string]string {
	return c.instance.GetStringMapString(key)
}
func (c *config) GetStringSlice(key string) []string {
	return c.instance.GetStringSlice(key)
}
func (c *config) GetTime(key string) time.Time {
	return c.instance.GetTime(key)
}
func (c *config) GetDuration(key string) time.Duration {
	return c.instance.GetDuration(key)
}
func (c *config) IsSet(key string) bool {
	return c.instance.IsSet(key)
}
func (c *config) AllSettings() map[string]interface{} {
	return c.instance.AllSettings()
}

func (c *config) Set(key string, value interface{}) {
	c.instance.Set(key, value)
}
func (c *config) RegisterAlias(alias string, key string) {
	c.instance.RegisterAlias(alias, key)
}

func (c *config) AutomaticEnv() {
	c.instance.AutomaticEnv()
}
func (c *config) BindEnv(input string) error {
	return c.instance.BindEnv(input)
}
func (c *config) SetEnvPrefix(envPrefix string) {
	c.instance.SetEnvPrefix(envPrefix)
}
func (c *config) AllowEmptyEnvVar(b bool) {
	c.instance.AllowEmptyEnv(b)
}

func (c *config) Unmarshal(rawVal interface{}) error {
	return c.instance.Unmarshal(rawVal)
}

func (c *config) Viper() *viper.Viper {
	return c.instance
}
