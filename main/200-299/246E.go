package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF246E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, t, q, d int
	Fscan(in, &n)
	g := make([][]int, n)
	roots := []int{}
	a := make([]string, n)
	for w := 0; w < n; w++ {
		Fscan(in, &a[w], &v)
		if v > 0 {
			g[v-1] = append(g[v-1], w)
		} else {
			roots = append(roots, w)
		}
	}

	type info struct{ in, out, d int }
	is := make([]info, n)
	depT := make([][]int, n)
	rows := make([][]string, n)
	var f func(v, p, d int)
	f = func(v, p, d int) {
		t++
		is[v].in = t
		is[v].d = d
		depT[d] = append(depT[d], t)
		rows[d] = append(rows[d], a[v])
		for _, w := range g[v] {
			if w != p {
				f(w, v, d+1)
			}
		}
		is[v].out = t
	}
	for _, root := range roots {
		f(root, -1, 0)
	}

	Fscan(in, &q)
	type query struct{ l, r, i int }
	qus := make([][]query, n)
	for i := 0; i < q; i++ {
		Fscan(in, &v, &d)
		nf := is[v-1]
		if d += nf.d; d < n {
			l := sort.SearchInts(depT[d], nf.in)
			r := sort.SearchInts(depT[d], nf.out+1)
			qus[d] = append(qus[d], query{l, r, i})
		}
	}
	ans := make([]int, q)
	var tree []int
	add := func(i, v int) {
		for i++; i < len(tree); i += i & -i {
			tree[i] += v
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	for i, qs := range qus {
		row := rows[i]
		tree = make([]int, len(row)+1)
		j, posR := 0, map[string]int{}
		sort.Slice(qs, func(i, j int) bool { return qs[i].r < qs[j].r })
		for _, q := range qs {
			for ; j < q.r; j++ {
				if p, ok := posR[row[j]]; ok {
					add(p, -1)
				}
				add(j, 1)
				posR[row[j]] = j
			}
			ans[q.i] = sum(q.r) - sum(q.l)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF246E(os.Stdin, os.Stdout) }
