#### 提示 1

由于价值都是正数，因此价值和最小的一条路径一定**只有一个点**。

#### 提示 2

根据提示 1，「价值和最大的一条路径与最小的一条路径的差值」等价于「去掉路径的一个端点」。

#### 提示 3

由于价值都是正数，一条路径能延长就尽量延长，这样路径和就越大，那么最优是延长到叶子。

根据提示 2，问题转换成**去掉一个叶子**后的**最大路径和**（这里的叶子严格来说是度为 $1$ 的点，因为根的度数也可能是 $1$）。

#### 提示 4

最大路径和是一个经典树形 DP 问题，类似「树的直径」。由于我们需要去掉一个叶子，那么可以让子树返回两个值：

- 带叶子的最大路径和；
- 不带叶子的最大路径和。

对于当前节点，它有多棵子树，我们一棵棵 DFS，假设当前 DFS 完了其中一棵子树，它返回了「当前带叶子的路径和」和「当前不带叶子的路径和」，那么答案有两种情况：

- 前面最大带叶子的路径和 + 当前不带叶子的路径和；
- 前面最大不带叶子的路径和 + 当前带叶子的路径和；

然后更新「最大带叶子的路径和」和「最大不带叶子的路径和」。

最后返回「最大带叶子的路径和」和「最大不带叶子的路径和」，用来供父节点计算。

附：[视频讲解](https://www.bilibili.com/video/BV1QT41127kJ/)。

```py [sol1-Python3]
class Solution:
    def maxOutput(self, n: int, edges: List[List[int]], price: List[int]) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建树

        ans = 0
        # 返回带叶子的最大路径和，不带叶子的最大路径和
        def dfs(x: int, fa: int) -> (int, int):
            nonlocal ans
            max_s1 = p = price[x]
            max_s2 = 0
            for y in g[x]:
                if y == fa: continue
                s1, s2 = dfs(y, x)
                # 前面最大带叶子的路径和 + 当前不带叶子的路径和
                # 前面最大不带叶子的路径和 + 当前带叶子的路径和
                ans = max(ans, max_s1 + s2, max_s2 + s1)
                max_s1 = max(max_s1, s1 + p)
                max_s2 = max(max_s2, s2 + p)  # 这里加上 p 是因为 x 必然不是叶子
            return max_s1, max_s2
        dfs(0, -1)
        return ans
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private int[] price;
    private long ans;

    public long maxOutput(int n, int[][] edges, int[] price) {
        this.price = price;
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建树
        }
        dfs(0, -1);
        return ans;
    }

    // 返回带叶子的最大路径和，不带叶子的最大路径和
    private long[] dfs(int x, int fa) {
        long p = price[x], maxS1 = p, maxS2 = 0;
        for (var y : g[x])
            if (y != fa) {
                var res = dfs(y, x);
                long s1 = res[0], s2 = res[1];
                // 前面最大带叶子的路径和 + 当前不带叶子的路径和
                // 前面最大不带叶子的路径和 + 当前带叶子的路径和
                ans = Math.max(ans, Math.max(maxS1 + s2, maxS2 + s1));
                maxS1 = Math.max(maxS1, s1 + p);
                maxS2 = Math.max(maxS2, s2 + p); // 这里加上 p 是因为 x 必然不是叶子
            }
        return new long[]{maxS1, maxS2};
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long maxOutput(int n, vector<vector<int>> &edges, vector<int> &price) {
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建树
        }

        long ans = 0;
        // 返回带叶子的最大路径和，不带叶子的最大路径和
        function<pair<long, long>(int, int)> dfs = [&](int x, int fa) -> pair<long, long> {
            long p = price[x], max_s1 = p, max_s2 = 0;
            for (int y : g[x])
                if (y != fa) {
                    auto[s1, s2] = dfs(y, x);
                    // 前面最大带叶子的路径和 + 当前不带叶子的路径和
                    // 前面最大不带叶子的路径和 + 当前带叶子的路径和
                    ans = max(ans, max(max_s1 + s2, max_s2 + s1));
                    max_s1 = max(max_s1, s1 + p);
                    max_s2 = max(max_s2, s2 + p); // 这里加上 p 是因为 x 必然不是叶子
                }
            return {max_s1, max_s2};
        };
        dfs(0, -1);
        return ans;
    }
};
```

```go [sol1-Go]
func maxOutput(n int, edges [][]int, price []int) int64 {
	ans := 0
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}
	// 返回带叶子的最大路径和，不带叶子的最大路径和
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		p := price[x]
		maxS1, maxS2 := p, 0
		for _, y := range g[x] {
			if y != fa {
				s1, s2 := dfs(y, x)
				// 前面最大带叶子的路径和 + 当前不带叶子的路径和
				// 前面最大不带叶子的路径和 + 当前带叶子的路径和
				ans = max(ans, max(maxS1+s2, maxS2+s1))
				maxS1 = max(maxS1, s1+p)
				maxS2 = max(maxS2, s2+p) // 这里加上 p 是因为 x 必然不是叶子
			}
		}
		return maxS1, maxS2
	}
	dfs(0, -1)
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$。
- 空间复杂度：$O(n)$。

#### 相似题目

- [124. 二叉树中的最大路径和](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)
- [1245. 树的直径](https://leetcode-cn.com/problems/tree-diameter/)
- [2246. 相邻字符不同的最长路径](https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/)
