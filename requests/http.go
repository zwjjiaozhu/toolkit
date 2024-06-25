/*
* @Author: 西园公子
* @File: request.go
* @Date: 2024/1/21 14:00
* @IDE:  GoLand
 */

package requests

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func PostJson(url string, body any, client *http.Client, header http.Header) (respBody []byte, err error) {
	respBody, err = PostContent(
		url, "application/json", body, client, header,
	)
	if err != nil {
		return nil, err
	}
	return
}

func PostUrlForm(url_ string, body url.Values, client *http.Client, header http.Header) (respBody []byte, err error) {
	resp, err := http.PostForm(url_, body)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	respBody, err = io.ReadAll(resp.Body)
	return
}

func PostContent(url, contentType string, body any, client *http.Client, header http.Header) (respBody []byte,
	err error) {
	//if client
	var jsonBytes []byte
	switch body.(type) {
	case []byte:
		jsonBytes = body.([]byte)
	default:
		jsonBytes, err = json.Marshal(body)
		if err != nil {
			return
		}
	}
	var (
		request *http.Request
		resp    *http.Response
	)
	// 设置header时，需创建client
	if header != nil && client == nil {
		client = &http.Client{}
	}
	// 客户端模式
	if client != nil {
		request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			return
		}
		if header != nil {
			request.Header = header
		}
		if _, ok := request.Header["Content-Type"]; !ok {
			request.Header["Content-Type"] = []string{contentType}
		}
		resp, err = client.Do(request)
	} else {
		// 默认模式
		resp, err = http.Post(url, contentType, bytes.NewBuffer(jsonBytes))
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 解析结果
	if resp.StatusCode != http.StatusOK {
		c, _ := io.ReadAll(resp.Body)
		err = errors.New(strconv.Itoa(resp.StatusCode) + "\nJsonParam:" + string(jsonBytes) + "\nResponse:" + string(c))
		return
	}
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// PostMultipart 文件上传
func PostMultipart(url_ string, bodyBuf io.Reader, header, params map[string]string,
	timeout int) (resp *Response,
	err error) {

	//header["Content-Type"] = "multipart/form-data"

	req, err := http.NewRequest("POST", url_, bodyBuf)
	if err != nil {
		return
	}

	req.Header = ToStandardHeader(header)

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	httpResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	resp = &Response{
		Header:     ToMapHeader(httpResp.Header),
		Body:       respBody,
		StatusCode: httpResp.StatusCode,
	}

	return
}

func Post(url_ string, body map[string]any, header, params map[string]string, timeout int) (resp *Response,
	err error) {
	var bodyReader io.Reader
	switch header["Content-Type"] {
	case "application/x-www-form-urlencoded":
		bodyReader = strings.NewReader(ToUrlValues(body).Encode())
	case "application/json":
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBytes)
	default:
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequest("POST", url_, bodyReader)
	if err != nil {
		return
	}
	if params != nil {
		url_ = url_ + ToQueryParams(params)
	}
	if header != nil {
		req.Header = ToStandardHeader(header)
	}
	// 发送请求
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second, // 为0则不设置超时
	}
	httpResp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(httpResp.Body)
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	resp = &Response{
		Header:     ToMapHeader(httpResp.Header),
		Body:       respBody,
		StatusCode: httpResp.StatusCode,
	}
	return
}

func Get(url_ string, params, header map[string]string, timeout int) (resp *Response, err error) {
	if params != nil {
		url_ = url_ + ToQueryParams(params)
	}
	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		return
	}
	if header != nil {
		req.Header = ToStandardHeader(header)
	}
	// 发送请求
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second, // 为0则不设置超时
	}
	httpResp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(httpResp.Body)
	respBody, err := io.ReadAll(httpResp.Body)
	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("err: statusCode %d", resp.StatusCode)
		return
	}
	resp = &Response{
		Header:     ToMapHeader(httpResp.Header),
		Body:       respBody,
		StatusCode: httpResp.StatusCode,
	}
	return
}

func GetWithParams(url_ string, params url.Values) (respBody []byte, err error) {
	if params != nil {
		url_ = url_ + "?" + params.Encode()
	}
	resp, err := http.Get(url_)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	respBody, err = io.ReadAll(resp.Body)
	return
}

// SSE Server-Sent Events
func SSE(url_ string, body map[string]any, header, params map[string]string,
	timeout int, callback func(name, text, mode string), keyName, mode string) (err error) {

	if params != nil {
		url_ += ToQueryParams(params)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url_, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return
	}
	req.Header = ToStandardHeader(header)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// 持续读取数据
	// 就用文新一言做一个例子
	//wg := &sync.WaitGroup{}
	//wg.Add(1)
	log.Println("chat ...")
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		dataStr := strings.TrimSpace(line)
		callback(keyName, dataStr, mode)
	}
	return
}
