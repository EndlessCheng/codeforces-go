package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
const mod17 = 1_000_000_007

type matrix17 [][]int

func newMatrix17(n, m int) matrix17 {
	a := make(matrix17, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix17) mul(b matrix17) matrix17 {
	c := newMatrix17(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod17
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix17) powMul(n int, f0 matrix17) matrix17 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf1117D(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
 
	m := newMatrix17(k, k)
	m[0][0] = 1
	m[0][k-1] = 1
	for i := range k - 1 {
		m[i+1][i] = 1
	}

	f0 := newMatrix17(k, 1)
	f0[0][0] = 1
 
	fn := m.powMul(n, f0)
	Fprint(out, fn[0][0])
}

//func main() { cf1117D(os.Stdin, os.Stdout) }
