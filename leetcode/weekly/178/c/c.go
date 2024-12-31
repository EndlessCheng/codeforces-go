package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"math/rand"
)

// https://space.bilibili.com/206214
func isSubPath(head *ListNode, root *TreeNode) bool {
	// mod 和 base 随机其中一个就行，无需两个都随机
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack

	n := 0        // 链表长度
	powBase := 1  // base^(n-1)
	listHash := 0 // 多项式哈希 s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	for ; head != nil; head = head.Next {
		n++
		if n > 1 {
			powBase = powBase * base % mod
		}
		listHash = (listHash*base + head.Val) % mod // 秦九韶算法计算多项式哈希
	}

	st := []int{}
	var dfs func(*TreeNode, int) bool
	dfs = func(t *TreeNode, hash int) bool {
		if t == nil { // 无法继续匹配
			return false
		}
		st = append(st, t.Val)
		hash = (hash*base + t.Val) % mod // 移入窗口
		if len(st) >= n {
			if hash == listHash {
				return true
			}
			hash = (hash - powBase*st[len(st)-n]%mod + mod) % mod // 移出窗口
		}
		defer func() { st = st[:len(st)-1] }() // 恢复现场
		return dfs(t.Left, hash) || dfs(t.Right, hash)
	}
	return dfs(root, 0)
}

func isSubPath2(head *ListNode, root *TreeNode) bool {
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(s *ListNode, t *TreeNode) bool {
		if s == nil { // 整个链表匹配完毕
			return true
		}
		// 否则需要继续匹配
		if t == nil { // 无法继续匹配
			return false
		}
		// 节点值相同则继续匹配，否则从 head 开始重新匹配
		return s.Val == t.Val && (dfs(s.Next, t.Left) || dfs(s.Next, t.Right)) ||
			s == head && (dfs(head, t.Left) || dfs(head, t.Right))
	}
	return dfs(head, root)
}
