package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func modifiedList(nums []int, head *ListNode) *ListNode {
	has := make(map[int]bool, len(nums)) // 预分配空间
	for _, x := range nums {
		has[x] = true
	}

	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		nxt := cur.Next
		if has[nxt.Val] {
			cur.Next = nxt.Next // 从链表中删除 nxt 节点
		} else {
			cur = nxt // 不删除 nxt，继续向后遍历链表
		}
	}
	return dummy.Next
}
