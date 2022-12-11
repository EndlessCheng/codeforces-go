package main

// https://space.bilibili.com/206214
func longestSquareStreak(nums []int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	for x := range set {
		cnt := 1
		for x *= x; set[x]; x *= x {
			cnt++
		}
		ans = max(ans, cnt)
	}
	if ans == 1 {
		return -1
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
