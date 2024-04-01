package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

// https://space.bilibili.com/206214
func minimumDistance(points [][]int) int {
	const inf = math.MaxInt
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf
	var maxXi, minXi, maxYi, minYi int

	for i, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2 = maxX1
			maxX1 = x
			maxXi = i
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2 = minX1
			minX1 = x
			minXi = i
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2 = maxY1
			maxY1 = y
			maxYi = i
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2 = minY1
			minY1 = y
			minYi = i
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, i := range []int{maxXi, minXi, maxYi, minYi} {
		dx := f(i != maxXi, maxX1, maxX2) - f(i != minXi, minX1, minX2)
		dy := f(i != maxYi, maxY1, maxY2) - f(i != minYi, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(b bool, v1, v2 int) int {
	if b {
		return v1
	}
	return v2
}

func minimumDistance2(points [][]int) int {
	xs := redblacktree.New[int, int]()
	ys := redblacktree.New[int, int]()
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		put(xs, x)
		put(ys, y)
	}
	ans := math.MaxInt
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		remove(xs, x)
		remove(ys, y)
		ans = min(ans, max(xs.Right().Key-xs.Left().Key, ys.Right().Key-ys.Left().Key))
		put(xs, x)
		put(ys, y)
	}
	return ans
}

func put(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	t.Put(v, c+1)
}

func remove(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	if c == 1 {
		t.Remove(v)
	} else {
		t.Put(v, c-1)
	}
}
