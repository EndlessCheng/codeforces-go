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
		if has[cur.Next.Val] {
			cur.Next = cur.Next.Next // 删除
		} else {
			cur = cur.Next // 向后移动
		}
	}
	return dummy.Next
}
