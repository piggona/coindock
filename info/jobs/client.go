package jobs

import (
	"io"
	"net/http"
)

// BinanceClient Packed Binance restful api http client.
type BinanceClient struct {
	*http.Client
}

// BinanceTransport Packed Binance single request roundtrip.
type BinanceTransport struct {
	*http.Transport
	apiKey string
}

// RoundTrip Binance single request roundTrip.
func (t *BinanceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-MBX-APIKEY", t.apiKey)
	return t.Transport.RoundTrip(req)
}

// NewBinanceClient Get a packed binance API client.
func NewBinanceClient(key string) *BinanceClient {
	t := http.Transport{}
	return &BinanceClient{
		Client: &http.Client{
			Transport: &BinanceTransport{
				Transport: &t,
				apiKey:    key,
			},
		},
	}
}

// GetBinance Get Request from BinanceClient
func (c *BinanceClient) GetBinance(urlPath string) (*http.Response, error) {
	resp, err := c.Get(urlPath)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PostBinance Post Request from BinanceClient
func (c *BinanceClient) PostBinance(urlPath string, body io.Reader) (*http.Response, error) {
	resp, err := c.Post(urlPath, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
