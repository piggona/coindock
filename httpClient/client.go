package httpClient

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func Setup(isSecure, nop bool) *http.Client {
	c := &http.Client{}
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}

	if nop {
		c.Transport = &NopTransport{}
	}
	fmt.Println(http.DefaultClient == c)
	return c
}

// NopTransport 没有任何操作的传输
type NopTransport struct {
}

// RoundTripper接口
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
