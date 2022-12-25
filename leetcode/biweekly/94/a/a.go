package main

// https://space.bilibili.com/206214
func captureForts(forts []int) (ans int) {
	for i, x := range forts {
		if x != 1 {
			continue
		}
		j := i - 1
		for j >= 0 && forts[j] == 0 {
			j--
		}
		if j >= 0 && forts[j] < 0 {
			ans = max(ans, i-j-1)
		}
		j = i + 1
		for j < len(forts) && forts[j] == 0 {
			j++
		}
		if j < len(forts) && forts[j] < 0 {
			ans = max(ans, j-i-1)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }

