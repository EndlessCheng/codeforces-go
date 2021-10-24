package main

/* 拓扑排序 + 动态规划

定义 $f[i]$ 表示完成第 $i$ 门课程需要花费的最少月份数。根据题意，只有当 $i$ 的所有先修课程都完成时，才可以开始 $i$，并且我们可以立即开始 $i$。

因此 $f[i]=\textit{time}[i] + \max f[j]$，这里 $j$ 是 $i$ 的所有先修课程。

由于题目保证图是一个有向无环图，所以一定存在拓扑排序。我们可以在求拓扑排序的同时，计算状态转移。

代码实现时，设当前节点为 $v$，我们可以在计算出 $f[v]$ 后，更新 $f[w]$ 的最大值，这里 $v$ 是 $w$ 的先修课程。

答案就是 $\max f[i]$。

相似题目：

- [1857. 有向图中最大颜色值](https://leetcode-cn.com/problems/largest-color-value-in-a-directed-graph/)

 */

// github.com/EndlessCheng/codeforces-go
func minimumTime(n int, relations [][]int, time []int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range relations {
		v, w := e[0]-1, e[1]-1
		g[v] = append(g[v], w)
		deg[w]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	f := make([]int, n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		f[v] += time[v]
		ans = max(ans, f[v])
		for _, w := range g[v] {
			f[w] = max(f[w], f[v])
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
