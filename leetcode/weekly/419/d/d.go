package main

import (
	"cmp"
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

// https://space.bilibili.com/206214
type pair struct{ c, x int }

func less(p, q pair) int {
	return cmp.Or(p.c-q.c, p.x-q.x)
}

func findXSum(nums []int, k, x int) []int64 {
	L := redblacktree.NewWith[pair, struct{}](less)
	R := redblacktree.NewWith[pair, struct{}](less)

	sumL := 0
	cnt := map[int]int{}
	add := func(x int) {
		p := pair{cnt[x], x}
		if p.c == 0 {
			return
		}
		if !L.Empty() && less(p, L.Left().Key) > 0 { // p 比 L 中最小的还大
			sumL += p.c * p.x
			L.Put(p, struct{}{})
		} else {
			R.Put(p, struct{}{})
		}
	}
	del := func(x int) {
		p := pair{cnt[x], x}
		if p.c == 0 {
			return
		}
		if _, ok := L.Get(p); ok {
			sumL -= p.c * p.x
			L.Remove(p)
		} else {
			R.Remove(p)
		}
	}

	l2r := func() {
		p := L.Left().Key
		sumL -= p.c * p.x
		L.Remove(p)
		R.Put(p, struct{}{})
	}
	r2l := func() {
		p := R.Right().Key
		sumL += p.c * p.x
		R.Remove(p)
		L.Put(p, struct{}{})
	}

	ans := make([]int64, len(nums)-k+1)
	for r, in := range nums {
		// 添加 in
		del(in)
		cnt[in]++
		add(in)

		l := r + 1 - k
		if l < 0 {
			continue
		}

		// 维护大小
		for !R.Empty() && L.Size() < x {
			r2l()
		}
		for L.Size() > x {
			l2r()
		}
		ans[l] = int64(sumL)

		// 移除 out
		out := nums[l]
		del(out)
		cnt[out]--
		add(out)
	}
	return ans
}
