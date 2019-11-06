package account

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

type OrderDetail struct {
	Price           string
	Qty             string
	Commission      string
	CommissionAsset string
}

type NewOrderContainer struct {
	Symbol              string
	OrderId             int
	ClientOrderId       string
	TransactTime        int64
	Price               string
	OrigQty             string
	ExecuteQty          string
	CummulativeQuoteQty string
	Status              string
	TimeInForce         string
	Type                string
	Side                string
	Fills               []OrderDetail
}

type Conf struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	Type             string `json:"type"`
	TimeInForce      string `json:"timeInForce"`
	Quantity         string `json:"quantity"`
	Price            string `json:"price"`
	NewClientOrderId string `json:"newClientOrderId"`
	StopPrice        string `json:"stopPrice"`
	IcebergQty       string `json:"icebergQty"`
	NewOrderRespType string `json:"newOrderRespType"`
	RecvWindow       string `json:"recvWindow"`
	TimeStamp        string `json:"timestamp"`
}

func (o *NewOrderContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	con, ok := conf.(Conf)
	if !ok {
		err := fmt.Errorf("Error occurs in RequestCompiler: input is not in configuration format")
		return nil, err
	}
	endPoint := utils.EncodeQuery(con)
	signatured := utils.NewSignature(endPoint)

	body := utils.EncodeBody(signatured)
	// 构造CallID,使用uuid算法
	id, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// 构造CallData
	data := &defs.CallData{
		CallID:   id,
		Method:   "Post",
		EndPoint: "/api/v3/order/test",
		Type:     "Full",
		Body:     body,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *NewOrderContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
