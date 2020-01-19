package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, tar int) *TreeNode {
	var del func(o *TreeNode) bool
	del = func(o *TreeNode) bool {
		if o == nil {
			return false
		}
		if del(o.Left) {
			o.Left = nil
		}
		if del(o.Right) {
			o.Right = nil
		}
		return o.Left == nil && o.Right == nil && o.Val == tar
	}
	if del(root) {
		return nil
	}
	return root
}
