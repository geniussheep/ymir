package db

import (
	"gorm.io/gorm/logger"
	"time"
)

type Option func(*options)

type options struct {
	Dsn       string
	Driver    string
	logConfig logger.Config
}

func setDefault() options {
	return options{
		Dsn:    "",
		Driver: "",
		logConfig: logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	}
}

func SetDsn(dsn string) Option {
	return func(o *options) {
		o.Dsn = dsn
	}
}

func SetDriver(driver string) Option {
	return func(o *options) {
		o.Driver = driver
	}
}

func SetLogLevel(l logger.LogLevel) Option {
	return func(o *options) {
		o.logConfig.LogLevel = l
	}
}

func SetLogColorful(b bool) Option {
	return func(o *options) {
		o.logConfig.Colorful = b
	}
}

func SetLogSlowThreshold(t time.Duration) Option {
	return func(o *options) {
		o.logConfig.SlowThreshold = t
	}
}

func SetLogIgnoreRecordNotFoundError(b bool) Option {
	return func(o *options) {
		o.logConfig.IgnoreRecordNotFoundError = b
	}
}
