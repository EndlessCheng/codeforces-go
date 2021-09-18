package main

import "sort"

/* 排序+去重+二分

考虑最多可以保留多少个元素不变。由于元素的位置不影响答案，且要求所有元素互不相同，我们可以将 $\textit{nums}$ 排序，并去掉重复元素。

记原数组长度为 $n$。对排序去重后的 $\textit{nums}'$ 中的一段区间 $[l,r]$，若要保留这段区间内的所有元素，我们需要替换区间外的元素，填充到 $[\textit{nums}'[l],\textit{nums}'[r]]$ 内缺失的元素上。

需要填充的元素个数为

$$\textit{nums}'[r]-\textit{nums}'[l]+1-(r-l+1)$$

区间外有 $n-(r-l+1)$ 个元素，由于区间外的元素个数不能少于需要填充的元素个数，所以有

$$
\textit{nums}'[r]-\textit{nums}'[l]+1-(r-l+1) \le n-(r-l+1)
$$

上式可化简为

$$
\textit{nums}'[l]\ge\textit{nums}'[r]-n+1
$$

根据该式，我们可以遍历 $\textit{nums}'[r]$，二分得到最小的满足该式的 $l$，此时 $[l,r]$ 区间内的元素均可以保留。用 $n$ 减去最多可以保留的元素个数，就是答案。

时间复杂度：$O(n\log n)$。

空间复杂度：$O(1)$。只需要常数的空间存放若干变量。

*/

// github.com/EndlessCheng/codeforces-go
func minOperations(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	nums = unique(nums)
	for r, v := range nums {
		l := sort.SearchInts(nums[:r], v-n+1)
		ans = max(ans, r-l+1) // [l,r] 内的元素均可以保留
	}
	return n - ans
}

// 原地去重
func unique(a []int) []int {
	k := 0
	for _, v := range a[1:] {
		if a[k] != v {
			k++
			a[k] = v
		}
	}
	return a[:k+1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
