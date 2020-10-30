package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func goodNodes(root *TreeNode) (ans int) {
	var f func(*TreeNode, int)
	f = func(o *TreeNode, mx int) {
		if o == nil {
			return
		}
		v := o.Val
		if v >= mx {
			ans++
			mx = v
		}
		f(o.Left, mx)
		f(o.Right, mx)
	}
	f(root, -1e9)
	return
}
