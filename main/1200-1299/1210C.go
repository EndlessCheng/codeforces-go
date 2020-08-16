package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1210C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, v, w int
	Fscan(in, &n)
	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type pair struct{ g, c int64 }
	cnt := make([][]pair, n+1)
	var f func(v, fa int)
	f = func(v, fa int) {
		x := a[v]
		for _, p := range cnt[fa] {
			if gg := gcd(p.g, x); cnt[v] == nil || gg != cnt[v][len(cnt[v])-1].g {
				cnt[v] = append(cnt[v], pair{gg, p.c})
			} else {
				cnt[v][len(cnt[v])-1].c += p.c
			}
		}
		if cnt[v] == nil || x != cnt[v][len(cnt[v])-1].g {
			cnt[v] = append(cnt[v], pair{x, 1})
		} else {
			cnt[v][len(cnt[v])-1].c++
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
	}
	f(1, 0)
	ans := int64(0)
	for _, ps := range cnt {
		for _, p := range ps {
			ans = (ans + p.g*p.c) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { CF1210C(os.Stdin, os.Stdout) }
