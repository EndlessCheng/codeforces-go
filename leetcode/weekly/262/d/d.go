package main

import (
	"math"
	"sort"
)

/* 折半枚举+排序+二分

两个数组和之差可以视作从 $\textit{nums}$ 中选 $n$ 个数取正号，其余 $n$ 个数取负号，然后求元素和。

我们可以使用折半枚举的方法，枚举 $\textit{nums}$ 的前 $n$ 个元素取正或取负的所有情况，按取正个数分组，并按照元素和排序。然后枚举 $\textit{nums}$ 的后 $n$ 个元素取正或取负的所有情况，然后去对应的组里找元素和最近的数，答案即为所有情况中最小的差值。

*/

// github.com/EndlessCheng/codeforces-go
func minimumDifference(nums []int) int {
	n := len(nums) / 2
	a := nums[:n]
	res := make([][]int, n+1)
	for i := 0; i < 1<<n; i++ {
		sum, cnt := 0, 0
		for j, v := range a {
			if i>>j&1 > 0 { // 1 视作取正
				sum += v
				cnt++
			} else { // 0 视作取负
				sum -= v
			}
		}
		res[cnt] = append(res[cnt], sum) // 按照取正的个数将元素和分组
	}

	for _, b := range res {
		sort.Ints(b) // 排序，方便下面二分
	}

	ans := math.MaxInt64
	a = nums[n:]
	for i := 0; i < 1<<n; i++ {
		sum, cnt := 0, 0
		for j, v := range a {
			if i>>j&1 == 0 { // 0 视作取正
				sum += v
				cnt++
			} else { // 1 视作取负
				sum -= v
			}
		}
		// 在对应的组里二分最近的数
		b := res[cnt]
		j := sort.SearchInts(b, sum)
		if j < len(b) {
			ans = min(ans, b[j]-sum)
		}
		if j > 0 {
			ans = min(ans, sum-b[j-1])
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
