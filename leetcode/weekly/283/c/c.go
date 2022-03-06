package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	in := map[int]bool{}
	for _, d := range descriptions {
		x, y := d[0], d[1]
		if nodes[x] == nil {
			nodes[x] = &TreeNode{Val: x}
		}
		if nodes[y] == nil {
			nodes[y] = &TreeNode{Val: y}
		}
		if d[2] == 1 {
			nodes[x].Left = nodes[y]
		} else {
			nodes[x].Right = nodes[y]
		}
		in[y] = true
	}

	for i, node := range nodes {
		if !in[i] { // 入度为 0 则为根
			return node
		}
	}
	return nil
}
