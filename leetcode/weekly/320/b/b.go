package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"slices"
)

// https://space.bilibili.com/206214
func closestNodes(root *TreeNode, queries []int) [][]int {
	a := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		a = append(a, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		mn, mx := -1, -1
		j, ok := slices.BinarySearch(a, q)
		if j < len(a) {
			mx = a[j]
		}
		if !ok {
			j--
		}
		if j >= 0 {
			mn = a[j]
		}
		ans[i] = []int{mn, mx}
	}
	return ans
}
