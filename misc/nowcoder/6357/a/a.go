package main

const mod int = 1e9 + 7

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

func (a matrix) pow(k int) matrix {
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
func solve(a0 int, a1 int, n int) int {
	n--
	if n == 0 {
		return (a0%mod + mod) % mod
	}
	if n == 1 {
		return (a1%mod + mod) % mod
	}
	m := matrix{
		{1, -1},
		{1, 0},
	}.pow(n - 1)
	return (m[0][0]*a1 + m[0][1]*a0%mod + mod) % mod
}
