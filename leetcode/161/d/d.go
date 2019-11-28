package main

func isGoodArray(nums []int) bool {
	if len(nums) == 1 {
		return nums[0] == 1
	}
	calcGCD := func(a, b int) int {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	val := nums[0]
	for _, v := range nums {
		val = calcGCD(val, v)
	}
	return val == 1
}
