package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	maxVal := int(-2e9)
	maxPos := 1
	q := [2][]*TreeNode{{}, {root}}
	for i := 1; len(q[i&1]) > 0; i++ {
		s := 0
		q[(i+1)&1] = []*TreeNode{}
		for _, o := range q[i&1] {
			s += o.Val
			if o.Left != nil {
				q[(i+1)&1] = append(q[(i+1)&1], o.Left)
			}
			if o.Right != nil {
				q[(i+1)&1] = append(q[(i+1)&1], o.Right)
			}
		}
		if s > maxVal {
			maxVal = s
			maxPos = i
		}
	}
	return maxPos
}
