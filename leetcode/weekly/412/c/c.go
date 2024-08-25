package main

import (
	"container/heap"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	clone := slices.Clone(h)

	// 打表，计算出最小的 e 满足 multiplier^e >= 2^i
	mxLen := bits.Len(uint(mx))
	type ep struct{ e, powM int }
	ePowM := make([]ep, 0, mxLen)
	for pow2, powM, e := 1, 1, 0; pow2 <= mx; pow2 <<= 1 {
		if powM < pow2 {
			powM *= multiplier
			e++
		}
		ePowM = append(ePowM, ep{e, powM})
	}

	// 每个数操作到 >= mx
	left := k
	for i := range h {
		x := h[i].x
		p := ePowM[mxLen-bits.Len(uint(x))]
		e, powM := p.e, p.powM
		if powM/multiplier*x >= mx { // 多操作了一次
			powM /= multiplier
			e--
		} else if x*powM < mx { // 少操作了一次
			powM *= multiplier
			e++
		}
		left -= e
		if left < 0 {
			break
		}
		h[i].x *= powM
	}

	if left < 0 {
		// 暴力模拟
		h = clone
		heap.Init(&h)
		for ; k > 0; k-- {
			h[0].x *= multiplier
			heap.Fix(&h, 0)
		}
		sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
		for _, p := range h {
			nums[p.i] = p.x % mod
		}
		return nums
	}

	// 剩余的操作可以直接用公式计算
	k = left
	pow1 := pow(multiplier, k/n)
	pow2 := pow1 * multiplier % mod
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		pw := pow1
		if i < k%n {
			pw = pow2
		}
		nums[p.i] = p.x % mod * pw % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
