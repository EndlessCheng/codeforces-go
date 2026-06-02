package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil { // 只有一个节点
		return nil
	}
	// 876. 链表的中间结点
	// 本题先让兔子走两步，这样乌龟少走一步，刚好落在中间节点的前一个节点
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next // 删除 slow 的下一个节点
	return head
}
