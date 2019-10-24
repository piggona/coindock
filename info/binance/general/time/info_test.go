package time

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	_ "coindock/info/jobs"
)

func TestServerTime(t *testing.T) {
	t.SkipNow()
	// time := Exec("/api/v1/exchangeInfo")
	// fmt.Println(time)
}

func TestCaller(t *testing.T) {
	t.SkipNow()
	// GetData()
}

var time *TimeContainer
var w *io.PipeWriter
var r *io.PipeReader

func TestMain(m *testing.M) {
	time = &TimeContainer{}
	r, w = io.Pipe()
	m.Run()
}

func TestTime(t *testing.T) {
	t.Run("Compiler", testCompiler)
	t.Run("Extract", testExtract)
}

func testCompiler(t *testing.T) {
	call, err := time.RequestCompiler()
	if err != nil {
		t.Errorf("Error occurs in RequestCompiler: %v\n", err)
	}
	call.Data = w
}

func testExtract(t *testing.T) {
	// 合成一个json流
	url := "https://api.binance.com/api/v1/exchangeInfo"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Error ocurrs in http.Get(): %v\n", err)
	}
	defer resp.Body.Close()
	go func() {
		io.Copy(w, resp.Body)
		w.Close()
	}()
	err = time.ExtractData(r)
	if err != nil {
		t.Errorf("Error occurs in ExtractData: %v\n", err)
	}
	fmt.Printf("Got Server Time: %v\n", time)
}
