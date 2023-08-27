package main

import "math/bits"

// https://space.bilibili.com/206214
func getMaxFunctionValue(receiver []int, K int64) int64 {
	type pair struct{ pa, sum int }
	n, m := len(receiver), bits.Len(uint(K))
	pa := make([][]pair, n)
	for i, p := range receiver {
		pa[i] = make([]pair, m)
		pa[i][0] = pair{p, p}
	}
	for i := 0; i+1 < m; i++ {
		for x := range pa {
			p := pa[x][i]
			pp := pa[p.pa][i]
			pa[x][i+1] = pair{pp.pa, p.sum + pp.sum}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		x := i
		sum := i
		for k := uint(K); k > 0; k &= k - 1 {
			p := pa[x][bits.TrailingZeros(k)]
			sum += p.sum
			x = p.pa
		}
		ans = max(ans, sum)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }

//func getMaxFunctionValue2(g []int, k int64) int64 {
//	ans := 0
//	n := len(g)
//	rg := make([][]int, n) // g 的反图（外向基环树）
//	deg := make([]int, n)  // g 上每个节点的入度
//	for v, w := range g {
//		rg[w] = append(rg[w], v)
//		deg[w]++
//	}
//
//	// 拓扑排序，剪掉 g 上的所有树枝
//	// 拓扑排序后 deg 值为 1 的点必定在基环上，为 0 的点必定在树枝上
//	// 注：拓扑排序时还可以做 DP，比如给树枝上的每个点标记反向深度
//	q := []int{}
//	for i, d := range deg {
//		if d == 0 {
//			q = append(q, i)
//		}
//	}
//	//maxD := make([]int, n)
//	for len(q) > 0 {
//		v := q[0]
//		q = q[1:]
//		//maxD[v]++
//		w := g[v] // v 只有一条出边
//		//maxD[w] = max(maxD[w], maxD[v])
//		if deg[w]--; deg[w] == 0 {
//			q = append(q, w)
//		}
//	}
//
//	for i, d := range deg {
//		if d <= 0 {
//			continue
//		}
//		// 遍历基环上的点（拓扑排序后入度大于 0）
//		//m := 0
//		ring := []int{}
//		for v := i; ; v = g[v] {
//			deg[v] = -1 // 将基环上的点的入度标记为 -1，避免重复访问
//			ring = append(ring, v)
//			//m = max(m, maxD[v])
//			if g[v] == i {
//				break
//			}
//		}
//		// do ring ...
//		// 特别注意基环大小小于 3 的特殊情况
//
//		na := len(ring)
//		sum := make([]int64, na+1)
//		for i, v := range ring {
//			sum[i+1] = sum[i] + int64(v)
//		}
//		pre := func(p int) int64 {
//			return sum[na]*int64(p/na) + sum[p%na]
//		}
//		query := func(l, r int) int64 {
//			return pre(r) - pre(l)
//		}
//
//		// 在反图上遍历树枝
//		path := []int{}
//		sp := 0
//		var rdfs func(int, int)
//		rdfs = func(v, d int) {
//			path = append(path, v)
//			sp += v
//			for _, w := range rg[v] {
//				if deg[w] == 0 { // 树枝上的点在拓扑排序后，入度均为 0
//					rdfs(w, d+1)
//				}
//			}
//			sp -= v
//			path = path[:len(path)-1]
//		}
//		for _, v := range ring {
//			rdfs(v, 0)
//		}
//	}
//
//	return int64(ans)
//}
