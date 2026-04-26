package main

import "math/bits"

// https://space.bilibili.com/206214
func evenSumSubgraphs(nums []int, edges [][]int) (ans int) {
	n := len(nums)
	g := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] |= 1 << y
		g[y] |= 1 << x
	}

	ones := 0
	for i, x := range nums {
		ones |= x << i
	}

	// 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
	u := 1<<n - 1
	for sub := 1; sub <= u; sub++ {
		// 计算子图的点权和
		sum := bits.OnesCount(uint(sub & ones))
		if sum%2 != 0 {
			continue
		}

		// 判断子图是否连通
		vis := u ^ sub  // 技巧：把不在子图中的节点都标记为已访问
		q := sub & -sub // 随便选一个在子图中的节点，开始 BFS
		vis |= q
		for q > 0 {
			x := q & -q // 出队
			q ^= x
			to := g[bits.TrailingZeros(uint(x))] &^ vis // 访问 x 的（尚未访问过的）邻居
			vis |= to
			q |= to // x 的邻居入队
		}

		if vis == u { // 所有节点都已访问，子图是连通的
			ans++
		}
	}
	return
}
