package main

import "slices"

// 图一乐：记忆化搜索套记忆化搜索
func countOfPairs(a []int) (ans int) {
	const mod = 1_000_000_007
	n := len(a)
	m := slices.Max(a)
	dp0 := make([][]int, n)
	for i := range dp0 {
		dp0[i] = make([]int, m+1)
		for j := range dp0[i] {
			dp0[i][j] = -1
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var sum func(int, int) int
	var f func(int, int) int

	sum = func(i, j int) int {
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		if j == 0 {
			if i == 0 {
				*dv = 1
			} else {
				*dv = f(i, j)
			}
		} else {
			*dv = (sum(i, j-1) + f(i, j)) % mod
		}
		return *dv
	}

	f = func(i, j int) int {
		if i == 0 {
			return 1
		}
		dv := &dp0[i][j]
		if *dv != -1 {
			return *dv
		}
		maxK := min(j, a[i-1]-a[i]+j)
		if maxK < 0 {
			*dv = 0
		} else {
			*dv = sum(i-1, maxK)
		}
		return *dv
	}
	for j := 0; j <= a[n-1]; j++ {
		ans += f(n-1, j)
	}
	return ans % mod
}

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 3001 // maxN + maxM = 2000 + 1000 = 3000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func countOfPairsComb(nums []int) int {
	n := len(nums)
	m := nums[n-1]
	for i := 1; i < n; i++ {
		m -= max(nums[i]-nums[i-1], 0)
		if m < 0 {
			return 0
		}
	}
	return comb(m+n, n)
}

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

func countOfPairs1(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := nums[n-1]
	f := make([]int, m+1)
	for j := range f[:min(nums[0], m)+1] {
		f[j] = 1
	}
	for i := 1; i < n; i++ {
		j0 := max(nums[i]-nums[i-1], 0)
		m2 := min(nums[i], m)
		if j0 > m2 {
			return 0
		}
		for j := 1; j <= m2-j0; j++ {
			f[j] = (f[j] + f[j-1]) % mod // 计算前缀和
		}
		copy(f[j0:m2+1], f)
		clear(f[:j0])
	}
	for _, v := range f {
		ans += v
	}
	return ans % mod
}

func countOfPairs2(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := nums[n-1]
	f := make([]int, m+1)
	for j := range f[:min(nums[0], m)+1] {
		f[j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		j0 := max(nums[i]-nums[i-1], 0)
		if j0 > m {
			return 0
		}
		copy(f[j0:], f[:min(nums[i], m)+1])
		clear(f[:min(j0, m+1)])
	}
	for _, v := range f {
		ans += v
	}
	return ans % mod
}

func countOfPairs3(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	s := make([]int, m+1)

	for j := 0; j <= nums[0]; j++ {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		s[0] = f[i-1][0]
		for k := 1; k <= m; k++ {
			s[k] = s[k-1] + f[i-1][k] // f[i-1] 的前缀和
		}
		for j := 0; j <= nums[i]; j++ {
			maxK := j + min(nums[i-1]-nums[i], 0)
			if maxK >= 0 {
				f[i][j] = s[maxK] % mod
			}
		}
	}

	for _, v := range f[n-1][:nums[n-1]+1] {
		ans += v
	}
	return ans % mod
}
