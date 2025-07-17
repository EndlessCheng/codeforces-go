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
	sum := make([]int, n+1)
	f := make([]int, n+1)
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

	pa := make([]int, n+1)
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

	for _, e := range es {
		x, y := find(e.v), find(e.w)
		if x == y {
			continue
		}
		f[y] = max(min(f[y], e.wt)-sum[x], min(f[x], e.wt)-sum[y])
		sum[y] += sum[x]
		pa[x] = y
	}

	ans := f[find(1)]
	if ans <= 0 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf1578L(bufio.NewReader(os.Stdin), os.Stdout) }
