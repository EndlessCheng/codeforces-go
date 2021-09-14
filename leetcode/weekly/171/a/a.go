package main

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) (ans []int) {
	for i := 1; i < n; i++ {
		if strings.IndexByte(strconv.Itoa(i), '0')+strings.IndexByte(strconv.Itoa(n-i), '0') == -2 {
			return []int{i, n - i}
		}
	}
	return
}
