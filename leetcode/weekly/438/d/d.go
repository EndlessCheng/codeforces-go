package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxDistance(side int, points [][]int, k int) int {
	// 正方形边上的点，按照顺时针映射到一维数轴上
	a := make([]int, len(points)+1)
	for i, p := range points {
		x, y := p[0], p[1]
		if x == 0 {
			a[i] = y
		} else if y == side {
			a[i] = side + x
		} else if x == side {
			a[i] = side*3 - y
		} else {
			a[i] = side*4 - x
		}
	}
	slices.Sort(a)
	a[len(points)] = math.MaxInt / 2 // 哨兵

	// 本题保证 k >= 4，所以最远距离不会超过 side
	ans := sort.Search(side, func(low int) bool {
		// 如果 low+1 不满足要求，但 low 满足要求，那么答案就是 low
		low++
	next:
		for i, start := range a { // 枚举第一个点
			cur := start
			for range k - 1 { // 还需要找 k-1 个点
				i += sort.Search(len(a)-i, func(j int) bool { return a[i+j] >= cur+low })
				if i == len(a) || a[i] > start+side*4-low { // 不能离第一个点太近
					continue next
				}
				cur = a[i]
			}
			return false
		}
		return true
	})
	return ans
}
