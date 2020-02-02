package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mod int = 1e9 + 7

	all := 0
	var f1 func(o *TreeNode)
	f1 = func(o *TreeNode) {
		if o != nil {
			all += o.Val
			f1(o.Left)
			f1(o.Right)
		}
	}
	f1(root)

	var f2 func(o *TreeNode) int
	f2 = func(o *TreeNode) int {
		sum := o.Val
		if o.Left != nil {
			s := f2(o.Left)
			ans = max(ans, s*(all-s))
			sum += s
		}
		if o.Right != nil {
			s := f2(o.Right)
			ans = max(ans, s*(all-s))
			sum += s
		}
		return sum
	}
	f2(root)
	return ans % mod
}
