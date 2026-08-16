package main

// https://space.bilibili.com/206214
func elevatorRequests(n int, requests []int) (ans int) {
	pre := 0
	for _, req := range requests {
		ans += abs(req - pre)
		pre = req
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
