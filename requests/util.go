package requests

import (
	"fmt"
	"net/http"
	"net/url"
)

func ToStandardHeader(header map[string]string) http.Header {
	newHeader := make(http.Header)
	for k, v := range header {
		newHeader[k] = []string{v}
	}
	return newHeader
}

func ToMapHeader(header http.Header) map[string]string {
	newHeader := make(map[string]string)
	for k, v := range header {
		if len(v) > 0 {
			newHeader[k] = v[0]
		}
	}
	return newHeader
}

func ToQueryParams(params map[string]string) (query string) {
	query = "?"
	for k, v := range params {
		query += fmt.Sprintf("%s=%s&", k, url.QueryEscape(v))
	}
	return
}

func ToUrlValues(body map[string]any) (data url.Values) {
	data = make(url.Values)
	for k, v := range body {
		data[k] = []string{v.(string)}
	}
	return
}
