[视频讲解](https://www.bilibili.com/video/BV1Yr4y1o7aP/)第三题。

建议结合视频中画的图来理解。

1. 从所有 $1$ 出发，写一个多源 BFS，计算出每个格子 $(i,j)$ 到最近的 $1$ 的曼哈顿距离 $\textit{dis}[i][j]$。注意题目保证至少有一个 $1$。
2. 答案不会超过 $\textit{dis}[i][j]$ 的最大值，我们可以倒序枚举答案。
3. 如何判断我们能从左上角 $(0,0)$ 走到右下角 $(n-1,n-1)$ 呢？并查集！
4. 假设答案为 $d$，我们可以把所有 $\textit{dis}[i][j]=d$ 的格子与其四周 $\ge d$ 的格子用并查集连起来，在答案为 $d$ 的情况下，这些格子之间是可以互相到达的。
5. 用并查集判断 $(0,0)$ 和 $(n-1,n-1)$ 是否连通，只要连通就立刻返回 $d$ 作为答案。

```py [sol-Python3]
class Solution:
    def maximumSafenessFactor(self, grid: List[List[int]]) -> int:
        n = len(grid)
        q = []
        dis = [[-1] * n for _ in range(n)]
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x:
                    q.append((i, j))
                    dis[i][j] = 0

        groups = [q]
        while q:  # 多源 BFS
            tmp = q
            q = []
            for i, j in tmp:
                for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                    if 0 <= x < n and 0 <= y < n and dis[x][y] < 0:
                        q.append((x, y))
                        dis[x][y] = len(groups)
            groups.append(q)  # 相同 dis 分组记录

        # 并查集模板
        fa = list(range(n * n))
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]

        for d in range(len(groups) - 2, 0, -1):
            for i, j in groups[d]:
                for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                    if 0 <= x < n and 0 <= y < n and dis[x][y] >= dis[i][j]:
                        fa[find(x * n + y)] = find(i * n + j)
            if find(0) == find(n * n - 1):  # 写这里判断更快些
                return d
        return 0
```

```java [sol-Java]
class Solution {
    private final static int[][] DIRS = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int maximumSafenessFactor(List<List<Integer>> grid) {
        int n = grid.size();
        var q = new ArrayList<int[]>();
        var dis = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (grid.get(i).get(j) > 0) {
                    q.add(new int[]{i, j});
                } else {
                    dis[i][j] = -1;
                }
            }
        }

        var groups = new ArrayList<List<int[]>>();
        groups.add(q);
        while (!q.isEmpty()) { // 多源 BFS
            var tmp = q;
            q = new ArrayList<>();
            for (var p : tmp) {
                for (var d : DIRS) {
                    int x = p[0] + d[0], y = p[1] + d[1];
                    if (0 <= x && x < n && 0 <= y && y < n && dis[x][y] < 0) {
                        q.add(new int[]{x, y});
                        dis[x][y] = groups.size();
                    }
                }
            }
            groups.add(q); // 相同 dis 分组记录
        }

        // 初始化并查集
        fa = new int[n * n];
        for (int i = 0; i < n * n; i++)
            fa[i] = i;

        for (int ans = groups.size() - 2; ans > 0; ans--) {
            var g = groups.get(ans);
            for (var p : groups.get(ans)) {
                int i = p[0], j = p[1];
                for (var d : DIRS) {
                    int x = i + d[0], y = j + d[1];
                    if (0 <= x && x < n && 0 <= y && y < n && dis[x][y] >= dis[i][j])
                        fa[find(x * n + y)] = find(i * n + j);
                }
            }
            if (find(0) == find(n * n - 1)) // 写这里判断更快些
                return ans;
        }
        return 0;
    }

    // 并查集模板
    private int[] fa;

    private int find(int x) {
        if (fa[x] != x) fa[x] = find(fa[x]);
        return fa[x];
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    int maximumSafenessFactor(vector<vector<int>> &grid) {
        int n = grid.size();
        vector<pair<int, int>> q;
        vector<vector<int>> dis(n, vector<int>(n, -1));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j]) {
                    q.emplace_back(i, j);
                    dis[i][j] = 0;
                }
            }
        }

        vector<vector<pair<int, int>>> groups = {q};
        while (!q.empty()) { // 多源 BFS
            vector<pair<int, int>> nq;
            for (auto &[i, j]: q) {
                for (auto &d: dirs) {
                    int x = i + d[0], y = j + d[1];
                    if (0 <= x && x < n && 0 <= y && y < n && dis[x][y] < 0) {
                        nq.emplace_back(x, y);
                        dis[x][y] = groups.size();
                    }
                }
            }
            groups.push_back(nq); // 相同 dis 分组记录
            q = move(nq);
        }

        // 并查集模板
        vector<int> fa(n * n);
        iota(fa.begin(), fa.end(), 0);
        function<int(int)> find = [&](int x) -> int { return fa[x] == x ? x : fa[x] = find(fa[x]); };

        for (int ans = (int) groups.size() - 2; ans > 0; ans--) {
            for (auto &[i, j]: groups[ans]) {
                for (auto &d: dirs) {
                    int x = i + d[0], y = j + d[1];
                    if (0 <= x && x < n && 0 <= y && y < n && dis[x][y] >= dis[i][j])
                        fa[find(x * n + y)] = find(i * n + j);
                }
            }
            if (find(0) == find(n * n - 1)) // 写这里判断更快些
                return ans;
        }
        return 0;
    }
};
```

```go [sol-Go]
func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	type pair struct{ x, y int }
	q := []pair{}
	dis := make([][]int, n)
	for i, row := range grid {
		dis[i] = make([]int, n)
		for j, x := range row {
			if x > 0 {
				q = append(q, pair{i, j})
			} else {
				dis[i][j] = -1
			}
		}
	}

	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	groups := [][]pair{q}
	for len(q) > 0 { // 多源 BFS
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] < 0 {
					q = append(q, pair{x, y})
					dis[x][y] = len(groups)
				}
			}
		}
		groups = append(groups, q) // 相同 dis 分组记录
	}

	// 并查集模板
	fa := make([]int, n*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for ans := len(groups) - 2; ans > 0; ans-- {
		for _, p := range groups[ans] {
			i, j := p.x, p.y
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] >= dis[i][j] {
					fa[find(x*n+y)] = find(i*n + j)
				}
			}
		}
		if find(0) == find(n*n-1) { // 写这里判断更快些
			return ans
		}
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$ 或者 $\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的长度。时间复杂度取决于并查集的实现，加上按秩合并的话，均摊地说并查集的操作可以视作是 $O(1)$ 的。
- 空间复杂度：$\mathcal{O}(n^2)$。
