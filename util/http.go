package util

import (
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/enum"
)

func getPostBody(params map[string]string) string {
	var list []string
	query := ""
	for key, val := range params {
		query = key + "=" + Urlencode(val)
		list = append(list, query)
	}
	sort.Strings(list)
	return strings.Join(list, "&")
}

func Post(url string, queryBody map[string]string, params map[string]string, headers map[string]string, logger yeepayGo.YeepayLoggerInterface) (int, []byte, error) {
	postBody := ""
	if queryBody != nil {
		postBody = getPostBody(queryBody)
	}

	req, err := http.NewRequest(enum.HTTP_METHOD_POST, url, strings.NewReader(postBody))
	if err != nil {
		if logger != nil {
			logger.Error(req.URL.String(), "|", headers, "|", postBody, "|", err)
		}
		return 0, []byte(""), err
	}
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		if logger != nil {
			logger.Error(req.URL.String(), "|", headers, "|", postBody, "|", err)
		}
		return 0, []byte(""), err
	}

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	if logger != nil {
		logger.Info(req.URL.String(), "|", headers, "|", postBody, "|", resp.StatusCode, "|", string(bodyBytes))
	}

	return resp.StatusCode, bodyBytes, nil
}
