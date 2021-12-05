package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"strings"
)

/* 一题双解：BFS / 最近公共祖先

#### 解法一：DFS + BFS

我们可以从起点出发，通过 BFS 找到终点，同时记录每个点的来源节点和方向，在找到终点后，顺着来源节点往回走，同时记录答案。

由于要往父节点方向走，我们需要先通过一次 DFS 记录每个节点的父节点，这样就可以在 BFS 中往父节点和左右节点三个方向前进了。

DFS 的过程中也可以顺带找到起点。

#### 解法二：最近公共祖先

按照 [236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/) 给出的方法，我们可以从起点出发，找到到起点和终点的路径，然后去掉前缀相同的部分。剩下即为从起点和终点的最近公共祖先出发，到起点和终点的路径，分别记作 $\textit{pathToStart}$ 和 $\textit{pathToDest}$。

我们要找的最短路径即为：起点 => 起点和终点的最近公共祖先 => 终点。

对于起点到最近公共祖先这一段，可以看成长度为 $\textit{pathToStart}$ 的向父节点走的路径；对于最近公共祖先到终点这一段就是 $\textit{pathToDest}$。将这两段路径拼起来即为答案。

*/

// github.com/EndlessCheng/codeforces-go
func getDirections(root *TreeNode, startValue, destValue int) string {
	var path []byte
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}
		if node.Val == target {
			return true
		}
		path = append(path, 'L')
		if dfs(node.Left, target) {
			return true
		}
		path[len(path)-1] = 'R'
		if dfs(node.Right, target) {
			return true
		}
		path = path[:len(path)-1]
		return false
	}
	dfs(root, startValue)
	pathToStart := path

	path = nil
	dfs(root, destValue)
	pathToDest := path

	for len(pathToStart) > 0 && len(pathToDest) > 0 && pathToStart[0] == pathToDest[0] {
		pathToStart = pathToStart[1:]
		pathToDest = pathToDest[1:]
	}

	return strings.Repeat("U", len(pathToStart)) + string(pathToDest)
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
