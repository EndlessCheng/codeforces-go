package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func reContruct(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val%2 == 0 {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
