package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
// 线性基模板
type xorBasis []int

func (b xorBasis) insert(x int) {
	for x > 0 {
		i := bits.Len(uint(x)) - 1 // x 的最高位
		if b[i] == 0 {             // x 和之前的基是线性无关的
			b[i] = x // 新增一个基，最高位为 i
			return
		}
		x ^= b[i] // 保证参与 maxXor 的基的最高位是互不相同的，方便我们贪心
	}
	// 正常循环结束，此时 x=0，说明一开始的 x 可以被已有基表出，不是一个线性无关基
}

func (b xorBasis) maxXor() (res int) {
	// 从高到低贪心：越高的位，越必须是 1
	// 由于每个位的基至多一个，所以每个位只需考虑异或一个基，若能变大，则异或之
	for i := len(b) - 1; i >= 0; i-- {
		res = max(res, res^b[i])
	}
	return
}

func maximizeXorAndXor(nums []int) int64 {
	n := len(nums)
	type pair struct{ and, xor, or int } // 多算一个子集 OR，用于剪枝
	subSum := make([]pair, 1<<n)
	subSum[0].and = -1
	for i, x := range nums {
		highBit := 1 << i
		for mask, p := range subSum[:highBit] {
			subSum[highBit|mask] = pair{p.and & x, p.xor ^ x, p.or | x}
		}
	}
	subSum[0].and = 0

	sz := bits.Len(uint(slices.Max(nums)))
	b := make(xorBasis, sz)
	maxXor2 := func(sub uint) (res int) {
		clear(b)
		xor := subSum[sub].xor
		for ; sub > 0; sub &= sub - 1 {
			x := nums[bits.TrailingZeros(sub)]
			b.insert(x &^ xor) // 只考虑有偶数个 1 的比特位（xor 在这些比特位上是 0）
		}
		return xor + b.maxXor()*2
	}

	ans := 0
	u := 1<<n - 1
	for i, p := range subSum {
		if p.and+subSum[u^i].or*2 > ans { // 有机会让 ans 变得更大
			ans = max(ans, p.and+maxXor2(uint(u^i)))
		}
	}
	return int64(ans)
}
