package main

// github.com/EndlessCheng/codeforces-go
func minKnightMoves(x int, y int) (ans int) {
	x, y = abs(x), abs(y)
	if x+y == 1 {
		return 3
	}
	ans = max(max((x+1)/2, (y+1)/2), (x+y+2)/3)
	ans += (ans ^ x ^ y) & 1
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
