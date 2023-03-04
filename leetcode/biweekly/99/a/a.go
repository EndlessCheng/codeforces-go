package main

import (
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
func splitNum(num int) int {
	s := []byte(strconv.Itoa(num))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	a := [2][]byte{}
	for i, c := range s {
		a[i&1] = append(a[i&1], c)
	}
	x, _ := strconv.Atoi(string(a[0]))
	y, _ := strconv.Atoi(string(a[1]))
	return x + y
}
