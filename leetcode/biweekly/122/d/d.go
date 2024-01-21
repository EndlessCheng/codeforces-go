package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

// https://space.bilibili.com/206214
func minimumCost(nums []int, k, dist int) int64 {
	k--
	L := redblacktree.NewWithIntComparator()
	R := redblacktree.NewWithIntComparator()
	add := func(t *redblacktree.Tree, x int) {
		c, ok := t.Get(x)
		if ok {
			t.Put(x, c.(int)+1)
		} else {
			t.Put(x, 1)
		}
	}
	del := func(t *redblacktree.Tree, x int) {
		c, _ := t.Get(x)
		if c.(int) > 1 {
			t.Put(x, c.(int)-1)
		} else {
			t.Remove(x)
		}
	}

	sumL := nums[0]
	for _, x := range nums[1 : dist+2] {
		sumL += x
		add(L, x)
	}
	sizeL := dist + 1

	l2r := func() {
		x := L.Right().Key.(int)
		sumL -= x
		sizeL--
		del(L, x)
		add(R, x)
	}
	r2l := func() {
		x := R.Left().Key.(int)
		sumL += x
		sizeL++
		del(R, x)
		add(L, x)
	}
	for sizeL > k {
		l2r()
	}

	ans := sumL
	for i := dist + 2; i < len(nums); i++ {
		// 移除 out
		out := nums[i-dist-1]
		if _, ok := L.Get(out); ok {
			sumL -= out
			sizeL--
			del(L, out)
		} else {
			del(R, out)
		}

		// 添加 in
		in := nums[i]
		if in < L.Right().Key.(int) {
			sumL += in
			sizeL++
			add(L, in)
		} else {
			add(R, in)
		}

		// 维护大小
		if sizeL == k-1 {
			r2l()
		} else if sizeL == k+1 {
			l2r()
		}

		ans = min(ans, sumL)
	}
	return int64(ans)
}
