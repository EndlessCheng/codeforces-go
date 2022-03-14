package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func isP(a []int) bool {
	for i, n := 0, len(a); i < n/2; i++ {
		v, w := a[i], a[n-1-i]
		if v != w {
			return false
		}
	}
	return true
}

func isPalindrome(head *ListNode) (ans bool) {
	a := []int{}
	for o := head; o != nil; o = o.Next {
		a = append(a, o.Val)
	}
	i, j := 0, len(a)-1
	for i < j && a[i] == a[j] {
		i++
		j--
	}
	return i >= j || isP(a[i:j]) || isP(a[i+1:j+1])
}
