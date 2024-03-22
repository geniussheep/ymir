package zookeeper

import (
	"fmt"
	"github.com/geniussheep/ymir/component/zookeeper"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/geniussheep/ymir/sdk/pkg"
	"os"
)

func Setup() {
	for n, cfg := range config.ZookeeperConfig {
		zk, err := zookeeper.New(
			zookeeper.SetServers(cfg.Servers),
			zookeeper.SetSessionTimeout(cfg.SessionTimeout),
		)

		if err != nil {
			logger.Fatal(pkg.Red(fmt.Sprintf("zk-%s:[%v] connect error: %s", n, cfg, err)))
			os.Exit(0)
		}
		sdk.Runtime.SetZookeeper(n, zk)
	}
}
