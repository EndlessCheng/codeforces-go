package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxDistance(side int, points [][]int, k int) int {
	a := make([]int, len(points))
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

	ans := sort.Search(side, func(low int) bool {
		low++
		idx := make([]int, k)
		cur := a[0]
		for j, i := 1, 0; j < k; j++ {
			i += sort.Search(len(a)-i, func(j int) bool { return a[i+j] >= cur+low })
			if i == len(a) {
				return true
			}
			idx[j] = i
			cur = a[i]
		}
		if cur-a[0] <= side*4-low {
			return false
		}

		// 第一个指针移动到第二个指针的位置，就不用继续枚举了
		end0 := idx[1]
		for idx[0]++; idx[0] < end0; idx[0]++ {
			for j := 1; j < k; j++ {
				for a[idx[j]] < a[idx[j-1]]+low {
					idx[j]++
					if idx[j] == len(a) {
						return true
					}
				}
			}
			if a[idx[k-1]]-a[idx[0]] <= side*4-low {
				return false
			}
		}
		return true
	})
	return ans
}

func maxDistance2(side int, points [][]int, k int) int {
	a := make([]int, len(points))
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

	ans := sort.Search(side, func(low int) bool {
		low++
		idx := make([]int, k)
		for {
			for j := 1; j < k; j++ {
				for a[idx[j]] < a[idx[j-1]]+low {
					idx[j]++
					if idx[j] == len(a) {
						return true
					}
				}
			}
			if a[idx[k-1]]-a[idx[0]] <= side*4-low {
				return false
			}
			idx[0]++
		}
	})
	return ans
}
