package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func closeLampInTree(root *TreeNode) (ans int) {
	type tuple struct {
		node             *TreeNode
		switch2, switch3 bool
	}
	memo := map[tuple]int{}
	var dfs func(*TreeNode, bool, bool) int
	dfs = func(node *TreeNode, switch2, switch3 bool) int {
		if node == nil {
			return 0
		}
		p := tuple{node, switch2, switch3}
		if res, ok := memo[p]; ok {
			return res
		}
		if node.Val == 1 == (switch2 == switch3) {
			res1 := dfs(node.Left, switch2, false) + dfs(node.Right, switch2, false) + 1
			res2 := dfs(node.Left, !switch2, false) + dfs(node.Right, !switch2, false) + 1
			res3 := dfs(node.Left, switch2, true) + dfs(node.Right, switch2, true) + 1
			r123 := dfs(node.Left, !switch2, true) + dfs(node.Right, !switch2, true) + 3
			memo[p] = min(res1, res2, res3, r123)
		} else {
			res0 := dfs(node.Left, switch2, false) + dfs(node.Right, switch2, false)
			res12 := dfs(node.Left, !switch2, false) + dfs(node.Right, !switch2, false) + 2
			res13 := dfs(node.Left, switch2, true) + dfs(node.Right, switch2, true) + 2
			res23 := dfs(node.Left, !switch2, true) + dfs(node.Right, !switch2, true) + 2
			memo[p] = min(res0, res12, res13, res23)
		}
		return memo[p]
	}
	return dfs(root, false, false)
}
