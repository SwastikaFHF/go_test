package controller

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"strings"
)

const (
	HEADER_SIGN_KEY = "6847E4D3-1A6B-4AA5-B504-15B1A7C01490"
)

func CheckoutHead(head, body string) bool {
	h := md5.New()
	h.Write([]byte(body + HEADER_SIGN_KEY)) // 需要加密的字符串为 sharejs.com
	sec := hex.EncodeToString(h.Sum(nil))
	return strings.EqualFold(sec, head)
}

func GetEncryptedBody(body string) string {
	h := md5.New()
	h.Write([]byte(body + HEADER_SIGN_KEY))
	sec := hex.EncodeToString(h.Sum(nil))
	return sec
}
