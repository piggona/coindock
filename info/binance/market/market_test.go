package market

import (
	"fmt"
	"testing"
)

func TestOrderBook(t *testing.T) {
	o := &OrderContainer{}
	data, err := o.RequestCompiler("BTCZEC")
	if err != nil {
		t.Errorf("Error occurs in RequestCompiler: %v", err)
	}
	fmt.Println(data)
}
