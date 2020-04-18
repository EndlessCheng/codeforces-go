package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

func minimalExecTime(o *TreeNode) float64 {
	var f func(o *TreeNode) (a, b float64)
	f = func(o *TreeNode) (a, b float64) {
		if o.Left == nil || o.Right == nil {
			if o.Left != nil {
				a, b = f(o.Left)
			} else if o.Right != nil {
				a, b = f(o.Right)
			}
			a += float64(o.Val)
			return
		}
		a = float64(o.Val)
		c, d := f(o.Left)
		e, f := f(o.Right)
		if c < e {
			c, d, e, f = e, f, c, d
		}
		b = d + e + f
		c -= e
		if c <= 2*f {
			b += c / 2
		} else {
			b += f
			a += c - 2*f
		}
		return
	}
	a, b := f(o)
	return a + b
}
