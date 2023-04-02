package main

import "sort"

// https://space.bilibili.com/206214
type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u *uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	u.fa[x] = y
}

func minReverseOperations(n, p int, banned []int, k int) []int {
	ban := map[int]bool{p: true}
	for _, v := range banned {
		ban[v] = true
	}
	notBanned := [2][]int{}
	for i := 0; i < n; i++ {
		if !ban[i] {
			notBanned[i%2] = append(notBanned[i%2], i)
		}
	}
	notBanned[0] = append(notBanned[0], n)
	notBanned[1] = append(notBanned[1], n) // 哨兵
	ufs := [2]uf{newUnionFind(len(notBanned[0])), newUnionFind(len(notBanned[1]))}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	q := []int{p}
	for step := 0; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			ans[i] = step
			mn := max(i+k-(i*2+1), i-k+1)
			mx := min(i-k+((n-1-i)*2+1), i+k-1)
			a, u := notBanned[mn%2], ufs[mn%2]
			for j := u.find(sort.SearchInts(a, mn)); a[j] <= mx; j = u.find(j + 1) {
				q = append(q, a[j])
				u.merge(j, j+1)
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
