package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data [30][2]uint8

type seg []data

func (t seg) set(o, x int) {
	for i := range 30 {
		t[o][i][0] = 1
		t[o][i][1] = uint8(x>>i&1 ^ 1)
	}
}

func (seg) merge(a, b data) (c data) {
	for i := range 30 {
		c[i][0] = b[i][a[i][0]]
		c[i][1] = b[i][a[i][1]]
	}
	return
}

func (t seg) maintain(o int) {
	t[o] = t.merge(t[o*2], t[o*2+1])
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t.set(o, a[l])
		return
	}
	m := (l + r) >> 1
	t.build(a, o*2, l, m)
	t.build(a, o*2+1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, i, x int) {
	if l == r {
		t.set(o, x)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o*2, l, m, i, x)
	} else {
		t.update(o*2+1, m+1, r, i, x)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r, ql, qr int) data {
	if ql <= l && r <= qr {
		return t[o]
	}
	m := (l + r) >> 1
	if qr <= m {
		return t.query(o*2, l, m, ql, qr)
	}
	if ql > m {
		return t.query(o*2+1, m+1, r, ql, qr)
	}
	return t.merge(t.query(o*2, l, m, ql, qr), t.query(o*2+1, m+1, r, ql, qr))
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			a[l] = r
			t.update(1, 1, n, l, r)
			continue
		}
		ans := 0
		if l == r {
			ans = a[l]
		} else {
			mp := t.query(1, 1, n, l+1, r)
			for i := range 30 {
				ans |= int(mp[i][a[l]>>i&1]) << i
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
