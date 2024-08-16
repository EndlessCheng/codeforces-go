package main

import (
	"math"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minimumValueSum(nums, andValues []int) int {
	const inf = math.MaxInt / 2
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = inf
	}
	newF := make([]int, n+1)
	for _, target := range andValues {
		nums := slices.Clone(nums)
		left, right := 0, 0
		q := []int{} // 单调队列，保存 f 的下标
		qi := 0      // 单调队列目前处理到 f[qi]

		newF[0] = inf
		for i, x := range nums {
			for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
				nums[j] &= x
			}
			for left <= i && nums[left] < target {
				left++
			}
			for right <= i && nums[right] <= target {
				right++
			}

			// 上面这段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值
			// 下面是单调队列的滑窗过程

			if left < right {
				// 单调队列：右边入
				for ; qi < right; qi++ {
					for len(q) > 0 && f[qi] <= f[q[len(q)-1]] {
						q = q[:len(q)-1]
					}
					q = append(q, qi)
				}

				// 单调队列：左边出
				for q[0] < left {
					q = q[1:]
				}

				// 单调队列：计算答案
				newF[i+1] = f[q[0]] + x // 队首就是最小值
			} else {
				newF[i+1] = inf
			}
		}
		f, newF = newF, f
	}
	if f[n] < inf {
		return f[n]
	}
	return -1
}

func minimumValueSum22(nums, andValues []int) int {
	const inf = math.MaxInt / 2
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = inf
	}
	newF := make([]int, n+1)
	for _, target := range andValues {
		newF[0] = inf
		type pair struct{ and, l int }
		a := []pair{}
		q := []int{}
		qi := 0
		for i, x := range nums {
			for j := range a {
				a[j].and &= x
			}
			a = append(a, pair{x, i})

			// 原地去重
			j := 1
			for k := 1; k < len(a); k++ {
				if a[k].and != a[k-1].and {
					a[j] = a[k]
					j++
				}
			}
			a = a[:j]

			// 去掉无用数据
			for len(a) > 0 && a[0].and < target {
				a = a[1:]
			}

			if len(a) > 0 && a[0].and == target {
				r := i
				if len(a) > 1 {
					r = a[1].l - 1
				}
				for ; qi <= r; qi++ {
					for len(q) > 0 && f[qi] <= f[q[len(q)-1]] {
						q = q[:len(q)-1]
					}
					q = append(q, qi)
				}
				for q[0] < a[0].l {
					q = q[1:]
				}
				newF[i+1] = f[q[0]] + x
			} else {
				newF[i+1] = inf
			}
		}
		f, newF = newF, f
	}
	if f[n] < inf {
		return f[n]
	}
	return -1
}

func minimumValueSumDP(nums, andValues []int) int {
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
