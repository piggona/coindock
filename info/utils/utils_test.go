package utils

import (
	"fmt"
	"testing"
)

type Conf struct {
	Symbol    string
	Limit     string
	FromId    string
	StartTime string
	EndTime   string
}

func TestEncodeQuery(t *testing.T) {
	var confs Conf = Conf{
		Symbol:    "ETHBTC",
		FromId:    "",
		StartTime: "1572588059",
		EndTime:   "1572588059",
		Limit:     "2",
	}
	fmt.Println(EncodeQuery(confs))
}

type ConfBody struct {
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

func TestEncodeBody(t *testing.T) {
	// data := &ConfBody{
	// 	Symbol:      "BTCZEC",
	// 	Side:        "left",
	// 	Type:        "resll",
	// 	TimeInForce: "keep",
	// 	Quantity:    "respect",
	// }
	// reader, err := EncodeBody(data)
	// if err != nil {
	// 	fmt.Printf("Error occurs in Encoding body: %v\n", err)
	// }
	// var output string
	// output = reader.(*bytes.Buffer).String()
	// fmt.Println(output)
}

func TestSignature(t *testing.T) {
	data := "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"
	fmt.Println(NewSignature(data))
}
