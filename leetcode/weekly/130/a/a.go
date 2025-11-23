package main

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	x := 0
	for i, bit := range nums {
		x = (x<<1 | bit) % 5
		ans[i] = x == 0
	}
	return ans
}
