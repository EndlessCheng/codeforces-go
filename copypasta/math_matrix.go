package copypasta

import (
	. "fmt"
	"io"
)

// 一些题目：https://oi-wiki.org/math/matrix/

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

func readMatrix(in io.Reader, n, m int) matrix {
	mat := make(matrix, n)
	for i := range mat {
		mat[i] = make([]int64, m)
		//mat[i] = make([]int64, m, m+1) // 方便高斯消元
		for j := range mat[i] {
			Fscan(in, &mat[i][j])
		}
	}
	return mat
}

func copyMatrix(a matrix) matrix {
	mat := make(matrix, len(a))
	for i, ai := range a {
		mat[i] = make([]int64, len(ai))
		copy(mat[i], ai)
	}
	return mat
}

func (a matrix) swapRows(i, j int) {
	for k := range a[0] {
		a[i][k], a[j][k] = a[j][k], a[i][k]
	}
}

func (a matrix) swapCols(i, j int) {
	for k := range a {
		a[k][i], a[k][j] = a[k][j], a[k][i]
	}
}

func (a matrix) mulRow(i int, k int64) {
	for j := range a[i] {
		a[i][j] *= k // % mod
	}
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
				// 小心爆 int64，必要时用快速乘
				c[i][j] += aik * b[k][j] // % mod
			}
		}
	}
	return c
}

// NxN 矩阵快速幂
// 模板题 https://www.luogu.com.cn/problem/P3390
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

func (a matrix) trace() (sum int64) {
	for i, ai := range a {
		sum += ai[i]
	}
	return
}

// NxN 矩阵求逆
// 模板题 https://www.luogu.com.cn/problem/P4783
func (matrix) inv(in io.Reader, out io.Writer, n int) matrix {
	const mod int64 = 1e9 + 7
	modInv := func(x int64) int64 {
		x %= mod
		res := int64(1)
		for n := mod - 2; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	m := 2 * n
	f := make(matrix, n)
	for i := range f {
		f[i] = make([]int64, 2*n)
		for j := range f {
			Fscan(in, &f[i][j])
		}
		f[i][i+n] = 1 // 单位矩阵
	}

	for i := range f {
		for j := i; j < n; j++ {
			if f[j][i] != 0 {
				// swapRows(i,j)
				for k := range f[0] {
					f[i][k], f[j][k] = f[j][k], f[i][k]
				}
				break
			}
		}
		if f[i][i] == 0 {
			// 矩阵不是满秩的
			return nil
		}
		inv := modInv(f[i][i])
		for j := i; j < m; j++ {
			f[i][j] = f[i][j] * inv % mod
		}
		for j := range f {
			if j != i {
				inv := f[j][i]
				for k := i; k < m; k++ {
					f[j][k] = (f[j][k] - inv*f[i][k]%mod + mod) % mod
				}
			}
		}
	}

	// 结果保存在右侧
	ans := make(matrix, n)
	for i, fi := range f {
		ans[i] = fi[n:]
	}
	return ans
}

// 高斯消元
// TODO https://oi-wiki.org/math/gauss/
