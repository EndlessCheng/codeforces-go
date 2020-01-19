package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 我的憨憨写法
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

// 优雅的写法
func removeLeafNodes2(o *TreeNode, tar int) *TreeNode {
	if o == nil {
		return nil
	}
	o.Left = removeLeafNodes2(o.Left, tar)
	o.Right = removeLeafNodes2(o.Right, tar)
	if o.Left == nil && o.Right == nil && o.Val == tar {
		return nil
	}
	return o
}
