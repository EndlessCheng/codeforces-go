package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func mergeNodes(head *ListNode) *ListNode {
	for node, ans, sum := head.Next, head, 0; node != nil; node = node.Next {
		if node.Val > 0 {
			sum += node.Val
		} else {
			ans.Next = &ListNode{sum, nil}
			ans = ans.Next
			sum = 0
		}
	}
	return head.Next
}
