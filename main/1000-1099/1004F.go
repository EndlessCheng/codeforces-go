package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
var x04 int
type pair04 struct{ v, cnt int }
type data04 struct {
	pre, suf []pair04
	cnt      int
}
type seg04 []struct {
	l, r int
	val  data04
}

func (seg04) merge(a, b []pair04) []pair04 {
	c := append(make([]pair04, 0, len(a)+len(b)), a...)
	for j := range c {
		c[j].v |= b[0].v
	}
	c = append(c, b...)
	j := 0
	for _, t := range c[1:] {
		if c[j].v != t.v {
			j++
			c[j] = t
		} else {
			c[j].cnt += t.cnt
		}
	}
	return c[:j+1]
}

func (t seg04) mergeInfo(ld, rd data04) data04 {
	cnt := ld.cnt + rd.cnt
	a, b := ld.suf, rd.pre // 0: max, -1: min
	s, i := 0, 0
	for j := len(b) - 1; j >= 0; j-- {
		t := b[j]
		for ; i < len(a) && a[i].v|t.v >= x04; i++ {
			s += a[i].cnt
		}
		cnt += s * t.cnt
	}
	return data04{t.merge(rd.pre, ld.pre), t.merge(ld.suf, rd.suf), cnt}
}

func (t seg04) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		p := pair04{a[l-1], 1}
		t[o].val.pre = []pair04{p}
		t[o].val.suf = []pair04{p}
		if a[l-1] >= x04 {
			t[o].val.cnt = 1
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg04) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].val.pre[0].v = val
		t[o].val.suf[0].v = val
		if val >= x04 {
			t[o].val.cnt = 1
		} else {
			t[o].val.cnt = 0
		}
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

func (t seg04) maintain(o int) {
	t[o].val = t.mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg04) query(o, l, r int) (res data04) {
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
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1004F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, op, l, r int
	Fscan(in, &n, &m, &x04)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg04, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			t.update(1, l, r)
		} else {
			Fprintln(out, t.query(1, l, r).cnt)
		}
	}
}

//func main() { cf1004F(os.Stdin, os.Stdout) }
