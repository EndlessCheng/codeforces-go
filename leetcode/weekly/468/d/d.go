package main

import (
	"fmt"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type data struct{ sumMin, sumMax, lMin, lMax int }
type todo struct{ todoMin, todoMax int }
type lazySeg []struct {
	l, r int
	data
	todo
}

var todoInit = todo{-1, -1}

func merge(l, r data) data {
	return data{l.sumMin + r.sumMin, l.sumMax + r.sumMax, l.lMin, l.lMax}
}

func (t lazySeg) apply(o int, f todo) {
	cur := &t[o]
	sz := cur.r - cur.l + 1
	if f.todoMin >= 0 {
		cur.lMin = f.todoMin
		cur.sumMin = f.todoMin * sz
		cur.todoMin = f.todoMin
	}
	if f.todoMax >= 0 {
		cur.lMax = f.todoMax
		cur.sumMax = f.todoMax * sz
		cur.todoMax = f.todoMax
	}
	fmt.Println(o, f, cur)
}

func (t lazySeg) maintain(o int) {
	t[o].data = merge(t[o<<1].data, t[o<<1|1].data)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t lazySeg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t lazySeg) update(o, l, r int, f todo) {
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

func (t lazySeg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	lRes := t.query(o<<1, l, r)
	rRes := t.query(o<<1|1, l, r)
	return merge(lRes, rRes)
}

// 查询 [l,r] 内最后一个满足 f 的下标
func (t lazySeg) findLast(o, l, r int, f func(data) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].data) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findLast(o<<1|1, l, r, f)
	if idx < 0 {
		idx = t.findLast(o<<1, l, r, f)
	}
	return idx
}

func maxTotalValue(nums []int, k int) (ans int64) {
	// 二分 + 滑动窗口 + 单调队列
	lowD := sort.Search(slices.Max(nums)-slices.Min(nums), func(lowD int) bool {
		lowD++
		// 1438. 绝对差不超过限制的最长连续子数组（改成求子数组个数）
		var minQ, maxQ []int
		cnt, left := 0, 0
		for right, x := range nums {
			for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
				minQ = minQ[:len(minQ)-1]
			}
			minQ = append(minQ, right)

			for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
				maxQ = maxQ[:len(maxQ)-1]
			}
			maxQ = append(maxQ, right)

			for nums[maxQ[0]]-nums[minQ[0]] >= lowD {
				left++
				if minQ[0] < left {
					minQ = minQ[1:]
				}
				if maxQ[0] < left {
					maxQ = maxQ[1:]
				}
			}

			cnt += left
			if cnt >= k {
				return false
			}
		}
		return true
	})

	// 单调栈
	n := len(nums)
	leftLessEq := make([]int, n)
	leftGreatEq := make([]int, n)
	st1 := []int{-1}
	st2 := []int{-1}
	for i, x := range nums {
		for len(st1) > 1 && nums[st1[len(st1)-1]] > x {
			st1 = st1[:len(st1)-1]
		}
		leftLessEq[i] = st1[len(st1)-1]
		st1 = append(st1, i)

		for len(st2) > 1 && nums[st2[len(st2)-1]] < x {
			st2 = st2[:len(st2)-1]
		}
		leftGreatEq[i] = st2[len(st2)-1]
		st2 = append(st2, i)
	}

	// Lazy 线段树
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	cnt, sum := 0, 0
	for i, x := range nums {
		t.update(1, leftLessEq[i]+1, i, todo{x, -1})
		t.update(1, leftGreatEq[i]+1, i, todo{-1, x})
		l := t.findLast(1, 0, i, func(d data) bool { return d.lMax-d.lMin >= lowD })
		if l >= 0 {
			cnt += l + 1
			d := t.query(1, 0, l)
			sum += d.sumMax - d.sumMin
		}
	}

	return int64(sum - (cnt-k)*lowD) // 减掉多算的
}
