package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func findGCD1(nums []int) int {
	return gcd(slices.Min(nums), slices.Max(nums))
}

func findGCD(nums []int) int {
	n := len(nums)
	mn, mx := nums[0], nums[0]
	if n%2 == 0 {
		if nums[0] < nums[1] {
			mx = nums[1]
		} else {
			mn = nums[1]
		}
	}

	for i := 2 - n%2; i < n; i += 2 {
		x, y := nums[i], nums[i+1]
		if x > y {
			x, y = y, x
		}
		if x < mn {
			mn = x
		}
		if y > mx {
			mx = y
		}
	}

	return gcd(mn, mx)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
