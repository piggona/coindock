package httpClient

import (
	"fmt"
	"net/http"
)

func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.baidu.com")
	if err != nil {
		return err
	}
	fmt.Println("result of DoOps：", resp.StatusCode)

	return nil
}

func DefaultGetGolang() error {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		return err
	}
	fmt.Println("results of DefaultGetGolang:", resp.StatusCode)
	return nil
}
