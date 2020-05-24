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
func pseudoPalindromicPaths(root *TreeNode) (ans int) {
	cnt := [10]int{}
	var f func(*TreeNode)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}
		v := o.Val
		cnt[v]++
		defer func() { cnt[v]-- }()
		if o.Left == nil && o.Right == nil {
			odd := 0
			for _, c := range cnt {
				if c&1 == 1 {
					odd++
				}
			}
			if odd <= 1 {
				ans++
			}
		}
		f(o.Left)
		f(o.Right)
	}
	f(root)
	return
}
