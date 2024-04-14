package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func numberOfSubarrays(nums []int) int64 {
	ans := len(nums)
	type pair struct{ x, cnt int }
	st := []pair{{math.MaxInt, 0}} // 哨兵
	for _, x := range nums {
		for x > st[len(st)-1].x {
			st = st[:len(st)-1]
		}
		if x == st[len(st)-1].x {
			ans += st[len(st)-1].cnt
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{x, 1})
		}
	}
	return int64(ans)
}

func numberOfSubarrays2(a []int) int64 {
	n := len(a)
	ans := n
	right := make([]int, n)
	st := []int{n}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for len(st) > 1 && a[st[len(st)-1]] <= v {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	pos := map[int][]int{}
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}
	keys := make([]int, 0, len(pos))
	for k := range pos {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, key := range keys {
		ps := pos[key]
		for i, n := 0, len(ps); i < n; {
			st := i
			upp := right[ps[i]]
			for ; i < n && ps[i] < upp; i++ {
			}
			ans += (i - st) * (i - st - 1) / 2
		}
	}
	return int64(ans)
}
