package main

// https://space.bilibili.com/206214/dynamic
func arrayChange(nums []int, operations [][]int) []int {
	mp := map[int]int{}
	for i := len(operations) - 1; i >= 0; i-- {
		p := operations[i]
		x, y := p[0], p[1]
		if mpY, ok := mp[y]; ok {
			y = mpY
		}
		mp[x] = y
	}
	for i, num := range nums {
		if m, ok := mp[num]; ok {
			nums[i] = m
		}
	}
	return nums
}
