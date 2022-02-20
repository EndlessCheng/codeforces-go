package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func mergeNodes(head *ListNode) *ListNode {
	ans := head
	for node, sum := head.Next, 0; node != nil; node = node.Next {
		if node.Val > 0 {
			sum += node.Val
		} else {
			ans.Next.Val = sum // 原地修改
			ans = ans.Next
			sum = 0
		}
	}
	ans.Next = nil
	return head.Next
}
