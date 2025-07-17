package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1578L(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	sum := make([]int, n*2)
	f := make([]int, n*2)
	for i := 1; i <= n; i++ {
		Fscan(in, &sum[i])
		f[i] = 1e18
	}
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
	}
	slices.SortFunc(es, func(a, b edge) int { return b.wt - a.wt })

	pa := make([]int, n*2)
	for i := range pa {
		pa[i] = i
	}
	find := func(x int) int {
		rt := x
		for pa[rt] != rt {
			rt = pa[rt]
		}
		for pa[x] != rt {
			pa[x], x = rt, pa[x]
		}
		return rt
	}

	k := n
	for _, e := range es {
		x, y := find(e.v), find(e.w)
		if x == y {
			continue
		}
		k++
		pa[x] = k
		pa[y] = k
		f[k] = max(min(f[y], e.wt)-sum[x], min(f[x], e.wt)-sum[y])
		sum[k] = sum[x] + sum[y]
	}
	if f[k] > 0 {
		Fprint(out, f[k])
	} else {
		Fprint(out, -1)
	}
}

//func main() { cf1578L(bufio.NewReader(os.Stdin), os.Stdout) }
