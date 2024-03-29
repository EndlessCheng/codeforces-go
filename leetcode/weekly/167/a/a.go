package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(o *ListNode) int {
	s := ""
	for ; o != nil; o = o.Next {
		s += string('0' + o.Val)
	}
	v, _ := strconv.ParseInt(s, 2, 64)
	return int(v)
}
