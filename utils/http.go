package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

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

func Get(callUrl string, header map[string]string, query map[string]string, response interface{}) error {
	if query == nil {
		query = make(map[string]string)
	}

	for retries := 0; retries < 2; retries++ {

		ret, err := HttpRequest(http.MethodGet, callUrl, header, query, nil)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(ret, response); err != nil {
			return err
		}
		break
	}
	return nil
}

func Post(callUrl string, headers map[string]string, query map[string]string, body interface{}, response interface{}) error {
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

	for retries := 0; retries < 2; retries++ {

		ret, err := HttpRequest(http.MethodPost, callUrl, headers, query, b)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(ret, &response); err != nil {
			return err
		}
		break
	}
	return nil
}

func Put(callUrl string, headers map[string]string, query map[string]string, body interface{}, response interface{}) error {
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

	for retries := 0; retries < 2; retries++ {

		ret, err := HttpRequest(http.MethodPut, callUrl, headers, query, b)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(ret, response); err != nil {
			return err
		}
		break
	}
	return nil
}

func Delete(callUrl string, headers map[string]string, query map[string]string, body interface{}, response interface{}) error {
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

	for retries := 0; retries < 2; retries++ {

		ret, err := HttpRequest(http.MethodDelete, callUrl, headers, query, b)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(ret, response); err != nil {
			return err
		}
		break
	}
	return nil
}
