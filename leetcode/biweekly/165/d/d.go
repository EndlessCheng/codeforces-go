package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
type xorBasis []int

// n 为值域最大值 U 的二进制长度，例如 U=1e9 时 n=30
func newXorBasis(n int) xorBasis {
	return make(xorBasis, n)
}

func (b xorBasis) insert(x int) {
	// 从高到低遍历，保证计算 maxXor 的时候，参与 XOR 的基的最高位（或者说二进制长度）是互不相同的
	for i := len(b) - 1; i >= 0; i-- {
		if x>>i == 0 { // 由于大于 i 的位都被我们异或成了 0，所以 x>>i 的结果只能是 0 或 1
			continue
		}
		if b[i] == 0 { // x 和之前的基是线性无关的
			b[i] = x // 新增一个基，最高位为 i
			return
		}
		x ^= b[i] // 保证每个基的二进制长度互不相同
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

func maxXorSubsequences(nums []int) int {
	u := slices.Max(nums)
	m := bits.Len(uint(u))
	b := newXorBasis(m)
	for _, x := range nums {
		b.insert(x)
	}
	return b.maxXor()
}
