package main

// https://space.bilibili.com/206214
func canMakeEqual(nums []int, k int) bool {
	check := func(target int) bool {
		left := k
		mul := 1
		for i, x := range nums {
			if x*mul == target {
				mul = 1 // 下一个数不用乘 -1
				continue
			}
			if left == 0 || i == len(nums)-1 {
				return false
			}
			left--
			mul = -1 // 下一个数要乘 -1
		}
		return true
	}
	return check(-1) || check(1)
}
