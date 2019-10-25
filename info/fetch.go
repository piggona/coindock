package info

import (
	"coindock/info/defs"
	"coindock/info/jobs"
	"io"
)

// APIData 抽象所有API数据格式，及获取API数据的方式：如serverTime（见Client结构算法）
type APIData interface {
	RequestCompiler(conf ...interface{}) (*defs.CallData, error)
	ExtractData(r io.Reader) error
}

// Fetch 输入一个空API格式，使用RequestCompiler可以获取其对应的请求定义，
// 使用ExtractData来获取worker传来的数据流，并解析填充进API数据结构中.
// 建议使用单个goroutine来调用Fetch,发挥其最佳的并发性能.
func Fetch(resDat APIData, conf ...interface{}) (APIData, error) {
	callData, err := resDat.RequestCompiler(conf...)
	if err != nil {
		return nil, err
	}
	r, w := io.Pipe()
	callData.Data = w

	// 与Caller连接
	ch := make(chan *defs.CallData)
	jobs.Call <- ch
	ch <- callData
	// 等待worker返回流
	err = resDat.ExtractData(r)
	if err != nil {
		return nil, err
	}
	defer func() { jobs.Ok <- ch }()
	return resDat, nil
}
