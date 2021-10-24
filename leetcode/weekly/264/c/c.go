package main

/* 两次 DFS

首先，我们以 $0$ 号节点为根节点，跑一次 DFS，求出所有子树的大小，记在 $\textit{size}$ 数组中。

对于一个点 $v$，删除与 $v$ 相连的边，剩余部分可以分为两类：

- 以 $v$ 的子节点为根的子树；
- 整棵树去掉以 $v$ 为根的子树后，剩余的部分。

我们可以从 $0$ 出发，对整棵树再跑一次 DFS。DFS 除了传入当前节点外，还要传入上述第二类的大小。然后就可以直接计算出当前节点的分数了。

*/

// github.com/EndlessCheng/codeforces-go
func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := parents[w]
		g[v] = append(g[v], w)
	}

	size := make([]int, n)
	var initSize func(int)
	initSize = func(v int) {
		size[v]++
		for _, w := range g[v] {
			initSize(w)
			size[v] += size[w]
		}
	}
	initSize(0)

	maxMul := 0
	var dfs func(int, int)
	dfs = func(v, otherSize int) {
		mul := otherSize
		for _, w := range g[v] {
			mul *= size[w] // 由于是二叉树所以 mul 最大约为 (1e5/3)^3，在 64 位整数范围内
		}
		if mul > maxMul {
			maxMul, ans = mul, 1
		} else if mul == maxMul {
			ans++
		}
		for _, w := range g[v] {
			dfs(w, n-size[w])
		}
	}
	dfs(0, 1)
	return
}
