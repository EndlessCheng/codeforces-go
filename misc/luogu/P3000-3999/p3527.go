package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func p3527(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, v int
	Fscan(in, &n, &m)
	pos := make([][]int, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v)
		v--
		pos[v] = append(pos[v], i)
	}
	need := make([]int, n)
	c := make([]int, n)
	for i := range need {
		Fscan(in, &need[i])
		c[i] = i
	}
	Fscan(in, &q)
	qs := make([]struct{ l, r, v int }, q)
	for i := range qs {
		Fscan(in, &qs[i].l, &qs[i].r, &qs[i].v)
		qs[i].r++
	}

	ans := make([]int, n)
	tree := make([]int, m+1)
	add := func(i, v int) {
		for ; i <= m; i += i & -i {
			tree[i] += v
		}
	}
	addRange := func(l, r, v int) {
		if l < r {
			add(l, v)
			add(r, -v)
		} else {
			add(1, v)
			add(r, -v)
			add(l, v)
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	var f func([]int, int, int)
	f = func(c []int, l, r int) {
		if len(c) == 0 {
			return
		}
		if l+1 == r {
			for _, c := range c {
				ans[c] = r
			}
			return
		}

		m := (l + r) / 2
		for _, q := range qs[l:m] {
			addRange(q.l, q.r, q.v)
		}

		var lc, rc []int
	o:
		for _, c := range c {
			s := 0
			for _, p := range pos[c] {
				if s += sum(p); s >= need[c] {
					lc = append(lc, c)
					continue o
				}
			}
			need[c] -= s
			rc = append(rc, c)
		}

		for _, q := range qs[l:m] {
			addRange(q.l, q.r, -q.v)
		}
		f(lc, l, m)
		f(rc, m, r)
	}
	f(c, 0, q+1) // q+1 用来标记无法达标的国家
	for _, v := range ans {
		if v > q {
			Fprintln(out, "NIE")
		} else {
			Fprintln(out, v)
		}
	}
}

//func main() { p3527(os.Stdin, os.Stdout) }
