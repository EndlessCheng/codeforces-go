package main

// https://space.bilibili.com/206214
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
type unionFind struct {
	fa  []int // fa[x] 是 x 的代表元
	dis []int // dis[x] = 从 x 到 fa[x] 的路径异或和
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, dis}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		root := u.find(u.fa[x])
		u.dis[x] ^= u.dis[u.fa[x]]
		u.fa[x] = root
	}
	return u.fa[x]
}

func (u unionFind) merge(from, to, value int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return u.dis[from]^u.dis[to] == value
	}
	u.dis[x] = value ^ u.dis[to] ^ u.dis[from]
	u.fa[x] = y
	return true
}

func numberOfEdgesAdded(n int, edges [][]int) (ans int) {
	uf := newUnionFind(n)
	for _, e := range edges {
		if uf.merge(e[0], e[1], e[2]) {
			ans++
		}
	}
	return
}
