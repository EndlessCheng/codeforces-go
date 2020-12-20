package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func sumEvenGrandparent(root *TreeNode) (ans int) {
	var dfs func(o *TreeNode, fa, pa int)
	dfs = func(o *TreeNode, fa, pa int) {
		if o == nil {
			return
		}
		v := o.Val
		if pa&1 == 0 {
			ans += v
		}
		dfs(o.Left, v, fa)
		dfs(o.Right, v, fa)
	}
	dfs(root, 1, 1)
	return
}
