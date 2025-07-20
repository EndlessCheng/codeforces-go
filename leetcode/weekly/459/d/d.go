package main

import "math"

// https://space.bilibili.com/206214
func countTrapezoids(points [][]int) (ans int) {
	cnt := map[float32]map[float32]int{}
	type pair struct{ x, y int }
	cnt2 := map[pair]map[float32]int{}

	for i, p := range points {
		x, y := p[0], p[1]
		for _, q := range points[:i] {
			x2, y2 := q[0], q[1]
			dy := y - y2
			dx := x - x2
			k := float32(math.MaxFloat32)
			b := float32(x)
			if dx != 0 {
				k = float32(dy) / float32(dx)
				b = float32(y*dx-dy*x) / float32(dx)
			}

			if _, ok := cnt[k]; !ok {
				cnt[k] = map[float32]int{}
			}
			cnt[k][b]++ // 按照斜率和截距分组

			mid := pair{x + x2, y + y2}
			if _, ok := cnt2[mid]; !ok {
				cnt2[mid] = map[float32]int{}
			}
			cnt2[mid][k]++ // 按照中点和斜率分组
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
