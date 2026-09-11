package main

import "strings"

func maxScore(s string) (ans int) {
	score := strings.Count(s, "1")
	for _, c := range s[:len(s)-1] {
		if c == '0' {
			score++
		} else {
			score--
		}
		ans = max(ans, score)
	}
	return
}

func maxScore2(s string) int {
	n := len(s)
	total1, delta := 0, 0
	maxDelta := -1
	for _, c := range s[:n-1] {
		if c == '0' {
			delta++
		} else {
			total1++
			delta--
		}
		maxDelta = max(maxDelta, delta)
	}
	total1 += int(s[n-1] - '0')
	return total1 + maxDelta
}
