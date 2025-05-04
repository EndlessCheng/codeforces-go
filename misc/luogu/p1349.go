package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
var mod1349 int

type matrix1349 [][]int

func newMatrix1349(n, m int) matrix1349 {
	a := make(matrix1349, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix1349) mul(b matrix1349) matrix1349 {
	c := newMatrix1349(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod1349
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix1349) powMul(n int, f0 matrix1349) matrix1349 {
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
	Fscan(in, &p, &q, &a1, &a2, &n, &mod1349)
	if n == 1 {
		Fprint(out, a1%mod1349)
		return
	}
	m := matrix1349{
		{p, q},
		{1, 0},
	}
	f2 := matrix1349{
		{a2},
		{a1},
	}
	Fprint(out, m.powMul(n-2, f2)[0][0])
}

//func main() { p1349(bufio.NewReader(os.Stdin), os.Stdout) }
