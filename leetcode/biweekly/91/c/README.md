下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

1. DFS，求出 $\textit{bob}$ 到 $0$ 的路径上，Bob 经过每个点的时间 $\textit{bobTime}$。
2. DFS，从 $0$ 到每个叶节点，按照题目要求累加 $\textit{amount}$，在叶子节点上更新答案的最大值。

```py [sol1-Python3]
class Solution:
    def mostProfitablePath(self, edges: List[List[int]], bob: int, amount: List[int]) -> int:
        n = len(edges) + 1
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建树

        bob_time = [n] * n  # bobTime[x] 表示 bob 访问节点 x 的时间
        def dfs_bob(x: int, fa: int, t: int) -> bool:
            if x == 0:
                bob_time[x] = t
                return True
            found0 = False
            for y in g[x]:
                if y != fa and dfs_bob(y, x, t + 1):
                    found0 = True
            if found0:
                bob_time[x] = t  # 只有可以到达 0 才标记访问时间
            return found0
        dfs_bob(bob, -1, 0)

        g[0].append(-1)  # 防止把根节点当作叶子
        ans = -inf
        def dfs_alice(x: int, fa: int, alice_time: int, tot: int) -> None:
            if alice_time < bob_time[x]:
                tot += amount[x]
            elif alice_time == bob_time[x]:
                tot += amount[x] // 2
            if len(g[x]) == 1:  # 叶子
                nonlocal ans
                ans = max(ans, tot)  # 更新答案
                return
            for y in g[x]:
                if y != fa:
                    dfs_alice(y, x, alice_time + 1, tot)
        dfs_alice(0, -1, 0, 0)
        return ans
```

```go [sol1-Go]
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	bobTime := make([]int, n) // bobTime[x] 表示 bob 访问节点 x 的时间
	for i := range bobTime {
		bobTime[i] = n // 也可以初始化成 inf
	}
	var dfsBob func(int, int, int) bool
	dfsBob = func(x, fa, t int) bool {
		if x == 0 {
			bobTime[x] = t
			return true
		}
		found0 := false
		for _, y := range g[x] {
			if y != fa && dfsBob(y, x, t+1) {
				found0 = true
			}
		}
		if found0 {
			bobTime[x] = t // 只有可以到达 0 才标记访问时间
		}
		return found0
	}
	dfsBob(bob, -1, 0)

	g[0] = append(g[0], -1) // 防止把根节点当作叶子
	ans := math.MinInt32
	var dfsAlice func(int, int, int, int)
	dfsAlice = func(x, fa, aliceTime, sum int) {
		if aliceTime < bobTime[x] {
			sum += amount[x]
		} else if aliceTime == bobTime[x] {
			sum += amount[x] / 2
		}
		if len(g[x]) == 1 { // 叶子
			ans = max(ans, sum) // 更新答案
			return
		}
		for _, y := range g[x] {
			if y != fa {
				dfsAlice(y, x, aliceTime+1, sum)
			}
		}
	}
	dfsAlice(0, -1, 0, 0)
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为树的节点个数，即 $\textit{edges}$ 的长度加一。
- 空间复杂度：$O(n)$。
