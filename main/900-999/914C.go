package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf914C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s string
	var k int
	Fscan(in, &s, &k)
	if k == 0 {
		Fprint(out, 1)
		return
	}
	if k == 1 {
		Fprint(out, len(s)-1)
		return
	}

	n := len(s)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool) int
	dfs = func(i, left1 int, isLimit bool) (res int) {
		if left1 == 0 {
			return 1
		}
		if i == n {
			return
		}
		if !isLimit {
			p := &memo[i][left1]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		up := 1
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= up; d++ {
			res += dfs(i+1, left1-d, isLimit && d == up)
		}
		return res % mod
	}

	ans := 0
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[bits.OnesCount(uint(i))] + 1
		if f[i] == k {
			ans += dfs(0, i, true)
		}
	}
	Fprint(out, ans%mod)
}

//func main() { cf914C(os.Stdin, os.Stdout) }
