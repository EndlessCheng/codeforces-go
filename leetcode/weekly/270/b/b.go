package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func deleteMiddle(head *ListNode) *ListNode {
	dummy := ListNode{Next: head}

	// 876. 链表的中间结点
	slow := &dummy
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next // 删除 slow 的下一个节点
	return dummy.Next
}
