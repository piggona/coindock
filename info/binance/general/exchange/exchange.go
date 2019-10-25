package exchange

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

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
	// TimeZone        string
	// ServerTime      int64
	// RateLimits      []RateLimiter
	// ExchangeFilters []ExchangeFilter
	Symbols []SymbolInfo
}

func (t *ExchangeContainer) RequestCompiler(conf ...interface{}) (*defs.CallData, error) {
	// 构造CallID,使用uuid算法
	id, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// 构造CallData
	data := &defs.CallData{
		CallID:   id,
		Method:   "Get",
		EndPoint: "/api/v1/exchangeInfo",
		Type:     "None",
		Body:     "",
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (t *ExchangeContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(t); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
