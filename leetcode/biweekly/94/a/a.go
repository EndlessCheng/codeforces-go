package main

// https://space.bilibili.com/206214
func captureForts(forts []int) (ans int) {
	pre := -1 // 表示不存在
	for i, x := range forts {
		if x != 0 {
			if pre >= 0 && forts[i] != forts[pre] { // 一个是 1，另一个是 -1
				ans = max(ans, i-pre-1)
			}
			pre = i
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
