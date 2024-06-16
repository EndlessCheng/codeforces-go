package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	f := make([]int, n+1)
	j := 0
	for i, x := range a {
		for a[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[n])
}

func maximumTotalDamage2(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		x := a[i]
		j := i
		for j > 0 && a[j-1] >= x-2 {
			j--
		}
		*p = max(dfs(i-1), dfs(j-1)+x*cnt[x])
		return *p
	}
	return int64(dfs(n - 1))
}
