package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var nodes []*ListNode
	for node, size := head, 1; node != nil; node = node.Next {
		nodes = append(nodes, node)
		if len(nodes) == size || node.Next == nil {
			if n := len(nodes); n%2 == 0 {
				for i := 0; i < n/2; i++ {
					nodes[i].Val, nodes[n-1-i].Val = nodes[n-1-i].Val, nodes[i].Val // 直接交换元素值
				}
			}
			nodes = nil
			size++
		}
	}
	return head
}
