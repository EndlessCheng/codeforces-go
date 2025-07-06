package main

import (
	"sort"
)

// https://space.bilibili.com/206214
func minStable1(nums []int, maxC int) int {
	n := len(nums)
	leftMin := make([]int, n)
	type interval struct{ gcd, l int } // 子数组 GCD，最小左端点
	intervals := []interval{{1, 0}}    // 哨兵
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

		// 由于我们添加了哨兵，intervals[1] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
		if len(intervals) > 1 {
			leftMin[i] = intervals[1].l
		} else {
			leftMin[i] = n
		}
	}

	ans := sort.Search(n/(maxC+1), func(upper int) bool {
		c := maxC
		i := upper
		for i < n {
			if i-leftMin[i]+1 > upper {
				if c == 0 {
					return false
				}
				c--
				i += upper + 1
			} else {
				i++
			}
		}
		return true
	})
	return ans
}

func minStable(nums []int, maxC int) int {
	n := len(nums)
	leftMin := make([]int, n)
	var left, bottom, rightGcd int
	for i, x := range nums {
		rightGcd = gcd(rightGcd, x)
		for left <= i && gcd(nums[left], rightGcd) == 1 {
			if bottom <= left {
				// 重新构建一个栈
				// 由于 left 即将移出窗口，只需计算到 left+1
				for j := i - 1; j > left; j-- {
					nums[j] = gcd(nums[j], nums[j+1])
				}
				bottom = i
				rightGcd = 0
			}
			left++
		}
		leftMin[i] = left
	}

	ans := sort.Search(n/(maxC+1), func(upper int) bool {
		c := maxC
		i := upper
		for i < n {
			if i-leftMin[i]+1 > upper {
				if c == 0 {
					return false
				}
				c--
				i += upper + 1
			} else {
				i++
			}
		}
		return true
	})
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
