package main

// github.com/EndlessCheng/codeforces-go
func countGoodRectangles(rectangles [][]int) (ans int) {
	mx := 0
	for _, p := range rectangles {
		if v := min(p[0], p[1]); v > mx {
			mx, ans = v, 1
		} else if v == mx {
			ans++
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
