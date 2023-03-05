package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF219D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, rev1 int
	Fscan(in, &n)
	type edge struct{ to, dir int }
	g := make([][]edge, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], edge{w, 1})
		g[w] = append(g[w], edge{v, -1})
	}

	var f func(int, int)
	f = func(v, fa int) {
		for _, e := range g[v] {
			if e.to != fa {
				if e.dir < 0 {
					rev1++
				}
				f(e.to, v)
			}
		}
	}
	f(1, 0)

	mn, ans := rev1, []int{}
	var reroot func(int, int, int)
	reroot = func(v, fa, rev int) {
		if rev < mn {
			mn = rev
			ans = []int{v}
		} else if rev == mn {
			ans = append(ans, v)
		}
		for _, e := range g[v] {
			if e.to != fa {
				reroot(e.to, v, rev+e.dir)
			}
		}
	}
	reroot(1, 0, rev1)

	Fprintln(out, mn)
	sort.Ints(ans)
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF219D(os.Stdin, os.Stdout) }
