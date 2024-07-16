package k8s

import (
	"fmt"
	"github.com/geniussheep/ymir/k8s"
	"github.com/geniussheep/ymir/logger"
	"github.com/geniussheep/ymir/sdk"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/geniussheep/ymir/sdk/pkg"
	"os"
)

func Setup() {
	for n, cfg := range config.K8SConfig {
		client, err := k8s.New(
			k8s.SetOutOfCluster(cfg.OutOfCluster),
			k8s.SetKubeConfigPath(cfg.KubeConfigPath),
			k8s.SetWebBaseURL(cfg.WebBaseURL),
		)
		if err != nil {
			logger.Fatal(pkg.Red(fmt.Sprintf("k8s-%s:[%v] connect error: %s", n, cfg, err)))
			os.Exit(0)
		}
		sdk.Runtime.SetK8S(n, client)
	}

}
