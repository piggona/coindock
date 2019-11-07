package account

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

type QueryOrderContainer struct {
	Symbol              string
	OrderId             int
	OrderListId         int
	ClientOrderId       string
	Price               string
	OrigQty             string
	ExecuteQty          string
	CummulativeQuoteQty string
	Status              string
	TimeInForce         string
	Type                string
	Side                string
	StopPrice           string
	IceBergQty          string
	time                int64
	updateTime          int64
	isWorking           bool
}

type QueryOrderConf struct {
	Symbol            string `json:"symbol"`
	OrderId           string `json:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId"`
	RecvWindow        string `json:"recvWindow"`
	TimeStamp         string `json:"timestamp"`
}

func (o *QueryOrderContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	con, ok := conf.(QueryOrderConf)
	if !ok {
		err := fmt.Errorf("Error occurs in RequestCompiler: input is not in configuration format")
		return nil, err
	}
	endPoint := utils.EncodeQuery(con)
	signatured := utils.NewSignature(endPoint)

	// body := utils.EncodeBody(signatured)
	// 构造CallID,使用uuid算法
	id, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// 构造CallData
	data := &defs.CallData{
		CallID:   id,
		Method:   "Post",
		EndPoint: "/api/v3/order?" + signatured,
		Type:     "Full",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *QueryOrderContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
