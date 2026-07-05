package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
const mx = 1_000_001

var primeDivisors [mx][]int32

// 预处理每个数的质因子
func init() {
	for i := int32(2); i < mx; i++ {
		if primeDivisors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // 枚举 i 的倍数 j
				primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
			}
		}
	}
}

type data struct {
	sum, pre, suf, ans int
}

type seg []data

func (t seg) set(node, v int) {
	t[node] = data{v, v, v, v}
}

func (t seg) maintain(node int) {
	lo, ro := t[node*2], t[node*2+1]
	t[node].sum = lo.sum + ro.sum
	t[node].pre = max(lo.pre, lo.sum+ro.pre)
	t[node].suf = max(ro.suf, ro.sum+lo.suf)
	t[node].ans = max(lo.ans, ro.ans, lo.suf+ro.pre)
}

func (t seg) build(a []int, node, l, r int) {
	if l == r {
		t.set(node, -a[l])
		return
	}
	m := (l + r) >> 1
	t.build(a, node*2, l, m)
	t.build(a, node*2+1, m+1, r)
	t.maintain(node)
}

func (t seg) update(node, l, r, i, val int) {
	if l == r {
		t.set(node, val)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(node*2, l, m, i, val)
	} else {
		t.update(node*2+1, m+1, r, i, val)
	}
	t.maintain(node)
}

func divisibleGame(nums []int) (ans int) {
	const mod = 1_000_000_007

	primeDivisorsToIndices := map[int32][]int{}
	for i, x := range nums {
		for _, d := range primeDivisors[x] {
			primeDivisorsToIndices[d] = append(primeDivisorsToIndices[d], i)
		}
	}

	if len(primeDivisorsToIndices) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2
		return mod - 2
	}

	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	maxDiff, bestK := math.MinInt, int32(0)

	// 枚举质因子作为 k，计算最大子数组和
	for k, indices := range primeDivisorsToIndices {
		for _, i := range indices {
			// nums[i] 是质因子 k 的倍数
			t.update(1, 0, n-1, i, nums[i])
		}

		diff := t[1].ans
		if diff > maxDiff || diff == maxDiff && k < bestK {
			maxDiff, bestK = diff, k
		}

		for _, i := range indices {
			t.update(1, 0, n-1, i, -nums[i])
		}
	}

	return maxDiff * int(bestK) % mod
}
