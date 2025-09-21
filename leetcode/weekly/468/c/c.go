package main

import "slices"

// https://space.bilibili.com/206214
func minSplitMerge(nums1, nums2 []int) (ans int) {
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
