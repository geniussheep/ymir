package config

import (
	"time"
)

type Config interface {
	//Init(path string, configName string, configType string)
	//AddConfigPath(path ...string)

	Init(configPath string)

	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
	AllSettings() map[string]interface{}

	Set(key string, value interface{})
	RegisterAlias(alias string, key string)

	AutomaticEnv()
	BindEnv(input string) error
	SetEnvPrefix(envPrefix string)
	AllowEmptyEnvVar(b bool)

	Unmarshal(rawVal interface{})
}

var (
	Cfg Config
)

func New(configPath string) Config {
	return initConfig(configPath)
}

func Get(key string) interface{} {
	return Cfg.Get(key)
}

func GetBool(key string) bool {
	return Cfg.GetBool(key)
}

func GetFloat64(key string) float64 {
	return GetFloat64(key)
}

func GetInt(key string) int {
	return Cfg.GetInt(key)
}

func GetString(key string) string {
	return Cfg.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return Cfg.GetStringMap(key)
}
func GetStringMapString(key string) map[string]string {
	return Cfg.GetStringMapString(key)
}
func GetStringSlice(key string) []string {
	return Cfg.GetStringSlice(key)
}
func GetTime(key string) time.Time {
	return Cfg.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return Cfg.GetDuration(key)
}

func IsSet(key string) bool {
	return Cfg.IsSet(key)
}

func AllSettings() map[string]interface{} {
	return Cfg.AllSettings()
}

func Set(key string, value interface{}) {
	Cfg.Set(key, value)
}

func RegisterAlias(alias string, key string) {
	Cfg.RegisterAlias(alias, key)
}

func AutomaticEnv() {
	Cfg.AutomaticEnv()
}

func BindEnv(input string) error {
	return Cfg.BindEnv(input)
}

func SetEnvPrefix(envPrefix string) {
	Cfg.SetEnvPrefix(envPrefix)
}

func AllowEmptyEnvVar(b bool) {
	Cfg.AllowEmptyEnvVar(b)
}

func Unmarshal(rawVal interface{}) {
	Cfg.Unmarshal(rawVal)
}
