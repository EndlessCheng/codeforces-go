package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"slices"
)

// https://space.bilibili.com/206214
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	a := []int64{}
	for len(q) > 0 {
		sum := int64(0)
		tmp := q
		q = nil
		for _, node := range tmp {
			sum += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		a = append(a, sum)
	}
	n := len(a)
	if k > n {
		return -1
	}
	slices.Sort(a)
	return a[n-k]
}
