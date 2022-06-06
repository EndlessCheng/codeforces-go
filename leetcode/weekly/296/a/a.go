package main

// https://space.bilibili.com/206214/dynamic
func minMaxGame(a []int) int {
	for len(a) > 1 {
		b := make([]int, len(a)/2)
		for i := len(a)/2-1; i >= 0; i-- {
			if i&1 == 0 {
				a[i] = min(a[i*2], a[i*2+1])
			} else {
				b[i] = max(a[i*2], a[i*2+1])
			}
		}
		for i := range b {
			if i&1 == 0 {
				b[i] = min(a[i*2], a[i*2+1])
			} else {
				b[i] = max(a[i*2], a[i*2+1])
			}
		}
		a = b
	}
	return a[0]
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
