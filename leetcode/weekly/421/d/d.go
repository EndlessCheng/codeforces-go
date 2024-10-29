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

func lengthAfterTransformations(s string, t int, nums []int) (ans int) {
	const size = 26
	f0 := newMatrix(size, 1)
	for i := range f0 {
		f0[i][0] = 1
	}
	m := newMatrix(size, size)
	for i, c := range nums {
		for j := i + 1; j <= i+c; j++ {
			m[i][j%size] = 1
		}
	}
	m = m.powMul(t, f0)

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, row := range m {
		ans += row[0] * cnt[i]
	}
	return ans % mod
}
