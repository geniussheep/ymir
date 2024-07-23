package pkg

import "encoding/json"

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
