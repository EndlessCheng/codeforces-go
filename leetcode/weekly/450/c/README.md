**前置题目**：[2290. 到达角落需要移除障碍物的最小数目](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/)

本质上是计算如下图的最短路：

- 所有相同字母之间都有一条边权为 $0$ 的边。
- 所有相邻格子之间都有一条边权为 $1$ 的边。

这可以用 Dijkstra 算法解决。不过，对于边权只有 $0$ 和 $1$ 的特殊图，可以用 **0-1 BFS**（双端队列）解决，做到线性时间：

- 如果当前格子是非字母格子，那么像普通 BFS 那样遍历四方向的相邻格子。
- 如果当前格子是字母格子，那么除了像普通 BFS 那样遍历四方向的相邻格子以外，还需要遍历该字母的所有传送门，传送过去（边权为 $0$）。使用所有传送门后，清空传送门的位置列表，避免反复使用传送门。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Z3JGzwEU9/?t=19m54s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minMoves(self, matrix: List[str]) -> int:
        if matrix[-1][-1] == '#':
            return -1

        m, n = len(matrix), len(matrix[0])
        pos = defaultdict(list)
        for i, row in enumerate(matrix):
            for j, c in enumerate(row):
                if c.isupper():
                    pos[c].append((i, j))

        DIRS = [(0, -1), (0, 1), (-1, 0), (1, 0)]
        dis = [[inf] * n for _ in range(m)]
        dis[0][0] = 0
        q = deque([(0, 0)])

        while q:
            x, y = q.popleft()
            d = dis[x][y]

            if x == m - 1 and y == n - 1:  # 到达终点
                return d

            c = matrix[x][y]
            if c in pos:
                # 使用所有传送门
                for px, py in pos[c]:
                    if d < dis[px][py]:
                        dis[px][py] = d
                        q.appendleft((px, py))
                del pos[c]  # 避免重复使用传送门

            # 下面代码和普通 BFS 是一样的
            for dx, dy in DIRS:
                nx, ny = x + dx, y + dy
                if 0 <= nx < m and 0 <= ny < n and matrix[nx][ny] != '#' and d + 1 < dis[nx][ny]:
                    dis[nx][ny] = d + 1
                    q.append((nx, ny))

        return -1
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = new int[][]{{0, -1}, {0, 1}, {-1, 0}, {1, 0}};

    public int minMoves(String[] matrix) {
        int m = matrix.length, n = matrix[0].length();
        if (matrix[m - 1].charAt(n - 1) == '#') {
            return -1;
        }

        List<int[]>[] pos = new ArrayList[26];
        Arrays.setAll(pos, i -> new ArrayList<>());
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                char c = matrix[i].charAt(j);
                if (Character.isUpperCase(c)) {
                    pos[c - 'A'].add(new int[]{i, j});
                }
            }
        }

        int[][] dis = new int[m][n];
        for (int[] row : dis) {
            Arrays.fill(row, Integer.MAX_VALUE);
        }
        dis[0][0] = 0;

        Deque<int[]> q = new ArrayDeque<>();
        q.addFirst(new int[]{0, 0});

        while (!q.isEmpty()) {
            int[] p = q.pollFirst();
            int x = p[0], y = p[1];
            int d = dis[x][y];
            
            if (x == m - 1 && y == n - 1) {
                return d;
            }

            char c = matrix[x].charAt(y);
            if (c != '.') {
                // 使用所有传送门
                for (int[] portal : pos[c - 'A']) {
                    int px = portal[0], py = portal[1];
                    if (d < dis[px][py]) {
                        dis[px][py] = d;
                        q.addFirst(new int[]{px, py});
                    }
                }
                pos[c - 'A'].clear(); // 避免重复使用传送门
            }

            // 下面代码和普通 BFS 是一样的
            for (int[] dir : DIRS) {
                int nx = x + dir[0], ny = y + dir[1];
                if (0 <= nx && nx < m && 0 <= ny && ny < n && matrix[nx].charAt(ny) != '#' && d + 1 < dis[nx][ny]) {
                    dis[nx][ny] = d + 1;
                    q.addLast(new int[]{nx, ny});
                }
            }
        }

        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
public:
    int minMoves(vector<string>& matrix) {
        int m = matrix.size(), n = matrix[0].size();
        if (matrix[m - 1][n - 1] == '#') {
            return -1;
        }

        vector<pair<int, int>> pos[26]{};
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                char c = matrix[i][j];
                if (isupper(c)) {
                    pos[c - 'A'].emplace_back(i, j);
                }
            }
        }

        vector dis(m, vector<int>(n, INT_MAX));
        dis[0][0] = 0;
        deque<pair<int, int>> q{{0, 0}};

        while (!q.empty()) {
            auto [x, y] = q.front();
            q.pop_front();
            int d = dis[x][y];
            
            if (x == m - 1 && y == n - 1) {
                return d;
            }

            char c = matrix[x][y];
            if (c != '.') {
                // 使用所有传送门
                for (auto& [px, py] : pos[c - 'A']) {
                    if (d < dis[px][py]) {
                        dis[px][py] = d;
                        q.emplace_front(px, py);
                    }
                }
                pos[c - 'A'].clear(); // 避免重复使用传送门
            }

            // 下面代码和普通 BFS 是一样的
            for (auto& [dx, dy] : DIRS) {
                int nx = x + dx, ny = y + dy;
                if (0 <= nx && nx < m && 0 <= ny && ny < n && matrix[nx][ny] != '#' && d + 1 < dis[nx][ny]) {
                    dis[nx][ny] = d + 1;
                    q.emplace_back(nx, ny);
                }
            }
        }

        return -1;
    }
};
```

```go [sol-Go]
func minMoves(matrix []string) int {
	m, n := len(matrix), len(matrix[0])
	if matrix[m-1][n-1] == '#' {
		return -1
	}

	type pair struct{ x, y int }
	pos := ['Z' + 1][]pair{}
	for i, row := range matrix {
		for j, c := range row {
			if unicode.IsUpper(c) {
				pos[c] = append(pos[c], pair{i, j})
			}
		}
	}

	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	// 两个 slice 头对头，模拟 deque
	q0 := []pair{{}}
	q1 := []pair{}

	for len(q0) > 0 || len(q1) > 0 {
		// 弹出队首
		var p pair
		if len(q0) > 0 {
			p, q0 = q0[len(q0)-1], q0[:len(q0)-1]
		} else {
			p, q1 = q1[0], q1[1:]
		}
		d := dis[p.x][p.y]

		if p.x == m-1 && p.y == n-1 {
			return d
		}

		if c := matrix[p.x][p.y]; c != '.' {
			// 使用所有传送门
			for _, q := range pos[c] {
				x, y := q.x, q.y
				if d < dis[x][y] {
					dis[x][y] = d
					q0 = append(q0, pair{x, y}) // 加到队首
				}
			}
			pos[c] = nil // 避免重复使用传送门
		}

		// 下面代码和普通 BFS 是一样的
		for _, dir := range dirs {
			x, y := p.x+dir.x, p.y+dir.y
			if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] != '#' && d+1 < dis[x][y] {
				dis[x][y] = d + 1
				q1 = append(q1, pair{x, y}) // 加到队尾
			}
		}
	}

	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{matrix}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

更多相似题目，见下面网格图题单的「**三、网格图 0-1 BFS**」。

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
