package pkg

import (
	"encoding/json"
	"fmt"
)

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func MapToModel[T interface{}](jsonMap map[string]string) T {
	js, _ := json.Marshal(jsonMap)
	var s T
	json.Unmarshal(js, &s)
	return s
}

func MaptoJson[T interface{}](t T) map[string]string {
	js, _ := json.Marshal(t)
	s := map[string]string{}
	json.Unmarshal(js, &s)
	return s
}

func KeyInMap[K comparable, V any](key K, m map[K]V) bool {
	for k := range m {
		if k == key {
			return true
		}
	}
	return false
}

func ValueInMap[K comparable, V comparable](value V, m map[K]V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func ItemInSlice[T comparable](item T, s []T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

// Paginate 分页函数
func Paginate[T any](list []T, page, pageSize int) ([]T, error) {
	if page < 1 || pageSize < 1 {
		return nil, fmt.Errorf("page and pageSize must be greater than 0")
	}

	// 计算起始索引和结束索引
	start := (page - 1) * pageSize
	end := start + pageSize

	// 如果起始索引超出范围，返回空切片
	if start >= len(list) {
		return []T{}, nil
	}

	// 确保结束索引不超出范围
	if end > len(list) {
		end = len(list)
	}

	// 返回分页后的切片
	return list[start:end], nil
}
