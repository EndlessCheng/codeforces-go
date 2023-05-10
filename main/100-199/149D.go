package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF149D(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	var s string
	Fscan(in, &s)
	n := len(s)
	right := make([]int, n)
	st := []int{}
	for i, c := range s {
		if c == '(' {
			st = append(st, i)
		} else {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
	}

	dp := make([][][3][3]int64, n)
	for i := range dp {
		dp[i] = make([][3][3]int64, n)
		for j := range dp[i] {
			dp[i][j] = [3][3]int64{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
		}
	}
	var f func(int, int, int, int) int64
	f = func(l, r, lc, rc int) (res int64) {
		if l > r {
			return 1
		}
		ptr := &dp[l][r][lc][rc]
		if *ptr != -1 {
			return *ptr
		}
		mid := right[l]
		if mid < r {
			res += f(l+1, mid-1, 0, 1) * f(mid+1, r, 1, rc)
			res += f(l+1, mid-1, 0, 2) * f(mid+1, r, 2, rc)
			if lc != 1 {
				res += f(l+1, mid-1, 1, 0) * f(mid+1, r, 0, rc)
			}
			if lc != 2 {
				res += f(l+1, mid-1, 2, 0) * f(mid+1, r, 0, rc)
			}
		} else {
			if lc != 1 {
				res += f(l+1, r-1, 1, 0)
			}
			if lc != 2 {
				res += f(l+1, r-1, 2, 0)
			}
			if rc != 1 {
				res += f(l+1, r-1, 0, 1)
			}
			if rc != 2 {
				res += f(l+1, r-1, 0, 2)
			}
		}
		*ptr = res % mod
		return *ptr
	}
	Fprint(out, f(0, n-1, 0, 0))
}

//func main() { CF149D(os.Stdin, os.Stdout) }
