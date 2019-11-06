package utils

import (
	"coindock/config"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
	"time"
)

func NewSignature(data string) string {
	var key = config.Binance.SecretKey
	var recvWin = strconv.Itoa(config.Binance.RecvWindow)
	fmt.Println(recvWin)
	now := strconv.FormatInt((time.Now().UnixNano()/1e6)+int64(config.Binance.TimeOffset), 10)
	// fmt.Printf("!!debug time: %s\n", now)
	// now = "1573028367164"

	encodeStr := data + "&recvWindow=" + recvWin + "&timestamp=" + now
	digest := getHmacCodes(encodeStr, key)
	return encodeStr + "&signature=" + digest
}

func getHmacCodes(s, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
