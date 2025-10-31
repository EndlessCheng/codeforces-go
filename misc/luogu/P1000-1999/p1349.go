package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
var mod349 int

type matrix349 [][]int

func newMatrix349(n, m int) matrix349 {
	a := make(matrix349, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix349) mul(b matrix349) matrix349 {
	c := newMatrix349(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod349
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix349) powMul(n int, f0 matrix349) matrix349 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func p1349(in io.Reader, out io.Writer) {
	var p, q, a1, a2, n int
	Fscan(in, &p, &q, &a1, &a2, &n, &mod349)
	if n == 1 {
		Fprint(out, a1%mod349)
		return
	}
	m := matrix349{
		{p, q},
		{1, 0},
	}
	f2 := matrix349{
		{a2},
		{a1},
	}
	Fprint(out, m.powMul(n-2, f2)[0][0])
}

//func main() { p1349(bufio.NewReader(os.Stdin), os.Stdout) }
