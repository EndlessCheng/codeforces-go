请看 [视频讲解](https://www.bilibili.com/video/BV1Fc411R7xA/) 第三题。

## 前置知识：树形 DP

请看视频讲解 [树形 DP【基础算法精讲 24】](https://www.bilibili.com/video/BV1vu4y1f7dn/)

## 思路

正难则反，先把所有 $\textit{values}[i]$ 加到答案中，然后考虑哪些 $\textit{values}[i]$ 不能选（撤销，不加入答案）。

设当前节点为 $x$，计算以 $x$ 为根的子树是健康时，失去的最小分数。那么答案就是 $\textit{values}$ 的元素和，减去「以 $0$ 为根的子树是健康时，**失去**的最小分数」。

用「**选或不选**」分类讨论：

- 第一种情况：失去 $\textit{values}[x]$，也就是不加入答案，那么 $x$ 的所有子孙节点都可以加入答案，失去的最小分数就是 $\textit{values}[x]$。
- 第二种情况：$\textit{values}[x]$ 加入答案，问题变成「以 $y$ 为根的子树是健康时，失去的最小分数」，这里 $y$ 是 $x$ 的儿子。如果有多个儿子，累加失去的最小分数。

这两种情况取最小值。注意第一种情况是不会往下递归的，所以当我们递归到叶子的时候，叶子一定不能加入答案，此时直接返回 $\textit{values}[x]$。

代码实现时，为了方便判断 $x$ 是否为叶子节点，可以假设还有一条 $0$ 到 $-1$ 的边，这样不会误把根节点 $0$ 当作叶子。 

```py [sol-Python3]
class Solution:
    def maximumScoreAfterOperations(self, edges: List[List[int]], values: List[int]) -> int:
        g = [[] for _ in values]
        g[0].append(-1)  # 避免误把根节点当作叶子
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        # dfs(x, fa) 计算以 x 为根的子树是健康时，失去的最小分数
        def dfs(x: int, fa: int) -> int:
            if len(g[x]) == 1:  # x 是叶子
                return values[x]
            loss = 0  # 第二种情况
            for y in g[x]:
                if y != fa:
                    loss += dfs(y, x)  # 计算以 y 为根的子树是健康时，失去的最小分数
            return min(values[x], loss)  # 两种情况取最小值
        return sum(values) - dfs(0, -1)
```

```java [sol-Java]
class Solution {
    public long maximumScoreAfterOperations(int[][] edges, int[] values) {
        List<Integer>[] g = new ArrayList[values.length];
        Arrays.setAll(g, e -> new ArrayList<>());
        g[0].add(-1); // 避免误把根节点当作叶子
        for (int[] e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }

        // 先把所有分数加入答案
        long ans = 0;
        for (int v : values) {
            ans += v;
        }
        return ans - dfs(0, -1, g, values);
    }

    // dfs(x) 计算以 x 为根的子树是健康时，失去的最小分数
    private long dfs(int x, int fa, List<Integer>[] g, int[] values) {
        if (g[x].size() == 1) { // x 是叶子
            return values[x];
        }
        long loss = 0; // 第二种情况
        for (int y : g[x]) {
            if (y != fa) {
                loss += dfs(y, x, g, values); // 计算以 y 为根的子树是健康时，失去的最小分数
            }
        }
        return Math.min(values[x], loss); // 两种情况取最小值
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumScoreAfterOperations(vector<vector<int>> &edges, vector<int> &values) {
        vector<vector<int>> g(values.size());
        g[0].push_back(-1); // 避免误把根节点当作叶子
        for (auto &e: edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        // dfs(x, fa) 计算以 x 为根的子树是健康时，失去的最小分数
        function<long long(int, int)> dfs = [&](int x, int fa) -> long long {
            if (g[x].size() == 1) { // x 是叶子
                return values[x];
            }
            long long loss = 0; // 第二种情况
            for (int y: g[x]) {
                if (y != fa) {
                    loss += dfs(y, x); // 计算以 y 为根的子树是健康时，失去的最小分数
                }
            }
            return min((long long) values[x], loss); // 两种情况取最小值
        };
        return accumulate(values.begin(), values.end(), 0LL) - dfs(0, -1);
    }
};
```

```go [sol-Go]
func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	g := make([][]int, len(values))
	g[0] = append(g[0], -1) // 避免误把根节点当作叶子
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	total := 0
	// dfs(x, fa) 计算以 x 为根的子树是健康时，失去的最小分数
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		total += values[x]
		if len(g[x]) == 1 { // x 是叶子
			return values[x]
		}
		loss := 0 // 第二种情况
		for _, y := range g[x] {
			if y != fa {
				loss += dfs(y, x) // 计算以 y 为根的子树是健康时，失去的最小分数
			}
		}
		return min(values[x], loss) // 两种情况取最小值
	}
	return int64(total - dfs(0, -1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{values}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目

- [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/)

更多题目见【基础算法精讲】视频简介中的课后题。
