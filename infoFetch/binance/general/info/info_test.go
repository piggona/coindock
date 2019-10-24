package info

import (
	"fmt"
	"testing"

	_ "vct/infoFetch/jobs"
)

func TestServerTime(t *testing.T) {
	t.SkipNow()
	time := Exec("/api/v1/exchangeInfo")
	fmt.Println(time)
}

func TestCaller(t *testing.T) {
	GetData()
}
