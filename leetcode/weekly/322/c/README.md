由于路径可以折返，节点 $1$ 所在连通块的每条边都可以在路径上，所以答案为节点 $1$ 所在连通块中的最小的 $\textit{distance}_i$。

```py [sol-Python3]
class Solution:
    def minScore(self, n: int, roads: List[List[int]]) -> int:
        g = [[] for _ in range(n + 1)]
        for x, y, dis in roads:
            g[x].append((y, dis))
            g[y].append((x, dis))

        vis = [False] * (n + 1)
        ans = inf

        def dfs(x: int) -> None:
            nonlocal ans
            vis[x] = True  # 避免重复访问
            for y, dis in g[x]:
                ans = min(ans, dis)
                if not vis[y]:
                    dfs(y)

        # 遍历节点 1 所在连通块
        dfs(1)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans = Integer.MAX_VALUE;

    public int minScore(int n, int[][] roads) {
        List<int[]>[] g = new ArrayList[n + 1];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : roads) {
            int x = e[0], y = e[1], dis = e[2];
            g[x].add(new int[]{y, dis});
            g[y].add(new int[]{x, dis});
        }

        boolean[] vis = new boolean[n + 1];
        // 遍历节点 1 所在连通块
        dfs(1, g, vis);
        return ans;
    }

    private void dfs(int x, List<int[]>[] g, boolean[] vis) {
        vis[x] = true; // 避免重复访问
        for (int[] e : g[x]) {
            ans = Math.min(ans, e[1]);
            if (!vis[e[0]]) {
                dfs(e[0], g, vis);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minScore(int n, vector<vector<int>>& roads) {
        vector<vector<pair<int, int>>> g(n + 1);
        for (auto& e : roads) {
            int x = e[0], y = e[1], dis = e[2];
            g[x].emplace_back(y, dis);
            g[y].emplace_back(x, dis);
        }

        vector<int8_t> vis(n + 1);
        int ans = INT_MAX;

        auto dfs = [&](this auto&& dfs, int x) -> void {
            vis[x] = true; // 避免重复访问
            for (auto& [y, dis] : g[x]) {
                ans = min(ans, dis);
                if (!vis[y]) {
                    dfs(y);
                }
            }
        };

        // 遍历节点 1 所在连通块
        dfs(1);
        return ans;
    }
};
```

```go [sol-Go]
func minScore(n int, roads [][]int) int {
	type edge struct{ to, dis int }
	g := make([][]edge, n+1)
	for _, e := range roads {
		x, y, dis := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, dis})
		g[y] = append(g[y], edge{x, dis})
	}

	vis := make([]bool, n+1)
	ans := math.MaxInt

	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true // 避免重复访问
		for _, e := range g[x] {
			ans = min(ans, e.dis)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}

	// 遍历节点 1 所在连通块
	dfs(1)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 是 $\textit{roads}$ 的长度。注意创建大小为 $n$ 的数组需要 $\mathcal{O}(n)$ 时间。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 思考题

如果每条边至多访问一次呢？

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面图论题单的「**§1.1 深度优先搜索**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
