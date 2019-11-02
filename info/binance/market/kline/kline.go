package kline

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

const (
	explain string = `
	Kline/Candlestick chart intervals:
	m -> minutes; h -> hours; d -> days; w -> weeks; M -> months
	1m
	3m
	5m
	15m
	30m
	1h
	2h
	4h
	6h
	8h
	12h
	1d
	3d
	1w
	1M
	`
)

type KlineContainer [][]interface{}

type Conf struct {
	Symbol    string
	Interval  string
	StartTime string
	EndTime   string
	Limit     string
}

func (a *KlineContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	con, ok := conf.(Conf)
	if !ok {
		err := fmt.Errorf("Error occurs in AggregateTrades.RequestCompiler: Incorrect Conf")
		return nil, err
	}
	endPoint := utils.EncodeQuery(con)
	// 构造CallID,使用uuid算法
	id, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// 构造CallData
	data := &defs.CallData{
		CallID:   id,
		Method:   "Get",
		EndPoint: "/api/v1/klines?" + endPoint,
		Type:     "Half",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *KlineContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
