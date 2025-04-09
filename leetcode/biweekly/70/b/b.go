package main

// github.com/EndlessCheng/codeforces-go
func numberOfArrays(differences []int, lower, upper int) int {
	var s, minS, maxS int // s[0] = 0
	for _, d := range differences {
		s += d
		minS = min(minS, s)
		maxS = max(maxS, s)
	}
	return max(upper-lower-maxS+minS+1, 0)
}
