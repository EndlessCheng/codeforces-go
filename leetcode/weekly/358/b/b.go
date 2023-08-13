package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func doubleIt(head *ListNode) *ListNode {
	if head.Val > 4 {
		head = &ListNode{0, head}
	}
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = cur.Val * 2 % 10
		if cur.Next != nil && cur.Next.Val > 4 {
			cur.Val++
		}
	}
	return head
}
