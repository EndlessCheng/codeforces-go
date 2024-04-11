分类讨论：

- $s$ 和 $t$ 不在同一个连通块中。答案是 $-1$。
- $s$ 和 $t$ 在同一个连通块中。由于 AND 的性质是 **AND 的数字越多，结果越小**。在可以重复经过边的前提下，最优方案是把 $s$ 所在连通块内的边都走一遍。

所以我们需要知道 $s$ 和 $t$ 在哪个连通块，以及连通块内边权的 AND 是多少。

这可以用 DFS 或者并查集实现。请看 [视频讲解](https://www.bilibili.com/video/BV1ut421H7Wv/) 第四题，欢迎点赞关注！

代码实现时，可以把 AND 的初始值设为 $-1$，因为其二进制中的数都是 $1$，与任何 $x$ 求 AND 的结果都是 $x$。

## 方法一：DFS

```py [sol-Python3]
class Solution:
    def minimumCost(self, n: int, edges: List[List[int]], query: List[List[int]]) -> List[int]:
        g = [[] for _ in range(n)]
        for x, y, w in edges:
            g[x].append((y, w))
            g[y].append((x, w))

        def dfs(x: int) -> int:
            and_ = -1
            ids[x] = len(cc_and)  # 记录每个点所在连通块的编号
            for y, w in g[x]:
                and_ &= w
                if ids[y] < 0:  # 没有访问过
                    and_ &= dfs(y)
            return and_

        ids = [-1] * n  # 记录每个点所在连通块的编号
        cc_and = []  # 记录每个连通块的边权的 AND
        for i in range(n):
            if ids[i] < 0:
                cc_and.append(dfs(i))

        return [-1 if ids[s] != ids[t] else cc_and[ids[s]]
                for s, t in query]
```

```java [sol-Java]
class Solution {
    public int[] minimumCost(int n, int[][] edges, int[][] query) {
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].add(new int[]{y, w});
            g[y].add(new int[]{x, w});
        }

        int[] ids = new int[n]; // 记录每个点所在连通块的编号
        Arrays.fill(ids, -1);
        List<Integer> ccAnd = new ArrayList<>(); // 记录每个连通块的边权的 AND
        for (int i = 0; i < n; i++) {
            if (ids[i] < 0) {
                ccAnd.add(dfs(i, ccAnd.size(), g, ids));
            }
        }

        int[] ans = new int[query.length];
        for (int i = 0; i < query.length; i++) {
            int s = query[i][0], t = query[i][1];
            ans[i] = ids[s] != ids[t] ? -1 : ccAnd.get(ids[s]);
        }
        return ans;
    }

    private int dfs(int x, int curId, List<int[]>[] g, int[] ids) {
        ids[x] = curId; // 记录每个点所在连通块的编号
        int and = -1;
        for (int[] e : g[x]) {
            and &= e[1];
            if (ids[e[0]] < 0) { // 没有访问过
                and &= dfs(e[0], curId, g, ids);
            }
        }
        return and;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<pair<int, int>>> g;
    vector<int> cc_and, ids;

    int dfs(int x) {
        ids[x] = cc_and.size(); // 记录每个点所在连通块的编号
        int and_ = -1;
        for (auto &[y, w]: g[x]) {
            and_ &= w;
            if (ids[y] < 0) { // 没有访问过
                and_ &= dfs(y);
            }
        }
        return and_;
    }

public:
    vector<int> minimumCost(int n, vector<vector<int>> &edges, vector<vector<int>> &query) {
        g.resize(n);
        for (auto &e: edges) {
            int x = e[0], y = e[1], w = e[2];
            g[x].emplace_back(y, w);
            g[y].emplace_back(x, w);
        }

        ids.resize(n, -1); // 记录每个点所在连通块的编号
        for (int i = 0; i < n; i++) {
            if (ids[i] < 0) { // 没有访问过
                cc_and.push_back(dfs(i)); // 记录每个连通块的边权的 AND
            }
        }

        vector<int> ans;
        ans.reserve(query.size()); // 预分配空间
        for (auto &q: query) {
            int s = q[0], t = q[1];
            ans.push_back(ids[s] != ids[t] ? -1 : cc_and[ids[s]]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(n int, edges, query [][]int) []int {
	type edge struct{ to, w int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	ids := make([]int, n) // 记录每个点所在连通块的编号
	for i := range ids {
		ids[i] = -1
	}
	ccAnd := []int{} // 记录每个连通块的边权的 AND
	var dfs func(int) int
	dfs = func(x int) int {
		ids[x] = len(ccAnd) // 记录每个点所在连通块的编号
		and := -1
		for _, e := range g[x] {
			and &= e.w
			if ids[e.to] < 0 { // 没有访问过
				and &= dfs(e.to)
			}
		}
		return and
	}
	for i, id := range ids {
		if id < 0 { // 没有访问过
			ccAnd = append(ccAnd, dfs(i)) // 记录每个连通块的边权的 AND
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if ids[s] != ids[t] {
			ans[i] = -1
		} else {
			ans[i] = ccAnd[ids[s]]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m+q)$，其中 $m$ 为 $\textit{edges}$ 的长度，$q$ 为 $\textit{query}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。返回值不计入。

## 方法二：并查集

```py [sol-Python3]
class Solution:
    def minimumCost(self, n: int, edges: List[List[int]], query: List[List[int]]) -> List[int]:
        fa = list(range(n))
        and_ = [-1] * n

        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]

        for x, y, w in edges:
            x = find(x)
            y = find(y)
            and_[y] &= w
            if x != y:
                and_[y] &= and_[x]
                fa[x] = y

        return [-1 if find(s) != find(t) else and_[find(s)]
                for s, t in query]
```

```java [sol-Java]
class Solution {
    public int[] minimumCost(int n, int[][] edges, int[][] query) {
        int[] fa = new int[n];
        for (int i = 0; i < n; i++) {
            fa[i] = i;
        }
        int[] and = new int[n];
        Arrays.fill(and, -1);

        for (int[] e : edges) {
            int x = find(e[0], fa);
            int y = find(e[1], fa);
            and[y] &= e[2];
            if (x != y) {
                and[y] &= and[x];
                fa[x] = y;
            }
        }

        int[] ans = new int[query.length];
        for (int i = 0; i < query.length; i++) {
            int s = query[i][0], t = query[i][1];
            ans[i] = find(s, fa) != find(t, fa) ? -1 : and[find(s, fa)];
        }
        return ans;
    }

    private int find(int x, int[] fa) {
        if (fa[x] != x) {
            fa[x] = find(fa[x], fa);
        }
        return fa[x];
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> fa, and_;

    int find(int x) {
        return fa[x] == x ? x : fa[x] = find(fa[x]);
    };

public:
    vector<int> minimumCost(int n, vector<vector<int>> &edges, vector<vector<int>> &query) {
        fa.resize(n);
        iota(fa.begin(), fa.end(), 0);
        and_.resize(n, -1);
        for (auto &e: edges) {
            int x = find(e[0]);
            int y = find(e[1]);
            and_[y] &= e[2];
            if (x != y) {
                and_[y] &= and_[x];
                fa[x] = y;
            }
        }

        vector<int> ans;
        ans.reserve(query.size()); // 预分配空间
        for (auto &q: query) {
            int s = q[0], t = q[1];
            ans.push_back(find(s) != find(t) ? -1 : and_[find(s)]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(n int, edges, query [][]int) []int {
	fa := make([]int, n)
	and := make([]int, n)
	for i := range fa {
		fa[i] = i
		and[i] = -1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for _, e := range edges {
		x, y := find(e[0]), find(e[1])
		and[y] &= e[2]
		if x != y {
			and[y] &= and[x]
			fa[x] = y
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if find(s) != find(t) {
			ans[i] = -1
		} else {
			ans[i] = and[find(s)]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m+q)\log n)$，其中 $m$ 为 $\textit{edges}$ 的长度，$q$ 为 $\textit{query}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 相关题目

[图论题单](https://leetcode.cn/circle/discuss/01LUak/) 中的 DFS。

## 其它题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
