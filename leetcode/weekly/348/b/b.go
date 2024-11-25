package main

// https://space.bilibili.com/206214
func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	var p, q int
	for i, v := range nums {
		if v == 1 {
			p = i
		} else if v == n {
			q = i
		}
	}
	if p < q {
		return p + n - 1 - q
	}
	return p + n - 2 - q // 1 向左移动的时候和 n 交换了一次
}
