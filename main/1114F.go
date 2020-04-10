package main

import (
	"bufio"
	. "fmt"
	"io"
)

const mod1114 int64 = 1e9 + 7

type lazyST1114 []struct {
	l, r         int
	mul, mulTodo int64
	or, orTodo   int64
}

func (lazyST1114) pow(x int64, n int) int64 {
	res := int64(1)
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod1114
		}
		x = x * x % mod1114
	}
	return res
}

func (t lazyST1114) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].mul = lo.mul * ro.mul % mod1114
	t[o].or = lo.or | ro.or
}

func (t lazyST1114) _build(a []int, ors []int64, o, l, r int) {
	t[o].l, t[o].r, t[o].mulTodo = l, r, 1
	if l == r {
		t[o].mul = int64(a[l-1])
		t[o].or = ors[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, ors, o<<1, l, m)
	t._build(a, ors, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t lazyST1114) _spread(o int) {
	lo, ro := &t[o<<1], &t[o<<1|1]
	if mul := t[o].mulTodo; mul != 1 {
		lo.mul = lo.mul * t.pow(mul, lo.r-lo.l+1) % mod1114
		ro.mul = ro.mul * t.pow(mul, ro.r-ro.l+1) % mod1114
		lo.mulTodo = lo.mulTodo * mul % mod1114
		ro.mulTodo = ro.mulTodo * mul % mod1114
		t[o].mulTodo = 1
	}
	if or := t[o].orTodo; or != 0 {
		lo.or |= or
		ro.or |= or
		lo.orTodo |= or
		ro.orTodo |= or
		t[o].orTodo = 0
	}
}

func (t lazyST1114) _update(o, l, r, mul int, orVal int64) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].mul = t[o].mul * t.pow(int64(mul), or-ol+1) % mod1114
		t[o].mulTodo = t[o].mulTodo * int64(mul) % mod1114
		t[o].or |= orVal
		t[o].orTodo |= orVal
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._update(o<<1, l, r, mul, orVal)
	}
	if m < r {
		t._update(o<<1|1, l, r, mul, orVal)
	}
	t._pushUp(o)
}

func (t lazyST1114) _query(o, l, r int) (mul, or int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mul, t[o].or
	}
	t._spread(o)
	mul = 1
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		mul, or = t._query(o<<1, l, r)
	}
	if m < r {
		a, b := t._query(o<<1|1, l, r)
		mul = mul * a % mod1114
		or |= b
	}
	return
}

func (t lazyST1114) init(a []int, ors []int64)      { t._build(a, ors, 1, 1, len(a)) }
func (t lazyST1114) update(l, r, mul int, or int64) { t._update(1, l, r, mul, or) }
func (t lazyST1114) query(l, r int) (int64, int64)  { return t._query(1, l, r) }

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

	t := make(lazyST1114, 4*n)
	t.init(a, ors)
	for ; q > 0; q-- {
		Fscan(in, &s, &l, &r)
		if s[0] == 'M' {
			Fscan(in, &x)
			t.update(l, r, x, factors[x])
		} else {
			ans, ps := t.query(l, r)
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
