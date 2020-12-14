package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func deepestLeavesSum(root *TreeNode) (ans int) {
	mx := -1
	var dfs func(*TreeNode, int)
	dfs = func(o *TreeNode, d int) {
		if o == nil {
			return
		}
		if v := o.Val; d > mx {
			mx = d
			ans = v
		} else if d == mx {
			ans += v
		}
		dfs(o.Left, d+1)
		dfs(o.Right, d+1)
	}
	dfs(root, 0)
	return
}
