package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

func getDecimalValue(head *ListNode) (ans int) {
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return
}
