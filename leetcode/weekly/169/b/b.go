package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {
	var f func(o *TreeNode)
	f = func(o *TreeNode) {
		if o != nil {
			ans = append(ans, o.Val)
			f(o.Left)
			f(o.Right)
		}
	}
	f(root1)
	f(root2)
	sort.Ints(ans)
	return
}
