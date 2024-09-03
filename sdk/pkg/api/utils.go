package api

import (
	"fmt"
	"github.com/geniussheep/ymir/sdk/pkg"
	"github.com/geniussheep/ymir/sdk/pkg/api/model"
)

func SetFullUrl(domain, apiPath string) string {
	return fmt.Sprintf("%s%s", domain, apiPath)
}

func ExecutePagePost[T any](domain, apiPath string, body any) (*model.PaginationResult[T], error) {
	resp := model.ResponseResult[[]T]{}
	apiUrl := SetFullUrl(domain, apiPath)
	err := pkg.PostJson(apiUrl, nil, nil, body, &resp)
	if err != nil {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiUrl, body, err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v failed, response errr msg:%s", apiUrl, body, resp.Msg)
	}
	result := model.PaginationResult[T]{Total: len(resp.Data), Result: resp.Data}
	return &result, nil
}

func ExecuteDefaultPagePost[T any](domain, apiPath string, body any) (*model.PaginationResult[T], error) {
	resp := model.ResponsePaginationResult[T]{}
	apiUrl := SetFullUrl(domain, apiPath)
	err := pkg.PostJson(apiUrl, nil, nil, body, &resp)
	if err != nil {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiUrl, body, err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v failed, response errr msg:%s", apiUrl, body, resp.Msg)
	}
	return &resp.Data, nil
}

func ExecuteGet[T any](domain, apiPath string, query map[string]string) (*model.ResponseResult[T], error) {
	resp := model.ResponseResult[T]{}
	apiUrl := SetFullUrl(domain, apiPath)
	err := pkg.GetJson(apiUrl, nil, query, &resp)
	if err != nil {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiUrl, query, err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v failed, response errr msg:%s", apiUrl, query, resp.Msg)
	}
	return &resp, nil
}

func ExecutePost[T any](domain, apiPath string, body any) (*model.ResponseResult[T], error) {
	resp := model.ResponseResult[T]{}
	apiUrl := SetFullUrl(domain, apiPath)
	err := pkg.PostJson(apiUrl, nil, nil, body, &resp)
	if err != nil {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiUrl, body, err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("call [post]api:%s reqData:%+v failed, response errr msg:%s", apiUrl, body, resp.Msg)
	}
	return &resp, nil
}

func ExecuteGetString[T any](domain, apiPath string, query map[string]string) (string, error) {
	apiUrl := SetFullUrl(domain, apiPath)
	resp, err := pkg.GetString(apiUrl, nil, query)
	if err != nil {
		return "", fmt.Errorf("call [get]api:%s reqData:%+v error, err:%s", apiUrl, query, err)
	}
	return resp, nil
}

func ExecutePostString[T any](domain, apiPath string, body any) (string, error) {
	apiUrl := SetFullUrl(domain, apiPath)
	resp, err := pkg.PostString(apiUrl, nil, nil, body)
	if err != nil {
		return "", fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiUrl, body, err)
	}
	return resp, nil
}

//
//func ExecutePostSOAResponse[T interface{}](domain, apiPath string, query interface{}) (*model.ResponseResult[T], error) {
//	resp := model.ResponseResult[T]{}
//
//	str, err := ExecutePostSOA(domain, apiPath, query)
//	if err != nil {
//		return nil, fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiPath, query, err)
//	}
//	if err := json.Unmarshal([]byte(str), &resp); err != nil {
//		return nil, fmt.Errorf("call [post]api:%s reqData:%+v error, err:%s", apiPath, query, err)
//	}
//	return &resp, nil
//}
//
//func ExecutePostSOA(domain, apiPath string, query interface{}) (string, error) {
//
//	apiUrl := SetFullUrl(domain, apiPath)
//
//	b, _ := json.Marshal(&query)
//	var m map[string]any
//	_ = json.Unmarshal(b, &m)
//	payload := url.Values{}
//	for k, v := range m {
//		payload.Add(k, fmt.Sprintf("%v", v))
//	}
//	req, err := http.NewRequest(http.MethodPost,
//		apiUrl,
//		strings.NewReader(payload.Encode()))
//	if err != nil {
//		return "", err
//	}
//
//	req.Header.Add("Content-Type",
//		"application/x-www-form-urlencoded; param=value")
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return "", err
//	}
//
//	defer resp.Body.Close()
//
//	data, _ := io.ReadAll(resp.Body)
//	return string(data), nil
//}
