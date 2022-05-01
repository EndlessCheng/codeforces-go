#### 提示 1

如果可以停留 $t$ 分钟，那么肯定也可以停留少于 $t$ 分钟；如果不能停留 $t$ 分钟，那么肯定也不能停留超过 $t$ 分钟。

因此可以二分最长停留时间。

#### 提示 2

为了避免遇到火，人需要尽可能快地到达安全屋，这可以用 BFS。

扩展火势也符合 BFS 的模型。

#### 提示 3

每过一分钟，将人能到达的位置向外扩充一圈，火势也向外扩充一圈。


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
                    for x, y in ((i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j)):
                        if 0 <= x < m and 0 <= y < n and grid[x][y] != 2 and (x, y) not in fire:
                            fire.add((x, y))
                            f.append((x, y))
            while t and f:
                spread_fire()  # 扩充至多 t 分钟的火势
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
                        for x, y in ((i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j)):
                            if 0 <= x < m and 0 <= y < n and grid[x][y] != 2 and (x, y) not in fire and (x, y) not in vis:
                                if x == m - 1 and y == n - 1:  # 我们安全了…暂时。
                                    return False
                                vis.add((x, y))
                                q.append((x, y))
                spread_fire()  # 扩充 1 分钟的火势
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
        // 扩充至多 t 分钟的火势
        while (t > 0 && f.size() > 0) {
            var tmp = f;
            f = new ArrayList<>();
            for (var p : tmp)
                for (var d : dirs) {
                    int x = p[0] + d[0], y = p[1] + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] != 2) {
                        fire[x][y] = true;
                        f.add(new int[]{x, y});
                    }
                }
            t--;
        }
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
                        if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && !vis[x][y] && grid[x][y] != 2) {
                            if (x == m - 1 && y == n - 1) return true; // 我们安全了…暂时。
                            vis[x][y] = true;
                            q.add(new int[]{x, y});
                        }
                    }
            // 扩充 1 分钟的火势
            tmp = f;
            f = new ArrayList<>();
            for (var p : tmp)
                for (var d : dirs) {
                    int x = p[0] + d[0], y = p[1] + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] != 2) {
                        fire[x][y] = true;
                        f.add(new int[]{x, y});
                    }
                }
        }
        return false;
    }
}
```

```C++ [sol1-C++]
static const int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

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
                for (auto &d : dirs) {
                    int x = i + d[0], y = j + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] != 2) {
                        fire[x][y] = true;
                        nf.emplace_back(x, y);
                    }
                }
            f = move(nf);
        };
        while (t-- && !f.empty()) spread_fire(); // 扩充至多 t 分钟的火势
        if (fire[0][0]) return false; // 起点着火，寄

        bool vis[m][n]; memset(vis, 0, sizeof(vis));
        vis[0][0] = true;
        q.emplace_back(0, 0);
        while (!q.empty()) {
            vector<pair<int, int>> nq;
            for (auto &[i, j] : q)
                if (!fire[i][j])
                    for (auto &d : dirs) {
                        int x = i + d[0], y = j + d[1];
                        if (0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && !vis[x][y] && grid[x][y] != 2) {
                            if (x == m - 1 && y == n - 1) return true; // 我们安全了…暂时。
                            vis[x][y] = true;
                            nq.emplace_back(x, y);
                        }
                    }
            q = move(nq);
            spread_fire(); // 扩充 1 分钟的火势
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
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !fire[x][y] && grid[x][y] != 2 {
						fire[x][y] = true
						f = append(f, pair{x, y})
					}
				}
			}
		}
		for ; t > 0 && len(f) > 0; t-- {
			spreadFire() // 扩充至多 t 分钟的火势
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
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && !fire[x][y] && grid[x][y] != 2 {
							if x == m-1 && y == n-1 { // 我们安全了…暂时。
								return false
							}
							vis[x][y] = true
							q = append(q, pair{x, y})
						}
					}
				}
			}
			spreadFire() // 扩充 1 分钟的火势
		}
		return true // 寄
	}) - 1
	if ans < m*n {
		return ans
	}
	return 1e9
}
```
