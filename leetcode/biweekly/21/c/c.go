package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := [2]map[*TreeNode]int{{}, {}}
	var f func(o *TreeNode, dir int) int
	f = func(o *TreeNode, dir int) (_ans int) {
		if v, ok := dp[dir][o]; ok {
			return v
		}
		_ans = 1
		if dir == 0 && o.Right != nil {
			_ans = max(_ans, f(o.Right, 1)+1)
		}
		if dir == 1 && o.Left != nil {
			_ans = max(_ans, f(o.Left, 0)+1)
		}
		dp[dir][o] = _ans
		return
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
