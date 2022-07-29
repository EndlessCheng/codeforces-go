package main

import "math"

// https://space.bilibili.com/206214/dynamic
func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	xor := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	clock := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		clock++
		in[x] = clock
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = clock
	}
	dfs(0, -1)

	ans := math.MaxInt32
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			var x, y, z int
			if in[i] < in[j] && in[j] <= out[i] { // i 是 j 的祖先节点
				x, y, z = xor[j], xor[i]^xor[j], xor[0]^xor[i]
			} else if in[j] < in[i] && in[i] <= out[j] { // j 是 i 的祖先节点
				x, y, z = xor[i], xor[i]^xor[j], xor[0]^xor[j]
			} else { // 删除的两条边分别属于两颗不相交的子树
				x, y, z = xor[i], xor[j], xor[0]^xor[i]^xor[j]
			}
			ans = min(ans, max(max(x, y), z)-min(min(x, y), z))
			if ans == 0 {
				return 0 // 提前退出
			}
		}
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
