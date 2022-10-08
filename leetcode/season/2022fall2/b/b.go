package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func expandBinaryTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = &TreeNode{-1, expandBinaryTree(root.Left), nil}
	}
	if root.Right != nil {
		root.Right = &TreeNode{-1, nil, expandBinaryTree(root.Right)}
	}
	return root
}
