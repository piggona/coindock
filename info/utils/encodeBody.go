package utils

import (
	"io"
	"strings"
)

func EncodeBody(data string) io.Reader {
	// str, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, err
	// }
	// return bytes.NewBuffer(str), nil
	return strings.NewReader(data)
}
