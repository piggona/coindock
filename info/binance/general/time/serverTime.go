package time

import (
	"coindock/info/utils"
	"encoding/json"
	"fmt"
	"io"

	"coindock/info/defs"
)

// TimeContainer binance restful API's server time.
type TimeContainer struct {
	ServerTime int64
}

// RequestCompiler 构造CallData
func (t *TimeContainer) RequestCompiler(conf ...interface{}) (*defs.CallData, error) {
	// 构造CallID,使用uuid算法
	id, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// 构造CallData
	data := &defs.CallData{
		CallID:   id,
		Method:   "Get",
		EndPoint: "/api/v1/time",
		Type:     "None",
		Body:     nil,
		Data:     nil,
		PlatForm: "binance",
	}
	return data, nil
}

// ExtractData 接收io.PipeReader传来的信息
func (t *TimeContainer) ExtractData(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(t); err != nil {
		fmt.Errorf("Response Body Decode Failed: %v .\n", err)
		return err
	}
	return nil
}

// var defaultEndPoint string = "/api/v1/exchangeInfo"

// // Exec timeExec get server time from binance restful API.
// func Exec(endPoint string) interface{} {
// 	if len(endPoint) == 0 {
// 		endPoint = defaultEndPoint
// 	}
// 	url := "https://api.binance.com" + endPoint
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
// 		os.Exit(1)
// 	}
// 	defer resp.Body.Close()
// 	var result interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		fmt.Fprintf(os.Stderr, "Response Body Decode Failed. %v \n", err)
// 	}
// 	return result
// }

// func GetData() {
// 	ch := make(chan *defs.CallData)

// 	jobs.Call <- ch
// 	send := &defs.CallData{
// 		"12345", "GET", "/api/vi/time", "NONE", "respect", "", "binance",
// 	}
// 	ch <- send
// 	jobs.Ok <- ch
// }
