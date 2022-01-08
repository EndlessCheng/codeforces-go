package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func pairSum(head *ListNode) (ans int) {
	a := []int{}
	for node := head; node != nil; node = node.Next {
		a = append(a, node.Val)
	}
	for i, n := 0, len(a); i < n/2; i++ {
		ans = max(ans, a[i]+a[n-1-i])
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
