package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1363E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type node struct{ c, d int }
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, v, w, d int
	Fscan(in, &n)
	a := make([]node, n)
	for i := range a {
		Fscan(in, &a[i].c, &v, &w)
		if v != w {
			a[i].d = v
			d += v - w
		} else {
			a[i].d = 2
		}
	}
	if d != 0 {
		Fprint(out, -1)
		return
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := int64(0)
	var f func(v, fa, c int) [3]int
	f = func(v, fa, c int) (cnt [3]int) {
		cnt[a[v].d]++
		for _, w := range g[v] {
			if w != fa {
				cc := f(w, v, min(c, a[w].c))
				cnt[0] += cc[0]
				cnt[1] += cc[1]
			}
		}
		md := min(cnt[0], cnt[1])
		ans += int64(c) * int64(2*md)
		cnt[0] -= md
		cnt[1] -= md
		return
	}
	f(0, -1, a[0].c)
	Fprint(out, ans)
}

//func main() { CF1363E(os.Stdin, os.Stdout) }
