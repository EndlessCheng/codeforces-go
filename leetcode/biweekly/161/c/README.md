## 转化

「最大化最小值」是二分答案的代名词，为什么？

设有效路径上的边权都大于等于下界 $\textit{lower}$。

- 如果下界等于 $\textit{lower}$ 时，存在有效路径，那么当下界小于 $\textit{lower}$ 时，约束**更加宽松**，更加存在有效路径。
- 如果下界等于 $\textit{lower}$ 时，不存在有效路径，那么当下界大于 $\textit{lower}$ 时，约束**更加苛刻**，更不存在有效路径。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定下界 $\textit{lower}$，能否找到一条有效路径，除了满足题目的两个要求，还满足路径上的边权都大于等于下界 $\textit{lower}$。

如果存在有效路径，说明答案 $\ge \textit{lower}$，否则答案 $< \textit{lower}$。

## 思路

由于题目保证图是一个 DAG（有向无环图），计算 DP 无后效性，可以用 DAG DP 解决。

定义 $\textit{dfs}(x)$ 表示从 $x$ 到 $n-1$ 的有效路径的总恢复成本的最小值，即 $x$ 到 $n-1$ 的最短路长度。

枚举 $x$ 的邻居 $y$，如果 $y$ 在线且边权 $\textit{wt}  \ge \textit{lower}$，那么问题变成从 $y$ 到 $n-1$ 的有效路径的总恢复成本的最小值，即 $\textit{dfs}(y)$，加上边权 $\textit{wt}$，更新 $\textit{dfs}(x)$ 的最小值，即

$$
\textit{dfs}(x) = \min_y \textit{dfs}(y) + \textit{wt}
$$

递归边界：$\textit{dfs}(n-1)=0$。

递归入口：$\textit{dfs}(0)$。

> **注**：也可以用 Dijkstra 算法解决，但那样做时间复杂度要多乘以一个 $\log m$，可能会超时。

## 细节

### 1)

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$-1$。人为规定 $-1$ 一定满足要求。如果二分结果为 $-1$，那么返回 $-1$。
- 开区间右端点初始值：边权的最大值加一。一定不满足要求。注意本题 $n\ge 2$，至少要走一条边。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

### 2)

建图的时候，如果两端点都在线才连边。这样可以减少图的边数，也无需在 DP 过程中判断节点是否在线。

## 答疑

**问**：如果答案不等于 $-1$，为什么二分结束后，答案 $\textit{ans}$ 一定等于某条边的边权？

**答**：反证法。假设 $\textit{ans}$ 不等于某条边的边权，这意味着下界加一后，仍然存在有效路径。换句话说，$\text{check}(\textit{ans}+1)=\texttt{true}$。但根据循环不变量，二分结束后 $\text{check}(\textit{ans}+1)=\texttt{false}$，矛盾。故原命题成立。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1R5g8zDEGY/?t=9m32s)，欢迎点赞关注~

## 写法一：记忆化搜索

```py [sol-Python3]
class Solution:
    def findMaxPathScore(self, edges: List[List[int]], online: List[bool], k: int) -> int:
        n = len(online)
        g = [[] for _ in range(n)]
        max_wt = 0
        for x, y, wt in edges:
            if online[x] and online[y]:
                g[x].append((y, wt))
                max_wt = max(max_wt, wt)

        def check(lower: int) -> bool:
            @cache
            def dfs(x: int) -> int:
                if x == n - 1:  # 到达终点
                    return 0
                res = inf
                for y, wt in g[x]:
                    if wt >= lower:
                        res = min(res, dfs(y) + wt)
                return res
            return dfs(0) <= k

        left, right = -1, max_wt + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```py [sol-Python3 库函数]
class Solution:
    def findMaxPathScore(self, edges: List[List[int]], online: List[bool], k: int) -> int:
        n = len(online)
        g = [[] for _ in range(n)]
        max_wt = 0
        for x, y, wt in edges:
            if online[x] and online[y]:
                g[x].append((y, wt))
                max_wt = max(max_wt, wt)

        def check(lower: int) -> bool:
            @cache
            def dfs(x: int) -> int:
                if x == n - 1:  # 到达终点
                    return 0
                res = inf
                for y, wt in g[x]:
                    if wt >= lower:
                        res = min(res, dfs(y) + wt)
                return res
            return dfs(0) > k  # 取反

        # 二分无法到达 n-1 的最小 lower，那么减一后，就是可以到达 n-1 的最大 lower
        return bisect_left(range(max_wt + 1), True, key=check) - 1
```

```java [sol-Java]
class Solution {
    public int findMaxPathScore(int[][] edges, boolean[] online, long k) {
        int n = online.length;
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        int maxWt = 0;
        for (int[] e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            if (online[x] && online[y]) {
                g[x].add(new int[]{y, wt});
                maxWt = Math.max(maxWt, wt);
            }
        }

        long[] memo = new long[n];
        int left = -1, right = maxWt + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            Arrays.fill(memo, -1L); // -1 表示没有计算过
            if (dfs(0, mid, g, memo) <= k) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private long dfs(int x, int lower, List<int[]>[] g, long[] memo) {
        if (x == g.length - 1) { // 到达终点
            return 0;
        }
        if (memo[x] != -1) { // 之前计算过
            return memo[x];
        }
        long res = Long.MAX_VALUE / 2; // 防止加法溢出
        for (int[] e : g[x]) {
            int y = e[0], wt = e[1];
            if (wt >= lower) {
                res = Math.min(res, dfs(y, lower, g, memo) + wt);
            }
        }
        return memo[x] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMaxPathScore(vector<vector<int>>& edges, vector<bool>& online, long long k) {
        int n = online.size();
        vector<vector<pair<int, int>>> g(n);
        int max_wt = 0;
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            if (online[x] && online[y]) {
                g[x].emplace_back(y, wt);
                max_wt = max(max_wt, wt);
            }
        }

        vector<long long> memo(n);
        auto check = [&](int lower) -> bool {
            ranges::fill(memo, -1); // -1 表示没有计算过

            auto dfs = [&](this auto&& dfs, int x) -> long long {
                if (x == n - 1) { // 到达终点
                    return 0;
                }
                auto& res = memo[x]; // 注意这里是引用
                if (res != -1) { // 之前计算过
                    return memo[x];
                }
                res = LLONG_MAX / 2; // 防止加法溢出
                for (auto& [y, wt] : g[x]) {
                    if (wt >= lower) {
                        res = min(res, dfs(y) + wt);
                    }
                }
                return res;
            };

            return dfs(0) <= k;
        };

        int left = -1, right = max_wt + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	maxWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if online[x] && online[y] {
			g[x] = append(g[x], edge{y, wt})
			maxWt = max(maxWt, wt)
		}
	}

	memo := make([]int, n)
	// 二分无法到达 n-1 的最小 lower，那么减一后，就是可以到达 n-1 的最大 lower
	ans := sort.Search(maxWt+1, func(lower int) bool {
		for i := range memo {
			memo[i] = -1 // -1 表示没有计算过
		}
		var dfs func(int) int
		dfs = func(x int) int {
			if x == n-1 { // 到达终点
				return 0
			}
			p := &memo[x]
			if *p != -1 { // 之前计算过
				return *p
			}
			res := math.MaxInt / 2 // 防止加法溢出
			for _, e := range g[x] {
				y := e.to
				if e.wt >= lower {
					res = min(res, dfs(y)+e.wt)
				}
			}
			*p = res // 记忆化
			return res
		}
		return dfs(0) > int(k)
	}) - 1
	return ans
}
```

## 写法二：拓扑排序

拓扑排序相当于记忆化搜索的 1:1 翻译版本。严格地翻译需要把每条边反向，再跑拓扑排序。

但也可以从起点开始正着计算，即刷表法。需要注意的是，为了能让我们在拓扑排序中，把入度为 $0$ 的点入队，在拓扑排序之前，先清理那些无法到达的边。比如在有 $0\to 1$ 的情况下，需要去掉 $2\to 1$ 这种无法到达的边。

```py [sol-Python3]
class Solution:
    def findMaxPathScore(self, edges: List[List[int]], online: List[bool], k: int) -> int:
        n = len(online)
        g = [[] for _ in range(n)]
        deg = [0] * n
        max_wt = 0
        for x, y, wt in edges:
            if online[x] and online[y]:
                g[x].append((y, wt))
                deg[y] += 1
                max_wt = max(max_wt, wt)

        # 先清理无法从 0 到达的边
        q = deque(i for i in range(1, n) if deg[i] == 0)
        while q:
            x = q.popleft()
            for y, _ in g[x]:
                deg[y] -= 1
                if y and deg[y] == 0:
                    q.append(y)

        def check(lower: int) -> bool:
            deg_copy = deg.copy()
            f = [inf] * n
            f[0] = 0

            q = deque([0])
            while q:
                x = q.popleft()
                if x == n - 1:
                    return f[x] > k
                for y, wt in g[x]:
                    if wt >= lower:
                        f[y] = min(f[y], f[x] + wt)
                    deg_copy[y] -= 1
                    if deg_copy[y] == 0:
                        q.append(y)
            return True

        # 二分无法到达 n-1 的最小 lower，那么减一后，就是可以到达 n-1 的最大 lower
        return bisect_left(range(max_wt + 1), True, key=check) - 1
```

```java [sol-Java]
class Solution {
    public int findMaxPathScore(int[][] edges, boolean[] online, long k) {
        int n = online.length;
        List<int[]>[] g = new ArrayList[n];
        Arrays.setAll(g, _ -> new ArrayList<>());
        int[] deg = new int[n];
        int maxWt = 0;
        for (int[] e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            if (online[x] && online[y]) {
                g[x].add(new int[]{y, wt});
                deg[y]++;
                maxWt = Math.max(maxWt, wt);
            }
        }

        // 先清理无法从 0 到达的边
        Queue<Integer> q = new ArrayDeque<>();
        for (int i = 1; i < n; i++) {
            if (deg[i] == 0) {
                q.offer(i);
            }
        }
        while (!q.isEmpty()) {
            int x = q.poll();
            for (int[] e : g[x]) {
                int y = e[0];
                if (--deg[y] == 0 && y > 0) {
                    q.offer(y);
                }
            }
        }

        long f[] = new long[n];
        int left = -1, right = maxWt + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(mid, g, deg.clone(), f, k)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(int lower, List<int[]>[] g, int[] deg, long f[], long k) {
        Arrays.fill(f, Long.MAX_VALUE / 2); // 除 2 防止加法溢出
        f[0] = 0;
        Queue<Integer> q = new ArrayDeque<>();
        q.offer(0);
        while (!q.isEmpty()) {
            int x = q.poll();
            if (x == g.length - 1) {
                return f[x] <= k;
            }
            for (int[] e : g[x]) {
                int y = e[0], wt = e[1];
                if (wt >= lower) {
                    f[y] = Math.min(f[y], f[x] + wt);
                }
                if (--deg[y] == 0) {
                    q.offer(y);
                }
            }
        }
        return false; // 无法到达 n-1
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMaxPathScore(vector<vector<int>>& edges, vector<bool>& online, long long k) {
        int n = online.size();
        vector<vector<pair<int, int>>> g(n);
        vector<int> deg(n);
        int max_wt = 0;
        for (auto& e : edges) {
            int x = e[0], y = e[1], wt = e[2];
            if (online[x] && online[y]) {
                g[x].emplace_back(y, wt);
                deg[y]++;
                max_wt = max(max_wt, wt);
            }
        }

        // 先清理无法从 0 到达的边
        queue<int> q;
        for (int i = 1; i < n; i++) {
            if (deg[i] == 0) {
                q.push(i);
            }
        }
        while (!q.empty()) {
            int x = q.front();
            q.pop();
            for (auto& [y, _] : g[x]) {
                if (--deg[y] == 0 && y) {
                    q.push(y);
                }
            }
        }

        vector<long long> f(n);
        auto check = [&](int lower) -> bool {
            auto deg_copy = deg;
            ranges::fill(f, LLONG_MAX / 2); // 除 2 防止加法溢出
            f[0] = 0;

            queue<int> q;
            q.push(0);
            while (!q.empty()) {
                int x = q.front();
                q.pop();
                if (x == n - 1) {
                    return f[x] <= k;
                }
                for (auto& [y, wt] : g[x]) {
                    if (wt >= lower) {
                        f[y] = min(f[y], f[x] + wt);
                    }
                    if (--deg_copy[y] == 0) {
                        q.push(y);
                    }
                }
            }
            return false; // 无法到达 n-1
        };

        int left = -1, right = max_wt + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	deg := make([]int, n)
	maxWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if online[x] && online[y] {
			g[x] = append(g[x], edge{y, wt})
			deg[y]++
			maxWt = max(maxWt, wt)
		}
	}

	// 先清理无法从 0 到达的边
	q := []int{}
	for i := 1; i < n; i++ {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, e := range g[x] {
			y := e.to
			deg[y]--
			if deg[y] == 0 && y > 0 {
				q = append(q, y)
			}
		}
	}

	f := make([]int, n)
	ans := sort.Search(maxWt+1, func(lower int) bool {
		deg := slices.Clone(deg)
		for i := 1; i < n; i++ {
			f[i] = math.MaxInt / 2 // 除 2 防止加法溢出
		}

		q := []int{0}
		for len(q) > 0 {
			x := q[0]
			if x == n-1 {
				return f[x] > int(k)
			}
			q = q[1:]
			for _, e := range g[x] {
				y := e.to
				wt := e.wt
				if wt >= lower {
					f[y] = min(f[y], f[x]+wt)
				}
				deg[y]--
				if deg[y] == 0 {
					q = append(q, y)
				}
			}
		}
		return true
	}) - 1
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n + m)\log U)$，其中 $n$ 是 $\textit{online}$ 的长度，$m$ 是 $\textit{edges}$ 的长度，$U$ 是边权的最大值。
- 空间复杂度：$\mathcal{O}(n + m)$。

## 专题训练

1. 二分题单的「**§2.5 最大化最小值**」。
2. 动态规划题单的「**十三、图 DP**」中标有 DAG 的题目。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
