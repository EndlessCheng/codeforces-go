package main

import "sort"

// github.com/EndlessCheng/codeforces-go
const mx int = 1e5

var lpf [mx + 1]int

func init() { // 预处理每个数的最小质因子，方便后面分解
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

func gcdSort(a []int) bool {
	n := len(a)
	u := newUnionFind(n + mx)
	for i, x := range a {
		// 分解出 x 的质因子，时间复杂度 O(log x)
		for x > 1 {
			p := lpf[x]
			for x /= p; lpf[x] == p; x /= p {
			}
			u.merge(i, n+p) // 用并查集将下标 i 和质数 p 合并，为了区分下标集合和质数集合，将 p 加上 n
		}
	}

	// 整理处于同一集合的元素及其下标
	groups := make([]struct{ vs, is []int }, mx)
	for i, v := range a {
		p := u.find(i) - n
		groups[p].vs = append(groups[p].vs, v)
		groups[p].is = append(groups[p].is, i)
	}

	// 每一组内的元素可以直接或间接交换，因此对每组元素直接排序，然后按照在原数组的顺序填回去
	for _, g := range groups {
		sort.Ints(g.vs)
		for j, v := range g.vs {
			a[g.is[j]] = v
		}
	}
	return sort.IntsAreSorted(a) // 判断数组是否有序
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
