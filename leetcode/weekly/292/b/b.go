package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func averageOfSubtree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		sum, cnt := node.Val, 1
		if node.Left != nil {
			s, c := dfs(node.Left)
			sum += s
			cnt += c
		}
		if node.Right != nil {
			s, c := dfs(node.Right)
			sum += s
			cnt += c
		}
		if node.Val == sum/cnt {
			ans++
		}
		return sum, cnt
	}
	dfs(root)
	return
}
