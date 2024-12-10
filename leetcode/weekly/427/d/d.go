package main

import (
	"cmp"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea(xCoord, ys []int) int64 {
	type pair struct{ x, y int }
	points := make([]pair, len(xCoord))
	for i := range xCoord {
		points[i] = pair{xCoord[i], ys[i]}
	}
	slices.SortFunc(points, func(a, b pair) int { return cmp.Or(a.x-b.x, a.y-b.y) })

	// 离散化用
	slices.Sort(ys)
	ys = slices.Compact(ys)

	ans := -1
	tree := make(fenwick, len(ys)+1)
	tree.add(sort.SearchInts(ys, points[0].y) + 1) // 离散化
	type tuple struct{ x, y, c int }
	pre := make([]tuple, len(ys))
	for i := 1; i < len(points); i++ {
		x1, y1 := points[i-1].x, points[i-1].y
		x2, y2 := points[i].x, points[i].y
		y := sort.SearchInts(ys, y2) // 离散化
		tree.add(y + 1)
		if x1 != x2 { // 两点不在同一列
			continue
		}
		cur := tree.query(sort.SearchInts(ys, y1)+1, y+1)
		p := pre[y]
		if p.c > 0 && p.c+2 == cur && p.y == y1 {
			ans = max(ans, (x2-p.x)*(y2-y1))
		}
		pre[y] = tuple{x1, y1, cur}
	}
	return int64(ans)
}
