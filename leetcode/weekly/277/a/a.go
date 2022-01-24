package main

/* O(n) 一次遍历

统计 $\textit{nums}$ 的最小值的个数和最大值的个数，答案即为 $\textit{nums}$ 的长度减去这两个个数。

注意当最小值和最大值相等时返回 0。

*/

// github.com/EndlessCheng/codeforces-go
func countElements(nums []int) int {
	min, max, cntMin, cntMax := nums[0], nums[0], 1, 1
	for _, v := range nums[1:] {
		if v < min {
			min, cntMin = v, 1
		} else if v == min {
			cntMin++
		}
		if v > max {
			max, cntMax = v, 1
		} else if v == max {
			cntMax++
		}
	}
	if min == max {
		return 0
	}
	return len(nums) - cntMin - cntMax
}
