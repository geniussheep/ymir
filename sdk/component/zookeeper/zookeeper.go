package zookeeper

import (
	"fmt"
	"gitlab.benlai.work/go/ymir/component/zookeeper"
	"gitlab.benlai.work/go/ymir/logger"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/config"
	"gitlab.benlai.work/go/ymir/sdk/pkg"
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
