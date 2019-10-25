package jobs

import (
	"coindock/config"
	"fmt"
	"io"
	"net/http"
)

var base_binance = config.Binance.BaseEndPoint

func binanceNoneWorker() {
	for call := range binanceNone {
		url := base_binance + call.EndPoint
		// fmt.Println(url)
		w := call.Data
		resp, err := http.Get(url)
		if err != nil {
			fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
		}
		go func() {
			io.Copy(w, resp.Body)
			w.Close()
			resp.Body.Close()
		}()
	}
}

func binanceHalfWorker() {
	for call := range binanceHalf {
		urlstr := base_binance + call.EndPoint
		w := call.Data
		resp, err := binanceClient.GetBinance(urlstr)
		if err != nil {
			fmt.Errorf("Error occurs in binanceClient.GetBinance: %v\n", err)
		}
		go func() {
			io.Copy(w, resp.Body)
			w.Close()
			resp.Body.Close()
		}()
	}
}

func binanceFullWorker() {
	for call := range binanceFull {
		urlstr := base_binance + call.EndPoint
		w := call.Data
		resp, err := binanceClient.GetBinance(urlstr)
		if err != nil {
			fmt.Errorf("Error occurs in binanceClient.GetBinance: %v\n", err)
		}
		go func() {
			io.Copy(w, resp.Body)
			w.Close()
			resp.Body.Close()
		}()
	}
}
