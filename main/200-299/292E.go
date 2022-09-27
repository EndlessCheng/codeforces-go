package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type seg92 []struct{ l, r, d int }

func (t seg92) build(o, l, r int) {
	t[o].l, t[o].r, t[o].d = l, r, 1e9
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg92) update(o, l, r, d int) {
	if l <= t[o].l && t[o].r <= r {
		t[o].d = d
		return
	}
	if d := t[o].d; d != 1e9 {
		t[o<<1].d = d
		t[o<<1|1].d = d
		t[o].d = 1e9
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, d)
	}
	if m < r {
		t.update(o<<1|1, l, r, d)
	}
}

func (t seg92) query(o, i int) int {
	if t[o].d != 1e9 || t[o].l == t[o].r {
		return t[o].d
	}
	if i <= (t[o].l+t[o].r)>>1 {
		return t.query(o<<1, i)
	}
	return t.query(o<<1|1, i)
}

func CF292E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, op, x, y, k int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
	}
	t := make(seg92, n*4)
	t.build(1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &op, &x)
		if op == 1 {
			Fscan(in, &y, &k)
			t.update(1, y, y+k-1, x-y)
		} else {
			d := t.query(1, x)
			if d == 1e9 {
				Fprintln(out, b[x])
			} else {
				Fprintln(out, a[x+d])
			}
		}
	}
}

//func main() { CF292E(os.Stdin, os.Stdout) }
