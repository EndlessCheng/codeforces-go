package main

// https://space.bilibili.com/206214
func longestMonotonicSubarray1(a []int) int {
	ans := 1
	i, n := 0, len(a)
	for i < n-1 {
		if a[i+1] == a[i] {
			i++ // 直接跳过
			continue
		}
		start := i           // 记录这一组的开始位置
		inc := a[i] < a[i+1] // 定下基调：是严格递增还是严格递减
		i += 2               // i 和 i+1 已经满足要求，从 i+2 开始判断
		for i < n && a[i] != a[i-1] && a[i] > a[i-1] == inc {
			i++
		}
		// 从 start 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-start)
		i--
	}
	return ans
}

func longestMonotonicSubarray(a []int) int {
	ans, inc, dec := 1, 1, 1
	for i := 1; i < len(a); i++ {
		if a[i-1] < a[i] {
			inc++
		} else {
			inc = 1 // 从 a[i] 开始重新计数
		}

		if a[i-1] > a[i] {
			dec++
		} else {
			dec = 1 // 从 a[i] 开始重新计数
		}

		ans = max(ans, inc, dec)
	}
	return ans
}
