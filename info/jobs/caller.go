package jobs

import (
	"coindock/config"
	"coindock/info/defs"
	"fmt"
)

// client ,Communication entity for fetcher and caller
type client chan *defs.CallData

var (
	// Call Fetcher use this channel to login and commit synchronic work.
	Call = make(chan client)
	// Resp Worker may return some informations.
	Resp = make(chan *defs.CallData)
	// Ok Fetcher use this channel to exit.
	Ok = make(chan client)
)

// Worker channel
var binanceNone = make(chan *defs.CallData)
var binanceHalf = make(chan *defs.CallData)
var binanceFull = make(chan *defs.CallData)

var binanceClient *BinanceClient

func init() {
	go caller()

	binanceClient = NewBinanceClient(config.Binance.APIKey)
}

func caller() {
	calls := make(map[string]client)
	// 部署worker
	establishWorkers()
	// 接收fetcher
	for {
		select {
		case cli := <-Call:
			// 首先读client中传来的数据（配合调用函数中先传入client再向client中传CallData数据）
			data := getData(cli)
			id := data.CallID
			// 之后注册入calls中
			calls[id] = cli
			// 将网络任务交给worker
			reqType := data.Type
			switch reqType {
			case "None":
				binanceNone <- data
			case "Half":
				binanceHalf <- data
			case "Full":
				binanceFull <- data
			}

		case result := <-Resp:
			// 收到worker的结果，传递给调用函数
			id := result.CallID
			calls[id] <- result
		case cli := <-Ok:
			// 调用函数完成（调用函数指的是binance或bithumb包里发起工作的函数）
			// fmt.Println("Close cli")
			close(cli)
		}
	}
}

func establishWorkers() {
	binanceCount := config.Binance.CallWorker
	fmt.Println(config.Binance)
	var (
		binanceNoneCount = binanceCount.None
		binanceHalfCount = binanceCount.Half
		binanceFullCount = binanceCount.Full
	)
	fmt.Printf("binanceNoneCount: %d", binanceNoneCount)
	// 建立worker
	for i := 0; i < binanceNoneCount; i++ {
		go binanceNoneWorker()
	}
	for i := 0; i < binanceHalfCount; i++ {
		go binanceHalfWorker()
	}
	for i := 0; i < binanceFullCount; i++ {
		go binanceFullWorker()
	}
}

func getData(ch <-chan *defs.CallData) *defs.CallData {
	msg := <-ch
	// fmt.Fprintln(os.Stdout, *msg)
	return msg
}
