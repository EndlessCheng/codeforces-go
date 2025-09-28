package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type mat50 [5][5]int32

type seg50 []struct {
	l, r int
	mat50
}

func (seg50) merge(a, b mat50) mat50 {
	c := mat50{}
	for i := range 5 {
		for j := i; j < 5; j++ {
			c[i][j] = 1e9
		}
	}
	for i := range 5 {
		for k := i; k < 5; k++ {
			if a[i][k] == 1e9 {
				continue
			}
			for j := k; j < 5; j++ {
				c[i][j] = min(c[i][j], a[i][k]+b[k][j])
			}
		}
	}
	return c
}

func (t seg50) maintain(o int) {
	t[o].mat50 = t.merge(t[o<<1].mat50, t[o<<1|1].mat50)
}

func (t seg50) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		m := mat50{}
		for i := range 5 {
			for j := i + 1; j < 5; j++ {
				m[i][j] = 1e9
			}
		}
		switch s[l] {
		case '2':
			m[0][0] = 1
			m[0][1] = 0
		case '0':
			m[1][1] = 1
			m[1][2] = 0
		case '1':
			m[2][2] = 1
			m[2][3] = 0
		case '7':
			m[3][3] = 1
			m[3][4] = 0
		case '6':
			m[3][3] = 1
			m[4][4] = 1
		}
		t[o].mat50 = m
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg50) query(o, l, r int) mat50 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mat50
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.merge(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf750E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, l, r int
	var s string
	Fscan(in, &n, &q, &s)
	t := make(seg50, 2<<bits.Len(uint(n-1)))
	t.build(s, 1, 0, n-1)
	for range q {
		Fscan(in, &l, &r)
		ans := t.query(1, l-1, r-1)[0][4]
		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf750E(bufio.NewReader(os.Stdin), os.Stdout) }
