package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

func maxProduct(root *TreeNode) (ans int) {
	var dfs1 func(*TreeNode) int
	dfs1 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val + dfs1(node.Left) + dfs1(node.Right)
	}
	total := dfs1(root)

	var dfs2 func(*TreeNode) int
	dfs2 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		s := node.Val + dfs2(node.Left) + dfs2(node.Right)
		ans = max(ans, s*(total-s))
		return s
	}
	dfs2(root)

	return ans % 1_000_000_007
}
