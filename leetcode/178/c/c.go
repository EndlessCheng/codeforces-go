package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) (ans bool) {
	a := []int{}
	for o := head; o != nil; o = o.Next {
		a = append(a, o.Val)
	}
	n := len(a)

	var checkPath func(o *TreeNode, p int) bool
	checkPath = func(o *TreeNode, p int) bool {
		return p == n || o != nil && o.Val == a[p] && (checkPath(o.Left, p+1) || checkPath(o.Right, p+1))
	}
	var f func(o *TreeNode) bool
	f = func(o *TreeNode) bool {
		return o != nil && (checkPath(o, 0) || f(o.Left) || f(o.Right))
	}
	return f(root)
}
