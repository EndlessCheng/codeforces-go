package main

import (
	"sort"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func maxDiff(num int) (ans int) {
	s := strconv.Itoa(num)
	a := []int{}
	for i := byte('0'); i <= '9'; i++ {
		for j := byte('0'); j <= '9'; j++ {
			if i != s[0] || j != '0' {
				v, _ := strconv.Atoi(strings.ReplaceAll(s, string(i), string(j)))
				a = append(a, v)
			}
		}
	}
	sort.Ints(a)
	return a[len(a)-1] - a[0]
}
