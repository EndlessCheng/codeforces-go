package main

import "math/bits"

// https://space.bilibili.com/206214
func evenSumSubgraphs(nums []int, edges [][]int) (ans int) {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 枚举节点集合 U = {0,1,2,...,n-1} 的非空子集 sub
	u := 1<<n - 1
	for sub := 1; sub <= u; sub++ {
		// 计算子图的点权异或和
		xor := 0
		for i, x := range nums {
			if sub>>i&1 > 0 { // i 在 sub 中
				xor ^= x
			}
		}
		if xor != 0 {
			continue
		}

		// 判断子图是否连通
		vis := u ^ sub // 技巧：把不在子图中的节点都标记为已访问
		var dfs func(int)
		dfs = func(x int) {
			vis |= 1 << x // 标记 x 已访问
			for _, y := range g[x] {
				if vis>>y&1 == 0 { // y 没有访问过
					dfs(y)
				}
			}
		}
		dfs(bits.TrailingZeros(uint(sub))) // 随便选一个在子图中的节点，开始 DFS

		if vis == u { // 所有节点都已访问，子图是连通的
			ans++
		}
	}
	return
}
