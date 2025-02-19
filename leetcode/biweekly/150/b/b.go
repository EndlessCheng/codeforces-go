package main

import (
	"maps"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func separateSquares3(squares [][]int) float64 {
	totArea := 0
	diff := map[int]int{}
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		totArea += l * l
		diff[y] += l
		diff[y+l] -= l
	}

	ys := slices.Sorted(maps.Keys(diff))
	area, sumL := 0, 0
	for i := 0; ; i++ {
		sumL += diff[ys[i]]              // 矩形底边长度之和
		area += sumL * (ys[i+1] - ys[i]) // 底边长 * 高 = 新增面积
		if area*2 >= totArea {
			return float64(ys[i+1]) - float64(area*2-totArea)/float64(sumL*2) // 这样写误差更小
		}
	}
}

func separateSquares(squares [][]int) float64 {
	totArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	calcArea := func(y int) (area int) {
		for _, sq := range squares {
			yi := sq[1]
			if yi < y {
				l := sq[2]
				area += l * min(y-yi, l)
			}
		}
		return
	}
	y := sort.Search(maxY, func(y int) bool { return calcArea(y)*2 >= totArea })

	areaY := calcArea(y)
	sumL := areaY - calcArea(y-1)
	return float64(y) - float64(areaY*2-totArea)/float64(sumL*2) // 这样写误差更小
}

func separateSquares2(squares [][]int) float64 {
	totArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	const m = 100_000
	multiY := sort.Search(maxY*m, func(multiY int) bool {
		area := 0
		for _, sq := range squares {
			y, l := sq[1], sq[2]
			if y*m < multiY {
				area += l * min(multiY-y*m, l*m)
			}
		}
		return area*2 >= totArea*m
	})
	return float64(multiY) / m
}

func separateSquares1(squares [][]int) float64 {
	totArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totArea += l * l
		maxY = max(maxY, sq[1]+l)
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
		return area >= float64(totArea)/2
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
