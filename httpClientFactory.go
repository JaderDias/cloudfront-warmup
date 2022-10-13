package warmup

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

const timeout = time.Duration(10 * time.Second)

type HttpClientFactory interface {
	Get(ip net.IP) HttpClient
}

type HttpClient interface {
	Get(uri string) (*http.Response, error)
}

type MyHttpClientFactory struct {
}

func (f *MyHttpClientFactory) Get(ip net.IP) HttpClient {
	dialer := &net.Dialer{
		Timeout:   timeout,
		KeepAlive: timeout,
	}
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				addr = fmt.Sprintf("[%s]:443", ip)
				return dialer.DialContext(ctx, network, addr)
			},
		},
	}
}
