package main

import "github.com/emirpasic/gods/trees/redblacktree"

// github.com/EndlessCheng/codeforces-go
type CountIntervals struct {
	*redblacktree.Tree[int, int]
	cnt int
}

func Constructor() CountIntervals {
	return CountIntervals{redblacktree.New[int, int](), 0}
}

func (t *CountIntervals) Add(left, right int) {
	// 遍历所有被 [left,right] 覆盖到的区间（部分覆盖也算）
	for node, _ := t.Ceiling(left); node != nil && node.Value <= right; node, _ = t.Ceiling(left) {
		l, r := node.Value, node.Key
		if l < left { left = l }   // 合并后的新区间，其左端点为所有被覆盖的区间的左端点的最小值
		if r > right { right = r } // 合并后的新区间，其右端点为所有被覆盖的区间的右端点的最大值
		t.cnt -= r - l + 1
		t.Remove(r)
	}
	t.cnt += right - left + 1
	t.Put(right, left) // 所有被覆盖到的区间与 [left,right] 合并成一个新区间
}

func (t *CountIntervals) Count() int { return t.cnt }
