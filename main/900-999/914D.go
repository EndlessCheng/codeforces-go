package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg914 []struct{ l, r, g int }

func (seg914) gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (t seg914) _pushUp(o int) { t[o].g = t.gcd(t[o<<1].g, t[o<<1|1].g) }

func (t seg914) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].g = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg914) _update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].g = v
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, i, v)
	} else {
		t._update(o<<1|1, i, v)
	}
	t._pushUp(o)
}

func (t seg914) _query(o, l, r, v int) (c int8) {
	if t[o].g%v == 0 {
		return
	}
	if t[o].l == t[o].r {
		return 1
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		c += t._query(o<<1, l, r, v)
	}
	if c <= 1 && r > m {
		c += t._query(o<<1|1, l, r, v)
	}
	return
}

func (t seg914) init(a []int)           { t._build(a, 1, 1, len(a)) }
func (t seg914) update(i, v int)        { t._update(1, i, v) }
func (t seg914) query(l, r, v int) bool { return t._query(1, l, r, v) <= 1 }

func CF914D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg914, 4*n)
	t.init(a)
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			if Fscan(in, &v); t.query(l, r, v) {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		} else {
			t.update(l, r)
		}
	}
}

//func main() { CF914D(os.Stdin, os.Stdout) }
