package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func minimizeResult(expression string) (ans string) {
	sp := strings.Split(expression, "+")
	l, r := sp[0], sp[1]
	for i, min := 0, math.MaxInt64; i < len(l); i++ {
		a := 1
		if i > 0 {
			a, _ = strconv.Atoi(l[:i])
		}
		b, _ := strconv.Atoi(l[i:])
		for j := 1; j <= len(r); j++ {
			c, _ := strconv.Atoi(r[:j])
			d := 1
			if j < len(r) {
				d, _ = strconv.Atoi(r[j:])
			}
			res := a * (b + c) * d
			if res < min {
				min = res
				ans = fmt.Sprintf("%s(%s+%s)%s", l[:i], l[i:], r[:j], r[j:])
			}
		}
	}
	return
}
