package main

// https://space.bilibili.com/206214
func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, x := range nums {
		if x%6 == 0 {
			sum += x
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
