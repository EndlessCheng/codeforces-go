package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func p4999(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T int
	var lowS, highS string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &lowS, &highS)
		n := len(highS)
		diffLH := n - len(lowS)
		type pair struct{ cnt, sum int }
		memo := make([]pair, n)
		for i := range memo {
			memo[i].cnt = -1
		}

		var dfs func(int, bool, bool) pair
		dfs = func(i int, limitLow, limitHigh bool) (res pair) {
			if i == n {
				return pair{1, 0}
			}
			if !limitLow && !limitHigh {
				dv := &memo[i]
				if dv.cnt >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}

			lo := 0
			if limitLow && i >= diffLH {
				lo = int(lowS[i-diffLH] - '0')
			}
			hi := 9
			if limitHigh {
				hi = int(highS[i] - '0')
			}

			for d := lo; d <= hi; d++ {
				sub := dfs(i+1, limitLow && d == lo, limitHigh && d == hi)
				res.cnt = (res.cnt + sub.cnt) % mod
				res.sum = (res.sum + sub.sum + d*sub.cnt) % mod
			}
			return
		}
		Fprintln(out, dfs(0, true, true).sum)
	}
}

func main() { p4999(bufio.NewReader(os.Stdin), os.Stdout) }
