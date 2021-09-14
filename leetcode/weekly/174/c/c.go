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
func maxProduct(root *TreeNode) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mod int = 1e9 + 7

	all := 0
	var f1 func(o *TreeNode)
	f1 = func(o *TreeNode) {
		if o != nil {
			all += o.Val
			f1(o.Left)
			f1(o.Right)
		}
	}
	f1(root)

	var f2 func(o *TreeNode) int
	f2 = func(o *TreeNode) int {
		if o == nil {
			return 0
		}
		sum := o.Val + f2(o.Left) + f2(o.Right)
		ans = max(ans, sum*(all-sum))
		return sum
	}
	f2(root)
	return ans % mod
}
