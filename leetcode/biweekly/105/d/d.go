package main

// https://space.bilibili.com/206214
const mx int = 1e5

var lpf [mx + 1]int

func init() {
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func canTraverseAllPairs(a []int) (ans bool) {
	n := len(a)
	u := newUnionFind(n + mx)
	for i, x := range a {
		for x > 1 {
			p := lpf[x]
			for x /= p; lpf[x] == p; x /= p {
			}
			u.merge(i, n+p)
		}
	}
	for i := range a {
		if u.find(i) != u.find(0) {
			return false
		}
	}
	return true
}


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

func (u uf) merge(from, to int) {
	u.fa[u.find(from)] = u.find(to)
}
