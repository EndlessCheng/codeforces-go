package main

import (
	"math"
	"slices"
	"sort"
)

func kthSmallest(par []int, vals []int, queries [][]int) []int {
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	a := make([]int, n)
	ranges := make([]struct{ l, r int }, n) // 左闭右开 [l,r)
	dfn := 0
	var dfs func(int, int)
	dfs = func(x, xor int) {
		ranges[x].l = dfn
		xor ^= vals[x]
		a[dfn] = xor
		dfn++
		for _, y := range g[x] {
			dfs(y, xor)
		}
		ranges[x].r = dfn
	}
	dfs(0, 0)

	// 排序去重
	b := slices.Clone(a)
	slices.Sort(b)
	b = slices.Compact(b)
	// 离散化
	for i, v := range a {
		a[i] = sort.SearchInts(b, v) // 从 0 开始
	}

	nq := len(queries)
	qBlockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(nq*2))))
	type query struct{ bid, l, r, k, qid int } // 左闭右开 [l,r)
	qs := make([]query, nq)
	for i, q := range queries {
		p := ranges[q[0]]
		qs[i] = query{p.l / qBlockSize, p.l, p.r, q[1], i}
	}
	slices.SortFunc(qs, func(a, b query) int {
		if a.bid != b.bid {
			return a.bid - b.bid
		}
		// 奇偶化排序
		if a.bid%2 == 0 {
			return a.r - b.r
		}
		return b.r - a.r
	})

	m := len(b)
	cBlockSize := int(math.Sqrt(float64(m)))
	blockUniqueCnt := make([]int, (m-1)/cBlockSize+1)
	cnt := make([]int, m+1)
	move := func(i, delta int) {
		v := a[i]
		if delta > 0 {
			if cnt[v] == 0 {
				blockUniqueCnt[v/cBlockSize]++
			}
			cnt[v]++
		} else {
			cnt[v]--
			if cnt[v] == 0 {
				blockUniqueCnt[v/cBlockSize]--
			}
		}
	}

	ans := make([]int, len(qs))
	l, r := 0, 0
	for _, q := range qs {
		for ; l < q.l; l++ {
			move(l, -1)
		}
		for l > q.l {
			l--
			move(l, 1)
		}
		for ; r < q.r; r++ {
			move(r, 1)
		}
		for r > q.r {
			r--
			move(r, -1)
		}

		k := q.k
		for i, c := range blockUniqueCnt {
			if k <= c {
				for j := i * cBlockSize; ; j++ {
					if cnt[j] == 0 {
						continue
					}
					k--
					if k == 0 {
						ans[q.qid] = b[j]
						break
					}
				}
				break
			}
			k -= c
		}
		if k > 0 {
			ans[q.qid] = -1
		}
	}
	return ans
}
