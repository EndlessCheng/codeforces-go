package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1073E(in io.Reader, out io.Writer) {
	const mod = 998244353
	var lowS, highS string
	var k int
	Fscan(in, &lowS, &highS, &k)
	n := len(highS)
	diffLH := n - len(lowS)
	type pair struct{ num, sum int }
	memo := make([][1 << 10]pair, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j].num = -1
		}
	}

	var dfs func(int, int, bool, bool) pair
	dfs = func(i, mask int, limitLow, limitHigh bool) (res pair) {
		if i == n {
			return pair{1, 0}
		}
		if !limitLow && !limitHigh {
			dv := &memo[i][mask]
			if dv.num >= 0 {
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

		d := lo
		if limitLow && i < diffLH {
			res = dfs(i+1, 0, true, false)
			d = 1
		}

		for ; d <= hi; d++ {
			newMask := mask | 1<<d
			if bits.OnesCount(uint(newMask)) > k {
				continue
			}
			sub := dfs(i+1, newMask, limitLow && d == lo, limitHigh && d == hi)
			res.num = (res.num + sub.num) % mod
			v := d * int(math.Pow10(n-1-i)) % mod
			res.sum = (res.sum + sub.sum + v*sub.num) % mod
		}
		return
	}
	Fprint(out, dfs(0, 0, true, true).sum)
}

//func main() { cf1073E(os.Stdin, os.Stdout) }
