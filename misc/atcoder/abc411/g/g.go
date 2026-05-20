package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	const inv2 = 499122177
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v][w]++
		g[w][v]++
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(mask, v int) int {
		p := &dp[mask][v]
		if *p != -1 {
			return *p
		}

		st := bits.TrailingZeros32(uint32(mask))
		res := 0
		if bits.OnesCount32(uint32(mask)) > 2 {
			res = g[v][st]
		}
		for w := st + 1; w < n; w++ {
			if mask>>w&1 == 0 {
				res += dfs(mask|1<<w, w) * g[v][w]
			}
		}
		res %= mod

		*p = res
		return res
	}

	ans := 0
	for i := range n {
		ans += dfs(1<<i, i)
	}
	for i, r := range g {
		for _, c := range r[:i] {
			ans += c * (c - 1)
		}
	}
	Fprint(out, ans%mod*inv2%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
