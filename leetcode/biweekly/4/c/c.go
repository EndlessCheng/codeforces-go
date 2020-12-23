package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func maximumAverageSubtree(root *TreeNode) (ans float64) {
	a, b := 0, 1
	var dfs func(*TreeNode) (int, int)
	dfs = func(o *TreeNode) (sum, sz int) {
		if o == nil {
			return
		}
		sum, sz = o.Val, 1
		x, y := dfs(o.Left)
		sum += x
		sz += y
		x, y = dfs(o.Right)
		sum += x
		sz += y
		if sum*b > sz*a {
			a, b = sum, sz
		}
		return
	}
	dfs(root)
	return float64(a) / float64(b)
}
