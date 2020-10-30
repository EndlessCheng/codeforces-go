package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func middleNode(o *ListNode) (ans *ListNode) {
	nodes := []*ListNode{}
	for ; o != nil; o = o.Next {
		nodes = append(nodes, o)
	}
	return nodes[len(nodes)/2]
}
