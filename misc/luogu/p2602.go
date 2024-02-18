package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2602(in io.Reader, out io.Writer) {
	var low, high string
	Fscan(in, &low, &high)
	n := len(high)
	diffLH := n - len(low)
	for tar := 0; tar < 10; tar++ {
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, n+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		var dfs func(int, int, bool, bool) int
		dfs = func(p, cnt int, limitLow, limitHigh bool) (res int) {
			if p == n {
				return cnt
			}
			if !limitLow && !limitHigh {
				dv := &memo[p][cnt]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}

			lo := 0
			if limitLow && p >= diffLH {
				lo = int(low[p-diffLH] - '0')
			}
			hi := 9
			if limitHigh {
				hi = int(high[p] - '0')
			}

			d := lo
			if limitLow && p < diffLH {
				// 什么也不填
				res = dfs(p+1, 0, true, false)
				d++
			}
			for ; d <= hi; d++ {
				c := cnt
				if d == tar {
					c++
				}
				res += dfs(p+1, c, limitLow && d == lo, limitHigh && d == hi)
			}
			return
		}
		Fprint(out, dfs(0, 0, true, true), " ")
	}
}

//func main() { p2602(os.Stdin, os.Stdout) }
