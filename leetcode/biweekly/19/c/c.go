package main

import "math"

// github.com/EndlessCheng/codeforces-go
func angleClock1(hour, minutes int) float64 {
	d := math.Abs(float64(hour*30) - float64(minutes)*5.5)
	return min(d, 360-d)
}

func angleClock(hour, minutes int) float64 {
	d := abs(hour*60 - minutes*11)
	return float64(min(d, 720-d)) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
