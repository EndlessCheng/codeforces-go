请看 [视频讲解](https://www.bilibili.com/video/BV1oC4y1o7Tz/) 第四题。

如果一条边左右两侧的点权和都是 $k$ 的倍数，那么这条边就可以删除。

由于题目保证 $\textit{values}$ 之和可以被 $k$ 整除。那么只需要看一侧的点权和是否为 $k$ 的倍数。

换言之，可以从任意点出发 DFS，只要发现子树的点权和是 $k$ 的倍数，就说明子树到上面父节点的这条边是可以删除的。

```py [sol-Python3]
class Solution:
    def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
        g = [[] for _ in range(n)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        ans = 0
        def dfs(x: int, fa: int) -> int:
            s = values[x]
            for y in g[x]:
                if y != fa:
                    s += dfs(y, x)
            nonlocal ans
            ans += s % k == 0
            return s
        dfs(0, -1)
        return ans
```

```java [sol-Java]
class Solution {
    private List<Integer>[] g;
    private int[] values;
    private int k, ans;

    public int maxKDivisibleComponents(int n, int[][] edges, int[] values, int k) {
        g = new ArrayList[n];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        this.values = values;
        this.k = k;
        dfs(0, -1);
        return ans;
    }

    private long dfs(int x, int fa) {
        long sum = values[x];
        for (int y : g[x]) {
            if (y != fa) {
                sum += dfs(y, x);
            }
        }
        ans += sum % k == 0 ? 1 : 0;
        return sum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxKDivisibleComponents(int n, vector<vector<int>>& edges, vector<int>& values, int k) {
        vector<vector<int>> g(n);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }

        int ans = 0;
        function<long long(int, int)> dfs = [&](int x, int fa) -> long long {
            long long sum = values[x];
            for (int y : g[x])
                if (y != fa)
                    sum += dfs(y, x);
            ans += sum % k == 0;
            return sum;
        };
        dfs(0, -1);
        return ans;
    }
};
```

```go [sol-Go]
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		s := values[x]
		for _, y := range g[x] {
			if y != fa {
				s += dfs(y, x)
			}
		}
		if s%k == 0 {
			ans++
		}
		return s
	}
	dfs(0, -1)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [2440. 创建价值相同的连通块](https://leetcode.cn/problems/create-components-with-same-value/)
