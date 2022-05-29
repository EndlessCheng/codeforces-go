#### 提示 1

把障碍物当作可以经过的单元格，经过它的「花费」为 $1$。

#### 提示 2

问题转化成从起点到终点的最短路。

#### 提示 3

我们可以用 Dijkstra，但还可以用 0-1 BFS 来将时间复杂度优化至 $O(mn)$。

具体见 [1368. 使网格图至少有一条有效路径的最小代价](https://leetcode.cn/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/) 官方题解的方法二，本文不再赘述。

#### 复杂度分析

- 时间复杂度：$O(mn)$。
- 空间复杂度：$O(mn)$。

```Python [sol1-Python3]
class Solution:
    def minimumObstacles(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dis = [[inf] * n for _ in range(m)]
        dis[0][0] = 0
        q = deque([(0, 0)])
        while q:
            x, y = q.popleft()
            for nx, ny in (x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1):
                if 0 <= nx < m and 0 <= ny < n:
                    g = grid[x][y]
                    if dis[x][y] + g < dis[nx][ny]:
                        dis[nx][ny] = dis[x][y] + g
                        if g == 0: q.appendleft((nx, ny))
                        else: q.append((nx, ny))
        return dis[m - 1][n - 1]
```

```java [sol1-Java]
class Solution {
    final static int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int minimumObstacles(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        var dis = new int[m][n];
        for (var i = 0; i < m; i++) Arrays.fill(dis[i], Integer.MAX_VALUE);
        dis[0][0] = 0;
        var q = new ArrayDeque<int[]>();
        q.addFirst(new int[]{0, 0});
        while (!q.isEmpty()) {
            var p = q.pollFirst();
            int x = p[0], y = p[1];
            for (var d : dirs) {
                int nx = x + d[0], ny = y + d[1];
                if (0 <= nx && nx < m && 0 <= ny && ny < n) {
                    var g = grid[nx][ny];
                    if (dis[x][y] + g < dis[nx][ny]) {
                        dis[nx][ny] = dis[x][y] + g;
                        if (g == 0) q.addFirst(new int[]{nx, ny});
                        else q.addLast(new int[]{nx, ny});
                    }
                }
            }
        }
        return dis[m - 1][n - 1];
    }
}
```

```C++ [sol1-C++]
class Solution {
    static constexpr int dirs[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

public:
    int minimumObstacles(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        int dis[m][n];
        memset(dis, 0x3f, sizeof(dis));
        dis[0][0] = 0;
        deque<pair<int, int>> q;
        q.emplace_front(0, 0);
        while (!q.empty()) {
            auto[x, y] = q.front();
            q.pop_front();
            for (auto &[dx, dy] : dirs) {
                int nx = x + dx, ny = y + dy;
                if (0 <= nx && nx < m && 0 <= ny && ny < n) {
                    int g = grid[nx][ny];
                    if (dis[x][y] + g < dis[nx][ny]) {
                        dis[nx][ny] = dis[x][y] + g;
                        g == 0 ? q.emplace_front(nx, ny) : q.emplace_back(nx, ny);
                    }
                }
            }
        }
        return dis[m - 1][n - 1];
    }
};
```

```go [sol1-Go]
type pair struct{ x, y int }
var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumObstacles(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = m * n
		}
	}
	dis[0][0] = 0
	q := [2][]pair{{{}}}
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		for _, d := range dir4 {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				g := grid[x][y]
				if dis[p.x][p.y]+g < dis[x][y] {
					dis[x][y] = dis[p.x][p.y] + g
					q[g] = append(q[g], pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1]
}
```

