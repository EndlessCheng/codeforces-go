package main

import (
	"math"
	"math/bits"
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
func concatenatedDivisibility(nums []int, k int) []int {
	slices.Sort(nums)
	n := len(nums)
	pow10 := make([]int, n)
	for i, x := range nums {
		pow10[i] = int(math.Pow10(len(strconv.Itoa(x))))
	}

	ans := make([]int, 0, n)
	vis := make([][]bool, 1<<n)
	for i := range vis {
		vis[i] = make([]bool, k)
	}
	var dfs func(int, int) bool
	dfs = func(s, x int) bool {
		if s == 0 {
			return x == 0
		}
		if vis[s][x] {
			return false
		}
		vis[s][x] = true
		// 枚举在 s 中的下标 i
		for t := uint(s); t > 0; t &= t - 1 {
			i := bits.TrailingZeros(t)
			if dfs(s^1<<i, (x*pow10[i]+nums[i])%k) {
				ans = append(ans, nums[i])
				return true
			}
		}
		return false
	}
	if !dfs(1<<n-1, 0) {
		return nil
	}
	slices.Reverse(ans) // nums[i] 是倒序加入答案的，所以要反转
	return ans
}
