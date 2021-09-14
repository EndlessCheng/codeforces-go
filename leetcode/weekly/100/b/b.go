package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var f func(o *TreeNode)
	f = func(o *TreeNode) {
		if o != nil {
			f(o.Left)
			vals = append(vals, o.Val)
			f(o.Right)
		}
	}
	f(root)

	n := len(vals)
	nodes := make([]*TreeNode, n)
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}
	for i := 1; i < n; i++ {
		nodes[i-1].Right = nodes[i]
	}
	return nodes[0]
}
