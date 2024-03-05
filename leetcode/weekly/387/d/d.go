package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

// 把下标为 i 的元素增加 v
func (f fenwick) add(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

// 返回下标在 [1,i] 的元素之和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func resultArray(nums []int) (ans []int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	a := nums[:1]
	b := []int{nums[1]}
	t := make(fenwick, m+1)
	t.add(m-sort.SearchInts(sorted, nums[0]), 1)
	t.add(m-sort.SearchInts(sorted, nums[1]), -1)
	for _, x := range nums[2:] {
		v := m - sort.SearchInts(sorted, x)
		d := t.pre(v - 1)
		if d > 0 || d == 0 && len(a) <= len(b) {
			a = append(a, x)
			t.add(v, 1)
		} else {
			b = append(b, x)
			t.add(v, -1)
		}
	}
	return append(a, b...)
}
