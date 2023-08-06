package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {
	cur := head
	for cur.Next != nil {
		cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
		cur = cur.Next.Next
	}
	return head
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
