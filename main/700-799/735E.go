package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf735E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k, ans int
	Fscan(in, &n, &k)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var dfs func(int, int) []int
	dfs = func(v, fa int) []int {
		f := make([]int, k*2+2)
		f[0] = 1
		f[k+1] = 1
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			fw := dfs(w, v)
			nf := make([]int, k*2+2)
			for i, pre := range f[:k*2+1] {
				for j, cur := range fw[:k*2+1] {
					p := min(j+1, i)
					if i+j > k*2 {
						p = max(j+1, i)
					}
					nf[p] = (nf[p] + pre*cur) % mod
				}
			}
			f = nf
		}
		return f
	}

	f := dfs(1, 0)
	for _, v := range f[:k+1] {
		ans += v
	}
	Fprint(out, ans%mod)
}

//func main() { cf735E(bufio.NewReader(os.Stdin), os.Stdout) }
