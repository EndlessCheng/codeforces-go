package main

import (
	"maps"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func findAllPeople1(_ int, meetings [][]int, firstPerson int) []int {
	// 按照 time 从小到大排序
	slices.SortFunc(meetings, func(a, b []int) int { return a[2] - b[2] })

	// 一开始 0 和 firstPerson 都知道秘密
	haveSecret := map[int]bool{0: true, firstPerson: true}

	// 分组循环
	m := len(meetings)
	for i := 0; i < m; {
		// 在同一时间发生的会议，建图
		g := map[int][]int{}
		time := meetings[i][2]
		for ; i < m && meetings[i][2] == time; i++ {
			x, y := meetings[i][0], meetings[i][1]
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
		}

		// 每个连通块只要有一个人知道秘密，那么整个连通块的人都知道秘密
		vis := map[int]bool{} // 避免重复访问节点
		var dfs func(int)
		dfs = func(x int) {
			vis[x] = true
			haveSecret[x] = true
			for _, y := range g[x] {
				if !vis[y] {
					dfs(y)
				}
			}
		}
		for x := range g { // 遍历在 time 时间点参加会议的专家
			if haveSecret[x] && !vis[x] { // 从知道秘密的专家出发，DFS 标记其余专家
				dfs(x)
			}
		}
	}

	// 可以按任何顺序返回答案
	return slices.Collect(maps.Keys(haveSecret))
}

// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
type unionFind struct {
	fa []int // 代表元
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 判断 x 和 y 是否在同一个集合
func (u unionFind) same(x, y int) bool {
	// 如果 x 的代表元和 y 的代表元相同，那么 x 和 y 就在同一个集合
	// 这就是代表元的作用：用来快速判断两个元素是否在同一个集合
	return u.find(x) == u.find(y)
}

// 把 from 所在集合合并到 to 所在集合中
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	u.fa[x] = y
}

func findAllPeople(n int, meetings [][]int, firstPerson int) (ans []int) {
	// 按照 time 从小到大排序
	slices.SortFunc(meetings, func(a, b []int) int { return a[2] - b[2] })

	uf := newUnionFind(n)
	// 一开始 0 和 firstPerson 都知道秘密
	uf.merge(firstPerson, 0)

	// 分组循环
	m := len(meetings)
	for i := 0; i < m; {
		start := i
		// 合并在同一时间发生的会议
		time := meetings[i][2]
		for ; i < m && meetings[i][2] == time; i++ {
			uf.merge(meetings[i][0], meetings[i][1])
		}

		// 如果节点不和 0 在同一个集合，那么撤销合并，恢复成初始值
		for j := start; j < i; j++ {
			x, y := meetings[j][0], meetings[j][1]
			if !uf.same(x, 0) {
				uf.fa[x] = x
			}
			if !uf.same(y, 0) {
				uf.fa[y] = y
			}
		}
	}

	// 和 0 在同一个集合的专家都知道秘密
	for i := range n {
		if uf.same(i, 0) {
			ans = append(ans, i)
		}
	}
	return
}
