package dto

import "gitlab.benlai.work/go/ymir/sdk/pkg"

type QueryByAppId struct {
	AppId int `msgpack:"appId" json:"appId" form:"appId" xml:"appId" uri:"appId"  binding:"required" example:"1"`
}

func (q *QueryByAppId) CheckArgs() error {
	return pkg.CheckIntArgs("appId", q.AppId)
}
