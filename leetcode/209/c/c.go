package main

import (
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func visiblePoints(points [][]int, angle int, location []int) (ans int) {
	arc := make([]float64, 0, 2*len(points))
	for _, p := range points {
		p[0] -= location[0]
		p[1] -= location[1]
		if p[0] == 0 && p[1] == 0 {
			ans++
		} else {
			arc = append(arc, math.Atan2(float64(p[1]), float64(p[0])))
		}
	}
	const eps = 1e-10
	sort.Float64s(arc)
	n := len(arc)
	for _, a := range arc {
		arc = append(arc, a+2*math.Pi)
	}
	cnt := 0
	ang := math.Pi / 180 * float64(angle)
	for i, a := range arc[:n] {
		if j := sort.SearchFloat64s(arc, a+ang+eps); j-i > cnt {
			cnt = j - i
		}
	}
	ans += cnt
	return
}
