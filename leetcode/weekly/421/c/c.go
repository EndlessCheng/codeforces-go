package main

import "slices"

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 201

var lcms [mx][mx]int
var pow2, pow3, mu [mx]int

func init() {
	for i := 1; i < mx; i++ {
		for j := 1; j < mx; j++ {
			lcms[i][j] = lcm(i, j)
		}
	}

	pow2[0], pow3[0] = 1, 1
	for i := 1; i < mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
		pow3[i] = pow3[i-1] * 3 % mod
	}

	mu[1] = 1
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			mu[j] -= mu[i]
		}
	}
}

func subsequencePairCount(nums []int) int {
	m := slices.Max(nums)
	// cnt[i] 表示 nums 中的 i 的倍数的个数
	cnt := make([]int, m+1)
	for _, x := range nums {
		cnt[x]++
	}
	for i := 1; i <= m; i++ {
		for j := i * 2; j <= m; j += i {
			cnt[i] += cnt[j] // 统计 i 的倍数的个数
		}
	}

	f := make([][]int, m+1)
	for g1 := 1; g1 <= m; g1++ {
		f[g1] = make([]int, m+1)
		for g2 := 1; g2 <= m; g2++ {
			l := lcms[g1][g2]
			c := 0
			if l <= m {
				c = cnt[l]
			}
			c1, c2 := cnt[g1], cnt[g2]
			f[g1][g2] = (pow3[c]*pow2[c1+c2-c*2] - pow2[c1] - pow2[c2] + 1) % mod
		}
	}

	// 倍数容斥
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= m/i; j++ {
			for k := 1; k <= m/i; k++ {
				ans += mu[j] * mu[k] * f[j*i][k*i]
			}
		}
	}
	return (ans%mod + mod) % mod // 保证 ans 非负
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func subsequencePairCount2(nums []int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, m+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 {
			if j == k {
				return 1
			}
			return 0
		}
		p := &memo[i][j][k]
		if *p < 0 {
			*p = (dfs(i-1, j, k) + dfs(i-1, gcd(j, nums[i]), k) + dfs(i-1, j, gcd(k, nums[i]))) % mod
		}
		return *p
	}
	// 减去两个子序列都是空的情况
	return (dfs(n-1, 0, 0) - 1 + mod) % mod
}
