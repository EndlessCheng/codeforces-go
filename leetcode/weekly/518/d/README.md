## 方法一：Dijkstra 算法

[Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)

在网格图最短路的基础上，额外增加两个参数 $k$ 和 $\textit{idx}$，定义 $\textit{dis}[k][i][j][\textit{idx}]$ 表示从起点 $(0,0)$ 到 $(i,j)$ 的最小路径代价，恰好发生 $k$ 次转向，且最后一步的方向为 $\textit{idx}$。这里 $\textit{idx}$ 是一个 $[0,3]$ 中的整数（分别对应左右上下四个方向）。

枚举左右上下四个方向：

- 如果前进方向与 $\textit{idx}$ 相同，那么 $k$ 不变。
- 否则，必须满足 $k>0$，然后把 $k$ 减少一。

```py [sol-Python3]
class Solution:
    def minCost(self, grid: list[list[int]], k0: int) -> int:
        dirs = ((0, -1), (0, 1), (-1, 0), (1, 0))  # 左右上下
        m, n = len(grid), len(grid[0])
        dis = [[[[inf] * 4 for _ in range(n)] for _ in range(m)] for _ in range(k0 + 1)]

        # 初始方向可以向右（1）或向下（3）
        h = [(grid[0][0], k0, 0, 0, 1), (grid[0][0], k0, 0, 0, 3)]
        dis[k0][0][0][1] = dis[k0][0][0][3] = grid[0][0]

        while h:
            d, k, i, j, idx = heappop(h)
            if i == m - 1 and j == n - 1:
                return d
            if d > dis[k][i][j][idx]:
                continue
            for new_idx, (dx, dy) in enumerate(dirs):
                x, y = i + dx, j + dy
                if 0 <= x < m and 0 <= y < n:
                    new_k = k
                    if new_idx != idx:
                        if k == 0:
                            continue
                        new_k -= 1
                    new_d = d + grid[x][y]
                    if new_d < dis[new_k][x][y][new_idx]:
                        dis[new_k][x][y][new_idx] = new_d
                        heappush(h, (new_d, new_k, x, y, new_idx))
        return -1
```

```java [sol-Java]
class Solution {
    private final static int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

    public int minCost(int[][] grid, int k0) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][][] dis = new int[k0 + 1][m][n][4];
        for (int[][][] a : dis) {
            for (int[][] b : a) {
                for (int[] c : b) {
                    Arrays.fill(c, Integer.MAX_VALUE / 2);
                }
            }
        }
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);

        // 初始方向可以向右（1）或向下（3）
        pq.offer(new int[]{grid[0][0], k0, 0, 0, 1});
        pq.offer(new int[]{grid[0][0], k0, 0, 0, 3});
        dis[k0][0][0][1] = dis[k0][0][0][3] = grid[0][0];

        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int d = top[0];
            int k = top[1];
            int i = top[2];
            int j = top[3];
            int idx = top[4];
            if (i == m - 1 && j == n - 1) {
                return d;
            }
            if (d > dis[k][i][j][idx]) {
                continue;
            }
            for (int newIdx = 0; newIdx < 4; newIdx++) {
                int x = i + DIRS[newIdx][0];
                int y = j + DIRS[newIdx][1];
                if (0 <= x && x < m && 0 <= y && y < n) {
                    int newK = k;
                    if (newIdx != idx) {
                        if (k == 0) {
                            continue;
                        }
                        newK--;
                    }
                    int newD = d + grid[x][y];
                    if (newD < dis[newK][x][y][newIdx]) {
                        dis[newK][x][y][newIdx] = newD;
                        pq.offer(new int[]{newD, newK, x, y, newIdx});
                    }
                }
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

public:
    int minCost(vector<vector<int>>& grid, int k0) {
        const int INF = INT_MAX / 2;
        int m = grid.size(), n = grid[0].size();
        vector dis(k0 + 1, vector(m, vector<array<int, 4>>(n, {INF, INF, INF, INF})));
        priority_queue<tuple<int, int, int, int, int>, vector<tuple<int, int, int, int, int>>, greater<>> pq;
        // 初始方向可以向右（1）或向下（3）
        pq.emplace(grid[0][0], k0, 0, 0, 1);
        pq.emplace(grid[0][0], k0, 0, 0, 3);
        dis[k0][0][0][1] = dis[k0][0][0][3] = grid[0][0];

        while (!pq.empty()) {
            auto [d, k, i, j, idx] = pq.top();
            pq.pop();
            if (i == m - 1 && j == n - 1) {
                return d;
            }
            if (d > dis[k][i][j][idx]) {
                continue;
            }
            for (int new_idx = 0; new_idx < 4; new_idx++) {
                int x = i + DIRS[new_idx][0];
                int y = j + DIRS[new_idx][1];
                if (0 <= x && x < m && 0 <= y && y < n) {
                    int new_k = k;
                    if (new_idx != idx) {
                        if (k == 0) {
                            continue;
                        }
                        new_k--;
                    }
                    int new_d = d + grid[x][y];
                    if (new_d < dis[new_k][x][y][new_idx]) {
                        dis[new_k][x][y][new_idx] = new_d;
                        pq.emplace(new_d, new_k, x, y, new_idx);
                    }
                }
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func minCost(grid [][]int, k0 int) int {
	m, n := len(grid), len(grid[0])
	dis := make([][][][4]int, k0+1)
	for k := range dis {
		dis[k] = make([][][4]int, m)
		for i := range dis[k] {
			dis[k][i] = make([][4]int, n)
			for j := range dis[k][i] {
				for idx := range dis[k][i][j] {
					dis[k][i][j][idx] = math.MaxInt / 2
				}
			}
		}
	}

	// 初始方向可以向右（1）或向下（3）
	v := grid[0][0]
	h := hp{{v, k0, 0, 0, 1}, {v, k0, 0, 0, 3}}
	dis[k0][0][0][1] = v
	dis[k0][0][0][3] = v

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d, k, i, j, idx := top.dis, top.k, top.i, top.j, top.idx
		if i == m-1 && j == n-1 {
			return d
		}
		if d > dis[k][i][j][idx] {
			continue
		}
		for newIdx, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < m && 0 <= y && y < n {
				newK := k
				if newIdx != idx {
					if k == 0 {
						continue
					}
					newK--
				}
				newD := d + grid[x][y]
				if newD < dis[newK][x][y][newIdx] {
					dis[newK][x][y][newIdx] = newD
					heap.Push(&h, tuple{newD, newK, x, y, newIdx})
				}
			}
		}
	}
	return -1
}

type tuple struct{ dis, k, i, j, idx int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kmn\log(kmn))$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(kmn)$。

## 方法二：动态规划

走回头路要转向，会消耗 $k$，所以我们不会回到完全一样的状态（回到相同格子会导致 $k$ 变小），所以添加参数 $k$ 后的分层图是一个**有向无环图**（DAG）。

所以可以直接跑 DP，无需 Dijkstra。

### 记忆化搜索

和 Dijkstra 算法的 $\textit{dis}$ 数组一样，定义 $\textit{dfs}(k,i,j,\textit{idx})$ 表示从起点 $(0,0)$ 到 $(i,j)$ 的最小路径代价，恰好发生 $k$ 次转向，且最后一步的方向为 $\textit{idx}$。我的写法是从终点 $(m-1,n-1)$ 倒着走的，这里的 $\textit{idx}$ 是倒着走的方向。

枚举左右上下四个方向：

- 如果前进方向与 $\textit{idx}$ 相同，那么 $k$ 不变。
- 否则，必须满足 $k>0$，然后把 $k$ 减少一。

递归边界：$\textit{dfs}(k,0,0,\textit{idx}) = \textit{grid}[0][0]$。

递归入口：$\min(\textit{dfs}(k_0, m - 1, n - 1, \textit{idx}_{左}), \textit{dfs}(k_0, m - 1, n - 1, \textit{idx}_{上}))$。

```py [sol-Python3]
class Solution:
    def minCost(self, grid: list[list[int]], k0: int) -> int:
        dirs = ((0, -1), (0, 1), (-1, 0), (1, 0))  # 左右上下
        m, n = len(grid), len(grid[0])

        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(k: int, i: int, j: int, idx: int) -> int:
            if i == j == 0:
                return grid[0][0]

            res = inf
            for new_idx, (dx, dy) in enumerate(dirs):
                x, y = i + dx, j + dy
                if 0 <= x < m and 0 <= y < n:
                    new_k = k
                    if new_idx != idx:
                        if k == 0:
                            continue
                        new_k -= 1
                    res = min(res, dfs(new_k, x, y, new_idx))
            return res + grid[i][j]

        ans = min(dfs(k0, m - 1, n - 1, 0), dfs(k0, m - 1, n - 1, 2))
        return -1 if ans == inf else ans
```

```java [sol-Java]
class Solution {
    private final static int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

    public int minCost(int[][] grid, int k0) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][][] memo = new int[k0 + 1][m][n][4];
        for (int[][][] a : memo) {
            for (int[][] b : a) {
                for (int[] c : b) {
                    Arrays.fill(c, -1); // -1 表示该状态没有计算过
                }
            }
        }

        int ans = Math.min(dfs(k0, m - 1, n - 1, 0, grid, memo), dfs(k0, m - 1, n - 1, 2, grid, memo));
        return ans < Integer.MAX_VALUE / 2 ? ans : -1;
    }

    private int dfs(int k, int i, int j, int idx, int[][] grid, int[][][][] memo) {
        if (i == 0 && j == 0) {
            return grid[0][0];
        }

        int[] p = memo[k][i][j];
        if (p[idx] != -1) { // 之前计算过
            return p[idx];
        }

        int m = grid.length;
        int n = grid[0].length;
        int res = Integer.MAX_VALUE / 2;
        for (int newIdx = 0; newIdx < 4; newIdx++) {
            int x = i + DIRS[newIdx][0];
            int y = j + DIRS[newIdx][1];
            if (0 <= x && x < m && 0 <= y && y < n) {
                int newK = k;
                if (newIdx != idx) {
                    if (k == 0) {
                        continue;
                    }
                    newK--;
                }
                res = Math.min(res, dfs(newK, x, y, newIdx, grid, memo));
            }
        }
        res += grid[i][j];

        p[idx] = res; // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}; // 左右上下

public:
    int minCost(vector<vector<int>>& grid, int k0) {
        int m = grid.size(), n = grid[0].size();
        vector memo(k0 + 1, vector(m, vector<array<int, 4>>(n, {-1, -1, -1, -1})));

        auto dfs = [&](this auto&& dfs, int k, int i, int j, int idx) -> int {
            if (i == 0 && j == 0) {
                return grid[0][0];
            }

            // 注意这里是引用
            int& res = memo[k][i][j][idx];
            if (res != -1) { // 之前计算过
                return res;
            }

            res = INT_MAX / 2;
            for (int new_idx = 0; new_idx < 4; new_idx++) {
                int x = i + DIRS[new_idx][0];
                int y = j + DIRS[new_idx][1];
                if (0 <= x && x < m && 0 <= y && y < n) {
                    int new_k = k;
                    if (new_idx != idx) {
                        if (k == 0) {
                            continue;
                        }
                        new_k--;
                    }
                    res = min(res, dfs(new_k, x, y, new_idx));
                }
            }
            res += grid[i][j];
            return res;
        };

        int ans = min(dfs(k0, m - 1, n - 1, 0), dfs(k0, m - 1, n - 1, 2));
        return ans < INT_MAX / 2 ? ans : -1;
    }
};
```

```go [sol-Go]
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func minCost(grid [][]int, k0 int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][][][4]int, k0+1)
	for i := range memo {
		memo[i] = make([][][4]int, m)
		for j := range memo[i] {
			memo[i][j] = make([][4]int, n)
			for p := range memo[i][j] {
				for q := range memo[i][j][p] {
					memo[i][j][p][q] = -1 // -1 表示该状态没有计算过
				}
			}
		}
	}

	var dfs func(int, int, int, int) int
	dfs = func(k, i, j, idx int) int {
		if i == 0 && j == 0 {
			return grid[0][0]
		}

		p := &memo[k][i][j][idx]
		if *p != -1 { // 之前计算过
			return *p
		}

		res := math.MaxInt / 2
		for newIdx, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < m && 0 <= y && y < n {
				newK := k
				if newIdx != idx {
					if k == 0 {
						continue
					}
					newK--
				}
				res = min(res, dfs(newK, x, y, newIdx))
			}
		}
		res += grid[i][j]

		*p = res // 记忆化
		return res
	}

	ans := min(dfs(k0, m-1, n-1, 0), dfs(k0, m-1, n-1, 2))
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kmn)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(kmn)$。

### 递推（空间优化）

分层图有 $k_0+1$ 层，每一层有 $m\times n$ 个格子，每个格子有 $4$ 个状态。

把移动过程细分为直走和原地转向：

- 直走：在同一层的网格图上移动到相邻格子中的同一方向。
- 原地转向：从 $k-1$ 层的格子 $(i,j)$ 中的 $4$ 个状态（取最小值）转移到 $k$ 层的格子 $(i,j)$ 中的 $4$ 个状态。

具体地，定义 $f[k][\textit{idx}][i][j]$ 表示从起点 $(0,0)$ 到 $(i,j)$ 的最小路径代价，恰好发生 $k$ 次转向，且最后一步的方向为 $\textit{idx}$。

由于我们可以上下左右走，递推写法需要保证转移来源是已经算出的值。例如从右到左直走，从 $(i,j+1)$ 移动到 $(i,j)$，必须先保证 $(i,j+1)$ 的状态已经算出来了。

一般地，计算转移要考虑从上到下、从下到上、从左到右、从右到左这四种移动方向。例如从右到左，有两种情况：

- 直走：从 $(i,j+1)$ 到 $(i,j)$，即 $f[k][\textit{idx}_{左}][i][j] = f[k][\textit{idx}_{左}][i][j+1] + \textit{grid}[i][j]$。
- 原地转向：按照上文的描述，$f[k][\textit{idx}_{左}][i][j] = \min\limits_{\textit{idx}=0}^{3} f[k-1][\textit{idx}][i][j]$。

二者取最小值。

初始值：$f[0][\textit{idx}_{下}][0][0] = f[0][\textit{idx}_{右}][0][0] = \textit{grid}[0][0]$。

答案：$\min\limits_{k=0}^{k_0} \min(f[k][\textit{idx}_{下}][m-1][n-1], f[k][\textit{idx}_{右}][m-1][n-1])$。

代码实现时，$f$ 的第一个维度可以优化掉。此时要先计算所有原地转向的状态转移，再计算所有直走的状态转移。

```py [sol-Python3]
class Solution:
    def minCost(self, grid: list[list[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])
        d = [[inf] * n for _ in range(m)]
        u = [[inf] * n for _ in range(m)]
        r = [[inf] * n for _ in range(m)]
        l = [[inf] * n for _ in range(m)]
        d[0][0] = r[0][0] = grid[0][0]
        ans = inf

        # 第一轮循环，把状态从 (0,0) 更新到第一行和第一列
        # 第二轮循环，把状态从第一行向下更新到其余行，从第一列向右更新到其余列，这样就算出了转向 1 次的所有情况
        # ……
        for _ in range(k + 1):
            # 原地转向
            for di, ui, ri, li in zip(d, u, r, l):
                for j, t in enumerate(zip(di, ui, ri, li)):
                    di[j] = ui[j] = ri[j] = li[j] = min(t)

            # 向下
            for i in range(1, m):
                for j, x in enumerate(grid[i]):
                    d[i][j] = min(d[i][j], d[i - 1][j] + x)

            # 向上
            for i in range(m - 2, -1, -1):
                for j, x in enumerate(grid[i]):
                    u[i][j] = min(u[i][j], u[i + 1][j] + x)

            # 向右
            for row, ri in zip(grid, r):
                for j in range(1, n):
                    ri[j] = min(ri[j], ri[j - 1] + row[j])

            # 向左
            for row, li in zip(grid, l):
                for j in range(n - 2, -1, -1):
                    li[j] = min(li[j], li[j + 1] + row[j])

            ans = min(ans, d[-1][-1], r[-1][-1])

        return -1 if ans == inf else ans
```

```java [sol-Java]
class Solution {
    public int minCost(int[][] grid, int k0) {
        final int INF = Integer.MAX_VALUE / 2;
        int m = grid.length;
        int n = grid[0].length;
        int[][] d = new int[m][n];
        int[][] u = new int[m][n];
        int[][] r = new int[m][n];
        int[][] l = new int[m][n];

        for (int i = 0; i < m; i++) {
            Arrays.fill(d[i], INF);
            Arrays.fill(u[i], INF);
            Arrays.fill(r[i], INF);
            Arrays.fill(l[i], INF);
        }
        d[0][0] = grid[0][0];
        r[0][0] = grid[0][0];
        int ans = INF;

        // 第一轮循环，把状态从 (0,0) 更新到第一行和第一列
        // 第二轮循环，把状态从第一行向下更新到其余行，从第一列向右更新到其余列，这样就算出了转向 1 次的所有情况
        // ……
        for (int k = 0; k <= k0; k++) {
            // 原地转向
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    int mn = Math.min(Math.min(Math.min(d[i][j], u[i][j]), r[i][j]), l[i][j]);
                    d[i][j] = u[i][j] = r[i][j] = l[i][j] = mn;
                }
            }

            // 向下
            for (int i = 1; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    d[i][j] = Math.min(d[i][j], d[i - 1][j] + grid[i][j]);
                }
            }

            // 向上
            for (int i = m - 2; i >= 0; i--) {
                for (int j = 0; j < n; j++) {
                    u[i][j] = Math.min(u[i][j], u[i + 1][j] + grid[i][j]);
                }
            }

            // 向右
            for (int i = 0; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    r[i][j] = Math.min(r[i][j], r[i][j - 1] + grid[i][j]);
                }
            }

            // 向左
            for (int i = 0; i < m; i++) {
                for (int j = n - 2; j >= 0; j--) {
                    l[i][j] = Math.min(l[i][j], l[i][j + 1] + grid[i][j]);
                }
            }

            ans = Math.min(ans, Math.min(d[m - 1][n - 1], r[m - 1][n - 1]));
        }

        return ans < INF ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<vector<int>>& grid, int k0) {
        constexpr int INF = INT_MAX / 2;
        int m = grid.size(), n = grid[0].size();
        vector d(m, vector<int>(n, INF));
        vector u(m, vector<int>(n, INF));
        vector r(m, vector<int>(n, INF));
        vector l(m, vector<int>(n, INF));
        d[0][0] = r[0][0] = grid[0][0];
        int ans = INF;

        // 第一轮循环，把状态从 (0,0) 更新到第一行和第一列
        // 第二轮循环，把状态从第一行向下更新到其余行，从第一列向右更新到其余列，这样就算出了转向 1 次的所有情况
        // ……
        for (int k = 0; k <= k0; k++) {
            // 原地转向
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    int mn = min(min(min(d[i][j], u[i][j]), r[i][j]), l[i][j]); // 这样写比 min({...}) 快
                    d[i][j] = u[i][j] = r[i][j] = l[i][j] = mn;
                }
            }

            // 向下
            for (int i = 1; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    d[i][j] = min(d[i][j], d[i - 1][j] + grid[i][j]);
                }
            }

            // 向上
            for (int i = m - 2; i >= 0; i--) {
                for (int j = 0; j < n; j++) {
                    u[i][j] = min(u[i][j], u[i + 1][j] + grid[i][j]);
                }
            }

            // 向右
            for (int i = 0; i < m; i++) {
                for (int j = 1; j < n; j++) {
                    r[i][j] = min(r[i][j], r[i][j - 1] + grid[i][j]);
                }
            }

            // 向左
            for (int i = 0; i < m; i++) {
                for (int j = n - 2; j >= 0; j--) {
                    l[i][j] = min(l[i][j], l[i][j + 1] + grid[i][j]);
                }
            }

            ans = min(ans, min(d[m - 1][n - 1], r[m - 1][n - 1]));
        }

        return ans < INF ? ans : -1;
    }
};
```

```go [sol-Go]
func minCost(grid [][]int, k int) int {
	const inf = math.MaxInt / 2
	m, n := len(grid), len(grid[0])
	d := make([][]int, m)
	u := make([][]int, m)
	r := make([][]int, m)
	l := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
		u[i] = make([]int, n)
		r[i] = make([]int, n)
		l[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = inf
			u[i][j] = inf
			r[i][j] = inf
			l[i][j] = inf
		}
	}
	d[0][0] = grid[0][0]
	r[0][0] = grid[0][0]
	ans := inf

	// 第一轮循环，把状态从 (0,0) 更新到第一行和第一列
	// 第二轮循环，把状态从第一行向下更新到其余行，从第一列向右更新到其余列，这样就算出了转向 1 次的所有情况
	// ……
	for range k + 1 {
		// 原地转向
		for i := range m {
			for j := range n {
				mn := min(d[i][j], u[i][j], r[i][j], l[i][j])
				d[i][j] = mn
				u[i][j] = mn
				r[i][j] = mn
				l[i][j] = mn
			}
		}

		// 向下
		for i := 1; i < m; i++ {
			for j, x := range grid[i] {
				d[i][j] = min(d[i][j], d[i-1][j]+x)
			}
		}

		// 向上
		for i := m - 2; i >= 0; i-- {
			for j, x := range grid[i] {
				u[i][j] = min(u[i][j], u[i+1][j]+x)
			}
		}

		// 向右
		for i, row := range grid {
			for j := 1; j < n; j++ {
				r[i][j] = min(r[i][j], r[i][j-1]+row[j])
			}
		}

		// 向左
		for i, row := range grid {
			for j := n - 2; j >= 0; j-- {
				l[i][j] = min(l[i][j], l[i][j+1]+row[j])
			}
		}

		ans = min(ans, d[m-1][n-1], r[m-1][n-1])
	}

	if ans < inf {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kmn)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 专题训练

1. 图论题单的「**§3.1 单源最短路：Dijkstra 算法**」。
2. 动态规划题单的「**二、网格图 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
