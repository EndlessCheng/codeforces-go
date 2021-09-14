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

func same(l *ListNode, t *TreeNode) bool {
	return l == nil || t != nil && t.Val == l.Val && (same(l.Next, t.Left) || same(l.Next, t.Right))
}
func isSubPath(head *ListNode, t *TreeNode) bool {
	return t != nil && (same(head, t) || isSubPath(head, t.Left) || isSubPath(head, t.Right))
}
