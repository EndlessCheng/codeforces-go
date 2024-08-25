package main

import (
	"container/heap"
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
	heap.Init(&h)

	// 模拟，直到堆顶是 mx
	for ; k > 0 && h[0].x < mx; k-- {
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}

	// 剩余的操作可以直接用公式计算
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.i] = p.x % mod * pow(multiplier, e) % mod
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
