package main

import "container/heap"

/* 暴搜+剪枝

根据题目的数据范围，至多只能走 $10$ 条边，这意味着爆搜的层数至多为 $10$；同时，题目保证每个节点至多有四条边与之相连，因此每次搜索时至多会递归 $4$ 次。因此计算量至多为 $4^{10}$，可以在时限内跑完。

本题的一个剪枝技巧是，先预处理起点 $0$ 到其余节点的最短路，在搜索时提前判断下一个节点在走最短路的前提下能否在 $\textit{maxTime}$ 时间内回到起点 $0$，若不能则不进行递归。

 */

// github.com/EndlessCheng/codeforces-go
func maximalPathQuality(values []int, edges [][]int, maxTime int) (ans int) {
	n := len(values)
	g := make([][]edge, n)
	for _, e := range edges {
		v, w, t := e[0], e[1], e[2]
		g[v] = append(g[v], edge{w, t}) // 建图
		g[w] = append(g[w], edge{v, t})
	}

	dis := dijkstra(g, 0) // 预处理从起点 0 到每个点的最短路

	vis := make([]bool, n)
	sum := 0
	var dfs func(int, int)
	dfs = func(v, time int) {
		if !vis[v] { // 没有访问时，更新价值之和
			vis[v] = true
			sum += values[v]
			if sum > ans {
				ans = sum // 更新答案
			}
			defer func() {
				sum -= values[v] // 恢复现场
				vis[v] = false
			}()
		}
		for _, e := range g[v] {
			if time+e.t+dis[e.to] <= maxTime { // 剪枝：下个节点在走最短路的情况下可以在 maxTime 时间内返回起点 0
				dfs(e.to, time+e.t)
			}
		}
	}
	dfs(0, 0)
	return
}

// 下面是求最短路的模板
type edge struct{ to, t int }
type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }

func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = 1e9
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		vd := h.pop()
		v := vd.v
		if dis[v] < vd.dis {
			continue
		}
		for _, e := range g[v] {
			w, wt := e.to, e.t
			if newD := dis[v] + wt; newD < dis[w] {
				dis[w] = newD
				h.push(pair{w, dis[w]})
			}
		}
	}
	return dis
}
