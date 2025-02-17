package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type segRect []struct {
	l, r        int
	minCoverLen int // 区间内被矩形覆盖次数最少的底边长之和
	minCover    int // 区间内被矩形覆盖的最小次数
	todo        int // 子树内的所有节点的 minCover 需要增加的量，注意这可能是负数
}

// 根据左右儿子的信息，更新当前节点的信息
func (t segRect) maintain(o int) {
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

// 仅更新节点信息，不下传懒标记
func (t segRect) do(o, v int) {
	t[o].minCover += v
	t[o].todo += v
}

// 下传懒标记
func (t segRect) spread(o int) {
	v := t[o].todo
	if v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

// 建树
func (t segRect) build(xs []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].minCoverLen = xs[l+1] - xs[l]
		return
	}
	m := (l + r) >> 1
	t.build(xs, o<<1, l, m)
	t.build(xs, o<<1|1, m+1, r)
	t.maintain(o)
}

// 区间更新
func (t segRect) update(o, l, r, v int) {
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

func p5490(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	xs := make([]int, 0, n*2)
	type event struct{ y, lx, rx, delta int }
	events := make([]event, 0, n*2)
	for ; n > 0; n-- {
		var lx, ly, rx, ry int
		Fscan(in, &lx, &ly, &rx, &ry)
		xs = append(xs, lx, rx)
		events = append(events, event{ly, lx, rx, 1}, event{ry, lx, rx, -1})
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)

	// 初始化线段树
	n = len(xs) - 1 // len(xs) 个横坐标有 len(xs)-1 个差值
	t := make(segRect, 2<<bits.Len(uint(n-1)))
	t.build(xs, 1, 0, n-1)

	// 模拟扫描线从下往上移动
	slices.SortFunc(events, func(a, b event) int { return a.y - b.y })
	for i, e := range events[:len(events)-1] {
		l := sort.SearchInts(xs, e.lx)     // 离散化
		r := sort.SearchInts(xs, e.rx) - 1 // 注意 r 对应着 xs[r] 与 xs[r+1]=e.rx 的差值
		t.update(1, l, r, e.delta)         // 更新被 [e.lx, e.rx] 覆盖的次数
		sumLen := xs[len(xs)-1] - xs[0]    // 总的底边长度
		if t[1].minCover == 0 {            // 需要去掉没被矩形覆盖的长度
			sumLen -= t[1].minCoverLen
		}
		ans += sumLen * (events[i+1].y - e.y) // 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度
	}
	Fprint(out, ans)
}

//func main() { p5490(bufio.NewReader(os.Stdin), os.Stdout) }
