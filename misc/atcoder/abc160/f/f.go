package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	g := make([][]int, n)
	fac := 1
	for i := 2; i <= n; i++ {
		fac = fac * i % mod
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	ans[0] = 1
	size := make([]int, n)
	var dfs func(int, int)
	dfs = func(x, fa int) {
		size[x] = 1
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				size[x] += size[y]
			}
		}
		ans[0] = ans[0] * size[x] % mod
	}
	dfs(0, -1)

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				ans[y] = ans[x] * pow(size[y], mod-2) % mod * (n - size[y]) % mod
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)

	for _, v := range ans {
		Fprintln(out, fac*pow(v, mod-2)%mod)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
