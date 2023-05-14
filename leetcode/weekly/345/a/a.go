package main

// https://space.bilibili.com/206214
func circularGameLosers(n int, k int) (ans []int) {
	vis := make([]bool, n)
	for i, d := 0, k; !vis[i]; d += k {
		vis[i] = true
		i = (i + d) % n
	}
	for i, b := range vis {
		if !b {
			ans = append(ans, i+1)
		}
	}
	return
}
