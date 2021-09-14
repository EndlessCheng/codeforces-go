package main

/* 前缀和优化 DP

根据题意，首次进入一个房间时，下一天是一定要回到 $\textit{nextVisit}[i]$ 房间的，下文简称为「回访」。

于是定义状态 $f[i]$ 表示从房间 $i$ 回访到房间 $\textit{nextVisit}[i]$（记为 $j$），再重新访问到房间 $i$ 时所需要的天数。

根据题意，如果从房间 $i$ 回访到房间 $j$，那我们需要回访 $[j,i-1]$ 范围内的每个房间。同时我们还需要从 $j$ 走到 $i$，花费 $i-j+1$ 天。于是有转移方程：

$$
f[i] = i-j+1 + \sum_{k=j}^{i-1} f[k]
$$

其中和式可以用前缀和优化，这样单次转移就是 $O(1)$ 的。

代码实现时，可以略去数组 $f$，直接将其记录在前缀和 $\textit{sum}$ 中。

最后，加上首次访问每个房间所需的天数 $n$，最后答案为 $\textit{sum}[n-1]+n-1$（减 $1$ 是因为天数从 $0$ 开始）

*/

// github.com/EndlessCheng/codeforces-go
func firstDayBeenInAllRooms(nextVisit []int) int {
	const mod int = 1e9 + 7
	n := len(nextVisit)
	sum := make([]int, n)
	for i, j := range nextVisit[:n-1] {
		v := (i - j + 1 + sum[i] - sum[j] + mod) % mod
		sum[i+1] = sum[i] + v
	}
	return (sum[n-1] + n - 1) % mod
}
