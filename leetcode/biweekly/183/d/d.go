package main

// https://space.bilibili.com/206214
func countValidSubsets(parent []int, nums []int, k int) int {
	const mod = 1_000_000_007
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	var dfs func(int) ([]int, []int)
	dfs = func(x int) ([]int, []int) {
		f0 := make([]int, k) // f0[i] 表示不选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
		f1 := make([]int, k) // f1[i] 表示选 x 时，子树 x 的子集点权和模 k 为 i 的方案数
		f0[0] = 1
		f1[nums[x]%k] = 1

		for _, y := range g[x] {
			fy0, fy1 := dfs(y)

			// 不选 x，那么 y 可选可不选
			nf0 := make([]int, k)
			for i := range k { // 枚举从子树 y 中选出的点权和模 k 为 i
				v := fy0[i] + fy1[i]
				if v == 0 { // 优化
					continue
				}
				for j, w := range f0 { // 枚举从之前的子树中选出的点权和模 k 为 j
					s := (i + j) % k
					nf0[s] = (nf0[s] + v*w) % mod
				}
			}

			// 选 x，那么 y 不能选
			nf1 := make([]int, k)
			for i, v := range fy0 { // 枚举从子树 y 中选出的点权和模 k 为 i
				if v == 0 { // 优化
					continue
				}
				for j, w := range f1 { // 枚举从 x 以及之前的子树中选出的点权和模 k 为 j
					s := (i + j) % k
					nf1[s] = (nf1[s] + v*w) % mod
				}
			}
			f0, f1 = nf0, nf1
		}

		return f0, f1
	}

	f0, f1 := dfs(0)
	// 恰好被 k 整除即模 k 为 0，注意减去空集的方案数 1
	return (f0[0] + f1[0] - 1 + mod) % mod
}
