package main

type uf struct {
	fa []int
}

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
func (u uf) merge(from, to int) (isNewMerge bool) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u.fa[x] = y
	return true
}
func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }
func (u uf) countRoots(st int) (cnt int) { // st = 0 or 1 ...
	for i := st; i < len(u.fa); i++ {
		if u.find(i) == i {
			cnt++
		}
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func maxNumEdgesToRemove(n int, es [][]int) (ans int) {
	a, b := newUnionFind(n+1), newUnionFind(n+1)
	ans = len(es)
	for _, e := range es {
		if tp, v, w := e[0], e[1], e[2]; tp == 3 && (!a.same(v, w) || !b.same(v, w)) {
			a.merge(v, w)
			b.merge(v, w)
			ans--
		}
	}
	us := [2]uf{a, b}
	for _, e := range es {
		if tp := e[0]; tp < 3 && us[tp-1].merge(e[1], e[2]) {
			ans--
		}
	}
	if a.countRoots(1) > 1 || b.countRoots(1) > 1 {
		return -1
	}
	return
}
