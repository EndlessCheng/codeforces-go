package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {
	var f  func(o *TreeNode)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}
		ans = append(ans, o.Val)
		f(o.Left)
		f(o.Right)
	}
	f(root1)
	f(root2)
	sort.Ints(ans)
	return
}
