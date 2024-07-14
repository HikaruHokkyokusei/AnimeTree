package MyAnimeListSDK

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	URL "net/url"
	"strings"
)

var (
	baseURL, _ = URL.Parse("https://api.myanimelist.net")
	pathPrefix = "v2"
)

func buildReqURL(path string, queryParams map[string]string) string {
	newUrl := baseURL.JoinPath(pathPrefix, path)

	qParams := URL.Values{}
	for key, value := range queryParams {
		qParams.Add(key, value)
	}

	newUrl.RawQuery = qParams.Encode()
	return newUrl.String()
}

func buildRequest(method string, path string, queryParams map[string]string, reqBody string) (*http.Request, error) {
	url := buildReqURL(path, queryParams)

	var bodyReader io.Reader
	if reqBody != "" {
		bodyReader = strings.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err == nil && bodyReader != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, err
}

func (c MyAnimeListClient) request(method string, path string, queryParams map[string]string, reqBody string) (*map[string]any, error) {
	req, err := buildRequest(method, path, queryParams, reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(resp *http.Response) {
		_ = resp.Body.Close()
	}(resp)

	body := &map[string]any{}
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return nil, err
	}

	if statusCode := resp.StatusCode; statusCode != 200 {
		errFormat := fmt.Sprint("Response status code: ", statusCode)
		if val, ok := (*body)["error"]; ok && val != "" {
			errFormat = fmt.Sprint(errFormat, ". Error: ", val)
		}
		if val, ok := (*body)["message"]; ok && val != "" {
			errFormat = fmt.Sprint(errFormat, ". Message: ", val)
		}
		return nil, fmt.Errorf(errFormat)
	}

	return body, nil
}
