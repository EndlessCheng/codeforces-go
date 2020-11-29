package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func mergeInBetween(l1 *ListNode, a int, b int, l2 *ListNode) (ans *ListNode) {
	ns := []*ListNode{}
	for o := l1; o != nil; o = o.Next {
		ns = append(ns, o)
	}
	ns[a-1].Next = l2
	for o := l2; ; o = o.Next {
		if o.Next == nil {
			o.Next = ns[b+1]
			break
		}
	}
	return l1
}
