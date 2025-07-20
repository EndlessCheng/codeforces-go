package main

import "math"

// https://space.bilibili.com/206214
func countTrapezoids(points [][]int) (ans int) {
	groups := map[float64][]float64{} // 斜率 -> 截距
	type pair struct{ x, y int }
	groups2 := map[pair][]float64{} // 中点 -> 斜率

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

			groups[k] = append(groups[k], b)
			mid := pair{x + x2, y + y2}
			groups2[mid] = append(groups2[mid], k)
		}
	}

	for _, g := range groups {
		if len(g) == 1 {
			continue
		}
		cnt := map[float64]int{}
		for _, b := range g {
			cnt[b]++
		}
		s := 0
		for _, c := range cnt {
			ans += s * c
			s += c
		}
	}

	for _, g := range groups2 {
		if len(g) == 1 {
			continue
		}
		cnt := map[float64]int{}
		for _, k := range g {
			cnt[k]++
		}
		s := 0
		for _, c := range cnt {
			ans -= s * c // 平行四边形会统计两次，减去多统计的一次
			s += c
		}
	}
	return
}
