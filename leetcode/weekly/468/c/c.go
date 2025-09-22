package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minSplitMerge1(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	t := [6]int{}
	for j, x := range nums1 {
		t[j] = x
	}
	vis := map[[6]int]bool{t: true}
	q := [][]int{nums1}
	for ; ; ans++ {
		tmp := q
		q = nil
		for _, a := range tmp {
			if slices.Equal(a, nums2) {
				return
			}
			for l := 0; l < n; l++ {
				for r := l + 1; r <= n; r++ {
					b := slices.Clone(a)
					sub := slices.Clone(b[l:r])
					b = append(b[:l], b[r:]...) // 从 b 中移除 sub
					for i := 0; i <= len(b); i++ {
						c := slices.Insert(slices.Clone(b), i, sub...)
						t := [6]int{}
						for j, x := range c {
							t[j] = x
						}
						if !vis[t] {
							vis[t] = true
							q = append(q, c)
						}
					}
				}
			}
		}
	}
}

func encode(nums, sorted []int) (res int) {
	for i, x := range nums {
		res |= sort.SearchInts(sorted, x) << (i * 3)
	}
	return
}

func minSplitMerge(nums1, nums2 []int) int {
	if slices.Equal(nums1, nums2) {
		return 0
	}

	n := len(nums1)
	sorted := slices.Clone(nums1)
	slices.Sort(sorted)

	val1 := encode(nums1, sorted)
	val2 := encode(nums2, sorted)

	vis := map[int]bool{val1: true}
	q := []int{val1}
	for ans := 1; ; ans++ {
		tmp := q
		q = nil
		for _, a := range tmp {
			for r := 1; r <= n; r++ { // 为方便实现，先枚举 r，再枚举 l
				t := a & (1<<(r*3) - 1)
				for l := range r {
					sub := t >> (l * 3)
					b := a&(1<<(l*3)-1) | a>>(r*3)<<(l*3) // 从 a 中移除 sub
					for i := range n - r + l + 1 {
						c := b&(1<<(i*3)-1) | sub<<(i*3) | b>>(i*3)<<((i+r-l)*3)
						if c == val2 {
							return ans
						}
						if !vis[c] {
							vis[c] = true
							q = append(q, c)
						}
					}
				}
			}
		}
	}
}
