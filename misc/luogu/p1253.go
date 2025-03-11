package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type pair1253 struct{ x, y int }

type lazySeg1253 []struct {
	l, r int
	mx   int
	todo pair1253
}

func (lazySeg1253) mergeInfo(a, b int) int {
	return max(a, b)
}

var todoInit1253 = pair1253{2e9, 0}

func (lazySeg1253) mergeTodo(f, old pair1253) pair1253 {
	if f.x != 2e9 {
		return f
	}
	old.y += f.y
	return old
}

func (t lazySeg1253) apply(o int, f pair1253) {
	cur := &t[o]

	if f.x == 2e9 {
		cur.mx += f.y
	} else {
		cur.mx = f.x + f.y
	}

	cur.todo = t.mergeTodo(f, cur.todo)
}

func (t lazySeg1253) maintain(o int) {
	t[o].mx = t.mergeInfo(t[o<<1].mx, t[o<<1|1].mx)
}

func (t lazySeg1253) spread(o int) {
	f := t[o].todo
	if f == todoInit1253 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit1253
}

func (t lazySeg1253) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit1253
	if l == r {
		t[o].mx = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazySeg1253) update(o, l, r int, f pair1253) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t lazySeg1253) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func p1253(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i, _n := 0, 0
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	n, q := rd(), rd()
	a := make([]int, n)
	for i := range a {
		a[i] = rd()
	}
	t := make(lazySeg1253, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	for ; q > 0; q-- {
		op, l, r := rd(), rd()-1, rd()-1
		if op == 1 {
			t.update(1, l, r, pair1253{rd(), 0})
		} else if op == 2 {
			t.update(1, l, r, pair1253{2e9, rd()})
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { p1253(os.Stdin, os.Stdout) }
