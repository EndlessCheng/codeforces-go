package main

// https://space.bilibili.com/206214
type uf struct {
	fa []int
	cc int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa, n}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return
	}
	u.fa[x] = y
	u.cc--
}

func numberOfComponents(properties [][]int, k int) int {
	sets := make([]map[int]bool, len(properties))
	for i, a := range properties {
		sets[i] = map[int]bool{}
		for _, x := range a {
			sets[i][x] = true
		}
	}

	u := newUnionFind(len(properties))
	for i, a := range sets {
		for j, b := range sets[:i] {
			cnt := 0
			for x := range b {
				if a[x] {
					cnt++
				}
			}
			if cnt >= k {
				u.merge(i, j)
			}
		}
	}
	return u.cc
}
