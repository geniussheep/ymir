package cli

import "github.com/gin-gonic/gin"

type CmdInfo struct {
	// Command
	Command string

	// [命令|程序]简短描述
	Short string

	// [命令|程序]完整描述
	Long string
}

// Program cli程序信息
type Program struct {
	// 程序名
	Program string

	//程序描述
	Desc string

	// 程序版本
	Version string

	// 应用的配置文件路径
	ConfigFilePath string

	// 应用的扩展配置
	ExtendConfig any

	// 应用的ApiRouter扫描
	AppRoutersScan []func()

	// gin的中间中间件加载
	MiddleWareFuncArray []gin.HandlerFunc

	// 初始化方法
	// 加载默认中间件方法 middleware.AppendDefault()
	InitFuncArray []func()
}
