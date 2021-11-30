package main

/* O(n) 做法：统计小于和等于 target 的元素个数

由于排序后，相同的值是连续的，记小于 $\textit{target}$ 的元素个数为 $\textit{less}$，等于 $\textit{target}$ 的元素个数为 $\textit{equal}$，那么答案即为

$$
\textit{less}, \textit{less}+1, \cdots, \textit{less}+$\textit{equal}$-1
$$

*/

// github.com/EndlessCheng/codeforces-go
func targetIndices(nums []int, target int) []int {
	less, equal := 0, 0
	for _, num := range nums {
		if num < target {
			less++
		} else if num == target {
			equal++
		}
	}
	ans := make([]int, equal)
	for i := range ans {
		ans[i] = less + i
	}
	return ans
}
