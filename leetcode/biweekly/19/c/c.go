package main

import "math"

// github.com/EndlessCheng/codeforces-go
func angleClock(hour, minutes int) float64 {
	d := math.Abs((float64(hour%12)+float64(minutes)/60)*30 - float64(minutes*6))
	return math.Min(d, 360-d)
}
