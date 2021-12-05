package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

/* 快慢指针

这题其实就是把 [876. 链表的中间结点](https://leetcode-cn.com/problems/middle-of-the-linked-list/) 和删除链表结点结合起来

我们只需要在 876 题的基础上，额外记录 $\textit{slow}$ 结点的上一个结点即可。

*/

// github.com/EndlessCheng/codeforces-go
func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre, slow, fast := dummyHead, head, head
	for fast != nil && fast.Next != nil {
		pre = slow // pre 记录了 slow 的上一个结点
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 循环结束后，slow 为待删除结点
	pre.Next = slow.Next // 删除 slow
	return dummyHead.Next
}
