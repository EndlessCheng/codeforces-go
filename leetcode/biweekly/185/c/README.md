按题意模拟即可，写一个自底向上的 DFS。

⚠**注意**：计算的是**儿子**的最小最大完成时间，不是子树中的所有节点。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def finishTime(self, n: int, edges: list[list[int]], baseTime: list[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)  # 题目保证 x 是 y 的父节点

        def dfs(x: int) -> int:
            if not g[x]:  # x 是叶子
                return baseTime[x]
            earliest = inf
            latest = 0
            for y in g[x]:
                t = dfs(y)
                earliest = min(earliest, t)
                latest = max(latest, t)
            return latest * 2 - earliest + baseTime[x]

        return dfs(0)
```

```java [sol-Java]
class Solution {
    public long finishTime(int n, int[][] edges, int[] baseTime) {
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        for (int[] e : edges) {
            g[e[0]].add(e[1]); // 题目保证 e[0] 是 e[1] 的父节点
        }
        return dfs(0, g, baseTime);
    }

    private long dfs(int x, List<Integer>[] g, int[] baseTime) {
        if (g[x].isEmpty()) { // x 是叶子
            return baseTime[x];
        }
        long earliest = Long.MAX_VALUE;
        long latest = 0;
        for (int y : g[x]) {
            long t = dfs(y, g, baseTime);
            earliest = Math.min(earliest, t);
            latest = Math.max(latest, t);
        }
        return latest * 2 - earliest + baseTime[x];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long finishTime(int n, vector<vector<int>>& edges, vector<int>& baseTime) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            g[e[0]].push_back(e[1]); // 题目保证 e[0] 是 e[1] 的父节点
        }

        auto dfs = [&](this auto&& dfs, int x) -> long long {
            if (g[x].empty()) { // x 是叶子
                return baseTime[x];
            }
            long long earliest = LLONG_MAX, latest = 0;
            for (int y : g[x]) {
                long long t = dfs(y);
                earliest = min(earliest, t);
                latest = max(latest, t);
            }
            return latest * 2 - earliest + baseTime[x];
        };

        return dfs(0);
    }
};
```

```go [sol-Go]
func finishTime(n int, edges [][]int, baseTime []int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y) // 题目保证 x 是 y 的父节点
	}

	var dfs func(int) int
	dfs = func(x int) int {
		if g[x] == nil { // x 是叶子
			return baseTime[x]
		}
		earliest, latest := math.MaxInt, 0
		for _, y := range g[x] {
			t := dfs(y)
			earliest = min(earliest, t)
			latest = max(latest, t)
		}
		return latest*2 - earliest + baseTime[x]
	}

	return int64(dfs(0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面树题单的「**§3.3 自底向上 DFS**」。

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
