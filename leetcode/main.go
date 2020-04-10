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

	var root *TreeNode

	var f func(*TreeNode)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}

		f(o.Left)
		f(o.Right)
	}
	f(root)

	_ = []interface{}{toBytes, ListNode{}, TreeNode{}}
}
