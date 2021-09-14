package main

// github.com/EndlessCheng/codeforces-go
type uf struct{ fa []int }

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}
func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}
func (u uf) merge(from, to int) { u.fa[u.find(from)] = u.find(to) }
func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }

func areConnected(n int, threshold int, queries [][]int) (ans []bool) {
	ans = make([]bool, len(queries))
	uf := newUnionFind(n + 1)
	// todo 线性筛 + 启发式合并
	vis := make([]bool, n+1)
	for i := threshold + 1; 2*i <= n; i++ {
		if vis[i] {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			vis[j] = true
			uf.merge(i, j)
		}
	}
	for i, q := range queries {
		ans[i] = uf.same(q[0], q[1])
	}
	return
}
