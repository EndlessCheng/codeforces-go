package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"strconv"
)

func main() {
	_ = strconv.FormatInt(18880, 2)

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

	var dfsTreeNode func(o *TreeNode, f func(o *TreeNode))
	dfsTreeNode = func(o *TreeNode, f func(*TreeNode)) {
		if o != nil {
			f(o)
			dfsTreeNode(o.Left, f)
			dfsTreeNode(o.Right, f)
		}
	}
	var root *TreeNode
	dfsTreeNode(root, func(o *TreeNode) {

	})

	_ = []interface{}{toBytes, ListNode{}, TreeNode{}}
}
