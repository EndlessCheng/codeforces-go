package main

import "slices"

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 70_001

var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return res % mod
}

func totalBeauty(nums []int) (ans int) {
	m := slices.Max(nums)

	// 计算 b 的严格递增子序列的个数
	countIncreasingSubsequence := func(b []int, g int) (res int) {
		t := newFenwickTree(m / g)
		for _, x := range b {
			x /= g
			// cnt 表示以 x 结尾的严格递增子序列的个数
			cnt := t.pre(x-1) + 1 // +1 是因为 x 可以一个数组成一个子序列
			res += cnt
			t.update(x, cnt) // 更新以 x 结尾的严格递增子序列的个数
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
		f[i] = countIncreasingSubsequence(groups[i], i)
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
