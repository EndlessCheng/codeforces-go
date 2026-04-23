package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3413(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var lowS, highS string
	Fscan(in, &lowS, &highS)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][11][11]int, n)

	var dfs func(int, int, int, bool, bool) int
	dfs = func(i, pp, pre int, limitLow, limitHigh bool) (res int) {
		if i == n {
			if pp == 0 && pre == 10 {
				return 1
			}
			return 0
		}
		if !limitLow && !limitHigh {
			p := &memo[i][pp][pre]
			if *p > 0 {
				return *p - 1
			}
			defer func() { *p = res + 1 }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		if limitLow && i < diffLH {
			res = dfs(i+1, pp, pre, true, false)
			d = 1
		}

		for ; d <= hi; d++ {
			if pp == 0 && pre == 10 {
				res += dfs(i+1, pp, pre, limitLow && d == lo, limitHigh && d == hi)
				continue
			}
			t1, t2 := pre, d
			if d == pre || d == pp {
				t1, t2 = 0, 10
			}
			res += dfs(i+1, t1, t2, limitLow && d == lo, limitHigh && d == hi)
		}
		res %= mod
		return
	}
	Fprint(out, dfs(0, 10, 10, true, true))
}

//func main() { p3413(bufio.NewReader(os.Stdin), os.Stdout) }
