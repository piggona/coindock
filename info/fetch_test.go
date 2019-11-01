package info

import (
	_ "coindock/config"
	"coindock/info/binance/market/orders"
	_ "coindock/info/jobs"
	"fmt"
	"sync"
	"testing"
)

// var API = &aggregateTrades.AggregateTradesContainer{}
// var confs aggregateTrades.Conf = aggregateTrades.Conf{
// 	Symbol:    "ETHBTC",
// 	FromId:    "",
// 	StartTime: "1572588059",
// 	EndTime:   "1572588059",
// 	Limit:     "2",
// }

var API = &orders.OrderContainer{}
var confs orders.Conf = orders.Conf{
	Symbol: "ETHBTC",
	Limit:  "10",
}
var isDisplay = false

func TestFetch(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		// API := &time.TimeContainer{}
		// API := &exchange.ExchangeContainer{}
		data, err := Fetch(API, confs)
		if err != nil {
			t.Errorf("Error occurs in Fetch: %v\n", err)
		}
		if isDisplay {
			fmt.Printf("Testfetch: %v\n", data)
		}
		ch <- struct{}{}
	}()
	<-ch
}

func TestFetchSerial(t *testing.T) {
	t.SkipNow()
	for i := 0; i < 20; i++ {
		// API := &time.TimeContainer{}
		data, err := Fetch(API, confs)
		if err != nil {
			t.Errorf("Error occurs in Fetch: %v\n", err)
		}
		if isDisplay {
			fmt.Printf("TestfetchSerial: %v\n", data)
		}
	}
}

func TestFetchParallel(t *testing.T) {
	t.SkipNow()
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			// API := &time.TimeContainer{}
			data, err := Fetch(API, confs)
			if err != nil {
				t.Errorf("Error occurs in Fetch: %v\n", err)
			}
			if isDisplay {
				fmt.Printf("Testfetch: %v\n", data)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
