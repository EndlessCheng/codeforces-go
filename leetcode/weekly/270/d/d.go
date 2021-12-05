package main

/* 有向图欧拉路径

本题与 [332. 重新安排行程](https://leetcode-cn.com/problems/reconstruct-itinerary/) 一样，都可以抽象成如下问题：

> 给你一张有向图，在这张有向图中找到一条欧拉路径。

对于此题，由于题目保证有解，所以有向图肯定是欧拉图或半欧拉图。

我们需要找到一个起点，如果是欧拉图，那么随便找一个点都行，如果是半欧拉图，那么找到一个出度比入度大 1 的点，把它当成起点，然后按照 332 给出的算法来求解。

*/

// github.com/EndlessCheng/codeforces-go
func validArrangement(pairs [][]int) [][]int {
	type edge struct{ to, idx int }
	g := map[int][]edge{}
	inDeg := map[int]int{}
	for i, p := range pairs {
		v, w := p[0], p[1]
		g[v] = append(g[v], edge{w, i}) // 建图（有向图）
		inDeg[w]++ // 统计入度
	}

	start := 0
	for i, es := range g {
		if len(es) == inDeg[i]+1 { // 如果存在出度比入度大 1 的点，那就把它当成起点
			start = i
			break
		}
		start = i // 或者随便找一个点当起点
	}

	m := len(pairs)
	ans := make([][]int, 0, m)
	var dfs func(int)
	dfs = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			dfs(e.to)
			ans = append(ans, pairs[e.idx]) // 保存所有边
		}
	}
	dfs(start)
	for i := 0; i < m/2; i++ {
		ans[i], ans[m-1-i] = ans[m-1-i], ans[i]
	}
	return ans
}
