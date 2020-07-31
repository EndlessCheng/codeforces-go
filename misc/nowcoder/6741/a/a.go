package main

const mod int64 = 1e9 + 7

type matrix [][]int64

func newMatrix(n, m int) matrix {
	mat := make(matrix, n)
	for i := range mat {
		mat[i] = make([]int64, m)
	}
	return mat
}

func newMatrixI(n int) matrix {
	mat := make(matrix, n)
	for i := range mat {
		mat[i] = make([]int64, n)
		mat[i][i] = 1
	}
	return mat
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i := range a {
		for j := range b[0] {
			for k, aik := range a[i] {
				c[i][j] += aik * b[k][j] % mod
			}
		}
	}
	return c
}

func (a matrix) pow(k int64) matrix {
	res := newMatrixI(len(a))
	for ; k > 0; k >>= 1 {
		if k&1 == 1 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

// github.com/EndlessCheng/codeforces-go
func nthElement(n int64, b int64, c int64) int64 {
	return matrix{{b, c}, {1, 0}}.pow(n - 1)[0][0] % mod
}
