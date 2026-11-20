package main

import "slices"

// https://space.bilibili.com/206214
func maxSelectedElements1(nums []int) (ans int) {
	slices.Sort(nums)
	f := map[int]int{}
	for _, x := range nums {
		f[x+1] = f[x] + 1
		f[x] = f[x-1] + 1
		ans = max(ans, f[x], f[x+1])
	}
	return
}

func maxSelectedElements(nums []int) (ans int) {
	slices.Sort(nums)
	f0, f1, f2 := 0, 0, 0
	pre := 0 // pre 初始值随意，不影响计算
	for _, x := range nums {
		if x == pre+1 {
			// 类似 70. 爬楼梯 的空间优化
			f0 = f1
			f1 = f2
		} else if x == pre+2 {
			f0 = f2
			f1 = 0
		} else if x > pre+2 {
			f0 = 0
			f1 = 0
		}

		f2 = f1 + 1
		f1 = f0 + 1
		ans = max(ans, f1, f2)
		pre = x
	}
	return
}
