package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type node struct{ l, r, v int }
var a []node
func del(i int) { a[a[i].l].r, a[a[i].r].l = a[i].r, a[i].l }

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]].v > a[h.IntSlice[j]].v }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func maxSizeSlices(A []int) (ans int) {
	n := len(A)
	a = make([]node, n)
	for i, v := range A {
		a[i] = node{(i - 1 + n) % n, (i + 1) % n, v}
	}
	vis := make([]bool, n)
	h := &hp{make([]int, n)}
	for i := 0; i < n; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	for k := n / 3; k > 0; {
		i := h.IntSlice[0]
		if vis[i] {
			heap.Pop(h)
			continue
		}
		ans += a[i].v
		l, r := a[i].l, a[i].r
		vis[l] = true
		vis[r] = true
		a[i].v = a[l].v + a[r].v - a[i].v
		heap.Fix(h, 0)
		del(l)
		del(r)
		k--
	}
	return
}

// O(n^2) 做法
func maxSizeSlices2(a []int) (ans int) {
	n := len(a)
	m := n / 3
	f := func(a []int) int {
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		for i := 1; i < n; i++ {
			for j := 1; j <= m; j++ {
				v := 0
				if i > 1 {
					v = dp[i-2][j-1]
				}
				dp[i][j] = max(dp[i-1][j], v+a[i-1])
			}
		}
		return dp[n-1][m]
	}
	return max(f(a[:n-1]), f(a[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
