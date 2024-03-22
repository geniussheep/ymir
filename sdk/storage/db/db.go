package db

import (
	"fmt"
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/geniussheep/ymir/sdk/pkg"
	"os"
	"time"

	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/storage/db"

	dbLogger "gorm.io/gorm/logger"
)

func getLogLevel(l string) dbLogger.LogLevel {
	switch l {
	case "silent":
		return dbLogger.Silent
	case "error":
		return dbLogger.Error
	case "warn":
		return dbLogger.Warn
	case "info":
		return dbLogger.Info
	default:
		return dbLogger.Info
	}
}

func getLogSlowThreshold(ms int) time.Duration {
	if ms <= 0 {
		return time.Second
	}
	return time.Millisecond * time.Duration(ms)
}

func Setup() {
	for n, cfg := range config.DatabaseConfig {
		yorm, err := db.New(
			db.SetDsn(cfg.Dsn),
			db.SetDriver(cfg.Driver),
			db.SetLogLevel(getLogLevel(cfg.LogConfig.LogLevel)),
			db.SetLogColorful(cfg.LogConfig.Colorful),
			db.SetLogIgnoreRecordNotFoundError(cfg.LogConfig.IgnoreRecordNotFoundError),
			db.SetLogSlowThreshold(getLogSlowThreshold(cfg.LogConfig.SlowThreshold)),
		)

		if err != nil {
			logger.Fatal(pkg.Red(fmt.Sprintf("db-%s:[%v] connect error: %s", n, cfg, err)))
			os.Exit(0)
		}
		sdk.Runtime.SetDb(n, yorm)
	}
}
