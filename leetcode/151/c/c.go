package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(o *ListNode) *ListNode {
	head := &ListNode{0, o}
	mp := map[int]*ListNode{}
	sum := 0
	for o = head; o != nil; o = o.Next {
		sum += o.Val
		mp[sum] = o
	}
	sum = 0
	for o = head; o != nil; o = o.Next {
		sum += o.Val
		o.Next = mp[sum].Next // 相同，直接映射到后方
	}
	return head.Next
}
