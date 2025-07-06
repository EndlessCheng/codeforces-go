package main

import (
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func concatHex36(n int) string {
	s := strconv.FormatInt(int64(n*n), 16) + strconv.FormatInt(int64(n*n*n), 36)
	return strings.ToUpper(s)
}
