package model

import "github.com/geniussheep/ymir/sdk/pkg"

type PageInfo struct {
	// PageNum 页码
	PageNum int `msgpack:"pageNum" form:"pageNum" xml:"pageNum" uri:"pageNum" json:"pageNum" example:"1"`
	// PageSize 每页大小
	PageSize int `msgpack:"pageSize" form:"pageSize" xml:"pageSize" uri:"pageSize" json:"pageSize" example:"10"`
}

type Base struct {
	// Environment 环境参数 Enums(branch, trunk, prepare, online)
	Environment string `msgpack:"environment" form:"environment" xml:"environment" uri:"environment" json:"environment" example:"trunk"`
}

type RequestPaginationBaseArgs struct {
	Base
	PageInfo
}

func (r *RequestPaginationBaseArgs) CheckArgs() error {
	if err := pkg.CheckIntArgs("pageNum", r.PageNum); err != nil {
		return err
	}
	if err := pkg.CheckIntArgs("pageSize", r.PageSize); err != nil {
		return err
	}
	if err := pkg.CheckStringArgs("environment", r.Environment); err != nil {
		return err
	}
	return nil
}

type RequestBaseArgs struct {
	Base
}

func (r *RequestBaseArgs) CheckArgs() error {
	if err := pkg.CheckStringArgs("environment", r.Environment); err != nil {
		return err
	}
	return nil
}
