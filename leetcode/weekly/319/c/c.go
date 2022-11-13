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

		b := append([]int{}, a...)
		sort.Ints(b)
		mp := make(map[int]int, n)
		for i, v := range b {
			mp[v] = i
		}

		ans += n
		vis := make([]bool, n)
		for _, v := range a {
			v = mp[v]
			if !vis[v] {
				for ; !vis[v]; v = mp[a[v]] {
					vis[v] = true
				}
				ans--
			}
		}
	}
	return
}
