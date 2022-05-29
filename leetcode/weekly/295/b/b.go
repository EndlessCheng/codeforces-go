package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214/dynamic
func discountPrices(sentence string, discount int) string {
	sp := strings.Split(sentence, " ")
	for i, s := range sp {
		if s[0] == '$' {
			if v, err := strconv.Atoi(s[1:]); err == nil {
				sp[i] = fmt.Sprintf("$%.2f", float64(v*(100-discount))/100)
			}
		}
	}
	return strings.Join(sp, " ")
}
