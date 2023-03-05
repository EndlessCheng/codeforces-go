package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

// https://space.bilibili.com/206214
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	sum := []int{}
	for len(q) > 0 {
		tmp, s := q, 0
		q = nil
		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		sum = append(sum, s)
	}
	n := len(sum)
	if n < k {
		return -1
	}
	sort.Ints(sum) // 也可以用快速选择
	return int64(sum[n-k])
}
