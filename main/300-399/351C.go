package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type matrix51 [][]int

func newMatrix51(n, m int) matrix51 {
	a := make(matrix51, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			a[i][j] = 1e18
		}
	}
	return a
}

func (a matrix51) mul(b matrix51) matrix51 {
	c := newMatrix51(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 1e18 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = min(c[i][j], x+y)
			}
		}
	}
	return c
}

func (a matrix51) powMul(n int, f0 matrix51) matrix51 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf351C(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l)
	}
	for i := range a {
		Fscan(in, &a[i].r)
	}

	mt := newMatrix51(n+1, n+1)
	for i, row := range mt {
		row[i] = 0
	}
	for _, p := range a {
		t := newMatrix51(n+1, n+1)
		for i := range n {
			t[i+1][i] = p.l
			t[i][i+1] = p.r
		}
		mt = t.mul(mt)
	}

	f0 := newMatrix51(n+1, 1)
	f0[0][0] = 0
	Fprint(out, mt.powMul(m, f0)[0][0])
}

//func main() { cf351C(os.Stdin, os.Stdout) }
