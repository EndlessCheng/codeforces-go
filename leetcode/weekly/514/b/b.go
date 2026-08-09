package main

// https://space.bilibili.com/206214
func weightedSum1(parent []int, nums []int) (ans int64) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	var getH func(int) int
	getH = func(x int) (h int) {
		for _, y := range g[x] {
			h = max(h, getH(y))
		}
		return h + 1
	}
	h := getH(0)

	var dfs func(int, int)
	dfs = func(x, weight int) {
		ans += int64(nums[x]) * int64(weight)
		for _, y := range g[x] {
			dfs(y, weight-1)
		}
	}
	dfs(0, h)
	return
}

func weightedSum(parent []int, nums []int) (ans int64) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	depth := make([]int, n)
	var dfs func(int) int
	dfs = func(x int) (h int) {
		for _, y := range g[x] {
			depth[y] = depth[x] + 1
			h = max(h, dfs(y))
		}
		return h + 1
	}
	h := dfs(0)

	for i, x := range nums {
		ans += int64(x) * int64(h-depth[i])
	}
	return
}
