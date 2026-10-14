package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1)
}

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return t.pre(r) - t.pre(l-1)
}

func cf1139F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	v := make([]int, 0, 2*n+2*m)
	type node struct{ x, o, i int }
	q := make([]node, 0, 2*n+m)

	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &p[i])
		q = append(q, node{p[i], -1, i})
	}
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		q = append(q, node{s[i], 1, i})
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
		v = append(v, b[i]+p[i], b[i]-p[i]+1)
	}
	x := make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &x[i])
		q = append(q, node{x[i], 0, i})
	}
	y := make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &y[i])
		v = append(v, x[i]+y[i], y[i]-x[i])
	}

	slices.SortFunc(q, func(a, b node) int { return cmp.Or(a.x-b.x, a.o-b.o) })
	slices.Sort(v)
	v = slices.Compact(v)
	val := func(u int) int { return sort.SearchInts(v, u) + 1 }

	t := [2]fenwick{newFenwickTree(len(v)), newFenwickTree(len(v))}
	ans := make([]int, m+1)
	for _, u := range q {
		if u.o != 0 {
			t[0].update(val(b[u.i]+p[u.i]), -u.o)
			t[1].update(val(b[u.i]-p[u.i]+1), -u.o)
		} else {
			ans[u.i] = t[0].pre(val(x[u.i]+y[u.i])) - t[1].pre(val(y[u.i]-x[u.i]))
		}
	}

	for i := 1; i <= m; i++ {
		Fprint(out, ans[i], " ")
	}
}

//func main() { cf1139F(bufio.NewReader(os.Stdin), os.Stdout) }
