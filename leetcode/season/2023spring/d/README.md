## 本题视频讲解

见[【力扣杯2023春·个人赛】](https://www.bilibili.com/video/BV1dg4y1j78A/)第四题。

## 思路

1. 遍历，找到 S 和 T 的位置。
2. BFS，计算 T 到其余点的最短距离。
3. 如果发现 S 无法到达 T，直接返回 $-1$。
4. 二分答案 $\textit{maxDis}$：看能否在「附加负面效果」的情况下，移动不超过 $\textit{maxDis}$ 步到达终点。我写的 DFS，如果 DFS 中的某个位置，守护者使用卷轴传送小扣，并将小扣传送到一个无法到达终点，或者无法在 $\textit{maxDis}$ 步内到达终点的位置，则不再 DFS（**相当于这种位置小扣不能走**）。在这种情况下，只要可以到达终点，则说明答案至多为  $\textit{maxDis}$，否则说明答案大于 $\textit{maxDis}$。

> 注：如果守护者无法发动传送，则答案为 $0$。

```py [sol1-Python3]
class Solution:
    def challengeOfTheKeeper(self, maze: List[str]) -> int:
        m, n = len(maze), len(maze[0])

        # 1. 找到起点终点坐标
        for i, row in enumerate(maze):
            for j, c in enumerate(row):
                if c == 'S':
                    sx, sy = i, j
                elif c == 'T':
                    tx, ty = i, j

        # 2. BFS 计算终点到其余点的最短距离
        dis_from_t = [[inf] * n for _ in range(m)]
        dis_from_t[tx][ty] = 0
        q = [(tx, ty)]
        step = 1
        while q:
            tmp = q
            q = []
            for i, j in tmp:
                for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                    if 0 <= x < m and 0 <= y < n and maze[x][y] != '#' and dis_from_t[x][y] == inf:
                        dis_from_t[x][y] = step
                        q.append((x, y))
            step += 1

        # 3. 剪枝：如果 S 无法到达 T，直接返回 -1
        if dis_from_t[sx][sy] == inf:
            return -1

        # 4. 二分答案 https://www.bilibili.com/video/BV1AP41137w7/
        vis = [[-1] * n for _ in range(m)]
        def check(max_dis: int) -> bool:
            # DFS，看能否在「附加负面效果」的情况下，移动不超过 max_dis 步到达终点
            def dfs(i: int, j: int) -> bool:
                if i < 0 or i >= m or j < 0 or j >= n or vis[i][j] == max_dis or maze[i][j] == '#':
                    return False
                if maze[i][j] == 'T':  # 到达终点
                    return True
                vis[i][j] = max_dis  # 为避免反复创建 vis，用一个每次二分都不一样的数来标记
                # 守护者使用卷轴传送小扣，如果小扣无法在 maxDis 步内到达终点，则返回 false
                if maze[i][j] == '.' and \
                   (maze[i][n - 1 - j] != '#' and dis_from_t[i][n - 1 - j] > max_dis or
                    maze[m - 1 - i][j] != '#' and dis_from_t[m - 1 - i][j] > max_dis):
                    return False
                for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                    if dfs(x, y):
                        return True
                return False
            return dfs(sx, sy)
        ans = bisect_left(range(m * n + 1), True, key=check)
        return -1 if ans > m * n else ans
```

```java [sol1-Java]
class Solution {
    private final static int[][] DIRS = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    private char[][] maze;
    private int[][] disFromT, vis;
    private int sx, sy, maxDis;

    public int challengeOfTheKeeper(String[] Maze) {
        int m = Maze.length, n = Maze[0].length(), tx = 0, ty = 0;
        maze = new char[m][];
        disFromT = new int[m][n];
        // 1. 找到起点终点坐标
        for (int i = 0; i < m; ++i) {
            maze[i] = Maze[i].toCharArray();
            for (int j = 0; j < n; ++j) {
                disFromT[i][j] = Integer.MAX_VALUE;
                if (maze[i][j] == 'S') {
                    sx = i;
                    sy = j;
                } else if (maze[i][j] == 'T') {
                    tx = i;
                    ty = j;
                }
            }
        }

        // 2. BFS 计算终点到其余点的最短距离
        disFromT[tx][ty] = 0;
        var q = new ArrayList<int[]>();
        q.add(new int[]{tx, ty});
        for (int step = 1; !q.isEmpty(); ++step) {
            var tmp = q;
            q = new ArrayList<>();
            for (var p : tmp) {
                for (var d : DIRS) {
                    int x = p[0] + d[0], y = p[1] + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && maze[x][y] != '#' && disFromT[x][y] == Integer.MAX_VALUE) {
                        disFromT[x][y] = step;
                        q.add(new int[]{x, y});
                    }
                }
            }
        }

        // 3. 剪枝：如果 S 无法到达 T，直接返回 -1
        if (disFromT[sx][sy] == Integer.MAX_VALUE)
            return -1;

        // 4. 二分答案 https://www.bilibili.com/video/BV1AP41137w7/
        vis = new int[m][n];
        int left = -1, right = m * n + 1;
        while (left + 1 < right) {
            maxDis = (left + right) >>> 1;
            if (dfs(sx, sy)) right = maxDis;
            else left = maxDis;
        }
        return right > m * n ? -1 : right;
    }

    private boolean dfs(int x, int y) {
        int m = maze.length, n = maze[0].length;
        if (x < 0 || x >= m || y < 0 || y >= n || vis[x][y] == maxDis + 1 || maze[x][y] == '#')
            return false;
        if (maze[x][y] == 'T') // 到达终点
            return true;
        vis[x][y] = maxDis + 1; // 为避免反复创建 vis，用一个每次二分都不一样的数来标记
        // 守护者使用卷轴传送小扣，如果小扣无法在 maxDis 步内到达终点，则返回 false
        if (maze[x][y] == '.' &&
            (maze[m - x - 1][y] != '#' && disFromT[m - 1 - x][y] > maxDis ||
             maze[x][n - 1 - y] != '#' && disFromT[x][n - 1 - y] > maxDis))
            return false;
        for (var d : DIRS)
            if (dfs(x + d[0], y + d[1]))
                return true;
        return false;
    }
}
```

```cpp [sol1-C++]
class Solution {
    static constexpr int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
public:
    int challengeOfTheKeeper(vector<string> &maze) {
        // 1. 找到起点终点坐标
        int m = maze.size(), n = maze[0].size(), sx, sy, tx, ty, disFromT[m][n];
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                disFromT[i][j] = INT_MAX;
                if (maze[i][j] == 'S')
                    sx = i, sy = j;
                else if (maze[i][j] == 'T')
                    tx = i, ty = j;
            }
        }

        // 2. BFS 计算终点到其余点的最短距离
        disFromT[tx][ty] = 0;
        vector<pair<int, int>> q = {{tx, ty}};
        for (int step = 1; !q.empty(); ++step) {
            vector<pair<int, int>> nq;
            for (auto &[i, j]: q) {
                for (auto &d: dirs) {
                    int x = i + d[0], y = j + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && maze[x][y] != '#' && disFromT[x][y] == INT_MAX) {
                        disFromT[x][y] = step;
                        nq.emplace_back(x, y);
                    }
                }
            }
            q = move(nq);
        }

        // 3. 剪枝：如果 S 无法到达 T，直接返回 -1
        if (disFromT[sx][sy] == INT_MAX)
            return -1;

        // 4. 二分答案 https://www.bilibili.com/video/BV1AP41137w7/
        int vis[m][n], maxDis;
        memset(vis, -1, sizeof(vis));
        auto dfs = [&](auto&& dfs, int x, int y) {
            if (x < 0 || x >= m || y < 0 || y >= n || vis[x][y] == maxDis || maze[x][y] == '#')
                return false;
            if (maze[x][y] == 'T') // 到达终点
                return true;
            vis[x][y] = maxDis; // 为避免反复创建 vis，用一个每次二分都不一样的数来标记
            // 守护者使用卷轴传送小扣，如果小扣无法在 maxDis 步内到达终点，则返回 false
            if (maze[x][y] == '.' &&
                (maze[m - x - 1][y] != '#' && disFromT[m - 1 - x][y] > maxDis ||
                 maze[x][n - 1 - y] != '#' && disFromT[x][n - 1 - y] > maxDis))
                return false;
            for (auto &d: dirs)
                if (dfs(dfs, x + d[0], y + d[1]))
                    return true;
            return false;
        };
        int left = -1, right = m * n + 1;
        while (left + 1 < right) {
            maxDis = left + (right - left) / 2;
            (dfs(dfs, sx, sy) ? right : left) = maxDis;
        }
        return right > m * n ? -1 : right;
    }
};
```

```go [sol1-Go]
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func challengeOfTheKeeper(maze []string) int {
	m, n := len(maze), len(maze[0])

	// 1. 找到起点终点坐标
	var sx, sy, tx, ty int
	for i, row := range maze {
		for j, c := range row {
			if c == 'S' {
				sx, sy = i, j
			} else if c == 'T' {
				tx, ty = i, j
			}
		}
	}

	// 2. BFS 计算终点到其余点的最短距离
	disFromT := make([][]int, m)
	for i := range disFromT {
		disFromT[i] = make([]int, n)
		for j := range disFromT[i] {
			disFromT[i][j] = math.MaxInt
		}
	}
	disFromT[tx][ty] = 0
	q := []pair{{tx, ty}}
	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && maze[x][y] != '#' && disFromT[x][y] == math.MaxInt {
					disFromT[x][y] = step
					q = append(q, pair{x, y})
				}
			}
		}
	}

	// 3. 剪枝：如果 S 无法到达 T，直接返回 -1
	if disFromT[sx][sy] == math.MaxInt {
		return -1
	}

	// 4. 二分答案 https://www.bilibili.com/video/BV1AP41137w7/
	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
	}
	ans := sort.Search(m*n+1, func(maxDis int) bool {
		// DFS，看能否在「附加负面效果」的情况下，移动不超过 maxDis 步到达终点
		var dfs func(int, int) bool
		dfs = func(i, j int) bool {
			if i < 0 || i >= m || j < 0 || j >= n || vis[i][j] == maxDis+1 || maze[i][j] == '#' {
				return false
			}
			if maze[i][j] == 'T' { // 到达终点
				return true
			}
			vis[i][j] = maxDis + 1 // 为避免反复创建 vis，用一个每次二分都不一样的数来标记
			if maze[i][j] == '.' {
				// 守护者使用卷轴传送小扣，如果小扣无法在 maxDis 步内到达终点，则返回 false
				if x, y := i, n-1-j; maze[x][y] != '#' && disFromT[x][y] > maxDis {
					return false
				}
				if x, y := m-1-i, j; maze[x][y] != '#' && disFromT[x][y] > maxDis {
					return false
				}
			}
			// 枚举四个方向
			for _, d := range dirs {
				if dfs(i+d.x, j+d.y) { // 到达终点
					return true
				}
			}
			return false // 无法到达终点
		}
		return dfs(sx, sy)
	})
	if ans > m*n { // 守护者使用卷轴传送小扣，可以把小扣传送到一个无法到达终点的位置
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log(mn))$，其中 $m$ 和 $n$ 分别为 $\textit{maze}$ 的行数和列数。本题 $m=n$。
- 空间复杂度：$\mathcal{O}(mn)$。

## 相似题目

- [778. 水位上升的泳池中游泳](https://leetcode.cn/problems/swim-in-rising-water/)

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

