package main

import "math"

// https://space.bilibili.com/206214
func countTrapezoids(points [][]int) (ans int) {
	cnt := map[float64]map[float64]int{}
	type pair struct{ x, y int }
	cnt2 := map[pair]map[float64]int{}

	for i, p := range points {
		x, y := p[0], p[1]
		for _, q := range points[:i] {
			x2, y2 := q[0], q[1]
			dy := y - y2
			dx := x - x2
			k := math.MaxFloat64
			b := float64(x)
			if dx != 0 {
				k = float64(dy) / float64(dx)
				b = float64(y*dx-dy*x) / float64(dx)
			}

			if _, ok := cnt[k]; !ok {
				cnt[k] = map[float64]int{}
			}
			cnt[k][b]++ // 按照截距和斜率分组

			t := pair{x + x2, y + y2}
			if _, ok := cnt2[t]; !ok {
				cnt2[t] = map[float64]int{}
			}
			cnt2[t][k]++ // 按照中点和斜率分组
		}
	}

	for _, ct := range cnt {
		s := 0
		for _, c := range ct {
			ans += s * c
			s += c
		}
	}

	for _, ct := range cnt2 {
		s := 0
		for _, c := range ct {
			ans -= s * c
			s += c
		}
	}
	return
}
