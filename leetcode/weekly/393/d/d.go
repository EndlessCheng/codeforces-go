package main

import (
	"math"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func minimumValueSum(nums, andValues []int) int {
	n, m := len(nums), len(andValues)
	type args struct{ i, j, and int }
	memo := map[args]int{}
	var dfs func(int, int, int) int
	dfs = func(i, j, and int) int {
		if m-j > n-i { // 剩余元素不足
			return math.MaxInt / 2
		}
		if j == m { // 分了 m 段
			if i == n {
				return 0
			}
			return math.MaxInt / 2
		}
		and &= nums[i]
		if and < andValues[j] { // 剪枝：无法等于 andValues[j]
			return math.MaxInt / 2
		}
		p := args{i, j, and}
		if res, ok := memo[p]; ok {
			return res
		}
		res := dfs(i+1, j, and)  // 不划分
		if and == andValues[j] { // 划分，nums[i] 是这一段的最后一个数
			res = min(res, dfs(i+1, j+1, -1)+nums[i])
		}
		memo[p] = res
		return res
	}
	ans := dfs(0, 0, -1)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

type ST [][]int

func NewST(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, v := range a {
		st[i] = make([]int, sz)
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = st[i][j-1] & st[i+1<<(j-1)][j-1]
		}
	}
	return st
}

// 查询区间 [l,r)    0 <= l < r <= n
func (st ST) Query(l, r int) int {
	k := bits.Len32(uint32(r-l)) - 1
	return st[l][k] & st[r-1<<k][k]
}

type seg []struct {
	l, r int
	val  int
}

func mergeInfo(a, b int) int {
	return min(a, b)
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = 1e18
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int, val int) {
	if t[o].l == t[o].r {
		t[o].val = mergeInfo(t[o].val, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].val = mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lv := t.query(o<<1, l, r)
	rv := t.query(o<<1|1, l, r)
	return mergeInfo(lv, rv)
}

func minimumValueSum2(a []int, tar []int) (ans int) {
	n := len(a)
	m := len(tar)

	f := make(seg, 2<<bits.Len(uint(n-1)))
	f.build(1, 0, n-1)

	st := NewST(a)
	and := -1
	for i, v := range a {
		and &= v
		if and == tar[0] {
			f.update(1, i, v)
		}
	}

	for sp := 1; sp < m; sp++ {
		nf := make(seg, 2<<bits.Len(uint(n-1)))
		nf.build(1, 0, n-1)
		for i := sp; i < n; i++ {
			l := sort.Search(i+1, func(j int) bool { return st.Query(j, i+1) >= tar[sp] })
			if l > i || st.Query(l, i+1) != tar[sp] {
				continue
			}
			r := sort.Search(i+1, func(j int) bool { return st.Query(j, i+1) > tar[sp] }) - 1
			l = max(l, 1)
			r = max(r, 1)
			nf.update(1, i, f.query(1, l-1, r-1)+a[i])
		}
		f = nf
	}
	ans = f.query(1, n-1, n-1)
	if ans == 1e18 {
		return -1
	}
	return
}
