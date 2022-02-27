package main

import "math"

// github.com/EndlessCheng/codeforces-go
func minimumFinishTime(tires [][]int, changeTime, numLaps int) int {
	minSec := [18]int{}
	for i := range minSec {
		minSec[i] = math.MaxInt32 / 2
	}
	for _, tire := range tires {
		f, r := tire[0], tire[1]
		for x, time, sum := 1, f, 0; time <= changeTime+f; x++ {
			sum += time
			minSec[x] = min(minSec[x], sum)
			time *= r
		}
	}

	f := make([]int, numLaps+1)
	f[0] = -changeTime
	for i := 1; i <= numLaps; i++ {
		f[i] = math.MaxInt32
		for j := 1; j <= 17 && j <= i; j++ {
			f[i] = min(f[i], f[i-j]+minSec[j])
		}
		f[i] += changeTime
	}
	return f[numLaps]
}

func min(a, b int) int { if a > b { return b }; return a }
