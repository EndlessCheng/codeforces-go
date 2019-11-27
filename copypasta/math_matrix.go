package copypasta

import (
	. "fmt"
	"io"
)

type matrix [][]int

func newMatrix(n, m int) matrix {
	mat := make(matrix, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}
	return mat
}

func newMatrixI(n int) matrix {
	mat := make(matrix, n)
	for i := range mat {
		mat[i] = make([]int, n)
		mat[i][i] = 1
	}
	return mat
}

func readMatrix(in io.Reader, n, m int) matrix {
	mat := make(matrix, n)
	for i := range mat {
		mat[i] = make([]int, m)
		for j := range mat[i] {
			Fscan(in, &mat[i][j])
		}
	}
	return mat
}

func copyMatrix(a matrix) matrix {
	mat := make(matrix, len(a))
	for i, ai := range a {
		mat[i] = make([]int, len(ai))
		copy(mat[i], ai)
	}
	return mat
}

func (a matrix) add(b matrix) matrix {
	c := newMatrix(len(a), len(a[0]))
	for i := range a {
		for j, aij := range a[i] {
			c[i][j] = aij + b[i][j] // % mod
		}
	}
	return c
}

func (a matrix) sub(b matrix) matrix {
	c := newMatrix(len(a), len(a[0]))
	for i := range a {
		for j, aij := range a[i] {
			c[i][j] = aij - b[i][j] // % mod) + mod) % mod
		}
	}
	return c
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i := range a {
		for j := range b[0] {
			for k, aik := range a[i] {
				c[i][j] += aik * b[k][j] // % mod
			}
		}
	}
	return c
}

// assert len(a) == len(a[0])
func (a matrix) pow(n int) matrix {
	res := newMatrixI(len(a))
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func (a matrix) trace() (sum int) {
	for i, ai := range a {
		sum += ai[i]
	}
	return
}
