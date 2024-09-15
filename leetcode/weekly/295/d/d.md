从 $(i,j)$ 移动到与其相邻的格子 $(i',j')$，视作一条从 $(i,j)$ 到 $(i',j')$ 的有向边，边权为 $\textit{grid}[i'][j']$。

问题变成计算从起点到终点的最短路。

这可以用 Dijkstra 算法解决，原理请看 [Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

由于本题的边权只有 $0$ 和 $1$，也可以用 **0-1 BFS** 解决。

0-1 BFS 本质是对 Dijkstra 算法的优化。因为边权只有 $0$ 和 $1$，我们可以把最小堆换成**双端队列**，遇到 $0$ 边权就加入**队首**，遇到 $1$ 边权就加入**队尾**，这样可以保证队首总是最小的，就不需要最小堆了。

[本题视频讲解](https://www.bilibili.com/video/BV1iF41157dG)（第四题）。

```Python [sol-Python3]
class Solution:
    def minimumObstacles(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dis = [[inf] * n for _ in range(m)]
        dis[0][0] = 0
        q = deque([(0, 0)])
        while q:
            i, j = q.popleft()
            for x, y in (i, j + 1), (i, j - 1), (i + 1, j), (i - 1, j):
                if 0 <= x < m and 0 <= y < n:
                    g = grid[x][y]
                    if dis[i][j] + g < dis[x][y]:
                        dis[x][y] = dis[i][j] + g
                        if g == 0:
                            q.appendleft((x, y))
                        else:
                            q.append((x, y))
        return dis[-1][-1]
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}};

    public int minimumObstacles(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] dis = new int[m][n];
        for (int[] row : dis) {
            Arrays.fill(row, Integer.MAX_VALUE);
        }
        dis[0][0] = 0;
        Deque<int[]> q = new ArrayDeque<>();
        q.addFirst(new int[]{0, 0});
        while (!q.isEmpty()) {
            int[] p = q.pollFirst();
            int i = p[0];
            int j = p[1];
            for (int[] d : DIRS) {
                int x = i + d[0];
                int y = j + d[1];
                if (0 <= x && x < m && 0 <= y && y < n) {
                    int g = grid[x][y];
                    if (dis[i][j] + g < dis[x][y]) {
                        dis[x][y] = dis[i][j] + g;
                        if (g == 0) {
                            q.addFirst(new int[]{x, y});
                        } else {
                            q.addLast(new int[]{x, y});
                        }
                    }
                }
            }
        }
        return dis[m - 1][n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
public:
    int minimumObstacles(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> dis(m, vector<int>(n, INT_MAX));
        dis[0][0] = 0;
        deque<pair<int, int>> q;
        q.emplace_front(0, 0);
        while (!q.empty()) {
            auto [i, j] = q.front();
            q.pop_front();
            for (auto& [dx, dy] : DIRS) {
                int x = i + dx, y = j + dy;
                if (0 <= x && x < m && 0 <= y && y < n) {
                    int g = grid[x][y];
                    if (dis[i][j] + g < dis[x][y]) {
                        dis[x][y] = dis[i][j] + g;
                        g == 0 ? q.emplace_front(x, y) : q.emplace_back(x, y);
                    }
                }
            }
        }
        return dis[m - 1][n - 1];
    }
};
```

```go [sol-Go]
func minimumObstacles(grid [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0
	q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		for _, d := range dirs {
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个点至多入队两次。
- 空间复杂度：$\mathcal{O}(mn)$。

## 思考题

构造一个 $\textit{grid}$，使得上述算法消耗的空间尽量多。

欢迎在评论区分享你的思路/代码。

更多相似题目，见下面网格图题单中的「**0-1 BFS**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
