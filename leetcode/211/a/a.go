package main

// github.com/EndlessCheng/codeforces-go
func maxLengthBetweenEqualCharacters(s string) (ans int) {
	pos := [26][2]int{}
	for i, b := range s {
		i++
		b -= 'a'
		if pos[b][1] == 0 {
			pos[b] = [2]int{i, i}
		} else {
			pos[b][1] = i
		}
	}
	ans = -1
	for _, ps := range pos[:] {
		ans = max(ans, ps[1]-ps[0]-1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
