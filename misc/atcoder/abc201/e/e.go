package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	var n, v, w, wt, ans int
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	cnt := [60]int{}
	var dfs func(int, int, int)
	dfs = func(v, fa, xor int) {
		for s := uint(xor); s > 0; s &= s - 1 {
			cnt[bits.TrailingZeros(s)]++
		}
		for _, e := range g[v] {
			if e.to != fa {
				dfs(e.to, v, xor^e.wt)
			}
		}
	}
	dfs(1, 0, 0)

	for i, c := range cnt {
		ans += 1 << i % mod * c % mod * (n - c)
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
