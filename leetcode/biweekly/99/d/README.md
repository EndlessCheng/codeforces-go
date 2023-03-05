如果只求以 $0$ 为根时的猜对次数 $\textit{cnt}_0$，那么把 $\textit{guesses}$ 转成哈希表，DFS 一次这棵树就可以算出来。

如果要枚举以每个点为根时的猜对次数，暴力做法就太慢了，怎么优化呢？

注意到，如果节点 $x$ 和 $y$ 之间有边，那么从「以 $x$ 为根的树」变成「以 $y$ 为根的树」，就只有 $[x,y]$ 和 $[y,x]$ 这两个猜测的正确性变了，其余猜测的正确性不变。

因此，从 $0$ 出发，再次 DFS 这棵树，从节点 $x$ 递归到节点 $y$ 时：

- 如果有猜测 $[x,y]$，那么猜对次数减一；
- 如果有猜测 $[y,x]$，那么猜对次数加一。

DFS 的同时，统计猜对次数 $\ge k$ 的节点个数，即为答案。

这个套路叫做「换根 DP」。

附：[视频讲解](https://www.bilibili.com/video/BV1dY4y1C77x/)

```py [sol1-Python3]
class Solution:
    def rootCount(self, edges: List[List[int]], guesses: List[List[int]], k: int) -> int:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)  # 建图

        s = {(x, y) for x, y in guesses}  # guesses 转成哈希表 s

        ans = cnt0 = 0
        def dfs(x: int, fa: int) -> None:
            nonlocal cnt0
            for y in g[x]:
                if y != fa:
                    cnt0 += (x, y) in s  # 以 0 为根时，猜对了
                    dfs(y, x)
        dfs(0, -1)

        def reroot(x: int, fa: int, cnt: int) -> None:
            nonlocal ans
            ans += cnt >= k  # 此时 cnt 就是以 x 为根时的猜对次数
            for y in g[x]:
                if y != fa:
                    reroot(y, x, cnt - ((x, y) in s) + ((y, x) in s))
        reroot(0, -1, cnt0)
        return ans
```

```java [sol1-Java]
class Solution {
    private List<Integer>[] g;
    private Set<Long> s = new HashSet<>();
    private int k, ans, cnt0;

    public int rootCount(int[][] edges, int[][] guesses, int k) {
        this.k = k;
        g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, e -> new ArrayList<>());
        for (var e : edges) {
            int x = e[0], y = e[1];
            g[x].add(y);
            g[y].add(x); // 建图
        }

        for (var e : guesses) // guesses 转成哈希表
            s.add((long) e[0] << 32 | e[1]); // 两个 4 字节数压缩成一个 8 字节数

        dfs(0, -1);
        reroot(0, -1, cnt0);
        return ans;
    }

    private void dfs(int x, int fa) {
        for (var y : g[x])
            if (y != fa) {
                if (s.contains((long) x << 32 | y)) // 以 0 为根时，猜对了
                    ++cnt0;
                dfs(y, x);
            }
    }

    private void reroot(int x, int fa, int cnt) {
        if (cnt >= k) ++ans; // 此时 cnt 就是以 x 为根时的猜对次数
        for (var y : g[x])
            if (y != fa) {
                int c = cnt;
                if (s.contains((long) x << 32 | y)) --c; // 原来是对的，现在错了
                if (s.contains((long) y << 32 | x)) ++c; // 原来是错的，现在对了
                reroot(y, x, c);
            }
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int rootCount(vector<vector<int>> &edges, vector<vector<int>> &guesses, int k) {
        vector<vector<int>> g(edges.size() + 1);
        for (auto &e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x); // 建图
        }

        unordered_set<long> s;
        for (auto &e : guesses) // guesses 转成哈希表
            s.insert((long) e[0] << 32 | e[1]); // 两个 4 字节数压缩成一个 8 字节数

        int ans = 0, cnt0 = 0;
        function<void(int, int)> dfs = [&](int x, int fa) {
            for (int y : g[x])
                if (y != fa) {
                    cnt0 += s.count((long) x << 32 | y); // 以 0 为根时，猜对了
                    dfs(y, x);
                }
        };
        dfs(0, -1);

        function<void(int, int, int)> reroot = [&](int x, int fa, int cnt) {
            ans += cnt >= k; // 此时 cnt 就是以 x 为根时的猜对次数
            for (int y : g[x])
                if (y != fa) {
                    reroot(y, x, cnt
                                 - s.count((long) x << 32 | y) // 原来是对的，现在错了
                                 + s.count((long) y << 32 | x)); // 原来是错的，现在对了
                }
        };
        reroot(0, -1, cnt0);
        return ans;
    }
};
```

```go [sol1-Go]
func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时，猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{edges}$ 的长度加一，$m$ 为 $\textit{guesses}$ 的长度。
- 空间复杂度：$O(n+m)$。

### 相似题目

- [834. 树中距离之和](https://leetcode.cn/problems/sum-of-distances-in-tree/)
- [310. 最小高度树](https://leetcode.cn/problems/minimum-height-trees/)

更多题目，可以看我的 [算法竞赛模板库](https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/dp.go#L2607) 中的「换根 DP」。

### 思考题

如果把「$u$ 是 $v$ 的父节点」改成「$u$ 是 $v$ 的**祖先节点**」，要怎么做呢？

如果改成「$\textit{guesses}[i]$ 猜对会得到 $\textit{score}[i]$ 分，计算的是以每个点为根时的得分之和」，要怎么做呢？（本题相当于 $\textit{score}[i]$ 均为 $1$）
