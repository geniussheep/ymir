package dto

import "gitlab.benlai.work/go/ymir/sdk/pkg"

type QueryByAppId struct {
	AppId int `msgpack:"appId" json:"appId" form:"appId" xml:"appId" uri:"appId"  binding:"required" example:"1"`
}

func (q *QueryByAppId) CheckArgs() error {
	return pkg.CheckIntArgs("appId", q.AppId)
}

type CreateBase struct {
	// 应用发布类型 [Normal:完整发布, BuildImage:制作镜像, GoLive:应用上线, RestartApp:重启发布, Config:配置发布, ConfigRestart:配置发布, ConfigWithoutRestart:配置发布, QuicklyDeploy:快速部署, QuicklyRollback:快速回滚, Rollback:Rollback]
	DeployType string `msgpack:"deployType" json:"deployType" form:"deployType" uri:"deployType" xml:"deployType" example:"normal"`
	// 应用发布环境
	Environment string `msgpack:"environment" json:"environment" form:"environment" uri:"environment" xml:"environment" example:"branch"`
	// 是否全自动发布
	IsAutoPublish string `msgpack:"isAutoPublish" json:"isAutoPublish" form:"isAutoPublish" uri:"isAutoPublish" xml:"isAutoPublish" example:"true"`
}

type CreateBuild struct {
	// 应用Id
	AppId int `msgpack:"appId" json:"appId" form:"appId" xml:"appId" uri:"appId" example:"1"`
	// 应用别名 -- 应用容器内名称
	AppName string `msgpack:"appName" json:"appName" form:"appName" uri:"appName" xml:"appName" example:"mon"`

	CreateBase
}
