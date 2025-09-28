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

// 返回 a*b
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

// 返回 a^n * f1
func (a matrix) powMul(n int, f1 matrix) matrix {
	res := f1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func zigZagArrays(n, l, r int) (ans int) {
	k := r - l + 1
	m := newMatrix(k, k)
	for i := range k {
		for j := range k - 1 - i {
			m[i][j] = 1
		}
	}

	f1 := newMatrix(k, 1)
	for i := range f1 {
		f1[i][0] = 1
	}

	fn := m.powMul(n-1, f1)
	for _, row := range fn {
		ans += row[0]
	}
	return ans * 2 % mod
}
