package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func swapNodes(head *ListNode, k int) *ListNode {
	a := []*ListNode{}
	for o := head; o != nil; o = o.Next {
		a = append(a, o)
	}
	k--
	a[k].Val, a[len(a)-1-k].Val = a[len(a)-1-k].Val, a[k].Val
	return head
}
