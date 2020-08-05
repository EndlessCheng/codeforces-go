package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct {
	l, r, sum int
	todo      bool
}

func (t seg) maintain(o int) { t[o].sum = t[o<<1].sum + t[o<<1|1].sum }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o int) {
	to := &t[o]
	to.todo = !to.todo
	to.sum = to.r - to.l + 1 - to.sum
}

func (t seg) spread(o int) {
	if t[o].todo {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].todo = false
	}
}

func (t seg) flip(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.flip(o<<1, l, r)
	}
	if m < r {
		t.flip(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg) query(o, k int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	if k < t[o<<1].sum {
		return t.query(o<<1, k)
	}
	return t.query(o<<1|1, k-t[o<<1].sum)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, k int
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &l, &r)
			t.flip(1, l+1, r)
		} else {
			Fscan(in, &k)
			Fprintln(out, t.query(1, k)-1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
