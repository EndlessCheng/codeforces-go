package main

import (
	"fmt"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func toHexspeak(num string) (ans string) {
	v, _ := strconv.Atoi(num)
	s := fmt.Sprintf("%X", v)
	for _, b := range s {
		if '1' < b && b <= '9' {
			return "ERROR"
		}
	}
	return strings.NewReplacer("0", "O", "1", "I").Replace(s)
}
