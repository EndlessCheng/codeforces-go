package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func getDirections(root *TreeNode, startValue, destValue int) string {
	var pathToStart, pathToDest string
	path := []byte{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val == startValue {
			pathToStart = string(path)
		} else if node.Val == destValue {
			pathToDest = string(path)
		}

		// 往左找
		path = append(path, 'L')
		dfs(node.Left)

		// 往右找
		path[len(path)-1] = 'R' // 直接覆盖，无需恢复现场
		dfs(node.Right)
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(root)

	i := 0
	for i < min(len(pathToStart), len(pathToDest)) && pathToStart[i] == pathToDest[i] {
		i++
	}

	return strings.Repeat("U", len(pathToStart)-i) + pathToDest[i:]
}

func getDirections1(root *TreeNode, startValue, destValue int) string {
	path := []byte{}
	// dfs 返回是否找到了 target
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}

		// 找到 target
		if node.Val == target {
			return true
		}

		// 往左找 target
		path = append(path, 'L')
		if dfs(node.Left, target) {
			return true
		}

		// 往右找 target
		path[len(path)-1] = 'R' // 直接覆盖，无需恢复现场
		if dfs(node.Right, target) {
			return true
		}

		path = path[:len(path)-1] // 恢复现场
		return false
	}

	dfs(root, startValue)
	pathToStart := path

	path = []byte{}
	dfs(root, destValue)
	pathToDest := path

	i := 0
	for i < min(len(pathToStart), len(pathToDest)) && pathToStart[i] == pathToDest[i] {
		i++
	}

	return strings.Repeat("U", len(pathToStart)-i) + string(pathToDest[i:])
}

func getDirections2(root *TreeNode, startValue, destValue int) string {
	q := []*TreeNode{nil}
	parents := map[*TreeNode]*TreeNode{}
	var dfs func(node, pa *TreeNode)
	dfs = func(node, pa *TreeNode) {
		if node == nil {
			return
		}
		parents[node] = pa
		if node.Val == startValue {
			q[0] = node // 只有一个起点
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	ans := []byte{}
	vis := map[*TreeNode]bool{nil: true, q[0]: true}
	type pair struct {
		from *TreeNode
		dir  byte
	}
	from := map[*TreeNode]pair{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Val == destValue {
			for ; from[node].from != nil; node = from[node].from {
				ans = append(ans, from[node].dir)
			}
			break
		}
		if !vis[node.Left] {
			vis[node.Left] = true
			from[node.Left] = pair{node, 'L'}
			q = append(q, node.Left)
		}
		if !vis[node.Right] {
			vis[node.Right] = true
			from[node.Right] = pair{node, 'R'}
			q = append(q, node.Right)
		}
		if !vis[parents[node]] {
			vis[parents[node]] = true
			from[parents[node]] = pair{node, 'U'}
			q = append(q, parents[node])
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
