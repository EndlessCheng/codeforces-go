package main

import (
	"bufio"
	. "fmt"
	"io"
)

const m1380 = 998244353

var s1380 []byte

type seg1380 []struct {
	l, r int
	mat  [2][2]int64
}

func (t seg1380) set(o, i int) {
	t[o].mat = [2][2]int64{{int64(s1380[i] + 1), 1}}
	if s1380[i-1] == 1 {
		t[o].mat[1][0] = int64(9 - s1380[i])
	}
}

func (seg1380) mul(a, b [2][2]int64) (c [2][2]int64) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
			c[i][j] %= m1380
		}
	}
	return
}

func (t seg1380) maintain(o int) { t[o].mat = t.mul(t[o<<1].mat, t[o<<1|1].mat) }

func (t seg1380) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, l)
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg1380) update(o, i int) {
	if t[o].l == t[o].r {
		t.set(o, i)
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i)
	} else {
		t.update(o<<1|1, i)
	}
	t.maintain(o)
}

// github.com/EndlessCheng/codeforces-go
func CF1380F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, i int
	var d byte
	Fscan(in, &n, &q, &s1380)
	for i := range s1380 {
		s1380[i] &= 15
	}
	s1380 = append([]byte{0}, s1380...)
	t := make(seg1380, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &i, &d)
		s1380[i] = d
		if i < n {
			t.update(1, i+1)
		}
		t.update(1, i)
		Fprintln(out, t[1].mat[0][0])
	}
}

//func main() { CF1380F(os.Stdin, os.Stdout) }
