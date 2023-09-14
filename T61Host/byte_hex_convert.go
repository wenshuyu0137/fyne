package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// 十六进制字符串转字节
func hex_string_to_bytes(s string) []byte {
	temp := strings.ReplaceAll(s, " ", "")

	decode, err := hex.DecodeString(temp)
	if err != nil {
		return []byte{0}
	}

	return decode
}

// 字节转十六进制字符串数组
func bytes_to_hex_string(b []byte) []string {
	s := make([]string, len(b))

	for i, j := range b {
		s[i] = fmt.Sprintf("%02X", j)
	}

	return s
}
