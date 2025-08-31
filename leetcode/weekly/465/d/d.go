package main

import "slices"

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 70_001

var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func totalBeauty(nums []int) (ans int) {
	m := slices.Max(nums)

	// 树状数组（时间戳优化）
	tree := make([]int, m+1)
	time := make([]int, m+1) // 避免反复初始化树状数组
	now := 0
	update := func(i, val int) {
		for ; i <= m; i += i & -i {
			if time[i] < now {
				time[i] = now
				tree[i] = 0 // 懒初始化
			}
			tree[i] = (tree[i] + val) % mod
		}
	}
	pre := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			if time[i] == now {
				res += tree[i]
			}
		}
		return res % mod
	}

	// 计算 b 的严格递增子序列的个数
	cntIS := func(b []int) (res int) {
		now++
		for _, x := range b {
			// cnt 表示以 x 结尾的严格递增子序列的个数
			cnt := pre(x-1) + 1 // +1 是因为 x 可以一个数组成一个子序列
			res += cnt
			update(x, cnt) // 更新以 x 结尾的严格递增子序列的个数
		}
		return res % mod
	}

	groups := make([][]int, m+1)
	for _, x := range nums {
		for _, d := range divisors[x] {
			groups[d] = append(groups[d], x)
		}
	}

	f := make([]int, m+1)
	for i := m; i > 0; i-- {
		f[i] = cntIS(groups[i])
		// 倍数容斥
		for j := i * 2; j <= m; j += i {
			f[i] -= f[j]
		}
		// 注意 |f[i]| * i < mod * (m / i) * i = mod * m
		// m 个 mod * m 相加，至多为 mod * m * m，不会超过 64 位整数最大值
		ans += f[i] * i
	}
	// 保证结果非负
	return (ans%mod + mod) % mod
}
