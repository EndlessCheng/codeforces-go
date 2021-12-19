package main

/* 分组循环

将 $\textit{prices}$ 按照平滑下降的定义分成若干组。例如 $[3,2,1,4]$ 分为 $[3,2,1]$ 和 $[4]$ 两组。

对于每一组的所有非空子数组，都是平滑下降的。设该组长度为 $m$，则该组的非空子数组个数为

$$
C_{m+1}^2 = \dfrac{m(m+1)}{2}
$$

累加每组的非空子区间个数即为答案。

- 时间复杂度：$O(n)$，其中 $n$ 是数组 $\textit{prices}$ 的长度。注意下面代码内外层循环共用同一个变量 $i$，时间复杂度就是 `i++` 的执行次数，即 $O(n)$。
- 空间复杂度：$O(1)$，我们只需要常数的空间保存若干变量。

*/

// github.com/EndlessCheng/codeforces-go
func getDescentPeriods(prices []int) (ans int64) {
	for i, n := 0, len(prices); i < n; {
		i0 := i
		for i++; i < n && prices[i] == prices[i-1]-1; i++ {
		}
		ans += int64(i-i0) * int64(i-i0+1) / 2
	}
	return
}
