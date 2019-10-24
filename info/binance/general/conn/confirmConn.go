package conn

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var defaultEndPoint string = "/api/v1/ping"

func Exec(endPoint string) {
	if len(endPoint) == 0 {
		endPoint = defaultEndPoint
	}
	url := "https://api.binance.com" + endPoint
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s %v \n", endPoint, err)
		os.Exit(1)
	}
	fmt.Printf("%s , %s", data, resp.Status)
}
