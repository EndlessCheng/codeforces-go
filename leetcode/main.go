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

// LC 41
func firstMissingPositive(a []int) int {
	n := len(a)
	for i, v := range a {
		for 0 < v && v <= n && v != a[v-1] {
			a[i], a[v-1] = a[v-1], a[i]
			v = a[i]
		}
	}
	for i, v := range a {
		if i+1 != v {
			return i + 1
		}
	}
	return n + 1
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
