package main

import (
	. "fmt"
	"io"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func CF431D(in io.Reader, out io.Writer) {
	search := func(l, r int64, f func(int64) bool) int64 {
		for l < r {
			m := (l + r) >> 1
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	var tar int64
	var k int
	Fscan(in, &tar, &k)
	calc := func(s string) int64 {
		const lowerC, upperC byte = '0', '1'
		n := len(s)
		dp := make([][]int64, n)
		for i := range dp {
			dp[i] = make([]int64, k+1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(p, c1 int, limitUp bool) int64
		f = func(p, c1 int, limitUp bool) (res int64) {
			if c1 > k {
				return
			}
			if p == n {
				if c1 == k {
					return 1
				}
				return
			}
			if !limitUp {
				dv := &dp[p][c1]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			up := upperC
			if limitUp {
				up = s[p]
			}
			for d := lowerC; d <= up; d++ {
				res += f(p+1, c1+int(d&1), limitUp && d == up)
			}
			return
		}
		return f(0, 0, true)
	}
	Fprint(out, search(1, 1e18, func(n int64) bool {
		return calc(strconv.FormatInt(n*2, 2))-calc(strconv.FormatInt(n, 2)) >= tar
	}))
}

//func main() { CF431D(os.Stdin, os.Stdout) }
