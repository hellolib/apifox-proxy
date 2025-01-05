package proxy

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func ReverseProxy() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Transport: &http.Transport{
			MaxIdleConns:        200,              // 控制最大连接数
			MaxIdleConnsPerHost: 10,               // 控制每个主机要保持的最大空闲（保持活动）连接。默认：2。
			IdleConnTimeout:     90 * time.Second, // 空闲连接超时
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			// 设置超时时间
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second, // 连接超时时间
			}).DialContext,
		},
		Director:       defaultDirector(),
		ModifyResponse: defaultModifyResponse(),
		ErrorHandler:   defaultErrorHandler(),
	}
}

func defaultDirector() func(*http.Request) {
	return func(req *http.Request) {
	}
}

func defaultModifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return nil
	}
}

func defaultErrorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
	}
}
