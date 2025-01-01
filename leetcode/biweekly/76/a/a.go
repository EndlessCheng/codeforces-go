package main

// github.com/EndlessCheng/codeforces-go
func findClosestNumber(nums []int) int {
	ans := nums[0]
	for _, x := range nums {
		if abs(x) < abs(ans) || abs(x) == abs(ans) && x > 0 {
			ans = x
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
