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

	Unmarshal(rawVal interface{}) error
}

var (
	cfg Config
)

func New(configPath string) Config {
	return initConfig(configPath)
}

func Get(key string) interface{} {
	return cfg.Get(key)
}

func GetBool(key string) bool {
	return cfg.GetBool(key)
}

func GetFloat64(key string) float64 {
	return GetFloat64(key)
}

func GetInt(key string) int {
	return cfg.GetInt(key)
}

func GetString(key string) string {
	return cfg.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return cfg.GetStringMap(key)
}
func GetStringMapString(key string) map[string]string {
	return cfg.GetStringMapString(key)
}
func GetStringSlice(key string) []string {
	return cfg.GetStringSlice(key)
}
func GetTime(key string) time.Time {
	return cfg.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return cfg.GetDuration(key)
}

func IsSet(key string) bool {
	return cfg.IsSet(key)
}

func AllSettings() map[string]interface{} {
	return cfg.AllSettings()
}

func Set(key string, value interface{}) {
	cfg.Set(key, value)
}

func RegisterAlias(alias string, key string) {
	cfg.RegisterAlias(alias, key)
}

func AutomaticEnv() {
	cfg.AutomaticEnv()
}

func BindEnv(input string) error {
	return cfg.BindEnv(input)
}

func SetEnvPrefix(envPrefix string) {
	cfg.SetEnvPrefix(envPrefix)
}

func AllowEmptyEnvVar(b bool) {
	cfg.AllowEmptyEnvVar(b)
}

func Unmarshal(rawVal interface{}) error {
	return cfg.Unmarshal(rawVal)
}
