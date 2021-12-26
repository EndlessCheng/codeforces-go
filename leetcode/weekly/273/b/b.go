package main

// github.com/EndlessCheng/codeforces-go
var dirs = [][2]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}

func executeInstructions(n int, startPos []int, s string) []int {
	ans := make([]int, len(s))
	for i := range s {
		ans[i] = len(s) - i
		x, y := startPos[0], startPos[1]
		for j, ch := range s[i:] {
			x += dirs[ch][0]
			y += dirs[ch][1]
			if x < 0 || x >= n || y < 0 || y >= n {
				ans[i] = j
				break
			}
		}
	}
	return ans
}
