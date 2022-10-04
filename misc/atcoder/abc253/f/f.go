package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"runtime/debug"
)

// https://space.bilibili.com/206214
type node struct {
	lo, ro *node
	sum    int
}

func build(l, r int) *node {
	o := &node{}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build(l, m)
	o.ro = build(m+1, r)
	return o
}

func (o *node) update(ol, or, l, r, add int) {
	if l <= ol && or <= r {
		o.sum += add
		return
	}
	lo, ro := *o.lo, *o.ro
	o.lo, o.ro = &lo, &ro
	if add := o.sum; add > 0 {
		o.lo.sum += add
		o.ro.sum += add
		o.sum = 0
	}
	m := (ol + or) / 2
	if l <= m {
		o.lo.update(ol, m, l, r, add)
	}
	if r > m {
		o.ro.update(m+1, or, l, r, add)
	}
}

func (o *node) query(ol, or, i int) int {
	if ol == or {
		return o.sum
	}
	m := (ol + or) / 2
	if i <= m {
		return o.sum + o.lo.query(ol, m, i)
	}
	return o.sum + o.ro.query(m+1, or, i)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, op, l, r, v int
	Fscan(in, &n, &m, &q)
	t := make([]*node, 1, q+1)
	t[0] = build(1, m)
	type pair struct{ i, v int }
	last := make([]pair, n+1)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			Fscan(in, &v)
			o := *t[len(t)-1]
			o.update(1, m, l, r, v)
			t = append(t, &o)
		} else if op == 2 {
			last[l] = pair{len(t) - 1, r}
		} else {
			p := last[l]
			Fprintln(out, p.v+t[len(t)-1].query(1, m, r)-t[p.i].query(1, m, r))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
func init() { debug.SetGCPercent(-1) }
