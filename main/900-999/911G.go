package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type node11 struct{ lo, ro *node11 }

func (o *node11) put(l, r, i int) *node11 {
	if o == nil {
		o = &node11{}
	}
	if l == r {
		return o
	}
	m := (l + r) / 2
	if i <= m {
		o.lo = o.lo.put(l, m, i)
	} else {
		o.ro = o.ro.put(m+1, r, i)
	}
	return o
}

func (o *node11) merge(b *node11) *node11 {
	if b == nil {
		return o
	}
	if o == nil {
		return b
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	return o
}

func move11(from, to **node11, l, r, ql, qr int) {
	if *from == nil {
		return
	}
	if ql <= l && r <= qr {
		*to = (*to).merge(*from)
		*from = nil
		return
	}
	if *to == nil {
		*to = &node11{}
	}
	m := (l + r) / 2
	if ql <= m {
		move11(&(*from).lo, &(*to).lo, l, m, ql, qr)
	}
	if qr > m {
		move11(&(*from).ro, &(*to).ro, m+1, r, ql, qr)
	}
}

func (o *node11) collect(ans []int, l, r, v int) {
	if o == nil {
		return
	}
	if l == r {
		ans[l] = v
		return
	}
	m := (l + r) / 2
	o.lo.collect(ans, l, m, v)
	o.ro.collect(ans, m+1, r, v)
}

func cf911G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, v, q, l, r, x, y int
	Fscan(in, &n)
	roots := [101]*node11{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		roots[v] = roots[v].put(1, n, i)
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r, &x, &y)
		if x != y {
			move11(&roots[x], &roots[y], 1, n, l, r)
		}
	}

	ans := make([]int, n+1)
	for v, rt := range roots {
		rt.collect(ans, 1, n, v)
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { cf911G(bufio.NewReader(os.Stdin), os.Stdout) }
