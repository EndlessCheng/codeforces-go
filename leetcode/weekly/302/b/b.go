package main

// https://space.bilibili.com/206214/dynamic
func maximumSum(nums []int) int {
	ans := -1
	mx := map[int]int{}
	for _, v := range nums {
		s := 0
		for x := v; x > 0; x /= 10 {
			s += x % 10
		}
		if mx[s] > 0 {
			ans = max(ans, mx[s] + v)
		}
		mx[s] = max(mx[s], v)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
