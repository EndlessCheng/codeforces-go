package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1095F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, mnI, v, w int
	var wt int64
	Fscan(in, &n, &m)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < a[mnI] {
			mnI = i
		}
	}
	type edge struct {
		v, w int
		wt   int64
	}
	es := make([]edge, 0, n-1+m)
	for i, x := range a {
		if i != mnI {
			es = append(es, edge{mnI, i, a[mnI] + x})
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		es = append(es, edge{v - 1, w - 1, wt})
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := int64(0)
	for _, e := range es {
		if fv, fw := find(e.v), find(e.w); fv != fw {
			fa[fv] = fw
			ans += e.wt
		}
	}
	Fprint(out, ans)
}

//func main() { CF1095F(os.Stdin, os.Stdout) }
