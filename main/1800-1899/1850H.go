package main

import (
	. "fmt"
	"io"
)

// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
type unionFind50 struct {
	fa  []int // 代表元
	dis []int // dis[x] 表示 x 到（x 所在集合的）代表元的距离
}

func newUnionFind50(n int) unionFind50 {
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己，自己到自己的距离是 0
	fa := make([]int, n)
	dis := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind50{fa, dis}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩
func (u unionFind50) find(x int) int {
	if u.fa[x] != x {
		root := u.find(u.fa[x])
		u.dis[x] += u.dis[u.fa[x]] // 递归更新 x 到其代表元的距离
		u.fa[x] = root
	}
	return u.fa[x]
}

// 判断 x 和 y 是否在同一个集合（同普通并查集）
func (u unionFind50) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

// 计算从 from 到 to 的相对距离
// 调用时需保证 from 和 to 在同一个集合中，否则返回值无意义
func (u unionFind50) getRelativeDistance(from, to int) int {
	u.find(from)
	u.find(to)
	// to-from = (x-from) - (x-to) = dis[from] - dis[to]
	return u.dis[from] - u.dis[to]
}

// 合并 from 和 to，新增信息 to - from = value
// 其中 to 和 from 表示未知量，下文的 x 和 y 也表示未知量
// 如果 from 和 to 不在同一个集合，返回 true，否则返回是否与已知信息矛盾
func (u unionFind50) merge(from, to int, value int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		// to-from = (x-from) - (x-to) = dis[from] - dis[to] = value
		return u.dis[from]-u.dis[to] == value
	}
	//    x --------- y
	//   /           /
	// from ------- to
	// 已知 x-from = dis[from] 和 y-to = dis[to]，现在合并 from 和 to，新增信息 to-from = value
	// 由于 y-from = (y-x) + (x-from) = (y-to) + (to-from)
	// 所以 y-x = (to-from) + (y-to) - (x-from) = value + dis[to] - dis[from]
	u.dis[x] = value + u.dis[to] - u.dis[from] // 计算 x 到其代表元 y 的距离
	u.fa[x] = y
	return true
}

func cf1850H(in io.Reader, out io.Writer) {
	var T, n, m, a, b, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		uf := newUnionFind50(n + 1)
		ok := true
		for range m {
			Fscan(in, &a, &b, &d)
			ok = ok && uf.merge(b, a, d)
		}
		if ok {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1850H(bufio.NewReader(os.Stdin), os.Stdout) }
