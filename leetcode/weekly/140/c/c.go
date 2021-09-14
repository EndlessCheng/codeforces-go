package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func f(o *TreeNode, sum int) (goodNode bool) {
	sum += o.Val
	if o.Left == nil && o.Right == nil {
		return sum >= 0
	}
	if o.Left != nil {
		if f(o.Left, sum) {
			goodNode = true
		} else {
			o.Left = nil
		}
	}
	if o.Right != nil {
		if f(o.Right, sum) {
			goodNode = true
		} else {
			o.Right = nil
		}
	}
	return
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if f(root, -limit) {
		return root
	}
	return nil
}
