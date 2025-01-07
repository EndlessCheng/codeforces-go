package main

import (
	"math"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxSubarraySum(nums []int) int64 {
	ans := math.MinInt
	var s, nonDelMinS, allMin int
	delMinS := map[int]int{}
	for _, x := range nums {
		s += x
		ans = max(ans, s-allMin)
		if x < 0 {
			delMinS[x] = min(delMinS[x], nonDelMinS) + x
			allMin = min(allMin, delMinS[x])
			nonDelMinS = min(nonDelMinS, s)
		}
	}
	return int64(ans)
}

func maxSubarraySumPS(nums []int) int64 {
	n := len(nums)
	f := math.MinInt / 2
	s := 0
	last := map[int]int{}

	update := func(x int) int {
		res := f          // f[i-1]
		f = max(f, 0) + x // f[i] = max(f[i-1], 0) + x
		if v, ok := last[x]; ok {
			res = max(res, v+s) // s[i]
		}
		s += x // s[i+1] = s[i] + x
		last[x] = res - s
		return res
	}

	pre := make([]int, n)
	for i, x := range nums {
		pre[i] = update(x)
	}

	ans := math.MinInt
	f = math.MinInt / 2
	s = 0
	clear(last)
	for i, x := range slices.Backward(nums) {
		suf := update(x)
		ans = max(ans, f, pre[i]+suf, pre[i], suf)
	}
	return int64(ans)
}

//

type info struct {
	ans, sum, pre, suf int
}

type seg []info

func (t seg) set(o, val int) {
	t[o] = info{val, val, val, val}
}

func (t seg) mergeInfo(a, b info) info {
	return info{
		max(max(a.ans, b.ans), a.suf+b.pre),
		a.sum + b.sum,
		max(a.pre, a.sum+b.pre),
		max(b.suf, b.sum+a.suf),
	}
}

func (t seg) maintain(o int) {
	t[o] = t.mergeInfo(t[o<<1], t[o<<1|1])
}

// 初始化线段树
func (t seg) build(nums []int, o, l, r int) {
	if l == r {
		t.set(o, nums[l])
		return
	}
	m := (l + r) >> 1
	t.build(nums, o<<1, l, m)
	t.build(nums, o<<1|1, m+1, r)
	t.maintain(o)
}

// 单点更新
func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t.set(o, val)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t.maintain(o)
}

// 区间询问（没用到）
func (t seg) query(o, l, r, L, R int) info {
	if L <= l && r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, L, R)
	}
	if m < L {
		return t.query(o<<1|1, m+1, r, L, R)
	}
	return t.mergeInfo(t.query(o<<1, l, m, L, R), t.query(o<<1|1, m+1, r, L, R))
}

func maxSubarraySumSeg(nums []int) int64 {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	ans := t[1].ans // 不删任何数
	if ans <= 0 {
		return int64(ans)
	}

	pos := map[int][]int{}
	for i, x := range nums {
		if x < 0 {
			pos[x] = append(pos[x], i)
		}
	}
	for _, idx := range pos {
		for _, i := range idx {
			t.update(1, 0, n-1, i, 0) // 删除
		}
		ans = max(ans, t[1].ans)
		for _, i := range idx {
			t.update(1, 0, n-1, i, nums[i]) // 复原
		}
	}
	return int64(ans)
}
