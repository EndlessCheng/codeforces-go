package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF543D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var n, v int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		g[v-1] = append(g[v-1], w)
	}

	type pair struct {
		k int
		x int64
	}
	add1 := func(p pair) pair {
		if p.k > 0 {
			return pair{0, 1}
		}
		if p.x == mod-1 {
			return pair{1, 1}
		}
		return pair{0, p.x + 1}
	}
	mul := func(p, q pair) pair { return pair{p.k + q.k, p.x * q.x % mod} }
	div := func(p, q pair) pair { return pair{p.k - q.k, p.x * pow(q.x, mod-2) % mod} }

	f := make([]pair, n)
	var dfs func(int)
	dfs = func(v int) {
		f[v].x = 1
		for _, w := range g[v] {
			dfs(w)
			f[v] = mul(f[v], add1(f[w]))
		}
	}
	dfs(0)

	ans := make([]pair, n)
	ans[0] = f[0]
	var reroot func(int)
	reroot = func(v int) {
		for _, w := range g[v] {
			fp := div(ans[v], add1(f[w]))
			ans[w] = mul(f[w], add1(fp))
			reroot(w)
		}
	}
	reroot(0)

	for _, p := range ans {
		if p.k > 0 {
			Fprint(out, "0 ")
		} else {
			Fprint(out, p.x, " ")
		}
	}
}

//func main() { CF543D(os.Stdin, os.Stdout) }
