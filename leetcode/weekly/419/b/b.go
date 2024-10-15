package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"slices"
)

// https://space.bilibili.com/206214
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	cnt := [10]int{}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH == -2 || leftH != rightH {
			return -2 // 不合法
		}
		cnt[leftH+1]++
		return leftH + 1
	}
	dfs(root)

	for i := len(cnt) - 1; i >= 0; i-- {
		c := cnt[i]
		if c >= k {
			return 1<<(i+1) - 1
		}
		k -= c
	}
	return -1
}

func kthLargestPerfectSubtree2(root *TreeNode, k int) int {
	hs := []int{}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH < 0 || leftH != rightH {
			return -1 // 不合法
		}
		hs = append(hs, leftH+1)
		return leftH + 1
	}
	dfs(root)

	slices.Sort(hs)
	if k > len(hs) {
		return -1
	}
	return 1<<hs[len(hs)-k] - 1
}
