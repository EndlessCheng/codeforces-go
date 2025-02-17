package main

import (
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214

// 线段树每个节点维护一段横坐标区间 [lx, rx]
type seg []struct {
	l, r        int
	minCoverLen int // 区间内被矩形覆盖次数最少的底边长之和
	minCover    int // 区间内被矩形覆盖的最小次数
	todo        int // 子树内的所有节点的 minCover 需要增加的量，注意这可以是负数
}

// 根据左右儿子的信息，更新当前节点的信息
func (t seg) maintain(o int) {
	lo, ro := &t[o<<1], &t[o<<1|1]
	mn := min(lo.minCover, ro.minCover)
	t[o].minCover = mn
	t[o].minCoverLen = 0
	if lo.minCover == mn { // 只统计等于 minCover 的底边长之和
		t[o].minCoverLen = lo.minCoverLen
	}
	if ro.minCover == mn {
		t[o].minCoverLen += ro.minCoverLen
	}
}

// 仅更新节点信息，不下传懒标记 todo
func (t seg) do(o, v int) {
	t[o].minCover += v
	t[o].todo += v
}

// 下传懒标记 todo
func (t seg) spread(o int) {
	v := t[o].todo
	if v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

// 建树
func (t seg) build(xs []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = 0
	if l == r {
		t[o].minCover = 0
		t[o].minCoverLen = xs[l+1] - xs[l]
		return
	}
	m := (l + r) >> 1
	t.build(xs, o<<1, l, m)
	t.build(xs, o<<1|1, m+1, r)
	t.maintain(o)
}

// 区间更新
func (t seg) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

// 代码逻辑同 850. 矩形面积 II，增加一个 records 数组记录关键数据
func separateSquares(squares [][]int) float64 {
	m := len(squares) * 2
	xs := make([]int, 0, m)
	type event struct{ y, lx, rx, delta int }
	events := make([]event, 0, m)
	for _, sq := range squares {
		lx, y, l := sq[0], sq[1], sq[2]
		rx := lx + l
		xs = append(xs, lx, rx)
		events = append(events, event{y, lx, rx, 1}, event{y + l, lx, rx, -1})
	}

	// 排序去重，方便离散化
	slices.Sort(xs)
	xs = slices.Compact(xs)

	// 初始化线段树
	n := len(xs) - 1 // len(xs) 个横坐标有 len(xs)-1 个差值
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(xs, 1, 0, n-1)

	// 模拟扫描线从下往上移动
	slices.SortFunc(events, func(a, b event) int { return a.y - b.y })
	type pair struct{ area, sumLen int }
	records := make([]pair, m-1)
	totArea := 0
	for i, e := range events[:m-1] {
		l := sort.SearchInts(xs, e.lx)
		r := sort.SearchInts(xs, e.rx) - 1 // 注意 r 对应着 xs[r] 与 xs[r+1]=e.rx 的差值
		t.update(1, l, r, e.delta)         // 更新被 [e.lx, e.rx] 覆盖的次数
		sumLen := xs[len(xs)-1] - xs[0]    // 总的底边长度
		if t[1].minCover == 0 {            // 需要去掉没被矩形覆盖的长度
			sumLen -= t[1].minCoverLen
		}
		records[i] = pair{totArea, sumLen} // 记录关键数据
		totArea += sumLen * (events[i+1].y - e.y) // 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度
	}

	// 二分找最后一个 < totArea / 2 的面积
	i := sort.Search(m-1, func(i int) bool { return records[i].area*2 >= totArea }) - 1
	return float64(events[i].y) + float64(totArea-records[i].area*2)/float64(records[i].sumLen*2)
}
