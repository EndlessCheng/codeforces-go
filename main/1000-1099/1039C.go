package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1039C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	var n, m, k, v, w int
	Fscan(in, &n, &m, &k)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type edge struct{ v, w int }
	mp := map[int64][]edge{}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		s := a[v] ^ a[w]
		mp[s] = append(mp[s], edge{v, w})
	}
	p2 := make([]int64, n+1)
	p2[0] = 1
	for i := 1; i <= n; i++ {
		p2[i] = p2[i-1] * 2 % mod
	}

	ans := (int64(1)<<k - int64(len(mp))) % mod * p2[n] % mod
	vis := make([]int64, n+1)
	for s, es := range mp {
		vs := map[int]struct{}{}
		g := map[int][]int{}
		for _, e := range es {
			v, w := e.v, e.w
			vs[v] = struct{}{}
			vs[w] = struct{}{}
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		cc := n - len(vs)
		var f func(int)
		f = func(v int) {
			vis[v] = s
			for _, w := range g[v] {
				if vis[w] != s {
					f(w)
				}
			}
		}
		for v := range vs {
			if vis[v] != s {
				cc++
				f(v)
			}
		}
		ans += p2[cc]
	}
	Fprint(out, ans%mod)
}

//func main() { CF1039C(os.Stdin, os.Stdout) }
