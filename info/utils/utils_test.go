package utils

import (
	"fmt"
	"testing"
)

type Conf struct {
	Symbol    string
	Limit     string
	FromId    string
	StartTime string
	EndTime   string
}

func TestEncodeQuery(t *testing.T) {
	var confs Conf = Conf{
		Symbol:    "ETHBTC",
		FromId:    "",
		StartTime: "1572588059",
		EndTime:   "1572588059",
		Limit:     "2",
	}
	fmt.Println(EncodeQuery(confs))
}
