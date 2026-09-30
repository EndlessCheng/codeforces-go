package main

// github.com/EndlessCheng/codeforces-go
func targetIndices(nums []int, target int) []int {
	less, equal := 0, 0
	for _, x := range nums {
		if x < target {
			less++
		} else if x == target {
			equal++
		}
	}

	ans := make([]int, equal)
	for i := range ans {
		ans[i] = less + i
	}
	return ans
}
