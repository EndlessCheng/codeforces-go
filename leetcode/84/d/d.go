package main

// github.com/EndlessCheng/codeforces-go
func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	sz := make([]int, n)
	var f func(int, int) int
	f = func(v, fa int) (sum int) { // sum 表示以 0 为根的子树 v 中的节点到 v 的距离之和
		sz[v] = 1
		for _, w := range g[v] {
			if w != fa {
				sum += f(w, v) + sz[w] // 子树 w 的每个节点都要经过 w-v，因此这条边对 sum 产生的贡献为 sz[w]
				sz[v] += sz[w]
			}
		}
		return
	}
	sum0 := f(0, -1)

	ans := make([]int, n)
	var reroot func(v, fa, sum int)
	reroot = func(v, fa, sum int) {
		ans[v] = sum
		for _, w := range g[v] {
			if w != fa {
				// 离子树 w 中的所有节点近了 1，又离不在子树 w 中的节点远了 1
				// 所以要减去 sz[w]，并加上 n-sz[w]
				reroot(w, v, sum+n-sz[w]*2)
			}
		}
	}
	reroot(0, -1, sum0)
	return ans
}
