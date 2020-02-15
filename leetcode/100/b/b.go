package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) (ans *TreeNode) {
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

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}
	ans = nodes[0]
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Right = nodes[i+1]
	}
	return
}
