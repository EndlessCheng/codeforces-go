package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type seg []struct{ l, r, minTodo, maxTodo int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].minTodo = l, r, 1e9
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) doUp(o, v int) {
	to := &t[o]
	to.minTodo = max(to.minTodo, v)
	to.maxTodo = max(to.maxTodo, v)
}
func (t seg) doDown(o, v int) {
	to := &t[o]
	to.minTodo = min(to.minTodo, v)
	to.maxTodo = min(to.maxTodo, v)
}
func (t seg) spread(o int) {
	t.doUp(o<<1, t[o].maxTodo)
	t.doUp(o<<1|1, t[o].maxTodo)
	t[o].maxTodo = 0
	t.doDown(o<<1, t[o].minTodo)
	t.doDown(o<<1|1, t[o].minTodo)
	t[o].minTodo = 1e9
}

func (t seg) update(o, l, r, v int, up bool) {
	if l <= t[o].l && t[o].r <= r {
		if up {
			t.doUp(o, v)
		} else {
			t.doDown(o, v)
		}
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v, up)
	}
	if m < r {
		t.update(o<<1|1, l, r, v, up)
	}
}

func (t seg) query(o int, out io.Writer) {
	if t[o].l == t[o].r {
		Fprintln(out, t[o].maxTodo)
		return
	}
	t.spread(o)
	t.query(o<<1, out)
	t.query(o<<1|1, out)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r, &v)
		t.update(1, l+1, r+1, v, op == 1)
	}
	t.query(1, out)
}

func main() { run(os.Stdin, os.Stdout) }
