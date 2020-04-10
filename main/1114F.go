package main

import (
	"bufio"
	. "fmt"
	"io"
)

const mod1114 int64 = 1e9 + 7

type mulST1114 []struct {
	l, r      int
	mul, todo int64
}

func (mulST1114) pow(x int64, n int) int64 {
	res := int64(1)
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod1114
		}
		x = x * x % mod1114
	}
	return res
}

func (t mulST1114) _pushUp(o int) {
	t[o].mul = t[o<<1].mul * t[o<<1|1].mul % mod1114
}

func (t mulST1114) _build(a []int, o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, 1
	if l == r {
		t[o].mul = int64(a[l-1])
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t mulST1114) _spread(o int) {
	if mul := t[o].todo; mul != 1 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.mul = lo.mul * t.pow(mul, lo.r-lo.l+1) % mod1114
		ro.mul = ro.mul * t.pow(mul, ro.r-ro.l+1) % mod1114
		lo.todo = lo.todo * mul % mod1114
		ro.todo = ro.todo * mul % mod1114
		t[o].todo = 1
	}
}

func (t mulST1114) _update(o, l, r, val int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].mul = t[o].mul * t.pow(int64(val), or-ol+1) % mod1114
		t[o].todo = t[o].todo * int64(val) % mod1114
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._update(o<<1, l, r, val)
	}
	if m < r {
		t._update(o<<1|1, l, r, val)
	}
	t._pushUp(o)
}

func (t mulST1114) _query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mul
	}
	t._spread(o)
	res = 1
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		res = t._query(o<<1, l, r)
	}
	if m < r {
		res = res * t._query(o<<1|1, l, r) % mod1114
	}
	return
}

func (t mulST1114) init(a []int)         { t._build(a, 1, 1, len(a)) }
func (t mulST1114) update(l, r, val int) { t._update(1, l, r, val) }
func (t mulST1114) query(l, r int) int64 { return t._query(1, l, r) }

type orST1114 []struct {
	l, r     int
	or, todo int64
}

func (t orST1114) _pushUp(o int) {
	t[o].or = t[o<<1].or | t[o<<1|1].or
}

func (t orST1114) _build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].or = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t orST1114) _spread(o int) {
	if or := t[o].todo; or != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.or |= or
		ro.or |= or
		lo.todo |= or
		ro.todo |= or
		t[o].todo = 0
	}
}

func (t orST1114) _update(o, l, r int, val int64) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].or |= val
		t[o].todo |= val
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._update(o<<1, l, r, val)
	}
	if m < r {
		t._update(o<<1|1, l, r, val)
	}
	t._pushUp(o)
}

func (t orST1114) _query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].or
	}
	t._spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		res = t._query(o<<1, l, r)
	}
	if m < r {
		res |= t._query(o<<1|1, l, r)
	}
	return
}

func (t orST1114) init(a []int64)             { t._build(a, 1, 1, len(a)) }
func (t orST1114) update(l, r int, val int64) { t._update(1, l, r, val) }
func (t orST1114) query(l, r int) int64       { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func CF1114F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod1114
			}
			x = x * x % mod1114
		}
		return res
	}
	div := func(p int64) int64 { return (p - 1) * pow(p, mod1114-2) % mod1114 }
	const mx = 300
	factors := [mx + 1]int64{}
	p1p := []int64{}
	for i := 2; i <= mx; i++ {
		if factors[i] == 0 {
			for j := i; j <= mx; j += i {
				factors[j] |= 1 << len(p1p)
			}
			p1p = append(p1p, div(int64(i)))
		}
	}

	var n, q, l, r, x int
	var s []byte
	Fscan(in, &n, &q)
	a := make([]int, n)
	ors := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		ors[i] = factors[a[i]]
	}

	mt := make(mulST1114, 4*n)
	mt.init(a)
	ot := make(orST1114, 4*n)
	ot.init(ors)
	for ; q > 0; q-- {
		Fscan(in, &s, &l, &r)
		if s[0] == 'M' {
			Fscan(in, &x)
			mt.update(l, r, x)
			ot.update(l, r, factors[x])
		} else {
			ans := mt.query(l, r)
			ps := ot.query(l, r)
			for i, v := range p1p {
				if ps>>i&1 == 1 {
					ans = ans * v % mod1114
				}
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { CF1114F(os.Stdin, os.Stdout) }
