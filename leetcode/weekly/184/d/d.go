package main

func numOfWays1(n int) int {
	const mod = 1_000_000_007
	f := make([]int, n+1)
	f[0] = 3
	f[1] = 12
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1]*5 - f[i-2]*2) % mod // 注意这里有减法，结果可能是负数
	}
	return (f[n] + mod) % mod // 保证结果非负
}

const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

// 返回矩阵 a 和矩阵 b 相乘的结果
func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 0 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = (c[i][j] + x*y) % mod
			}
		}
	}
	return c
}

// a^n * f0
func (a matrix) powMul(n int, f0 matrix) matrix {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func numOfWays(n int) int {
	m := matrix{
		{5, -2},
		{1, 0},
	}
	f1 := matrix{{12}, {3}}
	fn := m.powMul(n-1, f1)
	return (fn[0][0] + mod) % mod // 保证结果非负
}
