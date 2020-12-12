package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func longestZigZag(root *TreeNode) (ans int) {
	dp := [2]map[*TreeNode]int{{}, {}}
	var f func(o *TreeNode, dir int) int
	f = func(o *TreeNode, dir int) (res int) {
		if v, ok := dp[dir][o]; ok {
			return v
		}
		defer func() { dp[dir][o] = res }()
		if dir == 0 && o.Right != nil {
			res = max(res, f(o.Right, 1))
		}
		if dir == 1 && o.Left != nil {
			res = max(res, f(o.Left, 0))
		}
		return res + 1
	}
	var dfs func(o *TreeNode)
	dfs = func(o *TreeNode) {
		if o.Left != nil {
			ans = max(ans, f(o.Left, 0))
			dfs(o.Left)
		}
		if o.Right != nil {
			ans = max(ans, f(o.Right, 1))
			dfs(o.Right)
		}
	}
	dfs(root)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
