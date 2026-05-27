package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func pairSum1(head *ListNode) (ans int) {
	left := head

	var dfs func(*ListNode)
	dfs = func(right *ListNode) {
		// 「递」，先把 right 移到链表末尾
		if right.Next != nil {
			dfs(right.Next)
		}
		// 「归」的过程就是在从右到左遍历链表
		ans = max(ans, left.Val+right.Val)
		left = left.Next // left 往右走
		// 归，right 会往左走
	}

	dfs(head)
	return
}

//

// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func pairSum(head *ListNode) (ans int) {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2 != nil {
		ans = max(ans, head.Val+head2.Val)
		head = head.Next
		head2 = head2.Next
	}
	return
}
