package main

// https://space.bilibili.com/206214
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

// a^n * f
func (a matrix) powMul(n int64, f matrix) matrix {
	res := f
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func countGoodStrings(n int64) (ans int) {
	m := matrix{
		{1, 1},
		{1, 0},
	}
	f1 := matrix{{2}, {0}} // 这里初始化成 2，就不用把答案乘以 2 了
	fn := m.powMul(n-1, f1)
	return fn[0][0]
}
