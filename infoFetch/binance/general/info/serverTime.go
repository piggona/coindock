package info

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jobs "vct/infoFetch/jobs"
)

// TimeContainer binance restful API's server time.
type TimeContainer struct {
	ServerTime int64
}

type RateLimiter struct {
	RateLimitType string
	Interval      string
	IntervalNum   int
	Limit         int
}

type SymbolFilter struct {
	FilterType  string
	MinPrice    string
	MaxPrice    string
	TickSize    string
	MultipierUp string
}

type ExchangeFilter struct {
	FilterType       string
	MaxNumOrders     int
	MaxNumAlgoOrders int
}

type SymbolInfo struct {
	Symbol string
	Status string
}

type ExchangeContainer struct {
	TimeZone        string
	ServerTime      int64
	RateLimits      []RateLimiter
	ExchangeFilters []ExchangeFilter
	Symbols         []SymbolInfo
}

var defaultEndPoint string = "/api/v1/time"

// var defaultEndPoint string = "/api/v1/exchangeInfo"

// timeExec get server time from binance restful API.
func Exec(endPoint string) interface{} {
	if len(endPoint) == 0 {
		endPoint = defaultEndPoint
	}
	url := "https://api.binance.com" + endPoint
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Fprintf(os.Stderr, "Response Body Decode Failed. %v \n", err)
	}
	return result
}

func GetData() {
	ch := make(chan jobs.CallData)

	jobs.Call <- ch
	send := jobs.CallData{
		"12345", "GET", "/api/vi/time", "NONE", "respect", "",
	}
	ch <- send
	jobs.Ok <- ch
}
