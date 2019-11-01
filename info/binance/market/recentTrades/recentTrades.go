package recentTrades

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type SingleTrade struct {
	Id           int
	Price        string
	Qty          string
	QuoteQty     string
	Time         int64
	IsBuyerMaker bool
	IsBestMatch  bool
}

type RecentTradesContainer []SingleTrade

type Conf struct {
	Symbol string
	Limit  string
}

func (r *RecentTradesContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	params := url.Values{}
	con, ok := conf.(Conf)
	if !ok {
		err := fmt.Errorf("Error occurs in recentTrades.RequestCompiler: Incorrect Conf")
		return nil, err
	}
	params.Set("symbol", con.Symbol)
	params.Set("limit", con.Limit)
	endPoint := params.Encode()
	// 构造CallID,使用uuid算法
	id, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// 构造CallData
	data := &defs.CallData{
		CallID:   id,
		Method:   "Get",
		EndPoint: "/api/v1/trades?" + endPoint,
		Type:     "Half",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *RecentTradesContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
