package redis

import (
	"fmt"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/geniussheep/ymir/sdk/pkg"
	"github.com/geniussheep/ymir/storage/redis"
	"os"
)

func Setup() {
	for n, cfg := range config.RedisConfig {
		r, err := redis.New(
			redis.SetAddr(cfg.Addr),
			redis.SetUsername(cfg.Username),
			redis.SetPassword(cfg.Password),
			redis.SetDB(cfg.DB),
			redis.SetPoolSize(cfg.PoolSize),
			redis.SetMaxRetries(cfg.MaxRetries),
		)
		if err != nil {
			logger.Fatal(pkg.Red(fmt.Sprintf("redis-%s:[%v] connect error: %s", n, cfg, err)))
			os.Exit(0)
		}
		sdk.Runtime.SetRedis(n, r)
	}
}
