package info

import (
	_ "coindock/config"
	"coindock/info/binance/market/orders"
	_ "coindock/info/jobs"
	"sync"
	"testing"
)

func TestFetch(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		// API := &time.TimeContainer{}
		// API := &exchange.ExchangeContainer{}
		API := &orders.OrderContainer{}
		_, err := Fetch(API, "ETHBTC")
		if err != nil {
			t.Errorf("Error occurs in Fetch: %v", err)
		}
		// fmt.Printf("Testfetch: %v", data)
		ch <- struct{}{}
	}()
	<-ch
}

func TestFetchSerial(t *testing.T) {
	// t.SkipNow()
	for i := 0; i < 20; i++ {
		// API := &time.TimeContainer{}
		API := &orders.OrderContainer{}
		_, err := Fetch(API, "ETHBTC")
		if err != nil {
			t.Errorf("Error occurs in Fetch: %v", err)
		}
		// fmt.Printf("Testfetch: %v", data)
	}
}

func TestFetchParallel(t *testing.T) {
	// t.SkipNow()
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			// API := &time.TimeContainer{}
			API := &orders.OrderContainer{}
			_, err := Fetch(API, "ETHBTC")
			if err != nil {
				t.Errorf("Error occurs in Fetch: %v", err)
			}
			// fmt.Printf("Testfetch: %v", data)
			wg.Done()
		}()
	}
	wg.Wait()
}
