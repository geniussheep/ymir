package apiServer

import (
	"context"
	"fmt"
	"gitlab.benlai.work/go/ymir/sdk/common"
	"gitlab.benlai.work/go/ymir/sdk/component/zookeeper"
	"gitlab.benlai.work/go/ymir/sdk/middleware"
	"gitlab.benlai.work/go/ymir/sdk/storage/db"
	"gitlab.benlai.work/go/ymir/sdk/storage/redis"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gitlab.benlai.work/go/ymir/cli"
	"gitlab.benlai.work/go/ymir/sdk"
	"gitlab.benlai.work/go/ymir/sdk/api"
	"gitlab.benlai.work/go/ymir/sdk/config"
	"gitlab.benlai.work/go/ymir/sdk/pkg"
)

type apiServerFlag struct {
	cmd *cobra.Command
}

func New(p *cli.Program) cli.FlagSet {
	f := apiServerFlag{
		cmd: &cobra.Command{
			Use:          "server",
			Short:        "start app soa and esb web api server",
			Example:      fmt.Sprintf("%s server -c %s", p.Program, common.DEFAULT_CONFIG_FILE_PATH),
			SilenceUsage: true,
			PreRun: func(cmd *cobra.Command, args []string) {
				setup(p)
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return run(p)
			},
		},
	}
	f.cmd.PersistentFlags().StringVarP(&p.ConfigFilePath, "config", "c", common.DEFAULT_CONFIG_FILE_PATH, "start server with provided configuration file")
	return &f
}

func (f *apiServerFlag) Cmd() *cobra.Command {
	return f.cmd
}

func run(p *cli.Program) error {
	if config.ApplicationConfig.Mode == pkg.Prod.String() {
		gin.SetMode(gin.ReleaseMode)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ApplicationConfig.HttpPort),
		Handler: sdk.Runtime.GetEngine(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		// 服务连接
		log.Println(fmt.Sprintf("starting listen app port:%d...", config.ApplicationConfig.HttpPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}
	}()

	usageStr := `欢迎使用 ` + pkg.Green(fmt.Sprintf("%s - %s", p.Program, p.Version)) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
	fmt.Println(pkg.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", config.ApplicationConfig.HttpPort)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", pkg.GetLocalHost(), config.ApplicationConfig.HttpPort)
	fmt.Println(pkg.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/swagger/index.html \r\n", config.ApplicationConfig.HttpPort)
	fmt.Printf("-  Network: http://%s:%d/swagger/index.html \r\n", pkg.GetLocalHost(), config.ApplicationConfig.HttpPort)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", pkg.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", pkg.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return nil
}

func setup(p *cli.Program) {
	config.ExtendConfig = p.ExtendConfig

	config.Setup(p.ConfigFilePath, db.Setup, redis.Setup, zookeeper.Setup)

	for _, sf := range p.InitFuncArray {
		sf()
	}

	var r *gin.Engine
	var h http.Handler
	h = gin.New()
	sdk.Runtime.SetEngine(h)
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	r.Use(api.SetRequestLogger)

	if p.MiddleWareFuncArray != nil && len(p.MiddleWareFuncArray) > 0 {
		for _, mwf := range p.MiddleWareFuncArray {
			middleware.Append(mwf)
		}
	}

	// 初始化中间件
	middleware.InitMiddleware(r)

	if p.AppRoutersScan != nil && len(p.AppRoutersScan) > 0 {
		for _, srf := range p.AppRoutersScan {
			srf()
		}
	}

	sdk.Runtime.GetWebApi().RegisterRouters(r)
}
