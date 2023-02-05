package main

// https://space.bilibili.com/206214
func maxCount(banned []int, n, maxSum int) (ans int) {
	has := map[int]bool{}
	for _, v := range banned {
		has[v] = true
	}
	for i := 1; i <= n && i <= maxSum; i++ {
		if !has[i] {
			maxSum -= i
			ans++
		}
	}
	return
}
