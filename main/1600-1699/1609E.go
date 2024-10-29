package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type data09 struct{ a, b, c, ab, bc, abc int }

type seg09 []struct {
	l, r int
	data09
}

func (seg09) mergeInfo(l, r data09) data09 {
	return data09{
		l.a + r.a,
		l.b + r.b,
		l.c + r.c,
		min(l.a+r.ab, l.ab+r.b),
		min(l.b+r.bc, l.bc+r.c),
		min(l.a+r.abc, l.ab+r.bc, l.abc+r.c),
	}
}

func (seg09) newData(c byte) (d data09) {
	if c == 'a' {
		d.a = 1
	} else if c == 'b' {
		d.b = 1
	} else {
		d.c = 1
	}
	return
}

func (t seg09) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].data09 = t.newData(s[l])
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg09) update(o, i int, c byte) {
	if t[o].l == t[o].r {
		t[o].data09 = t.newData(c)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, c)
	} else {
		t.update(o<<1|1, i, c)
	}
	t.maintain(o)
}

func (t seg09) maintain(o int) {
	t[o].data09 = t.mergeInfo(t[o<<1].data09, t[o<<1|1].data09)
}

func cf1609E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, i int
	var s string
	Fscan(in, &n, &q, &s)
	t := make(seg09, 2<<bits.Len(uint(n-1)))
	t.build(s, 1, 0, n-1)
	for range q {
		Fscan(in, &i, &s)
		t.update(1, i-1, s[0])
		Fprintln(out, t[1].abc)
	}
}

//func main() { cf1609E(bufio.NewReader(os.Stdin), os.Stdout) }
