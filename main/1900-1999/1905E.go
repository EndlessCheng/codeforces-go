package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1905E(in io.Reader, out io.Writer) {
	const mod = 998244353
	f := map[int]int{1: 1}
	g := map[int]int{1: 0}
	p := map[int]int{1: 2}

	var dfs func(int)
	dfs = func(x int) {
		if f[x] != 0 {
			return
		}
		a := (x + 1) / 2
		b := x / 2
		dfs(a)
		dfs(b)
		f[x] = (f[a]*2 + f[b]*2 + (p[a]-1)*(p[b]-1)) % mod
		g[x] = (f[b] + g[a] + g[b]) % mod
		p[x] = p[a] * p[b] % mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		dfs(n)
		Fprintln(out, (f[n]+g[n])%mod)
	}
}

//func main() { cf1905E(os.Stdin, os.Stdout) }
