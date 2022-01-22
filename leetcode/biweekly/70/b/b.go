package main

// 还原差分数组

// github.com/EndlessCheng/codeforces-go
func numberOfArrays(differences []int, lower, upper int) int {
	num, min, max := 0, 0, 0
	for _, d := range differences {
		num += d // 根据差分数组还原原始数组的元素值 num
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	ans := upper - lower + 1 - (max - min)
	if ans > 0 {
		return ans
	}
	return 0
}
