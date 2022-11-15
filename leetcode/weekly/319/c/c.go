package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

// https://space.bilibili.com/206214
func minimumOperations(root *TreeNode) (ans int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		a := make([]int, n)
		tmp := q
		q = nil
		for i, node := range tmp {
			a[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		id := make([]int, n)
		for i := range id {
			id[i] = i
		}
		sort.Slice(id, func(i, j int) bool { return a[id[i]] < a[id[j]] })

		ans += n
		vis := make([]bool, n)
		for _, v := range id {
			if !vis[v] {
				for ; !vis[v]; v = id[v] {
					vis[v] = true
				}
				ans--
			}
		}
	}
	return
}
