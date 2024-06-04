枚举服务器 $c$，作为树的根，按照下图计算（下图 $c=0$）。

![b125C.png](https://pic.leetcode.cn/1709427910-nOCIAc-b125C.png)

**优化**：如果 $c$ 只有一个邻居，则答案为 $0$，不调用 $\text{dfs}$。

[视频讲解](https://www.bilibili.com/video/BV1AU411F7Fp/) 第三题。

```py [sol-Python3]
class Solution:
    def countPairsOfConnectableServers(self, edges: List[List[int]], signalSpeed: int) -> List[int]:
        n = len(edges) + 1
        g = [[] for _ in range(n)]
        for x, y, wt in edges:
            g[x].append((y, wt))
            g[y].append((x, wt))

        def dfs(x: int, fa: int, s: int) -> int:
            cnt = 0 if s % signalSpeed else 1
            for y, wt in g[x]:
                if y != fa:
                    cnt += dfs(y, x, s + wt)
            return cnt

        ans = [0] * n
        for i, gi in enumerate(g):
            if len(gi) == 1:
                continue
            s = 0
            for y, wt in gi:
                cnt = dfs(y, i, wt)
                ans[i] += cnt * s
                s += cnt
        return ans
```

```java [sol-Java]
class Solution {
    public int[] countPairsOfConnectableServers(int[][] edges, int signalSpeed) {
        int n = edges.length + 1;
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            int wt = e[2];
            g[x].add(new int[]{y, wt});
            g[y].add(new int[]{x, wt});
        }

        int[] ans = new int[n];
        for (int i = 0; i < n; i++) {
            if (g[i].size() == 1) {
                continue;
            }
            int sum = 0;
            for (int[] e : g[i]) {
                int cnt = dfs(e[0], i, e[1], g, signalSpeed);
                ans[i] += cnt * sum;
                sum += cnt;
            }
        }
        return ans;
    }

    private int dfs(int x, int fa, int sum, List<int[]>[] g, int signalSpeed) {
        int cnt = sum % signalSpeed == 0 ? 1 : 0;
        for (int[] e : g[x]) {
            int y = e[0];
            if (y != fa) {
                cnt += dfs(y, x, sum + e[1], g, signalSpeed);
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> countPairsOfConnectableServers(vector<vector<int>> &edges, int signalSpeed) {
        int n = edges.size() + 1;
        vector<vector<pair<int, int>>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            g[x].push_back({y, wt});
            g[y].push_back({x, wt});
        }

        function<int(int, int, int)> dfs = [&](int x, int fa, int sum) -> int {
            int cnt = sum % signalSpeed == 0;
            for (auto &[y, wt] : g[x]) {
                if (y != fa) {
                    cnt += dfs(y, x, sum + wt);
                }
            }
            return cnt;
        };

        vector<int> ans(n);
        for (int i = 0; i < n; i++) {
            if (g[i].size() == 1) {
                continue;
            }
            int sum = 0;
            for (auto &[y, wt] : g[i]) {
                int cnt = dfs(y, i, wt);
                ans[i] += cnt * sum;
                sum += cnt;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	ans := make([]int, n)
	for i, gi := range g {
		if len(gi) == 1 {
			continue
		}
		var cnt int
		var dfs func(int, int, int)
		dfs = func(x, fa, sum int) {
			if sum%signalSpeed == 0 {
				cnt++
			}
			for _, e := range g[x] {
				if e.to != fa {
					dfs(e.to, x, sum+e.wt)
				}
			}
			return
		}
		sum := 0
		for _, e := range gi {
			cnt = 0
			dfs(e.to, i, e.wt)
			ans[i] += cnt * sum
			sum += cnt
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{edges}$ 的长度加一。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [2867. 统计树中的合法路径数目](https://leetcode.cn/problems/count-valid-paths-in-a-tree/)

## 思考题

如果 $n=10^5,\ \textit{signalSpeed}=10$，你能想出一个更快的做法吗？

- 可以参考 [CF791D Bear and Tree Jumps](https://codeforces.com/problemset/problem/791/D)

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
