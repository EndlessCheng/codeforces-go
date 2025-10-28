package main

import (
	"slices"
)

// https://space.bilibili.com/206214
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

func lengthAfterTransformations1(s string, t int, nums []int) (ans int) {
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

//

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

// 给定数列的前 m 项 a，返回符合 a 的最短常系数齐次线性递推式的系数 coef（模 mod 意义下）
// 设 coef 长为 k，当 n >= k 时，有递推式 f(n) = coef[0] * f(n-1) + coef[1] * f(n-2) + ... + coef[k-1] * f(n-k)  （注意 coef 的顺序）
// 初始值 f(n) = a[n]  (0 <= n < k)
// 时间复杂度 O(m^2)，其中 m 是 a 的长度
func berlekampMassey(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0

	for i, v := range a {
		// d = a[i] - 递推式算出来的值
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod
		}
		if d == 0 { // 递推式正确
			continue
		}

		// 首次算错，初始化 coef 为 i+1 个 0
		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - preI
		oldLen := len(coef)
		newLen := bias + len(preC)
		var tmp []int
		if newLen > oldLen { // 递推式变长了
			tmp = slices.Clone(coef)
			coef = slices.Grow(coef, newLen-oldLen)[:newLen] // coef.resize(newLen)
		}

		// 历史错误为 preD = a[preI] - sum_j preC[j]*a[preI-1-j]
		// 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
		// 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
		// 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias-1 = i-preI-1 处
		delta := d * pow(preD, mod-2) % mod // pow(preD, mod-2) 为 preD 的逆元
		coef[bias-1] = (coef[bias-1] + delta) % mod
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod
		}

		if newLen > oldLen {
			preC = tmp
			preI, preD = i, d
		}
	}

	// 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
	// 比如数列 (1,2,4,2,4,2,4,...) 的系数为 [0,1,0]，表示 f(n) = 0*f(n-1) + 1*f(n-2) + 0*f(n-3) = f(n-2)   (n >= 3)
	// 如果把末尾的 0 去掉，变成 [0,1]，就表示 f(n) = 0*f(n-1) + f(n-2) = f(n-2)   (n >= 2)
	// 看上去一样，但按照这个式子算出来的数列是错误的 (1,2,1,2,1,2,...)

	// 手动找规律用
	for i, c := range coef {
		if c < -mod/2 {
			c += mod
		} else if c > mod/2 {
			c -= mod
		}
		coef[i] = c
	}

	return
}

// 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
// 以及初始值 f(i) = a[i] (0 <= i < k)
// 返回 f(n) % mod，其中参数 n 从 0 开始
// 注意 coef 的顺序
// 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
func kitamasa(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans + mod) % mod }() // 保证结果非负
	if n < len(a) {
		return a[n] % mod
	}

	k := len(coef)
	// 特判 k = 0, 1 的情况
	if k == 0 {
		return 0
	}
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

	return
}

// 计算 t = 1,2,...,52 的答案
func lengthAfterTransformationsInit(nums, cnt []int) []int {
	const k = 26
	f := [k]int{}
	for i := range f {
		f[i] = 1
	}
	sum := [k*2 + 1]int{}

	a := make([]int, k*2)
	for i := range a {
		// 计算 f + f 的前缀和
		for j := range k * 2 {
			sum[j+1] = sum[j] + f[j%k]
		}
		s := 0
		for j, c := range nums {
			f[j] = (sum[j+c+1] - sum[j+1]) % mod
			s += f[j] * cnt[j]
		}
		a[i] = s % mod
	}
	return a
}

func lengthAfterTransformations(s string, t int, nums []int) int {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	a := lengthAfterTransformationsInit(nums, cnt[:])
	coef := berlekampMassey(a)
	slices.Reverse(coef) // 注意 kitamasa 入参的顺序
	return kitamasa(coef, a, t-1)
}
