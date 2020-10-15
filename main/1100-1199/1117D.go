package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type matrix17 [][]int64

func newMatrix17(n int) matrix17 {
	mat := make(matrix17, n)
	for i := range mat {
		mat[i] = make([]int64, n)
	}
	return mat
}

func newMatrixI17(n int) matrix17 {
	mat := make(matrix17, n)
	for i := range mat {
		mat[i] = make([]int64, n)
		mat[i][i] = 1
	}
	return mat
}

func (a matrix17) mul(b matrix17) matrix17 {
	const mod int64 = 1e9 + 7
	c := newMatrix17(len(a))
	for i := range a {
		for j := range b[0] {
			for k, aik := range a[i] {
				c[i][j] += aik * b[k][j] % mod
			}
			c[i][j] %= mod
		}
	}
	return c
}

func (a matrix17) pow(k int64) matrix17 {
	res := newMatrixI17(len(a))
	for ; k > 0; k >>= 1 {
		if k&1 == 1 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func CF1117D(in io.Reader, out io.Writer) {
	n, m := int64(0), 0
	Fscan(in, &n, &m)
	if n < int64(m) {
		Fprint(out, 1)
		return
	}
	a := newMatrix17(m)
	a[0][0] = 1
	a[0][m-1] = 1
	for i := 1; i < m; i++ {
		a[i][i-1] = 1
	}
	Fprint(out, a.pow(n)[0][0])
}

//func main() { CF1117D(os.Stdin, os.Stdout) }
