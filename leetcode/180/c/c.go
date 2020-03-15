package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func build(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	m := len(a) / 2
	return &TreeNode{a[m], build(a[:m]), build(a[m+1:])}
}

func balanceBST(root *TreeNode) *TreeNode {
	a := []int{}
	var f func(*TreeNode)
	f = func(o *TreeNode) {
		if o != nil {
			a = append(a, o.Val)
			f(o.Left)
			f(o.Right)
		}
	}
	f(root)
	sort.Ints(a)
	return build(a)
}
