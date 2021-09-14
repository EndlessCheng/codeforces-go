package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

func isEvenOddTree(root *TreeNode) (ans bool) {
	q := []*TreeNode{root}
	for parity := 0; len(q) > 0; parity ^= 1 {
		tmp := q
		q = nil
		for i, o := range tmp {
			if o.Val&1 == parity || i > 0 && (o.Val == tmp[i-1].Val || o.Val < tmp[i-1].Val == (parity == 0)) {
				return
			}
			if o.Left != nil {
				q = append(q, o.Left)
			}
			if o.Right != nil {
				q = append(q, o.Right)
			}
		}
	}
	return true
}
