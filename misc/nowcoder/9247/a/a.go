package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func newMatrixI(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, aik := range row {
				c[i][j] += aik * b[k][j] % mod
			}
			c[i][j] %= mod
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

func calcFibonacci(p, q, a0, a1, n int) int {
	if n == 0 {
		return a0
	}
	if n == 1 {
		return a1
	}
	m := matrix{{p, q}, {1, 0}}.pow(n - 1)
	return (m[0][0]*a1 + m[0][1]*a0) % mod
}

func Answerforcn(N int64) int {
	n := int(N) - 1
	a := calcFibonacci(2, 3, 2, 6, n)
	b := calcFibonacci(3, 10, 7, 35, n)
	return a * b % mod
}
