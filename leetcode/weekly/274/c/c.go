package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, v := range asteroids {
		if mass < v {
			return false
		}
		mass += v
	}
	return true
}
