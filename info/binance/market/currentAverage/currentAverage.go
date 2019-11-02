package currentAverage

import (
	"coindock/info/defs"
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"
)

type CurrentAverageContainer struct {
	Mins  int
	Price string
}

type Conf struct {
	Symbol string
}

func (r *CurrentAverageContainer) RequestCompiler(conf interface{}) (*defs.CallData, error) {
	con, ok := conf.(Conf)
	if !ok {
		err := fmt.Errorf("Error occurs in recentTrades.RequestCompiler: Incorrect Conf")
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
		EndPoint: "/api/v3/avgPrice?" + endPoint,
		Type:     "Half",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (o *CurrentAverageContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}
