package main

import "sort"

// https://space.bilibili.com/206214
func numberOfPairs(a, nums2 []int, diff int) (ans int64) {
	for i, x := range nums2 {
		a[i] -= x
	}
	b := append(sort.IntSlice{}, a...)
	b.Sort() // 离散化

	t := make(BIT, len(a)+1)
	for _, x := range a {
		ans += int64(t.query(b.Search(x + diff + 1)))
		t.add(b.Search(x) + 1)
	}
	return
}

type BIT []int

func (t BIT) add(x int) {
	for x < len(t) {
		t[x]++
		x += x & -x
	}
}

func (t BIT) query(x int) (res int) {
	for x > 0 {
		res += t[x]
		x &= x - 1
	}
	return
}
