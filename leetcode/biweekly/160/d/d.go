package main

import "sort"

// https://space.bilibili.com/206214
func minStable(nums []int, maxC int) int {
	ans := sort.Search(len(nums)/(maxC+1), func(upper int) bool {
		type interval struct{ gcd, l int }
		intervals := []interval{}
		left := maxC
		for i, x := range nums {
			// 计算以 i 为右端点的子数组 GCD
			for j, p := range intervals {
				intervals[j].gcd = gcd(p.gcd, x)
			}
			// nums[i] 单独一个数作为子数组
			intervals = append(intervals, interval{x, i})

			// 去重（合并 GCD 相同的区间）
			idx := 1
			for j := 1; j < len(intervals); j++ {
				if intervals[j].gcd != intervals[j-1].gcd {
					intervals[idx] = intervals[j]
					idx++
				}
			}
			intervals = intervals[:idx]

			// intervals 的性质：越靠左，GCD 越小

			// 我们只关心 GCD >= 2 的子数组
			if intervals[0].gcd == 1 {
				intervals = intervals[1:]
			}

			// intervals[0] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
			if len(intervals) > 0 && i-intervals[0].l+1 > upper { // 必须修改 nums[i]=1
				if left == 0 {
					return false
				}
				left--
				intervals = intervals[:0] // 修改后 GCD 均为 1，直接清空
			}
		}
		return true
	})
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
