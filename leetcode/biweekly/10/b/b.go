package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func twoSumBSTs(root1, root2 *TreeNode, target int) (ans bool) {
	has := map[int]bool{}
	var f func(o *TreeNode)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}
		has[o.Val] = true
		f(o.Left)
		f(o.Right)
	}
	f(root1)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}
		if has[target-o.Val] {
			ans = true
		}
		f(o.Left)
		f(o.Right)
	}
	f(root2)
	return
}
