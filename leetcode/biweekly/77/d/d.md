## 方法一：二分答案 + BFS

#### 提示 1

如果可以停留 $t$ 分钟，那么肯定也可以停留少于 $t$ 分钟；如果不能停留 $t$ 分钟，那么肯定也不能停留超过 $t$ 分钟。

因此可以二分最长停留时间。

#### 提示 2

为了避免遇到火，人需要尽可能快地到达安全屋，这可以用 BFS。

蔓延火势也符合多源 BFS 的模型。

#### 提示 3

每过一分钟，将人能到达的位置向外扩充一圈，火势也向外蔓延一圈。

```Python [sol1-Python3]
class Solution:
    def maximumMinutes(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])

        def check(t: int) -> bool:
            f = [(i, j) for i, row in enumerate(grid) for j, v in enumerate(row) if v == 1]
            fire = set(f)
            def spread_fire():
                nonlocal f
                tmp = f
                f = []
                for i, j in tmp:
                    for x, y in (i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j):
                        if 0 <= x < m and 0 <= y < n and grid[x][y] == 0 and (x, y) not in fire:
                            fire.add((x, y))
                            f.append((x, y))
            while t and f:
                spread_fire()  # 蔓延至多 t 分钟的火势
                t -= 1
            if (0, 0) in fire:  # 起点着火，寄
                return True

            q = [(0, 0)]
            vis = set(q)
            while q:
                tmp = q
                q = []
                for i, j in tmp:
                    if (i, j) not in fire:
                        for x, y in (i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j):
                            if 0 <= x < m and 0 <= y < n and grid[x][y] == 0 and (x, y) not in fire and (x, y) not in vis:
                                if x == m - 1 and y == n - 1:  # 我们安全了…暂时。
                                    return False
                                vis.add((x, y))
                                q.append((x, y))
                spread_fire()  # 蔓延 1 分钟的火势
            return True  # 寄

        ans = bisect_left(range(m * n + 1), True, key=check) - 1
        return ans if ans < m * n else 10 ** 9
```

```java [sol1-Java]
class Solution {
    static int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int maximumMinutes(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int left = -1, right = m * n;
        while (left < right) {
            var mid = (left + right + 1) / 2;
            if (check(grid, mid)) left = mid;
            else right = mid - 1;
        }
        return left < m * n ? left : (int) 1e9;
    }

    boolean check(int[][] grid, int t) {
        int m = grid.length, n = grid[0].length;
        var fire = new boolean[m][n];
        var f = new ArrayList<int[]>();
        for (var i = 0; i < m; i++)
            for (var j = 0; j < n; j++)
                if (grid[i][j] == 1) {
                    fire[i][j] = true;
                    f.add(new int[]{i, j});
                }
        while (t-- > 0 && f.size() > 0)
            f = spreadFire(grid, fire, f); // 蔓延至多 t 分钟的火势
        if (fire[0][0]) return false; // 起点着火，寄

        var vis = new boolean[m][n];
        vis[0][0] = true;
        var q = new ArrayList<int[]>();
        q.add(new int[]{0, 0});
        while (q.size() > 0) {
            var tmp = q;
            q = new ArrayList<>();
            for (var p : tmp)
                if (!fire[p[0]][p[1]])
                    for (var d : dirs) {
                        int x = p[0] + d[0], y = p[1] + d[1];
                        if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && !vis[x][y] && grid[x][y] == 0) {
                            if (x == m - 1 && y == n - 1) return true; // 我们安全了…暂时。
                            vis[x][y] = true;
                            q.add(new int[]{x, y});
                        }
                    }
            f = spreadFire(grid, fire, f); // 蔓延 1 分钟的火势
        }
        return false; // 寄
    }

    ArrayList<int[]> spreadFire(int[][] grid, boolean[][] fire, ArrayList<int[]> f) {
        int m = grid.length, n = grid[0].length;
        var tmp = f;
        f = new ArrayList<>();
        for (var p : tmp)
            for (var d : dirs) {
                int x = p[0] + d[0], y = p[1] + d[1];
                if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] == 0) {
                    fire[x][y] = true;
                    f.add(new int[]{x, y});
                }
            }
        return f;
    }
}    
```

```C++ [sol1-C++]
const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

class Solution {
    bool check(vector<vector<int>> &grid, int t) {
        int m = grid.size(), n = grid[0].size();
        bool fire[m][n]; memset(fire, 0, sizeof(fire));
        vector<pair<int, int>> f, q;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (grid[i][j] == 1) {
                    fire[i][j] = true;
                    f.emplace_back(i, j);
                }
        auto spread_fire = [&]() {
            vector<pair<int, int>> nf;
            for (auto &[i, j] : f)
                for (auto [dx, dy] : dirs) {
                    int x = i + dx, y = j + dy;
                    if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] == 0) {
                        fire[x][y] = true;
                        nf.emplace_back(x, y);
                    }
                }
            f = move(nf);
        };
        while (t-- && !f.empty()) spread_fire(); // 蔓延至多 t 分钟的火势
        if (fire[0][0]) return false; // 起点着火，寄

        bool vis[m][n]; memset(vis, 0, sizeof(vis));
        vis[0][0] = true;
        q.emplace_back(0, 0);
        while (!q.empty()) {
            vector<pair<int, int>> nq;
            for (auto &[i, j] : q)
                if (!fire[i][j])
                    for (auto [dx, dy] : dirs) {
                        int x = i + dx, y = j + dy;
                        if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && !vis[x][y] && grid[x][y] == 0) {
                            if (x == m - 1 && y == n - 1) return true; // 我们安全了…暂时。
                            vis[x][y] = true;
                            nq.emplace_back(x, y);
                        }
                    }
            q = move(nq);
            spread_fire(); // 蔓延 1 分钟的火势
        }
        return false; // 寄
    }

public:
    int maximumMinutes(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        int left = -1, right = m * n;
        while (left < right) {
            int mid = (left + right + 1) / 2;
            if (check(grid, mid)) left = mid;
            else right = mid - 1;
        }
        return left < m * n ? left : 1e9;
    }
};
```

```go [sol1-Go]
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := sort.Search(m*n+1, func(t int) bool {
		fire := make([][]bool, m)
		for i := range fire {
			fire[i] = make([]bool, n)
		}
		f := []pair{}
		for i, row := range grid {
			for j, v := range row {
				if v == 1 {
					fire[i][j] = true
					f = append(f, pair{i, j})
				}
			}
		}
		spreadFire := func() {
			tmp := f
			f = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] == 0 {
						fire[x][y] = true
						f = append(f, pair{x, y})
					}
				}
			}
		}
		for ; t > 0 && len(f) > 0; t-- {
			spreadFire() // 蔓延至多 t 分钟的火势
		}
		if fire[0][0] { // 起点着火，寄
			return true
		}

		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		q := []pair{{}}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, p := range tmp {
				if !fire[p.x][p.y] {
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && !fire[x][y] && grid[x][y] == 0 {
							if x == m-1 && y == n-1 { // 我们安全了…暂时。
								return false
							}
							vis[x][y] = true
							q = append(q, pair{x, y})
						}
					}
				}
			}
			spreadFire() // 蔓延 1 分钟的火势
		}
		return true // 寄
	}) - 1
	if ans < m*n {
		return ans
	}
	return 1e9
}
```

#### 复杂度分析

- 时间复杂度：$O(mn\log mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(mn)$。

## 方法二：两次 BFS

首先通过 BFS 处理出人到每个格子的最短时间 $\textit{manTime}$，以及火到每个格子的最短时间 $\textit{fireTime}$。

如果 $\textit{manTime}[m-1][n-1]<0$，人无法到达终点，返回 $-1$。

如果 $\textit{fireTime}[m-1][n-1]<0$，火无法到达终点，返回 $10^9$。

记 $\textit{ans}=\textit{fireTime}[m-1][n-1]-\textit{manTime}[m-1][n-1]$。

如果 $\textit{ans} < 0$，说明火比人先到终点，返回 $-1$。

如果 $\textit{ans} > 0$，说明人比火先到终点。注意不会出现中途火把人烧到的情况，如果出现，那么火可以沿着人走的最短路到达终点，不会出现人比火先到的情况，与实际矛盾。

最后还需要细致分析一下：

- 如果火和人是从不同方向到达终点的（左侧和上侧），那么答案可以是 $\textit{ans}$，即人可以等待 $\textit{ans}$ 时间，最终与火同时到达终点。
- 如果火和人是从同一方向到达终点的，也就意味着火一直跟在人的后面，那么在中途不能出现人火重合的情况，所以答案应该是 $\textit{ans}-1$。

```Python [sol2-Python3]
class Solution:
    def maximumMinutes(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        def bfs(q: List[Tuple[int, int]]) -> (int, int, int):
            time = [[-1] * n for _ in range(m)]
            for i, j in q:
                time[i][j] = 0
            t = 1
            while q:
                tmp = q
                q = []
                for i, j in tmp:
                    for x, y in (i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j):
                        if 0 <= x < m and 0 <= y < n and grid[x][y] == 0 and time[x][y] < 0:
                            time[x][y] = t
                            q.append((x, y))
                t += 1
            return time[-1][-1], time[-1][-2], time[-2][-1]

        man_to_house_time, m1, m2 = bfs([(0, 0)])
        if man_to_house_time < 0: return -1  # 人无法到终点
        fire_to_house_time, f1, f2 = bfs([(i, j) for i, row in enumerate(grid) for j, v in enumerate(row) if v == 1])
        if fire_to_house_time < 0: return 10 ** 9  # 火无法到终点
        ans = fire_to_house_time - man_to_house_time
        if ans < 0: return -1  # 火比人先到终点
        if m1 < 0 or m2 < 0 or f1 - m1 == f2 - m2 == ans:
            return ans - 1  # 火只会跟在人的后面，在到达终点前，人和火不能重合
        return ans  # 人和火可以同时到终点
```

```java [sol2-Java]
class Solution {
    private static final int[][] DIRS = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int maximumMinutes(int[][] grid) {
        var res = bfs(grid, List.of(new int[]{0, 0}));
        int manToHouseTime = res[0], m1 = res[1], m2 = res[2];
        if (manToHouseTime < 0) return -1; // 人无法到终点

        var fires = new ArrayList<int[]>();
        for (var i = 0; i < grid.length; i++)
            for (var j = 0; j < grid[i].length; j++)
                if (grid[i][j] == 1) fires.add(new int[]{i, j});
        res = bfs(grid, fires);
        int fireToHouseTime = res[0], f1 = res[1], f2 = res[2];
        if (fireToHouseTime < 0) return (int) 1e9; // 火无法到终点

        int ans = fireToHouseTime - manToHouseTime;
        if (ans < 0) return -1; // 火比人先到终点
        if (m1 < 0 || m2 < 0 || f1 - m1 == ans && f2 - m2 == ans)
            return ans - 1; // 火只会跟在人的后面，在到达终点前，人和火不能重合
        return ans;// 人和火可以同时到终点
    }

    private int[] bfs(int[][] grid, List<int[]> q) {
        int m = grid.length, n = grid[0].length;
        var time = new int[m][n];
        for (int i = 0; i < m; i++)
            Arrays.fill(time[i], -1);
        for (var p : q)
            time[p[0]][p[1]] = 0;
        for (int t = 1; !q.isEmpty(); ++t) {
            var tmp = q;
            q = new ArrayList<>();
            for (var p : tmp)
                for (var d : DIRS) {
                    int x = p[0] + d[0], y = p[1] + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0) {
                        time[x][y] = t;
                        q.add(new int[]{x, y});
                    }
                }
        }
        return new int[]{time[m - 1][n - 1], time[m - 1][n - 2], time[m - 2][n - 1]};
    }
}
```

```C++ [sol2-C++]
const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

class Solution {
public:
    int maximumMinutes(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        auto bfs = [&](vector<pair<int, int>> &q) -> tuple<int, int, int> {
            int time[m][n];
            memset(time, -1, sizeof(time));
            for (auto &[i, j] : q)
                time[i][j] = 0;
            for (int t = 1; !q.empty(); ++t) {
                vector<pair<int, int>> nq;
                for (auto &[i, j] : q) {
                    for (auto[dx, dy] : dirs) {
                        int x = i + dx, y = j + dy;
                        if (0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0) {
                            time[x][y] = t;
                            nq.emplace_back(x, y);
                        }
                    }
                }
                q = move(nq);
            }
            return {time[m - 1][n - 1], time[m - 1][n - 2], time[m - 2][n - 1]};
        };

        vector<pair<int, int>> q = {{0, 0}};
        auto [man_to_house_time, m1, m2] = bfs(q);
        if (man_to_house_time < 0) return -1; // 人无法到终点

        vector<pair<int, int>> fires;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (grid[i][j] == 1) fires.emplace_back(i, j);
        auto [fire_to_house_time, f1, f2] = bfs(fires);
        if (fire_to_house_time < 0) return 1e9; // 火无法到终点

        int ans = fire_to_house_time - man_to_house_time;
        if (ans < 0) return -1; // 火比人先到终点
        if (m1 < 0 || m2 < 0 || f1 - m1 == ans && f2 - m2 == ans)
            return ans - 1; // 火只会跟在人的后面，在到达终点前，人和火不能重合
        return ans;// 人和火可以同时到终点
    }
};
```

```go [sol2-Go]
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	bfs := func(q []pair) (int, int, int) {
		time := make([][]int, m)
		for i := range time {
			time[i] = make([]int, n)
			for j := range time[i] {
				time[i][j] = -1
			}
		}
		for _, p := range q {
			time[p.x][p.y] = 0
		}
		for t := 1; len(q) > 0; t++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dirs {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 0 && time[x][y] < 0 {
						time[x][y] = t
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return time[m-1][n-1], time[m-1][n-2], time[m-2][n-1]
	}

	manToHouseTime, m1, m2 := bfs([]pair{{}})
	if manToHouseTime < 0 {
		return -1 // 人无法到终点
	}

	fires := []pair{}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				fires = append(fires, pair{i, j})
			}
		}
	}
	fireToHouseTime, f1, f2 := bfs(fires)
	if fireToHouseTime < 0 {
		return 1e9 // 火无法到终点
	}

	ans := fireToHouseTime - manToHouseTime
	if ans < 0 {
		return -1 // 火比人先到终点
	}
	if m1 < 0 || m2 < 0 || f1-m1 == ans && f2-m2 == ans {
		return ans - 1 // 火只会跟在人的后面，在到达终点前，人和火不能重合
	}
	return ans // 人和火可以同时到终点
}
```

#### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(mn)$。

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
