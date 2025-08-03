package main

import "math"

// https://space.bilibili.com/206214
func maxSumTrionic(nums []int) int64 {
	const negInf = math.MinInt / 2 // 除 2 防止下面加法（和负数相加）溢出
	ans, f1, f2, f3 := negInf, negInf, negInf, negInf
	for i := 1; i < len(nums); i++ {
		x, y := nums[i-1], nums[i]
		if x < y { // 第一段或者第三段
			f3 = max(f3, f2) + y
			ans = max(ans, f3)
			f2 = negInf
			f1 = max(f1, x) + y
		} else if x > y { // 第二段
			f2 = max(f2, f1) + y
			f1, f3 = negInf, negInf
		} else {
			f1, f2, f3 = negInf, negInf, negInf
		}
	}
	return int64(ans)
}

func maxSumTrionic1(nums []int) int64 {
	n := len(nums)
	ans := math.MinInt
	for i := 0; i < n; {
		// 第一段
		start := i
		for i++; i < n && nums[i-1] < nums[i]; i++ {
		}
		if i == start+1 { // 第一段至少要有两个数
			continue
		}

		// 第二段
		peak := i - 1
		res := nums[peak-1] + nums[peak] // 第一段的最后两个数必选
		for ; i < n && nums[i-1] > nums[i]; i++ {
			res += nums[i] // 第二段的所有元素必选
		}
		if i == peak+1 || i == n { // 第二段至少要有两个数，第三段至少要有两个数
			continue
		}

		// 第三段
		bottom := i - 1
		res += nums[i] // 第三段的前两个数必选（第一个数在上面的循环中加了）
		// 从第三段的第三个数往右，计算最大元素和
		maxS, s := 0, 0
		for i++; i < n && nums[i-1] < nums[i]; i++ {
			s += nums[i]
			maxS = max(maxS, s)
		}
		res += maxS

		// 从第一段的倒数第三个数往左，计算最大元素和
		maxS, s = 0, 0
		for j := peak - 2; j >= start; j-- {
			s += nums[j]
			maxS = max(maxS, s)
		}
		res += maxS
		ans = max(ans, res)

		i = bottom // 第三段的起点也是下一个极大三段式子数组的第一段的起点
	}
	return int64(ans)
}
