package main

import "slices"

// https://space.bilibili.com/206214
func maxProduct(nums []int, k, limit int) int {
	n := len(nums)
	mx := slices.Max(nums)
	if k >= 0 {
		if k > (n+1)/2*mx { // k 太大
			return -1
		}
	} else {
		if -k > n/2*mx { // k 太小（绝对值太大）
			return -1
		}
	}

	ans := -1
	type args struct {
		i, s, m    int
		odd, empty bool
	}
	vis := map[args]bool{}
	var dfs func(int, int, int, bool, bool)
	dfs = func(i, s, m int, odd, empty bool) {
		if m > limit || m < 0 {
			m = -1
		}

		if i == n {
			if !empty && s == k {
				ans = max(ans, m)
			}
			return
		}

		t := args{i, s, m, odd, empty}
		if vis[t] {
			return
		}
		vis[t] = true

		// 不选
		dfs(i+1, s, m, odd, empty)

		// 选
		x := nums[i]
		if odd {
			s -= x
		} else {
			s += x
		}
		dfs(i+1, s, m*x, !odd, false)
	}
	dfs(0, 0, 1, false, true)
	return ans
}
