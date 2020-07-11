package copypasta

import (
	. "fmt"
	"io"
)

// 一些题目：https://oi-wiki.org/math/matrix/

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

//

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
	const mod int64 = 1e9 + 7 // 998244353
	c := newMatrix(len(a), len(b[0]))
	for i := range a {
		for j := range b[0] {
			for k, aik := range a[i] {
				// 小心爆 int64，必要时用快速乘
				c[i][j] += aik * b[k][j] % mod
			}
		}
	}
	return c
}

// NxN 矩阵快速幂
// 注意并不会修改原矩阵的值
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

// 二阶递推数列第 n 项 - 矩阵快速幂
// a(n) = p*a(n-1) + q*a(n-2)
// 注意：数列从 0 开始，若题目从 1 开始则输入的 n 为 n-1
// https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97#%E7%B7%9A%E6%80%A7%E4%BB%A3%E6%95%B8%E8%A7%A3%E6%B3%95
// https://zhuanlan.zhihu.com/p/56444434
// 模板题 https://ac.nowcoder.com/acm/contest/6357/A
func calcFibonacci(p, q, a0, a1, n int64) int64 {
	const mod int64 = 1e9 + 7 // 998244353
	//n--
	if n == 0 {
		return (a0%mod + mod) % mod
	}
	if n == 1 {
		return (a1%mod + mod) % mod
	}
	m := matrix{
		{p, q},
		{1, 0},
	}.pow(n - 1)
	return ((m[0][0]*a1+m[0][1]*a0)%mod + mod) % mod
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

func (a matrix) trace() (sum int64) {
	for i, ai := range a {
		sum += ai[i]
	}
	return
}

// NxN 矩阵求逆
// 模板题 https://www.luogu.com.cn/problem/P4783
func (matrix) inv(in io.Reader, n int) matrix {
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
		f[i] = make([]int64, m)
		for j := range f {
			Fscan(in, &f[i][j])
		}
		f[i][n+i] = 1 // 单位矩阵
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

	// 结果保存在 f 右侧
	ans := make(matrix, n)
	for i, row := range f {
		ans[i] = row[n:]
	}
	return ans
}

// 行列式 高斯消元 Determinant
// TODO https://oi-wiki.org/math/gauss/
//      https://cp-algorithms.com/linear_algebra/determinant-gauss.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GaussianElimination.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GaussJordanElimination.java.html

// 线性基（子集异或和问题）
// todo

// 线性规划（单纯形算法）  linear programming (simplex)
// https://zh.wikipedia.org/zh-hans/%E5%8D%95%E7%BA%AF%E5%BD%A2%E6%B3%95
// https://oi-wiki.org/math/simplex/
// todo 算法第四版 https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/LinearProgramming.java.html
// https://zhuanlan.zhihu.com/p/31644892
// EXTRA: https://algs4.cs.princeton.edu/code/javadoc/edu/princeton/cs/algs4/TwoPersonZeroSumGame.html
