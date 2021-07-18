package main

// github.com/EndlessCheng/codeforces-go
func addRungs(rungs []int, dist int) (ans int) {
	pre := 0
	for _, h := range rungs {
		if d := h - pre; d > dist {
			ans += (d - 1) / dist // 等价于 ceil(d / dist) - 1
		}
		pre = h
	}
	return
}
