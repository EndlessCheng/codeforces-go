package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func amountOfTime(root *TreeNode, start int) int {
	var st *TreeNode
	parents := map[*TreeNode]*TreeNode{}
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node, pa *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == start {
			st = node
		}
		parents[node] = pa
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	ans := -1
	vis := map[*TreeNode]bool{nil: true, st: true}
	for q := []*TreeNode{st}; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, node := range tmp {
			if node != nil {
				if !vis[node.Left] {
					vis[node.Left] = true
					q = append(q, node.Left)
				}
				if !vis[node.Right] {
					vis[node.Right] = true
					q = append(q, node.Right)
				}
				if p := parents[node]; !vis[p] {
					vis[p] = true
					q = append(q, p)
				}
			}
		}
	}
	return ans
}
