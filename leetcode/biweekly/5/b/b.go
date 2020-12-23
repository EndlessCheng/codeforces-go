package main

import (
	"math"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func isArmstrong(n int) (ans bool) {
	s := strconv.Itoa(n)
	sum := 0.0
	for _, b := range s {
		sum += math.Pow(float64(b&15), float64(len(s)))
	}
	return int(math.Round(sum)) == n
}
