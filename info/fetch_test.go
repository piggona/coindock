package info

import (
	"coindock/info/binance/general/time"
	"fmt"
	"sync"
	"testing"
)

func TestFetch(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		API := &time.TimeContainer{}
		data, err := Fetch(API)
		if err != nil {
			t.Errorf("Error occurs in Fetch: %v", err)
		}
		fmt.Printf("Testfetch: %v", data)
		ch <- struct{}{}
	}()
	<-ch
}

func TestFetchSerial(t *testing.T) {
	for i := 0; i < 20; i++ {
		API := &time.TimeContainer{}
		data, err := Fetch(API)
		if err != nil {
			t.Errorf("Error occurs in Fetch: %v", err)
		}
		fmt.Printf("Testfetch: %v", data)
	}
}

func TestFetchParallel(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			API := &time.TimeContainer{}
			data, err := Fetch(API)
			if err != nil {
				t.Errorf("Error occurs in Fetch: %v", err)
			}
			fmt.Printf("Testfetch: %v", data)
			wg.Done()
		}()
	}
	wg.Wait()
}
