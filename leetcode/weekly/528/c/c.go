package main

import "math"

// https://space.bilibili.com/206214
func maxDifference1(nums []int) int64 {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = math.MinInt
	}

	// 返回在剩余元素为 nums[i:] 的情况下，(当前玩家得分 - 对手得分) 的最大值
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == n {
			return
		}

		p := &memo[i]
		if *p != math.MinInt {
			return *p
		}

		x := nums[i]
		if x%3 == 0 {
			res = x - dfs(i+1) // 下一回合是对手
		} else {
			res = x + dfs(i+1) // 下一回合还是我
		}

		if i < n-1 {
			s := x + nums[i+1]
			if s%3 == 0 {
				s -= dfs(i + 2) // 下一回合是对手
			} else {
				s += dfs(i + 2) // 下一回合还是我
			}
			res = max(res, s)
		}

		*p = res
		return
	}

	return int64(dfs(0))
}

func maxDifference2(nums []int) int64 {
	n := len(nums)
	f := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		if x%3 == 0 {
			f[i] = x - f[i+1] // 下一回合是对手
		} else {
			f[i] = x + f[i+1] // 下一回合还是我
		}

		if i < n-1 {
			s := x + nums[i+1]
			if s%3 == 0 {
				s -= f[i+2] // 下一回合是对手
			} else {
				s += f[i+2] // 下一回合还是我
			}
			f[i] = max(f[i], s)
		}
	}

	return int64(f[0])
}

func maxDifference(nums []int) int64 {
	n := len(nums)
	f1, f2 := nums[n-1], 0

	for i := n - 2; i >= 0; i-- {
		x := nums[i]
		f := 0
		if x%3 == 0 {
			f = x - f1 // 下一回合是对手
		} else {
			f = x + f1 // 下一回合还是我
		}

		s := x + nums[i+1]
		if s%3 == 0 {
			s -= f2 // 下一回合是对手
		} else {
			s += f2 // 下一回合还是我
		}
		f = max(f, s)

		f2 = f1
		f1 = f
	}

	return int64(f1)
}
