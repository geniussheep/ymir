package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var defaultHeader = map[string]string{"Content-Type": "application/json"}

func HttpRequest(method string, callUrl string, headers map[string]string, query map[string]string, body []byte) ([]byte, error) {
	apiURL, _ := url.Parse(callUrl)
	params := url.Values{}
	for k, v := range query {
		params.Add(k, v)
	}
	apiURL.RawQuery = params.Encode()

	u := apiURL.String()
	var b io.Reader
	if body != nil {
		b = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, u, b)

	for k, v := range defaultHeader {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpRequestJson(httpMethod string, callUrl string, headers map[string]string, query map[string]string, body interface{}, response interface{}) error {
	if query == nil {
		query = make(map[string]string)
	}

	var b []byte
	if body != nil {
		json, err := json.Marshal(body)

		if err != nil {
			return err
		}
		b = json
	}

	ret, err := HttpRequest(httpMethod, callUrl, headers, query, b)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(ret, response); err != nil {
		return fmt.Errorf("%s", ret)
	}
	return nil
}

func httpRequestString(httpMethod string, callUrl string, headers map[string]string, query map[string]string, body interface{}) (string, error) {
	if query == nil {
		query = make(map[string]string)
	}

	var b []byte
	if body != nil {
		json, err := json.Marshal(body)

		if err != nil {
			return "errot", err
		}
		b = json
	}

	response := ""

	ret, err := HttpRequest(httpMethod, callUrl, headers, query, b)
	if err != nil {
		return "error", err
	}
	response = string(ret)
	return response, nil
}

func GetJson(callUrl string, header map[string]string, query map[string]string, response interface{}) error {
	return httpRequestJson(http.MethodGet, callUrl, header, query, nil, response)
}

func GetString(callUrl string, header map[string]string, query map[string]string) (string, error) {
	return httpRequestString(http.MethodGet, callUrl, header, query, nil)
}

func PostJson(callUrl string, header map[string]string, query map[string]string, body interface{}, response interface{}) error {
	return httpRequestJson(http.MethodPost, callUrl, header, query, body, response)
}

func PostString(callUrl string, header map[string]string, query map[string]string, body interface{}) (string, error) {
	return httpRequestString(http.MethodPost, callUrl, header, query, body)
}

func PutJson(callUrl string, header map[string]string, query map[string]string, body interface{}, response interface{}) error {
	return httpRequestJson(http.MethodPut, callUrl, header, query, body, response)
}

func PutString(callUrl string, header map[string]string, query map[string]string, body interface{}) (string, error) {
	return httpRequestString(http.MethodPut, callUrl, header, query, body)
}

func DeleteJson(callUrl string, header map[string]string, query map[string]string, body interface{}, response interface{}) error {
	return httpRequestJson(http.MethodDelete, callUrl, header, query, body, response)
}

func DeleteString(callUrl string, header map[string]string, query map[string]string, body interface{}) (string, error) {
	return httpRequestString(http.MethodDelete, callUrl, header, query, body)
}
