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
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix {
	res := make(matrix, len(a))
	for i := range res {
		res[i] = make([]int, len(a))
		res[i][i] = 1
	}
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func lengthAfterTransformations(s string, t int, nums []int) (ans int) {
	const size = 26
	m := newMatrix(size, size)
	for i, c := range nums {
		for j := i + 1; j <= i+c; j++ {
			m[i][j%size] = 1
		}
	}
	m = m.pow(t)

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, row := range m {
		// m 第 i 行的和就是 f[t][i]
		fti := 0
		for _, x := range row {
			fti += x
		}
		ans += fti * cnt[i]
	}
	return ans % mod
}
