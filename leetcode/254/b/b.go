package main

import "sort"

/*

将数组排序后从中间分开，错开一位合并。

 */

// github.com/EndlessCheng/codeforces-go
func rearrangeArray(a []int) []int {
	sort.Ints(a)
	i, n := 0, len(a)
	ans := make([]int, n)
	for _, v := range a {
		ans[i] = v
		if i += 2; i >= n {
			i = 1
		}
	}
	return ans
}
