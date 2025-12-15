package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	. "slices"
)

// https://github.com/EndlessCheng
const mod11 = 95542721
const period11 = 48

func rotateLeft(a []int, k int) {
	Reverse(a[:k])
	Reverse(a[k:])
	Reverse(a)
}

type seg11 []struct {
	l, r int
	sum  [period11]int
	todo int
}

func (t seg11) apply(o, f int) {
	rotateLeft(t[o].sum[:], f%period11)
	t[o].todo += f
}

func (t seg11) maintain(o int) {
	for i := range period11 {
		t[o].sum[i] = (t[o<<1].sum[i] + t[o<<1|1].sum[i]) % mod11
	}
}

func (t seg11) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg11) build(in io.Reader, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		var x int
		Fscan(in, &x)
		s := &t[o].sum
		s[0] = x
		for i := 1; i < period11; i++ {
			x = x * x % mod11 * x % mod11
			s[i] = x
		}
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg11) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, 1)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg11) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum[0]
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func cf311D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n)
	t := make(seg11, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 1, n)
	Fscan(in, &q)
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			Fprintln(out, t.query(1, l, r)%mod11)
		} else {
			t.update(1, l, r)
		}
	}
}

//func main() { cf311D(bufio.NewReader(os.Stdin), os.Stdout) }
