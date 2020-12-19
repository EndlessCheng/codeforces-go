package main

import (
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func countSteppingNumbers(low, high int) (ans []int) {
	l, r := strconv.Itoa(low), strconv.Itoa(high)
	n := len(r)
	l = strings.Repeat("0", n-len(l)) + l
	var f func(p, v int, isL, isU bool)
	f = func(p, v int, isL, isU bool) {
		if p == n {
			ans = append(ans, v)
			return
		}
		if v == 0 {
			for i := l[p]; i <= '9'; i++ {
				f(p+1, int(i&15), i == l[p], false) // 此时 isU 一定不满足
			}
		} else {
			pre := v % 10
			if (!isL && pre > 0 || isL && pre > int(l[p]&15)) && (!isU || pre-1 <= int(r[p]&15)) {
				f(p+1, v*10+pre-1, isL && pre-1 == int(l[p]&15), isU && pre-1 == int(r[p]&15))
			}
			if (!isU && pre < 9 || isU && pre < int(r[p]&15)) && (!isL || pre+1 >= int(l[p]&15)) {
				f(p+1, v*10+pre+1, isL && pre+1 == int(l[p]&15), isU && pre+1 == int(r[p]&15))
			}
		}
	}
	for i := l[0]; i <= r[0]; i++ {
		f(1, int(i&15), i == l[0], i == r[0])
	}
	return
}
