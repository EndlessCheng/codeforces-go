## 方法一：BFS

暴力。每次加边后，重新跑一遍 BFS，求出从 $0$ 到 $n-1$ 的最短路。

### 细节

为避免反复创建 $\textit{vis}$ 数组，可以在 $\textit{vis}$ 中保存当前节点是第几次询问访问的。

```py [sol-Python3]
class Solution:
    def shortestDistanceAfterQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        g = [[i + 1] for i in range(n - 1)]
        vis = [-1] * (n - 1)

        def bfs(i: int) -> int:
            q = deque([0])
            for step in count(1):
                tmp = q
                q = []
                for x in tmp:
                    for y in g[x]:
                        if y == n - 1:
                            return step
                        if vis[y] != i:
                            vis[y] = i
                            q.append(y)
            return -1

        ans = [0] * len(queries)
        for i, (l, r) in enumerate(queries):
            g[l].append(r)
            ans[i] = bfs(i)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] shortestDistanceAfterQueries(int n, int[][] queries) {
        List<Integer>[] g = new ArrayList[n - 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 0; i < n - 1; i++) {
            g[i].add(i + 1);
        }

        int[] ans = new int[queries.length];
        int[] vis = new int[n - 1];
        for (int i = 0; i < queries.length; i++) {
            g[queries[i][0]].add(queries[i][1]);
            ans[i] = bfs(i + 1, g, vis, n);
        }
        return ans;
    }

    private int bfs(int i, List<Integer>[] g, int[] vis, int n) {
        List<Integer> q = new ArrayList<>();
        q.add(0);
        for (int step = 1; ; step++) {
            List<Integer> tmp = q;
            q = new ArrayList<>();
            for (int x : tmp) {
                for (int y : g[x]) {
                    if (y == n - 1) {
                        return step;
                    }
                    if (vis[y] != i) {
                        vis[y] = i;
                        q.add(y);
                    }
                }
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> shortestDistanceAfterQueries(int n, vector<vector<int>>& queries) {
        vector<vector<int>> g(n - 1);
        for (int i = 0; i < n - 1; i++) {
            g[i].push_back(i + 1);
        }
        vector<int> vis(n - 1, -1);

        auto bfs = [&](int i) -> int {
            vector<int> q = {0};
            for (int step = 1; ; step++) {
                vector<int> nxt;
                for (int x : q) {
                    for (int y : g[x]) {
                        if (y == n - 1) {
                            return step;
                        }
                        if (vis[y] != i) {
                            vis[y] = i;
                            nxt.push_back(y);
                        }
                    }
                }
                q = move(nxt);
            }
        };

        vector<int> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            g[queries[i][0]].push_back(queries[i][1]);
            ans[i] = bfs(i);
        }
        return ans;
    }
};
```

```go [sol-Go]
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    g := make([][]int, n-1)
    for i := range g {
        g[i] = append(g[i], i+1)
    }

    vis := make([]int, n-1)
    bfs := func(i int) int {
        q := []int{0}
        for step := 1; ; step++ {
            tmp := q
            q = nil
            for _, x := range tmp {
                for _, y := range g[x] {
                    if y == n-1 {
                        return step
                    }
                    if vis[y] != i {
                        vis[y] = i
                        q = append(q, y)
                    }
                }
            }
        }
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        g[q[0]] = append(g[q[0]], q[1])
        ans[i] = bfs(i + 1)
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q(n+q))$，其中 $q$ 是 $\textit{queries}$ 的长度。每次 BFS 的时间是 $\mathcal{O}(n+q)$。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 方法二：DP

定义 $f[i]$ 为从 $0$ 到 $i$ 的最短路。

用 $\textit{from}[i]$ 记录额外添加的边的终点是 $i$，起点列表是 $\textit{from}[i]$。

我们可以从 $i-1$ 到 $i$，也可以从 $\textit{from}[i][j]$ 到 $i$，这些位置作为转移来源，用其 $f$ 值 $+1$ 更新 $f[i]$ 的最小值。

初始值：$f[i] = i$。

答案：$f[n-1]$。

### 细节

设添加的边为 $l\to r$，只有当 $f[l]+1 < f[r]$ 时才更新 DP。

```py [sol-Python3]
class Solution:
    def shortestDistanceAfterQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        frm = [[] for _ in range(n)]
        f = list(range(n))
        ans = []
        for l, r in queries:
            frm[r].append(l)
            if f[l] + 1 < f[r]:
                f[r] = f[l] + 1
                for i in range(r + 1, n):
                    f[i] = min(f[i], f[i - 1] + 1, min((f[j] for j in frm[i]), default=inf) + 1)
            ans.append(f[-1])
        return ans
```

```java [sol-Java]
class Solution {
    public int[] shortestDistanceAfterQueries(int n, int[][] queries) {
        List<Integer>[] from = new ArrayList[n];
        Arrays.setAll(from, i -> new ArrayList<>());
        int[] f = new int[n];
        for (int i = 1; i < n; i++) {
            f[i] = i;
        }

        int[] ans = new int[queries.length];
        for (int qi = 0; qi < queries.length; qi++) {
            int l = queries[qi][0];
            int r = queries[qi][1];
            from[r].add(l);
            if (f[l] + 1 < f[r]) {
                f[r] = f[l] + 1;
                for (int i = r + 1; i < n; i++) {
                    f[i] = Math.min(f[i], f[i - 1] + 1);
                    for (int j : from[i]) {
                        f[i] = Math.min(f[i], f[j] + 1);
                    }
                }
            }
            ans[qi] = f[n - 1];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> shortestDistanceAfterQueries(int n, vector<vector<int>>& queries) {
        vector<vector<int>> from(n);
        vector<int> f(n);
        iota(f.begin(), f.end(), 0);

        vector<int> ans(queries.size());
        for (int qi = 0; qi < queries.size(); qi++) {
            int l = queries[qi][0], r = queries[qi][1];
            from[r].push_back(l);
            if (f[l] + 1 < f[r]) {
                f[r] = f[l] + 1;
                for (int i = r + 1; i < n; i++) {
                    f[i] = min(f[i], f[i - 1] + 1);
                    for (int j : from[i]) {
                        f[i] = min(f[i], f[j] + 1);
                    }
                }
            }
            ans[qi] = f[n - 1];
        }
        return ans;
    }
};
```

```go [sol-Go]
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    from := make([][]int, n)
    f := make([]int, n)
    for i := 1; i < n; i++ {
        f[i] = i
    }

    ans := make([]int, len(queries))
    for qi, q := range queries {
        l, r := q[0], q[1]
        from[r] = append(from[r], l)
        if f[l]+1 < f[r] {
            f[r] = f[l] + 1
            for i := r + 1; i < n; i++ {
                f[i] = min(f[i], f[i-1]+1)
                for _, j := range from[i] {
                    f[i] = min(f[i], f[j]+1)
                }
            }
        }
        ans[qi] = f[n-1]
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q(n+q))$，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
