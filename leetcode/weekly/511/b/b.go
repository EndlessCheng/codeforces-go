package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func countDominantNodes(root *TreeNode) (ans int) {
	// dfs(node) 返回 node 子树中的最大节点值
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		mx := max(dfs(node.Left), dfs(node.Right))
		if node.Val >= mx {
			// node.Val 是 node 子树中的最大节点值
			ans++
			mx = node.Val
		}
		return mx
	}

	dfs(root)
	return
}
