package main

/* 子集动态规划

定义状态 $f[s]$ 表示完成任务集合 $s$ 的所有任务所需要的最少数目的工作时间段。

枚举 $s$ 的子集 $\textit{sub}$，若完成 $\textit{sub}$ 的所有任务耗时不超过 $\textit{sessionTime}$，则可以将 $f[s\setminus\textit{sub}]+1$ 转移到 $f[s]$ 上，二者取最小值。这里 $s\setminus\textit{sub}$ 表示从 $s$ 中去掉 $\textit{sub}$ 的剩余集合。

通过预处理 $\textit{tasks}$ 的每个子集的子集和，可以做到 $O(1)$ 判断 $\textit{sub}$ 的所有任务耗时不超过 $\textit{sessionTime}$。

时间复杂度：$O(2^n+3^n)$。预处理子集和耗时 $O(2^n)$；状态转移次数为 $s$ 的每个子集的子集个数之和，由于元素个数为 $k$ 的集合有 $C(n,k)$ 个，其子集有 $2^k$ 个，根据二项式定理，$\sum C(n,k)2^k = (2+1)^n = 3^n$，故状态转移耗时 $O(3^n)$。

相似题目：

- [1494. 并行课程 II](https://leetcode-cn.com/problems/parallel-courses-ii/)
- [1655. 分配重复整数](https://leetcode-cn.com/problems/distribute-repeating-integers)

*/

// github.com/EndlessCheng/codeforces-go
func minSessions(tasks []int, sessionTime int) (ans int) {
	n := len(tasks)
	m := 1 << n
	// 预处理所有子集的子集和，复杂度 O(1+2+4+...+2^(n-1)) = O(2^n)
	sum := make([]int, m)
	for i, t := range tasks {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + t
		}
	}
	f := make([]int, m)
	for i := range f {
		f[i] = n
	}
	f[0] = 0
	for s := range f {
		// 枚举 s 的所有子集 sub，若 sub 耗时不超过 sessionTime，则将 f[s^sub] 转移到 f[s] 上
		for sub := s; sub > 0; sub = (sub - 1) & s {
			if sum[sub] <= sessionTime && f[s^sub]+1 < f[s] {
				f[s] = f[s^sub] + 1
			}
		}
	}
	return f[m-1]
}
