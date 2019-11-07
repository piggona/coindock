package market

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

type OldTradesContainer []SingleTrade

type OldTradesConf struct {
	Symbol string
	Limit  string
	FromId string
}

func (r *OldTradesContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	con, ok := conf.(OldTradesConf)
	if !ok {
		err := fmt.Errorf("Error occurs in orderBook.RequestCompiler: Incorrect Conf")
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
		EndPoint: "/api/v1/historicalTrades?" + endPoint,
		Type:     "Half",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *OldTradesContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
