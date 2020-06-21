package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func main() {
	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}

	_ = MustBuildTreeNode

	_ = []interface{}{toBytes, ListNode{}, TreeNode{}}
}

// LC 124
func maxPathSum(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := int(-1e18)
	var f func(*TreeNode) int
	f = func(o *TreeNode) int {
		if o == nil {
			return -1e18
		}
		l := max(f(o.Left), 0)
		r := max(f(o.Right), 0)
		ans = max(ans, o.Val+l+r)
		return o.Val + max(l, r)
	}
	f(root)
	return ans
}
