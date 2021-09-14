package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	has map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	f := FindElements{map[int]bool{}}
	f.dfs(root, 0)
	return f
}

func (f *FindElements) dfs(o *TreeNode, v int) {
	if o != nil {
		f.has[v] = true
		f.dfs(o.Left, v*2+1)
		f.dfs(o.Right, v*2+2)
	}
}

func (f *FindElements) Find(target int) bool {
	return f.has[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
