package main

/* 模拟

根据题意，缺失的投掷数据之和为 $\textit{miss}=(n+m)\cdot \textit{mean}-\sum \textit{rolls}$

$\textit{miss}$ 最小为 $n$，对应均投掷为 $1$ 的情况，最大为 $6n$，对应均投掷为 $6$ 的情况，超出该范围则不存在答案。

若答案存在，初始化答案数组为 $1$，将剩余 $\textit{miss}-n$ 填充至答案中即可。

*/

// github.com/EndlessCheng/codeforces-go
func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for _, v := range rolls {
		sum += v
	}
	miss := (n+len(rolls))*mean - sum
	if miss < n || miss > n*6 {
		return nil
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	miss -= n
	for i := range ans {
		if miss < 6 {
			ans[i] += miss
			break
		}
		ans[i] = 6
		miss -= 5
	}
	return ans
}
