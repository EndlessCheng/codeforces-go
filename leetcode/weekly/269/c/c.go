package main

/* 三种贪心策略取最小值

根据题意，可以得到如下三种贪心策略：

- 删除包含最小值和最大值的数组前缀；
- 删除包含最小值和最大值的数组后缀；
- 删除包含最小值的数组前缀（后缀），以及包含最大值的数组后缀（前缀）。

记数组长度为 $n$，最小值的位置为 $i$，最大值的位置为 $j$，假设 $i\le j$，那么上述三种策略所需的删除次数为：

- $j+1$
- $n-i$
- $(i+1)+(n-j)$

取三者最小值，即为答案。

*/

// github.com/EndlessCheng/codeforces-go
func minimumDeletions(nums []int) int {
	i, j := 0, 0
	for p, num := range nums {
		if num < nums[i] { i = p }
		if num > nums[j] { j = p }
	}
	if i > j {
		i, j = j, i // 保证 i <= j
	}
	return min(min(j+1, len(nums)-i), i+1+len(nums)-j)
}

func min(a, b int) int { if a > b { return b }; return a }
