package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func dfs(node1 *TreeNode, node2 *TreeNode, isOddLevel bool) {
	if node1 == nil {
		return
	}
	if isOddLevel {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
	dfs(node1.Left, node2.Right, !isOddLevel)
	dfs(node1.Right, node2.Left, !isOddLevel)
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	dfs(root.Left, root.Right, true)
	return root
}
