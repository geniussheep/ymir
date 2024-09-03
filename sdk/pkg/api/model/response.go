package model

import "github.com/geniussheep/ymir/sdk/api/response"

// PaginationResult 分页response
type PaginationResult[T any] struct {
	// 数据总数
	Total int `json:"total"`

	// 数据结果列表
	Result []T `json:"result" swaggertype:"object"`
}

type ResponsePaginationResult[T any] struct {
	response.Response
	Data PaginationResult[T] `json:"data" swaggertype:"object"`
}

type ResponseResult[T any] struct {
	response.Response
	Data T `json:"data" swaggertype:"object"`

	Value T `json:"value" swaggertype:"object"`
}
