package main

import "math/bits"

// https://space.bilibili.com/206214
func specialPerm(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	memo := make([][]int, u)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(s, i int) (res int) {
		if s == 0 {
			return 1 // 找到一个特别排列
		}
		p := &memo[s][i]
		if *p != -1 { // 之前计算过
			return *p
		}
		for j, x := range nums {
			if s>>j&1 > 0 && (nums[i]%x == 0 || x%nums[i] == 0) {
				res += dfs(s^(1<<j), j)
			}
		}
		*p = res // 记忆化
		return
	}
	for i := range nums {
		ans += dfs(u^(1<<i), i)
	}
	return ans % 1_000_000_007
}

func specialPerm3(nums []int) (ans int) {
	n := len(nums)
	d := make([]int, n)
	for i, x := range nums {
		for j, y := range nums {
			if j != i && (x%y == 0 || y%x == 0) {
				d[i] |= 1 << j
			}
		}
	}

	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for j := range f[0] {
		f[1<<j][j] = 1
	}
	for s, dr := range f {
		for _s := uint(s); _s > 0; _s &= _s - 1 {
			i := bits.TrailingZeros(_s)
			if dr[i] == 0 {
				continue
			}
			pre := nums[i]
			for cus, lb := (len(f)-1^s)&d[i], 0; cus > 0; cus ^= lb {
				lb = cus & -cus
				j := bits.TrailingZeros(uint(lb))
				cur := nums[j]
				if pre%cur == 0 || cur%pre == 0 {
					f[s|lb][j] += dr[i]
				}
			}
		}
	}
	for _, dv := range f[len(f)-1] {
		ans += dv
	}
	return ans % 1_000_000_007
}

func specialPerm2(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	f := make([][]int, u)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range nums {
		f[0][i] = 1
	}
	for s := 1; s < u; s++ {
		for i, pre := range nums {
			if s>>i&1 != 0 {
				continue
			}
			for j, x := range nums {
				if s>>j&1 != 0 && (pre%x == 0 || x%pre == 0) {
					f[s][i] += f[s^(1<<j)][j]
				}
			}
		}
	}
	for i := range nums {
		ans += f[u^(1<<i)][i]
	}
	return ans % 1_000_000_007
}
