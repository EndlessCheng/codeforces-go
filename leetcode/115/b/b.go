package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func isCompleteTree(root *TreeNode) (ans bool) {
	end := false
	q := []*TreeNode{root}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		if o == nil {
			end = true
		} else if end {
			return
		} else {
			q = append(q, o.Left, o.Right)
		}
	}
	return true
}

// 我的憨憨写法
func isCompleteTree2(root *TreeNode) (ans bool) {
	q := []*TreeNode{root}
	for {
		for i, o := range q {
			if o == nil {
				for _, o2 := range q[i:] {
					if o2 != nil {
						return false
					}
				}
				for _, o2 := range q[:i] {
					if o2.Left != nil || o2.Right != nil {
						return false
					}
				}
				break
			}
		}
		if q[0] == nil {
			return true
		}
		tmp := q
		q = nil
		for _, o := range tmp {
			if o != nil {
				q = append(q, o.Left, o.Right)
			}
		}
	}
}
