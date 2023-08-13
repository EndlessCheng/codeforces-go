package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

// https://space.bilibili.com/206214
func minAbsoluteDifference(nums []int, k int) int {
	ans := math.MaxInt
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt, nil) // 哨兵
	t.Put(math.MinInt/2, nil)
	for i, y := range nums[k:] {
		t.Put(nums[i], nil)
		c, _ := t.Ceiling(y)
		f, _ := t.Floor(y)
		ans = min(ans, min(c.Key.(int)-y, y-f.Key.(int)))
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }