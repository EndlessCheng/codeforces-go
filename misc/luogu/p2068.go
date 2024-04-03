package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg2068 []struct{ l, r, val int }

func (t seg2068) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg2068) update(o, i int, val int) {
	if t[o].l == t[o].r {
		t[o].val += val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg2068) maintain(o int) {
	t[o].val = t[o<<1].val + t[o<<1|1].val
}

func (t seg2068) query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func p2068(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	var op string
	Fscan(in, &n, &m)
	t := make(seg2068, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &op, &l, &r)
		if op == "x" {
			t.update(1, l, r)
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { p2068(os.Stdin, os.Stdout) }
