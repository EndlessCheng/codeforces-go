package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func longestBalanced1(nums []int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(n+1)))/2 + 1
	sum := make([]int, n+1)

	// === 分块模板开始 ===
	// 用分块维护 sum
	type block struct {
		l, r int // [l,r) 左闭右开
		todo int
		pos  map[int]int
	}
	blocks := make([]block, n/B+1)
	calcPos := func(l, r int) map[int]int {
		pos := map[int]int{}
		for j := r - 1; j >= l; j-- {
			pos[sum[j]] = j
		}
		return pos
	}
	for i := 0; i <= n; i += B {
		r := min(i+B, n+1)
		pos := calcPos(i, r)
		blocks[i/B] = block{i, r, 0, pos}
	}

	// sum[l:r] 增加 v
	rangeAdd := func(l, r, v int) {
		for i := range blocks {
			b := &blocks[i]
			if b.r <= l {
				continue
			}
			if b.l >= r {
				break
			}
			if l <= b.l && b.r <= r { // 完整块
				b.todo += v
			} else { // 部分块，直接重算
				for j := b.l; j < b.r; j++ {
					sum[j] += b.todo
					if l <= j && j < r {
						sum[j] += v
					}
				}
				b.pos = calcPos(b.l, b.r)
				b.todo = 0
			}
		}
	}

	// 返回 sum[:r] 中第一个 v 的下标
	// 如果没有 v，返回 n
	findFirst := func(r, v int) int {
		for i := range blocks {
			b := &blocks[i]
			if b.r <= r { // 完整块，直接查哈希表
				if j, ok := b.pos[v-b.todo]; ok {
					return j
				}
			} else { // 部分块，暴力查找
				for j := b.l; j < r; j++ {
					if sum[j] == v-b.todo {
						return j
					}
				}
				break
			}
		}
		return n
	}
	// === 分块模板结束 ===

	last := map[int]int{} // nums 的元素上一次出现的位置
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		v := x%2*2 - 1
		if j := last[x]; j == 0 { // 首次遇到 x
			rangeAdd(i, n+1, v) // sum[i:] 增加 v
		} else { // 再次遇到 x
			rangeAdd(j, i, -v) // 撤销之前对 sum[j:i] 的增加
		}
		last[x] = i

		s := sum[i] + blocks[i/B].todo // sum[i] 的实际值
		ans = max(ans, i-findFirst(i-ans, s))
	}
	return
}

type pair struct{ min, max int }
type lazySeg []struct {
	l, r int
	pair
	todo int
}

func merge(l, r pair) pair {
	return pair{min(l.min, r.min), max(l.max, r.max)}
}

func (t lazySeg) apply(o int, f int) {
	cur := &t[o]
	cur.min += f
	cur.max += f
	cur.todo += f
}

func (t lazySeg) maintain(o int) {
	t[o].pair = merge(t[o<<1].pair, t[o<<1|1].pair)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t lazySeg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t lazySeg) update(o, l, r int, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

// 查询 [l,r] 内第一个等于 target 的元素下标
func (t lazySeg) findFirst(o, l, r, target int) int {
	if t[o].l > r || t[o].r < l || target < t[o].min || target > t[o].max {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findFirst(o<<1, l, r, target)
	if idx < 0 {
		// 去右子树找
		idx = t.findFirst(o<<1|1, l, r, target)
	}
	return idx
}

func longestBalanced(nums []int) (ans int) {
	n := len(nums)
	t := make(lazySeg, 2<<bits.Len(uint(n)))
	t.build(1, 0, n)

	last := map[int]int{} // nums 的元素上一次出现的位置
	curSum := 0
	for i := 1; i <= n; i++ {
		x := nums[i-1]
		v := x%2*2 - 1
		if j := last[x]; j == 0 { // 首次遇到 x
			curSum += v
			t.update(1, i, n, v) // sum[i:] 增加 v
		} else { // 再次遇到 x
			t.update(1, j, i-1, -v) // 撤销之前对 sum[j:i] 的增加
		}
		last[x] = i

		j := t.findFirst(1, 0, i-1, curSum)
		if j >= 0 {
			ans = max(ans, i-j)
		}
	}
	return
}
