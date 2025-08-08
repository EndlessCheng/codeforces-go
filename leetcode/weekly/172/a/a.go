package main

import (
	"bytes"
	"strconv"
)

func maximum69Number1(num int) int {
	s := []byte(strconv.Itoa(num))
	i := bytes.IndexByte(s, '6')
	if i < 0 {
		return num
	}
	s[i] = '9'
	ans, _ := strconv.Atoi(string(s))
	return ans
}

func maximum69Number(num int) int {
	maxBase := 0
	base := 1
	for x := num; x > 0; x /= 10 {
		if x%10 == 6 {
			maxBase = base
		}
		base *= 10
	}
	return num + maxBase*3
}
