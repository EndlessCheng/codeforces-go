package main

import (
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func bestCoordinate(towers [][]int, radius int) (ans []int) {
	sort.Slice(towers, func(i, j int) bool { a, b := towers[i], towers[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
	maxS := -1
	for _, v := range towers {
		s := 0
		for _, w := range towers {
			if x, y := v[0]-w[0], v[1]-w[1]; x*x+y*y <= radius*radius {
				s += int(float64(w[2]) / (1 + math.Sqrt(float64(x*x+y*y))))
			}
		}
		if s > maxS {
			maxS = s
			ans = v[:2]
		}
	}
	return
}
