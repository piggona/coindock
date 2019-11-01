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
		method := call.Method
		var resp *http.Response
		var err error
		switch method {
		case "Get":
			resp, err = http.Get(url)
			if err != nil {
				fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
				return
			}
			break
		case "Post":
			body := call.Body
			resp, err = http.Post(url, "application/x-www-form-urlencoded", body)
			if err != nil {
				fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
				return
			}
			break
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
		// resp, err := binanceClient.GetBinance(urlstr)
		// if err != nil {
		// 	fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
		// }
		method := call.Method
		var resp *http.Response
		var err error
		switch method {
		case "Get":
			resp, err = binanceClient.GetBinance(urlstr)
			if err != nil {
				fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
				return
			}
		case "Post":
			body := call.Body
			resp, err = binanceClient.PostBinance(urlstr, body)
			if err != nil {
				fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
				return
			}
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
		method := call.Method
		var resp *http.Response
		var err error
		switch method {
		case "Get":
			resp, err = binanceClient.GetBinance(urlstr)
			if err != nil {
				fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
			}
			break
		case "Post":
			body := call.Body
			resp, err = binanceClient.PostBinance(urlstr, body)
			if err != nil {
				fmt.Errorf("Error occurs in binanceNonWorker: %v\n", err)
			}
			break
		}
		go func() {
			io.Copy(w, resp.Body)
			w.Close()
			resp.Body.Close()
		}()
	}
}
