package jobs

import (
	"fmt"
	"io"
	"net/http"
)

func binanceNoneWorker() {
	for call := range binanceNone {
		url := "https://api.binance.com" + call.EndPoint
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
