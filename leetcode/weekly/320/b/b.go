package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

// https://space.bilibili.com/206214
func closestNodes(root *TreeNode, queries []int) [][]int {
	a := []int{}
	var dfs func(*TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}
		dfs(o.Left)
		a = append(a, o.Val)
		dfs(o.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		min, max := -1, -1
		// 这是怎么转换的，可以看我上面贴的视频链接
		j := sort.SearchInts(a, q+1) - 1
		if j >= 0 {
			min = a[j]
		}
		j = sort.SearchInts(a, q)
		if j < len(a) {
			max = a[j]
		}
		ans[i] = []int{min, max}
	}
	return ans
}
