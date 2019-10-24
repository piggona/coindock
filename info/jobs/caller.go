package jobs

import (
	"coindock/info/defs"
	"fmt"
	"os"
)

type client chan *defs.CallData

var (
	Call = make(chan client)
	Resp = make(chan *defs.CallData)
	Ok   = make(chan client)
)

func init() {
	go caller()
}

func caller() {
	calls := make(map[string]client)
	for {
		select {
		case cli := <-Call:
			// 首先读client中传来的数据（配合调用函数中先传入client再向client中传CallData数据）
			data := getData(cli)
			id := data.CallID
			// 之后注册入calls中
			calls[id] = cli
			// 将网络任务交给worker
		case result := <-Resp:
			// 收到worker的结果，传递给调用函数
			id := result.CallID
			calls[id] <- result
		case cli := <-Ok:
			// 调用函数完成（调用函数指的是binance或bithumb包里发起工作的函数）
			fmt.Println("Close cli")
			close(cli)
		}
	}
}

func getData(ch <-chan *defs.CallData) *defs.CallData {
	msg := <-ch
	fmt.Fprintln(os.Stdout, *msg)
	return msg
}
