package main

// 简洁写法 + 一次遍历

// github.com/EndlessCheng/codeforces-go
func rearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))
	i, j := 0, 1
	for _, v := range nums {
		if v > 0 {
			ans[i] = v
			i += 2
		} else {
			ans[j] = v
			j += 2
		}
	}
	return ans
}
