package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	pos := map[int]int{}
	for i := -n; i < n; i++ {
		if i >= 0 {
			j := pos[nums[i]]
			left[i] = j
			// 对于左边的 j 来说，它的 right 就是 i
			if j >= 0 {
				right[j] = i
			} else {
				right[j+n] = i + n
			}
		}
		pos[nums[(i+n)%n]] = i
	}

	for qi, i := range queries {
		l := left[i]
		if i-l == n {
			queries[qi] = -1
		} else {
			queries[qi] = min(i-l, right[i]-i)
		}
	}
	return queries
}

func solveQueries2(nums []int, queries []int) []int {
	n := len(nums)
	left := make([]int, n)
	pos := map[int]int{}
	for i := -n; i < n; i++ {
		if i >= 0 {
			left[i] = pos[nums[i]]
		}
		pos[nums[(i+n)%n]] = i
	}

	right := make([]int, n)
	clear(pos)
	for i := n*2 - 1; i >= 0; i-- {
		if i < n {
			right[i] = pos[nums[i]]
		}
		pos[nums[i%n]] = i
	}

	for qi, i := range queries {
		l := left[i]
		if i-l == n {
			queries[qi] = -1
		} else {
			queries[qi] = min(i-l, right[i]-i)
		}
	}
	return queries
}

func solveQueries1(nums []int, queries []int) []int {
	indices := map[int][]int{}
	for i, x := range nums {
		indices[x] = append(indices[x], i)
	}

	n := len(nums)
	for x, p := range indices {
		// 前后各加一个哨兵
		i0 := p[0]
		p = slices.Insert(p, 0, p[len(p)-1]-n)
		indices[x] = append(p, i0+n)
	}

	for qi, i := range queries {
		p := indices[nums[i]]
		if len(p) == 3 {
			queries[qi] = -1
		} else {
			j := sort.SearchInts(p, i)
			queries[qi] = min(i-p[j-1], p[j+1]-i)
		}
	}
	return queries
}
