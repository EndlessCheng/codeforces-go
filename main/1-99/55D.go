package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf55D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcms := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 15, 18, 20, 21, 24, 28, 30, 35, 36, 40, 42, 45, 56, 60, 63, 70, 72, 84, 90, 105, 120, 126, 140, 168, 180, 210, 252, 280, 315, 360, 420, 504, 630, 840, 1260, 2520}
	idx := [2521]int{}
	for i, v := range lcms {
		idx[v] = i
	}
	lcmRes := [len(lcms)][10]int{}
	for i, v := range lcms {
		lcmRes[i][0] = i
		lcmRes[i][1] = i
		for d := 2; d < 10; d++ {
			lcmRes[i][d] = idx[v*d/gcd(v, d)]
		}
	}

	var T int
	var low, high string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &low, &high)
		n := len(high)
		low = strings.Repeat("0", n-len(low)) + low
		memo := make([][len(lcms)][2520]int, n)
		for i := range memo {
			for j := range memo[i] {
				for k := range memo[i][j] {
					memo[i][j][k] = -1
				}
			}
		}
		var f func(int, int, int, bool, bool) int
		f = func(i, j, rem int, limitLow, limitHigh bool) (res int) {
			if i == n {
				if rem%lcms[j] > 0 {
					return 0
				}
				return 1
			}
			if !limitLow && !limitHigh {
				p := &memo[i][j][rem]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}

			lo := 0
			if limitLow {
				lo = int(low[i] - '0')
			}
			hi := 9
			if limitHigh {
				hi = int(high[i] - '0')
			}

			for d := lo; d <= hi; d++ {
				res += f(i+1, lcmRes[j][d], (rem*10+d)%2520, limitLow && d == lo, limitHigh && d == hi)
			}
			return
		}
		Fprintln(out, f(0, 0, 0, true, true))
	}
}

//func main() { cf55D(os.Stdin, os.Stdout) }
