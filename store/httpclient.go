package store

import (
	"net"
	"net/http"
	"time"
)

func NewClient() *http.Client {
	tr := &http.Transport{
		ResponseHeaderTimeout: time.Duration(10) * time.Second,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 2,
			Timeout:   time.Duration(10) * time.Second,
		}).DialContext,
		MaxIdleConns:          2,
		IdleConnTimeout:       time.Duration(10) * time.Second,
		TLSHandshakeTimeout:   time.Duration(10) * time.Second,
		MaxIdleConnsPerHost:   2,
		ExpectContinueTimeout: time.Duration(10) * time.Second,
	}

	httpClient := http.Client{
		Transport: tr,
		Timeout:   time.Duration(10) * time.Second,
	}

	return &httpClient
}
