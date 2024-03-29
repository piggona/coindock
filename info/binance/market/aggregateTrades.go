package market

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

type AggregateTradesContainer []SingleCompressedTrade

type SingleCompressedTrade struct {
	A  int    `json:"a"`
	P  string `json:"p"`
	Q  string `json:"q"`
	F  int    `json:"f"`
	L  int    `json:"l"`
	T  int64  `json:"T"`
	M1 bool   `json:"m"`
	M2 bool   `json:"M"`
}

// [
//   {
//     "a": 26129,         // Aggregate tradeId
//     "p": "0.01633102",  // Price
//     "q": "4.70443515",  // Quantity
//     "f": 27781,         // First tradeId
//     "l": 27781,         // Last tradeId
//     "T": 1498793709153, // Timestamp
//     "m": true,          // Was the buyer the maker?
//     "M": true           // Was the trade the best price match?
//   }
// ]

type AggregateTradesConf struct {
	Symbol    string
	Limit     string
	FromId    string
	StartTime string
	EndTime   string
}

func (a *AggregateTradesContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	con, ok := conf.(AggregateTradesConf)
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
		EndPoint: "/api/v1/aggTrades?" + endPoint,
		Type:     "Half",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *AggregateTradesContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
