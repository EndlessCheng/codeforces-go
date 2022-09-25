package main

// https://space.bilibili.com/206214
func temperatureTrend(a, b []int) (ans int) {
	start := 0
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] != (b[i] == b[i-1]) || a[i] < a[i-1] != (b[i] < b[i-1]) {
			ans = max(ans, i-start-1)
			start = i
		}
	}
	return max(ans, len(a)-start-1)
}

func max(a, b int) int { if b > a { return b }; return a }
