package main

import (
	"sort"
)

/* 分类讨论+二分答案

先统计分负数乘积个数 $\textit{neg}$、正数乘积个数 $\textit{pos}$ 以及乘积为 $0$ 的个数 $\textit{zero}$，然后分三种情况讨论：

- $k\le \textit{neg}$，我们可以二分负数答案，统计不超过二分值的乘积个数；
- $\textit{neg}<k\le \textit{neg}+\textit{zero}$，此时返回 $0$；
- $k>\textit{neg}+\textit{zero}$，我们可以二分正数答案，统计不超过二分值的乘积个数；

以最后一种情况为例，记二分值为 $t$，我们可以遍历 $\textit{num1}$ 中的正数，并在 $\textit{num2}$ 的正数中二分（或双指针）不超过 $\dfrac{t}{\textit{num1}[i]}$ 的元素个数，然后遍历 $\textit{num1}$ 中的负数，方法同上。遍历结束后，若元素个数小于 $t$ 则说明二分值偏小，否则偏大。

*/

// github.com/EndlessCheng/codeforces-go
func kthSmallestProduct(nums1 []int, nums2 []int, K int64) int64 {
	n1, n2 := len(nums1), len(nums2)
	p10 := sort.SearchInts(nums1, 0)
	p11 := sort.SearchInts(nums1, 1)
	p20 := sort.SearchInts(nums2, 0)
	p21 := sort.SearchInts(nums2, 1)

	neg := p10*(n2-p21) + p20*(n1-p11) // 负数乘积个数
	pos := p10*p20 + (n1-p11)*(n2-p21) // 正数乘积个数
	zero := n1*n2 - neg - pos          // 乘积为 0 的个数

	pos1, pos2 := nums1[p11:], nums2[p21:]
	neg1, neg2 := nums1[:p10], nums2[:p20]
	for i := range neg2 {
		neg2[i] = -neg2[i]
	}
	for i, n := 0, len(neg2); i < n/2; i++ {
		neg2[i], neg2[n-1-i] = neg2[n-1-i], neg2[i]
	}

	k := int(K)
	if k <= neg {
		k = neg - k + 1
		return int64(-sort.Search(1e10, func(t int) bool {
			cnt := 0
			for _, v := range pos1 { cnt += sort.SearchInts(neg2, t/v+1) } // 也可以用双指针，这里为了方便直接二分
			for _, v := range neg1 { cnt += sort.SearchInts(pos2, t/-v+1) }
			return cnt >= k
		}))
	}

	if k <= neg+zero {
		return 0
	}

	k -= neg + zero
	return int64(sort.Search(1e10, func(t int) bool {
		cnt := 0
		for _, v := range pos1 { cnt += sort.SearchInts(pos2, t/v+1) }
		for _, v := range neg1 { cnt += sort.SearchInts(neg2, t/-v+1) }
		return cnt >= k
	}))
}
