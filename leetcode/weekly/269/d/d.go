package main

import "sort"

/* 按照相同时间分组 + 建图 DFS

首先将 $\textit{meetings}$ 按照时间 $\textit{time}$ 排序，然后遍历 $\textit{meetings}$，将时间相同的会议分成一组。

对于每组会议，由于专家可以互相共享秘密，我们可以将这些专家看成图上的点，专家 $\textit{x_i}$ 与专家 $\textit{y_i}$ 之间连边，我们从所有在这张图上且知道秘密的专家出发，DFS 这张图，将所有能访问到的专家标记成知道秘密的。

最后返回所有知道秘密的专家。

代码实现时，可以用哈希表来存储图上的点和边。

时间复杂度：$O(m\log m)$，这里 $m$ 为 $\textit{meetings}$ 数组的长度。对 $\textit{meetings}$ 的排序的耗时为 $O(m\log m)$，建图和 DFS 的复杂度为 $O(m)$，所以总体复杂度为 $O(m\log m)$。

*/

// github.com/EndlessCheng/codeforces-go
func findAllPeople(_ int, meetings [][]int, firstPerson int) (ans []int) {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] }) // 按照时间排序

	haveSecret := map[int]bool{0: true, firstPerson: true} // 一开始 0 和 firstPerson 都知道秘密
	for i, m := 0, len(meetings); i < m; {
		g := map[int][]int{}
		time := meetings[i][2]
		// 遍历时间相同的会议。注意这里的 i 和外层循环的 i 是同一个变量，所以整个循环部分的时间复杂度是线性的
		for ; i < m && meetings[i][2] == time; i++ {
			v, w := meetings[i][0], meetings[i][1]
			g[v] = append(g[v], w) // 建图
			g[w] = append(g[w], v)
		}

		vis := map[int]bool{} // 避免重复访问专家
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			haveSecret[v] = true
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		for v := range g {
			if haveSecret[v] && !vis[v] { // 从在图上且知道秘密的专家出发，DFS 标记所有能到达的专家
				dfs(v)
			}
		}
	}
	for i := range haveSecret {
		ans = append(ans, i) // 注意可以按任何顺序返回答案
	}
	return
}
