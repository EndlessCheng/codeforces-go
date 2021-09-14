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
func countPairs(root *TreeNode, distance int) (ans int) {
	g := make([][]int, 1024)
	isLeaf := make([]bool, 1024)
	cnt := 0
	leaves := []int{}
	var build func(o *TreeNode)
	build = func(o *TreeNode) {
		v := cnt
		if o.Left == nil && o.Right == nil {
			leaves = append(leaves, v)
			isLeaf[v] = true
		}
		if o.Left != nil {
			cnt++
			g[v] = append(g[v], cnt)
			g[cnt] = append(g[cnt], v)
			build(o.Left)
		}
		if o.Right != nil {
			cnt++
			g[v] = append(g[v], cnt)
			g[cnt] = append(g[cnt], v)
			build(o.Right)
		}
	}
	build(root)

	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		// 取等号是因为深度为 d 的情况已经在递归前检查了
		if d >= distance {
			return
		}
		for _, w := range g[v] {
			if w != fa {
				if isLeaf[w] {
					ans++
				}
				f(w, v, d+1)
			}
		}
	}
	for _, leaf := range leaves {
		f(leaf, -1, 0)
	}
	return ans / 2
}
