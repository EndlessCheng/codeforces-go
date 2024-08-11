package main

import "slices"

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

func countOfPairs(nums []int) int {
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
