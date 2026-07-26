package main

import (
	"bytes"
	"math"
	"strconv"
)

// https://space.bilibili.com/206214
func largestInteger1(n, s int) int {
	if s > n*9 {
		return -1
	}
	if s == 0 {
		return 0
	}

	res := bytes.Repeat([]byte{'0'}, n)
	for i := range res {
		if s <= 9 {
			res[i] += byte(s)
			break
		}
		res[i] = '9'
		s -= 9
	}
	ans, _ := strconv.Atoi(string(res))
	return ans
}

func largestInteger(n, s int) int {
	if s > n*9 {
		return -1
	}

	ans := int(math.Pow10(s/9)) - 1 // 填 9
	if s%9 > 0 {
		ans = ans*10 + s%9 // 填 s%9
		n--
	}
	return ans * int(math.Pow10(n-s/9)) // 填 0
}
