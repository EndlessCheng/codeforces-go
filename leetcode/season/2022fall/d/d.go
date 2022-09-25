package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func closeLampInTree(root *TreeNode) (ans int) {
	type tuple struct {
		node             *TreeNode
		switch2, switch3 bool
	}
	dp := map[tuple]int{}
	var f func(*TreeNode, bool, bool) int
	f = func(node *TreeNode, switch2, switch3 bool) int {
		if node == nil {
			return 0
		}
		p := tuple{node, switch2, switch3}
		if res, ok := dp[p]; ok {
			return res
		}
		if node.Val == 1 == (switch2 == switch3) {
			res1 := f(node.Left, switch2, false) + f(node.Right, switch2, false) + 1
			res2 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 1
			res3 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 1
			r123 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 3
			dp[p] = min(res1, res2, res3, r123)
		} else {
			res0 := f(node.Left, switch2, false) + f(node.Right, switch2, false)
			res12 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 2
			res13 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 2
			res23 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 2
			dp[p] = min(res0, res12, res13, res23)
		}
		return dp[p]
	}
	return f(root, false, false)
}

func min(a, b, c, d int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	if d < a {
		a = d
	}
	return a
}
