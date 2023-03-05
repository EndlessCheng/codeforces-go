package main

import (
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
func splitNum(num int) int {
	s := []byte(strconv.Itoa(num))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	a := [2]int{}
	for i, c := range s {
		a[i%2] = a[i%2]*10 + int(c-'0')
	}
	return a[0] + a[1]
}
