package main

import (
	"bufio"
	. "fmt"
	"os"
)

var fa []int

func initFa(n int) {
	fa = make([]int, n)
	for i := range fa {
		fa[i] = i
	}
}
func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}
func merge(from, to int) { fa[find(from)] = find(to) }
func same(x, y int) bool { return find(x) == find(y) }

// 题解 https://www.luogu.org/blog/endlesscheng/solution-cf1242b
func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	// 寻找 0-边最大的点 maxDeg0V
	maxDeg0, maxDeg0V := 0, 0
	for v, edges := range g {
		if deg0 := n - 1 - len(edges); deg0 > maxDeg0 {
			maxDeg0 = deg0
			maxDeg0V = v
		}
	}

	// 若图中没有 0-边，答案就是点的个数-1
	if maxDeg0 == 0 {
		Print(n - 1)
		return
	}

	mergeEdge0 := func(v int, edges []int) {
		// 将与点 v 以 0-边相连的点，合并到点 v 所属的连通分量上
		vs := map[int]bool{v: true}
		for _, w := range edges {
			vs[w] = true
		}
		for i := 0; i < n; i++ {
			if !vs[i] { // i-v 是 0-边
				merge(i, v)
			}
		}
	}
	initFa(n)
	mergeEdge0(maxDeg0V, g[maxDeg0V])
	for v, edges := range g {
		if !same(v, maxDeg0V) {
			// 暴力遍历剩余的点
			mergeEdge0(v, edges)
		}
	}

	// 计算联通分量个数-1
	ans := -1
	for i, faI := range fa {
		if i == faI {
			ans++
		}
	}
	Print(ans)
}
