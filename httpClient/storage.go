package httpClient

import (
	"fmt"
	"net/http"
)

type Controller struct {
	*http.Client
}

func (c *Controller) DoOps() error {
	resp, err := c.Client.Get("http://www.baidu.com")
	if err != nil {
		return err
	}
	fmt.Println("result of client.DoOps", resp.StatusCode)
	return nil
}
