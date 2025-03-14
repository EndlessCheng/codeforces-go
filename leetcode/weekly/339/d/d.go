package main

import "github.com/emirpasic/gods/v2/trees/redblacktree"

// https://space.bilibili.com/206214
type unionFind struct {
	fa []int
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa}
}

func (uf unionFind) find(x int) int {
	if uf.fa[x] != x {
		uf.fa[x] = uf.find(uf.fa[x])
	}
	return uf.fa[x]
}

func (uf unionFind) merge(from, to int) {
	uf.fa[uf.find(from)] = uf.find(to)
}

func minReverseOperations(n, p int, banned []int, k int) []int {
	indices := []unionFind{newUnionFind(n + 2), newUnionFind(n + 2)}
	indices[p%2].merge(p, p+2) // 删除 p
	for _, i := range banned {
		indices[i%2].merge(i, i+2) // 删除 i
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	ans[p] = 0
	q := []int{p}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		mn := max(i-k+1, k-i-1)
		mx := min(i+k-1, n*2-k-i-1)
		uf := indices[mn%2]
		for j := uf.find(mn); j <= mx; j = uf.find(j + 2) { // 快速跳到 >= j+2 的下一个下标
			ans[j] = ans[i] + 1
			q = append(q, j)
			uf.merge(j, mx+2) // 删除 j
		}
	}
	return ans
}

func minReverseOperations1(n int, p int, banned []int, k int) []int {
	ban := map[int]struct{}{p: {}}
	for _, b := range banned {
		ban[b] = struct{}{}
	}
	indices := [2]*redblacktree.Tree[int, struct{}]{
		redblacktree.New[int, struct{}](),
		redblacktree.New[int, struct{}](),
	}
	for i := range n {
		if _, ok := ban[i]; !ok {
			indices[i%2].Put(i, struct{}{})
		}
	}
	indices[0].Put(n, struct{}{}) // 哨兵，下面无需判断 node != nil
	indices[1].Put(n, struct{}{})

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	ans[p] = 0 // 起点
	q := []int{p}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		// indices[mn % 2] 中的从 mn 到 mx 的所有下标都可以从 i 翻转到
		mn := max(i-k+1, k-i-1)
		mx := min(i+k-1, n*2-k-i-1)
		t := indices[mn%2]
		for node, _ := t.Ceiling(mn); node.Key <= mx; node, _ = t.Ceiling(mn) {
			j := node.Key
			ans[j] = ans[i] + 1 // 移动一步
			q = append(q, j)
			t.Remove(j)
		}
	}
	return ans
}
