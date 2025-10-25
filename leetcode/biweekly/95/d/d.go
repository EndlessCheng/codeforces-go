package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	// 滑动窗口
	// 先计算 [0, r-1] 的发电量，为第一个窗口做准备
	sum := 0
	for _, x := range stations[:r] {
		sum += x
	}

	// 初始电量
	power := make([]int, n)
	mn := math.MaxInt
	for i := range power {
		// 右边进
		if i+r < n {
			sum += stations[i+r]
		}
		// 左边出
		if i-r-1 >= 0 {
			sum -= stations[i-r-1]
		}
		power[i] = sum
		mn = min(mn, sum)
	}

	// 二分答案
	left := mn + k/n
	right := mn + k
	ans := left + sort.Search(right-left, func(low int) bool {
		// 这里 +1 是为了二分最小的不满足要求的 low（符合库函数），这样最终返回的就是最大的满足要求的 low
		low += left + 1
		diff := make([]int, n+1) // 差分数组
		sumD, built := 0, 0
		for i, p := range power {
			sumD += diff[i] // 累加差分值
			m := low - (p + sumD)
			if m <= 0 {
				continue
			}
			// 需要在 i+r 额外建造 m 个供电站
			built += m
			if built > k { // 不满足要求
				return true
			}
			// 把区间 [i, i+2r] 加一
			sumD += m // 由于 diff[i] 后面不会再访问，我们直接加到 sumD 中
			diff[min(i+r*2+1, n)] -= m
		}
		return false
	})
	return int64(ans)
}
