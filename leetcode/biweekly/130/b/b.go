package main

import (
	"math"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func maxPointsInsideSquare(points [][]int, s string) (ans int) {
	minD := [26]int{}
	for i := range minD {
		minD[i] = math.MaxInt
	}
	min2 := math.MaxInt
	for i, p := range points {
		x, y, c := p[0], p[1], s[i]-'a'
		d := max(abs(x), abs(y))
		if d < minD[c] {
			// d 是目前最小的，那么 minD[c] 是次小的
			min2 = min(min2, minD[c])
			minD[c] = d
		} else {
			// d 可能是次小的
			min2 = min(min2, d)
		}
	}
	for _, d := range minD {
		if d < min2 {
			ans++
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }

func maxPointsInsideSquare2(points [][]int, s string) (ans int) {
	sort.Search(1e9+1, func(size int) bool {
		vis := 0
		for i, p := range points {
			if abs(p[0]) <= size && abs(p[1]) <= size {
				c := s[i] - 'a'
				if vis>>c&1 > 0 {
					return true
				}
				vis |= 1 << c
			}
		}
		ans = bits.OnesCount(uint(vis))
		return false
	})
	return
}
