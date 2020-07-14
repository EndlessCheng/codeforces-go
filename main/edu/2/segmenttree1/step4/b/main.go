package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

var mod int

type mat [2][2]int
type seg []struct {
	l, r int
	mat  mat
}

func mul(a, b mat) (c mat) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
			c[i][j] %= mod
		}
	}
	return
}

func (t seg) _pushUp(o int) { t[o].mat = mul(t[o<<1].mat, t[o<<1|1].mat) }

func (t seg) _build(a []mat, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mat = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _query(o, l, r int) mat {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mat
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	return mul(t._query(o<<1, l, r), t._query(o<<1|1, l, r))
}

func (t seg) init(a []mat)       { t._build(a, 1, 1, len(a)) }
func (t seg) query(l, r int) mat { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	Fscan(in, &mod, &n, &q)
	a := make([]mat, n)
	for i := range a {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				Fscan(in, &a[i][j][k])
				//a[i][j][k] %= mod
			}
		}
	}
	t := make(seg, 4*n)
	t.init(a)
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		mat := t.query(l, r) // all values are non-neg
		Fprintln(out, mat[0][0], mat[0][1])
		Fprintln(out, mat[1][0], mat[1][1])
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
