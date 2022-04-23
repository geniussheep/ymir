package db

import (
	"fmt"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/config"
	"gitlab.benlai.work/go/ymir/sdk/pkg"
	"os"
	"time"

	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/storage/db"

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
		return dbLogger.Warn
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
			db.SetUseDbms(cfg.UseDbms),
			db.SetLogLevel(getLogLevel(cfg.LogConfig.LogLevel)),
			db.SetLogColorful(cfg.LogConfig.Colorful),
			db.SetLogIgnoreRecordNotFoundError(cfg.LogConfig.IgnoreRecordNotFoundError),
			db.SetLogSlowThreshold(getLogSlowThreshold(cfg.LogConfig.SlowThreshold)),
		)

		if err != nil {
			logger.Fatal(pkg.Red(fmt.Sprintf("db:%s connect error: %s", cfg.Driver, err)))
			os.Exit(0)
		}
		sdk.Runtime.SetDb(n, yorm)
	}
}
