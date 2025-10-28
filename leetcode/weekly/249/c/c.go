package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func colorTheGrid1(m, n int) int {
	const mod = 1_000_000_007
	pow3 := make([]int, m)
	pow3[0] = 1
	for i := 1; i < m; i++ {
		pow3[i] = pow3[i-1] * 3
	}

	valid := []int{}
next:
	for color := range pow3[m-1] * 3 {
		for i := range m - 1 {
			if color/pow3[i+1]%3 == color/pow3[i]%3 { // 相邻颜色相同
				continue next
			}
		}
		valid = append(valid, color)
	}

	nv := len(valid)
	nxt := make([][]int, nv)
	for i, color1 := range valid {
	next2:
		for j, color2 := range valid {
			for _, p3 := range pow3 {
				if color1/p3%3 == color2/p3%3 { // 相邻颜色相同
					continue next2
				}
			}
			nxt[i] = append(nxt[i], j)
		}
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, nv)
	}
	for j := range f[0] {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := range f[i] {
			for _, k := range nxt[j] {
				f[i][j] += f[i-1][k]
			}
			f[i][j] %= mod
		}
	}

	ans := 0
	for _, fv := range f[n-1] {
		ans += fv
	}
	return ans % mod
}

const mod = 1_000_000_007

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
// 以及初始值 f(i) = a[i] (0 <= i < k)
// 返回 f(n) % mod，其中参数 n 从 0 开始
// 注意 coef 的顺序
// 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
func kitamasa(coef, a []int, n int) (ans int) {
	if n < len(a) {
		return a[n] % mod
	}

	k := len(coef)
	if k == 1 {
		return a[0] * pow(coef[0], n) % mod
	}

	// 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
	// 计算并返回 f(n+m) 的各项系数 c
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod
			}
			// 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
			// 倒序遍历，避免提前覆盖旧值
			bk1 := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk1*coef[i]) % mod
			}
			b[0] = bk1 * coef[0] % mod
		}
		return c
	}

	// 计算 resC，以表出 f(n) = resC[k-1] * a[k-1] + resC[k-2] * a[k-2] + ... + resC[0] * a[0]
	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = compose(c, resC)
		}
		// 由于会修改 compose 的第二个参数，这里把 c 复制一份再传入
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod
	}

	return (ans + mod) % mod // 保证结果非负
}

// 注意 kitamasa 入参的顺序
var coef = [][]int{
	{2},
	{3},
	{-2, 5},
	{6, -15, 9},
	{8, -48, 92, -65, 16},
}
var a = [][]int{
	{3},
	{6},
	{12, 54},
	{24, 162, 1122},
	{48, 486, 5118, 54450, 580986},
}

func colorTheGrid(m, n int) int {
	return kitamasa(coef[m-1], a[m-1], n-1)
}
