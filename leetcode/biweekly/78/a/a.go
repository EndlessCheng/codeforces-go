package main

import (
	"math"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func divisorSubstrings(num, k int) (ans int) {
	m := int(math.Pow10(k))
	for n := num; n >= m/10; n /= 10 {
		x := n % m
		if x > 0 && num%x == 0 {
			ans++
		}
	}
	return
}

func divisorSubstrings1(num, k int) (ans int) {
	s := strconv.Itoa(num)
	for i := k; i <= len(s); i++ {
		x, _ := strconv.Atoi(s[i-k : i]) // 长为 k 的子串
		if x > 0 && num%x == 0 {         // 子串能整除 num 
			ans++
		}
	}
	return
}
