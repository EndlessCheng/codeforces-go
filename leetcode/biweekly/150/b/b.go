package main

import (
	"maps"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func separateSquares(squares [][]int) float64 {
	totArea := 0.
	diff := map[int]int{}
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		totArea += float64(l * l)
		diff[y] += l
		diff[y+l] -= l
	}

	ys := slices.Sorted(maps.Keys(diff))
	area, sumL := 0., 0
	for i := 0; ; i++ {
		sumL += diff[ys[i]] // 矩形底边长度之和
		tmp := area + float64(sumL)*float64(ys[i+1]-ys[i]) // 底边长 * 高 = 新增面积
		if tmp >= totArea/2 {
			return float64(ys[i]) + (totArea/2-area)/float64(sumL)
		}
		area = tmp
	}
}

func separateSquares2(squares [][]int) float64 {
	totArea := 0.
	maxY := 0
	for _, sq := range squares {
		totArea += float64(sq[2] * sq[2])
		maxY = max(maxY, sq[1]+sq[2])
	}
	const m = 100_000 // 也可以调大，避免累计误差
	multiY := sort.Search(maxY*m, func(multiY int) bool {
		y := float64(multiY) / m
		area := 0.
		for _, sq := range squares {
			if float64(sq[1]) < y {
				l := float64(sq[2])
				area += l * min(y-float64(sq[1]), l)
			}
		}
		return area >= totArea/2
	})
	return float64(multiY) / m
}

func separateSquares1(squares [][]int) float64 {
	totArea := 0.
	maxY := 0
	for _, sq := range squares {
		totArea += float64(sq[2] * sq[2])
		maxY = max(maxY, sq[1]+sq[2])
	}

	check := func(y float64) bool {
		area := 0.
		for _, sq := range squares {
			yi := float64(sq[1])
			if yi < y {
				l := float64(sq[2])
				area += l * min(y-yi, l)
			}
		}
		return area >= totArea/2
	}

	left, right := 0., float64(maxY)
	for range bits.Len(uint(maxY * 1e5)) {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return (left + right) / 2 // 区间中点误差小
}
