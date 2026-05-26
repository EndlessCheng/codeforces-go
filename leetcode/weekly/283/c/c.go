package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// github.com/EndlessCheng/codeforces-go
func createBinaryTree1(descriptions [][]int) *TreeNode {
	n := len(descriptions)
	nodes := make(map[int]*TreeNode, n+1) // 预分配空间
	children := make(map[int]bool, n)

	// 建树
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
		children[y] = true // y 不是根节点
	}

	for x, node := range nodes {
		if !children[x] { // node 是根节点
			return node
		}
	}

	// 测试用例保证可以构造出有效的二叉树
	panic("不是有效的二叉树")
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode, len(descriptions)+1) // 预分配空间
	root := 0

	for _, d := range descriptions {
		x, y := d[0], d[1]
		if nodes[x] == nil {
			nodes[x] = &TreeNode{Val: x}
			root ^= x
		}
		if nodes[y] == nil {
			nodes[y] = &TreeNode{Val: y}
			root ^= y
		}
		if d[2] == 1 {
			nodes[x].Left = nodes[y]
		} else {
			nodes[x].Right = nodes[y]
		}
		root ^= y
	}

	return nodes[root]
}
