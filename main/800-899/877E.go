package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg []struct {
	l, r int
	ones int
	flip bool
}

func (t seg) maintain(o int) {
	t[o].ones = t[o<<1].ones + t[o<<1|1].ones
}

func (t seg) doFlip(O int) {
	o := &t[O]
	o.ones = o.r - o.l + 1 - o.ones
	o.flip = !o.flip
}

func (t seg) spread(o int) {
	if t[o].flip {
		t.doFlip(o << 1)
		t.doFlip(o<<1 | 1)
		t[o].flip = false
	}
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].ones = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) flip(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.doFlip(o)
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

func (t seg) onesCount(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].ones
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.onesCount(o<<1, l, r)
	}
	if m < l {
		return t.onesCount(o<<1|1, l, r)
	}
	return t.onesCount(o<<1, l, r) + t.onesCount(o<<1|1, l, r)
}

func cf877E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, p, time, q int
	var op string
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &p)
		g[p-1] = append(g[p-1], w)
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := make([]int, n)
	nodes := make([]struct{ l, r int }, n)
	var f func(int) int
	f = func(v int) int {
		b[time] = a[v]
		time++
		nodes[v].l = time
		sz := 1
		for _, w := range g[v] {
			sz += f(w)
		}
		nodes[v].r = nodes[v].l + sz - 1
		return sz
	}
	f(0)

	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(b, 1, 1, n)
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &p)
		o := nodes[p-1]
		if op[0] == 'p' {
			t.flip(1, o.l, o.r)
		} else {
			Fprintln(out, t.onesCount(1, o.l, o.r))
		}
	}
}

//func main() { cf877E(os.Stdin, os.Stdout) }
