package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1468J(in io.Reader, out io.Writer) {
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		fa := make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
		find := func(x int) int {
			rt := x
			for fa[rt] != rt {
				rt = fa[rt]
			}
			for fa[x] != rt {
				fa[x], x = rt, fa[x]
			}
			return rt
		}
		cc := n

		type edge struct{ v, w, wt int }
		es := []edge{}
		mx := 0
		for ; m > 0; m-- {
			var v, w, wt int
			Fscan(in, &v, &w, &wt)
			if wt > k {
				es = append(es, edge{v, w, wt})
				continue
			}
			mx = max(mx, wt)
			v = find(v)
			w = find(w)
			if v != w {
				cc--
				fa[v] = w
			}
		}
		slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

		if cc == 1 {
			ans := k - mx
			if len(es) > 0 {
				ans = min(ans, es[0].wt-k)
			}
			Fprintln(out, ans)
		} else {
			sum := 0
			for _, e := range es {
				v, w := find(e.v), find(e.w)
				if v != w {
					fa[v] = w
					sum += e.wt - k
				}
			}
			Fprintln(out, sum)
		}
	}
}

//func main() { cf1468J(bufio.NewReader(os.Stdin), os.Stdout) }
