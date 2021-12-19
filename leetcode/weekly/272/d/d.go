package main

import "sort"

/* 最长上升子序列

将 $\textit{arr}$ 每隔 $k$ 个数取一个元素，分为若干组，例如当 $k=3$ 时，分为如下三组：

- $arr[0],arr[3],arr[6],\cdots$
- $arr[1],arr[4],arr[7],\cdots$
- $arr[2],arr[5],arr[8],\cdots$

根据题意，组与组之间是互不影响的，我们需要让每一组内的元素单调不降。

题目要求需要修改的最少元素个数，这可以通过求不需要修改的最多元素个数来求出，也就是该组元素的最长不降子序列。（不了解的同学可以看看第 300 题）

需要注意的是：

- 求最长严格递增子序列需要二分找到大于或等于当前元素的元素位置（即 C++ 中的 `lower_bound`）；
- 求最长非降子序列需要二分找到大于当前元素的元素位置（即 C++ 中的 `upper_bound`）。

累加所有组的最长不降子序列的长度即为最多可以保留的元素个数，用 $\textit{arr}$ 的长度减去该个数即为答案。

*/

// github.com/EndlessCheng/codeforces-go
func kIncreasing(arr []int, k int) int {
	save := 0
	for i, n := 0, len(arr); i < k && i < n; i++ {
		f := []int{}
		for j := i; j < n; j += k {
			v := arr[j]
			if p := sort.SearchInts(f, v+1); p < len(f) {
				f[p] = v
			} else {
				f = append(f, v)
			}
		}
		save += len(f)
	}
	return len(arr) - save
}
