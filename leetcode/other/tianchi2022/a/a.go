package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func numberEvenListNode(head *ListNode) (ans int) {
	for o := head; o != nil; o = o.Next {
		ans += o.Val & 1
	}
	return
}
