package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func nodesBetweenCriticalPoints(head *ListNode) []int {
	first, pre := 0, math.MinInt/2
	minDis := math.MaxInt
	a, b, c := head, head.Next, head.Next.Next

	for i := 1; c != nil; i++ {
		if a.Val < b.Val && b.Val > c.Val || a.Val > b.Val && b.Val < c.Val {
			if first == 0 {
				first = i
			}
			minDis = min(minDis, i-pre)
			pre = i
		}
		a = b
		b = c
		c = c.Next
	}

	if first >= pre { // 临界点少于两个
		return []int{-1, -1}
	}
	return []int{minDis, pre - first}
}
