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

	var f func(o *TreeNode, p int) bool
	f = func(o *TreeNode, p int) bool {
		if p == n {
			return true
		}
		if o == nil {
			return false
		}
		if o.Val == a[p] {
			return f(o.Left, p+1) || f(o.Right, p+1)
		}
		if o.Val == a[0] {
			return f(o.Left, 1) || f(o.Right, 1)
		}
		return f(o.Left, 0) || f(o.Right, 0)
	}
	return f(root, 0)
}
