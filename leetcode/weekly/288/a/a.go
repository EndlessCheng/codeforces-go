package main

import (
	"sort"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func largestInteger(num int) int {
	s := []byte(strconv.Itoa(num))
	a := [2][]byte{}
	for _, v := range s {
		a[v&1] = append(a[v&1], v)
	}
	sort.Slice(a[0], func(i, j int) bool { return a[0][i] > a[0][j] })
	sort.Slice(a[1], func(i, j int) bool { return a[1][i] > a[1][j] })

	for i, ch := range s {
		j := ch & 1
		s[i] = a[j][0]
		a[j] = a[j][1:]
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
