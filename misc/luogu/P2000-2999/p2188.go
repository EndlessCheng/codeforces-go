package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2188(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var lowS, highS string
	var k int
	Fscan(in, &lowS, &highS, &k)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][10]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
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

		isFirst := limitLow && i <= diffLH
		for ; d <= hi; d++ {
			if isFirst || abs(d-pre) <= k {
				res += dfs(i+1, d, limitLow && d == lo, limitHigh && d == hi)
			}
		}
		return
	}

	Fprint(out, dfs(0, 0, true, true))
}

//func main() { p2188(bufio.NewReader(os.Stdin), os.Stdout) }
