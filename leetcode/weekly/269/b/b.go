package main

/* O(n) 滑动窗口
*/

// github.com/EndlessCheng/codeforces-go
func getAverages(nums []int, k int) []int {
	avgs := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		if i < k || i+k >= len(nums) { // 超过边界
			avgs[i] = -1
		}
		sum += v
		if i >= k*2 {
			avgs[i-k] = sum / (k*2 + 1)
			sum -= nums[i-k*2]
		}
	}
	return avgs
}
