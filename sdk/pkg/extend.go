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
