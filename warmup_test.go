package warmup

import (
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHttpClientFactory struct {
	ActualUris []string
}

type MockHttpClient struct {
	ActualUris []string
	ip         net.IP
	index      int
}

type MockNetLookup struct {
	ActualHosts []string
	index       int
}

func (f *MockHttpClientFactory) Get(ip net.IP) HttpClient {
	return &MockHttpClient{
		ActualUris: f.ActualUris,
	}
}

func (c *MockHttpClient) Get(uri string) (*http.Response, error) {
	c.ActualUris[c.index] = uri
	c.index++
	return &http.Response{
		StatusCode: 200,
	}, nil
}

func (l *MockNetLookup) LookupIP(host string) ([]net.IP, error) {
	l.ActualHosts[l.index] = host
	l.index++
	return []net.IP{
		net.ParseIP("127.0.0.1"),
	}, nil
}

func TestWarmup(t *testing.T) {
	tests := []struct {
		domainName       string
		pointsOfPresence []string
		event            LambdaEvent
		expectError      error
		expectUris       []string
		expectHosts      []string
	}{
		{
			domainName: "example",
			pointsOfPresence: []string{
				"AKL50-C1",
			},
			event: LambdaEvent{
				Records: []Record{
					{
						S3: S3{
							Object: Object{
								Key: "test",
							},
						},
					},
				},
			},
			expectError: nil,
			expectHosts: []string{
				"example.AKL50-C1.cloudfront.net",
			},
			expectUris: []string{
				"https://example/test",
			},
		},
	}

	for _, test := range tests {
		actualUris := make([]string, len(test.expectUris))
		actualHosts := make([]string, len(test.expectHosts))
		actualError := warmup(
			test.domainName,
			test.event,
			test.pointsOfPresence,
			&MockNetLookup{ActualHosts: actualHosts},
			&MockHttpClientFactory{ActualUris: actualUris},
		)
		if !assert.Equal(t, test.expectError, actualError) {
			return
		}
		if !assert.Equal(t, test.expectHosts, actualHosts) {
			return
		}
		if !assert.Equal(t, test.expectUris, actualUris) {
			return
		}
	}
}
