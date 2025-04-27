package main

import "slices"

// https://space.bilibili.com/206214
func calcPi(pattern string) []int {
	pi := make([]int, len(pattern))
	match := 0
	for i := 1; i < len(pi); i++ {
		b := pattern[i]
		for match > 0 && pattern[match] != b {
			match = pi[match-1]
		}
		if pattern[match] == b {
			match++
		}
		pi[i] = match
	}
	return pi
}

func kmpSearch(text []byte, pattern string, pi []int) []int {
	match := 0
	n := len(text)
	diff := make([]int, n+1)
	for i, b := range text {
		for match > 0 && pattern[match] != b {
			match = pi[match-1]
		}
		if pattern[match] == b {
			match++
		}
		if match == len(pi) {
			diff[i-len(pi)+1]++
			diff[i+1]--
			match = pi[match-1]
		}
	}
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff[:n]
}

func countCells(grid [][]byte, pattern string) (ans int) {
	hText := slices.Concat(grid...)
	m, n := len(grid), len(grid[0])
	vText := make([]byte, 0, m*n)
	for j := range n {
		for _, row := range grid {
			vText = append(vText, row[j])
		}
	}

	pi := calcPi(pattern)
	inPatternH := kmpSearch(hText, pattern, pi)
	inPatternV := kmpSearch(vText, pattern, pi)

	for i, x := range inPatternH {
		if x > 0 && inPatternV[i%n*m+i/n] > 0 {
			ans++
		}
	}
	return
}
