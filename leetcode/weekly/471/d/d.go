package main

// https://space.bilibili.com/206214

// 预处理平方剩余核
const mx = 100_001

var core = [mx]int{}

func init() {
	for i := 1; i < mx; i++ {
		if core[i] == 0 {
			for j := 1; i*j*j < mx; j++ {
				core[i*j*j] = i
			}
		}
	}
}

func sumOfAncestors(n int, edges [][]int, nums []int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := map[int]int{}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		c := core[nums[x]]
		// 本题 x 的祖先不包括 x 自己
		ans += int64(cnt[c])
		cnt[c]++
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}
		cnt[c]-- // 恢复现场
	}
	dfs(0, -1)
	return
}
