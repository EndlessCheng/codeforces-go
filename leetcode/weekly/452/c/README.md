本质是计算最短路，但需要一些额外的信息。

- 基本信息是当前位置，即行列下标 $(x,y)$。
- 需要额外知道当前能量值 $e$。
- 需要额外知道当前已收集的垃圾有哪些，即垃圾的编号集合 $\textit{mask}$。

从当前状态 $(x,y,e,\textit{mask})$ 移动到四方向相邻格子，新的状态为：

- 新的位置 $(x',y')$。
- 新的能量值：如果新的位置是 $\texttt{R}$，那么新的能量值为 $\textit{energy}$，否则为 $e-1$。
- 新的已收集垃圾编号集合：如果新的位置是 $\texttt{L}$，往 $\textit{mask}$ 中添加这个垃圾的编号。**注**：BFS 之前，给每个垃圾分配一个从 $0$ 开始的编号。

起点：$(\textit{sx},\textit{sy},\textit{energy},\varnothing)$，其中 $(\textit{sx},\textit{sy})$ 是学生的起始位置。

终点：$(x,y,e,U)$，所有垃圾清理完毕。其中 $U$ 是所有垃圾编号的集合。

关于 BFS 的原理和双列表写法，见[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)。

代码实现时，集合可以用二进制表示，集合相关运算可以用位运算代替，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

[本题视频讲解](https://www.bilibili.com/video/BV1Dz76zfEdi/?t=15m01s)，欢迎点赞关注~

## 答疑

**问**：能不能写类似旅行商问题（TSP）的状压 DP？

**答**：恐怕不行。考虑这样一种「接力」情况：$\texttt{L}\to \texttt{R}\to\texttt{R}\to\cdots \to\texttt{R} \to\texttt{L}$。有什么好的策略能快速判断在两个 $\texttt{L}$ 之间用哪些 $\texttt{R}$ 是最优的？

## 优化前

```py [sol-Python3]
class Solution:
    def minMoves(self, classroom: List[str], energy: int) -> int:
        m, n = len(classroom), len(classroom[0])
        idx = [[0] * n for _ in range(m)]
        cnt_l = sx = sy = 0
        for i, row in enumerate(classroom):
            for j, b in enumerate(row):
                if b == 'L':
                    idx[i][j] = 1 << cnt_l  # 给垃圾分配编号（提前计算左移）
                    cnt_l += 1
                elif b == 'S':
                    sx, sy = i, j
        if cnt_l == 0:
            return 0

        DIRS = (-1, 0), (1, 0), (0, -1), (0, 1)
        vis = [[[[False] * (1 << cnt_l) for _ in range(energy + 1)] for _ in range(n)] for _ in range(m)]
        vis[sx][sy][energy][0] = True
        q = [(sx, sy, energy, 0)]

        full_mask = (1 << cnt_l) - 1
        ans = 0
        while q:
            tmp = q
            q = []
            for x, y, e, mask in tmp:
                if mask == full_mask:  # 所有垃圾收集完毕
                    return ans
                if e == 0:  # 没能量了
                    continue
                for dx, dy in DIRS:
                    nx, ny = x + dx, y + dy
                    if 0 <= nx < m and 0 <= ny < n and classroom[nx][ny] != 'X':
                        new_e = energy if classroom[nx][ny] == 'R' else e - 1
                        new_mask = mask | idx[nx][ny]  # 添加垃圾（没有垃圾时 mask 不变）
                        if not vis[nx][ny][new_e][new_mask]:
                            vis[nx][ny][new_e][new_mask] = True
                            q.append((nx, ny, new_e, new_mask))
            ans += 1
        return -1
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}};

    private record Node(int x, int y, int e, int mask) {
    }

    public int minMoves(String[] classroom, int energy) {
        int m = classroom.length;
        int n = classroom[0].length();
        int[][] idx = new int[m][n];
        int cntL = 0, sx = 0, sy = 0;
        for (int i = 0; i < m; i++) {
            String row = classroom[i];
            for (int j = 0; j < n; j++) {
                char b = row.charAt(j);
                if (b == 'L') {
                    idx[i][j] = 1 << cntL++; // 给垃圾分配编号（提前计算左移）
                } else if (b == 'S') {
                    sx = i;
                    sy = j;
                }
            }
        }

        int u = 1 << cntL;
        boolean[][][][] vis = new boolean[m][n][energy + 1][u];
        vis[sx][sy][energy][0] = true;

        List<Node> q = new ArrayList<>();
        q.add(new Node(sx, sy, energy, 0));
        for (int ans = 0; !q.isEmpty(); ans++) {
            List<Node> tmp = q;
            q = new ArrayList<>();
            for (Node p : tmp) {
                if (p.mask == u - 1) { // 所有垃圾收集完毕
                    return ans;
                }
                if (p.e == 0) { // 走不动了
                    continue;
                }
                for (int[] d : DIRS) {
                    int x = p.x + d[0], y = p.y + d[1];
                    if (x >= 0 && x < m && y >= 0 && y < n && classroom[x].charAt(y) != 'X') {
                        int newE = classroom[x].charAt(y) == 'R' ? energy : p.e - 1;
                        int newMask = p.mask | idx[x][y]; // 添加垃圾（没有垃圾时 mask 不变）
                        if (!vis[x][y][newE][newMask]) {
                            vis[x][y][newE][newMask] = true;
                            q.add(new Node(x, y, newE, newMask));
                        }
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
    static constexpr int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

public:
    int minMoves(vector<string>& classroom, int energy) {
        int m = classroom.size(), n = classroom[0].size();
        vector idx(m, vector<int>(n));
        int cnt_l = 0, sx = 0, sy = 0;
        for (int i = 0; i < m; i++) {
            auto& row = classroom[i];
            for (int j = 0; j < n; j++) {
                char b = row[j];
                if (b == 'L') {
                    idx[i][j] = 1 << cnt_l++; // 给垃圾分配编号（提前计算左移）
                } else if (b == 'S') {
                    sx = i;
                    sy = j;
                }
            }
        }

        int u = 1 << cnt_l;
        vector vis(m, vector(n, vector(energy + 1, vector<int8_t>(u))));
        vis[sx][sy][energy][0] = true;
        struct Node { int x, y, e, mask; };
        vector<Node> q = {{sx, sy, energy, 0}};

        for (int ans = 0; !q.empty(); ans++) {
            auto tmp = q;
            q.clear();
            for (auto& [x, y, e, mask] : tmp) {
                if (mask == u - 1) { // 所有垃圾收集完毕
                    return ans;
                }
                if (e == 0) { // 走不动了
                    continue;
                }
                for (auto& [dx, dy] : dirs) {
                    int nx = x + dx, ny = y + dy;
                    if (0 <= nx && nx < m && 0 <= ny && ny < n && classroom[nx][ny] != 'X') {
                        int new_e = classroom[nx][ny] == 'R' ? energy : e - 1;
                        int new_mask = mask | idx[nx][ny]; // 添加垃圾（没有垃圾时 mask 不变）
                        if (!vis[nx][ny][new_e][new_mask]) {
                            vis[nx][ny][new_e][new_mask] = true;
                            q.emplace_back(nx, ny, new_e, new_mask);
                        }
                    }
                }
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minMoves(classroom []string, energy int) (ans int) {
	m, n := len(classroom), len(classroom[0])
	idx := make([][]int, m)
	for i := range idx {
		idx[i] = make([]int, n)
	}
	var cntL, sx, sy int
	for i, row := range classroom {
		for j, b := range row {
			if b == 'L' {
				idx[i][j] = 1 << cntL // 给垃圾分配编号（提前计算左移）
				cntL++
			} else if b == 'S' {
				sx, sy = i, j
			}
		}
	}

	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	u := 1 << cntL
	vis := make([][][][]bool, m)
	for i := range vis {
		vis[i] = make([][][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([][]bool, energy+1)
			for k := range vis[i][j] {
				vis[i][j][k] = make([]bool, u)
			}
		}
	}

	vis[sx][sy][energy][0] = true
	type tuple struct{ x, y, e, mask int }
	q := []tuple{{sx, sy, energy, 0}}

	for ; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.mask == u-1 { // 所有垃圾收集完毕
				return
			}
			if p.e == 0 { // 走不动了
				continue
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && classroom[x][y] != 'X' {
					newE := p.e - 1
					if classroom[x][y] == 'R' {
						newE = energy // 充满能量
					}
					newMask := p.mask | idx[x][y] // 添加垃圾（没有垃圾时 mask 不变）
					if !vis[x][y][newE][newMask] {
						vis[x][y][newE][newMask] = true
						q = append(q, tuple{x, y, newE, newMask})
					}
				}
			}
		}
	}
	return -1
}
```

## 优化

把 $\textit{vis}$ 改名为 $\textit{maxEnergy}$，其中的能量 $e$ 这一维度去掉，改为数组保存的值。

只有当状态 $(x,y,e,\textit{mask})$ 中的 $e > \textit{maxEnergy}[x][y][\textit{mask}]$，才入队，并更新 $\textit{maxEnergy}[x][y][\textit{mask}]=e$。

这样做的好处是，不会让相同 $(x,y,\textit{mask})$ 下的更小的能量入队，从而避免在两个相邻位置之间反复横跳，**避免无意义地消耗能量**。

```py [sol-Python3]
class Solution:
    def minMoves(self, classroom: List[str], energy: int) -> int:
        m, n = len(classroom), len(classroom[0])
        idx = [[0] * n for _ in range(m)]
        cnt_l = sx = sy = 0
        for i, row in enumerate(classroom):
            for j, b in enumerate(row):
                if b == 'L':
                    idx[i][j] = 1 << cnt_l
                    cnt_l += 1
                elif b == 'S':
                    sx, sy = i, j
        if cnt_l == 0:
            return 0

        DIRS = (-1, 0), (1, 0), (0, -1), (0, 1)
        max_energy = [[[-1] * (1 << cnt_l) for _ in range(n)] for _ in range(m)]
        max_energy[sx][sy][0] = energy
        q = [(sx, sy, energy, 0)]

        full_mask = (1 << cnt_l) - 1
        ans = 0
        while q:
            tmp = q
            q = []
            for x, y, e, mask in tmp:
                if mask == full_mask:
                    return ans
                if e == 0:
                    continue
                for dx, dy in DIRS:
                    nx, ny = x + dx, y + dy
                    if 0 <= nx < m and 0 <= ny < n and classroom[nx][ny] != 'X':
                        new_e = energy if classroom[nx][ny] == 'R' else e - 1
                        new_mask = mask | idx[nx][ny]
                        if new_e > max_energy[nx][ny][new_mask]:
                            max_energy[nx][ny][new_mask] = new_e
                            q.append((nx, ny, new_e, new_mask))
            ans += 1
        return -1
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}};

    private record Node(int x, int y, byte e, int mask) {
    }

    public int minMoves(String[] classroom, int energy) {
        int m = classroom.length;
        int n = classroom[0].length();
        char[][] grid = new char[m][]; // 把 String[] 转成 char[][]，读取效率更高
        int[][] idx = new int[m][n];
        int cntL = 0, sx = 0, sy = 0;
        for (int i = 0; i < m; i++) {
            grid[i] = classroom[i].toCharArray();
            for (int j = 0; j < n; j++) {
                char b = grid[i][j];
                if (b == 'L') {
                    idx[i][j] = 1 << cntL++;
                } else if (b == 'S') {
                    sx = i;
                    sy = j;
                }
            }
        }

        int u = 1 << cntL;
        byte[][][] maxEnergy = new byte[m][n][u]; // byte 空间小
        for (byte[][] mat : maxEnergy) {
            for (byte[] row : mat) {
                Arrays.fill(row, (byte) -1);
            }
        }
        maxEnergy[sx][sy][0] = (byte) energy;
        List<Node> q = new ArrayList<>();
        q.add(new Node(sx, sy, (byte) energy, 0));

        for (int ans = 0; !q.isEmpty(); ans++) {
            List<Node> tmp = q;
            q = new ArrayList<>();
            for (Node p : tmp) {
                if (p.mask == u - 1) {
                    return ans;
                }
                if (p.e == 0) {
                    continue;
                }
                for (int[] d : DIRS) {
                    int x = p.x + d[0], y = p.y + d[1];
                    if (x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != 'X') {
                        byte newE = (byte) (grid[x][y] == 'R' ? energy : p.e - 1);
                        int newMask = p.mask | idx[x][y];
                        if (newE > maxEnergy[x][y][newMask]) {
                            maxEnergy[x][y][newMask] = newE;
                            q.add(new Node(x, y, newE, newMask));
                        }
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
    static constexpr int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

public:
    int minMoves(vector<string>& classroom, int energy) {
        int m = classroom.size(), n = classroom[0].size();
        vector idx(m, vector<int>(n));
        int cnt_l = 0, sx = 0, sy = 0;
        for (int i = 0; i < m; i++) {
            auto& row = classroom[i];
            for (int j = 0; j < n; j++) {
                char b = row[j];
                if (b == 'L') {
                    idx[i][j] = 1 << cnt_l++;
                } else if (b == 'S') {
                    sx = i;
                    sy = j;
                }
            }
        }

        int u = 1 << cnt_l;
        vector max_energy(m, vector(n, vector<int8_t>(u, -1)));
        max_energy[sx][sy][0] = energy;
        struct Node { int x, y, e, mask; };
        vector<Node> q = {{sx, sy, energy, 0}};

        for (int ans = 0; !q.empty(); ans++) {
            auto tmp = q;
            q.clear();
            for (auto& [x, y, e, mask] : tmp) {
                if (mask == u - 1) {
                    return ans;
                }
                if (e == 0) {
                    continue;
                }
                for (auto& [dx, dy] : dirs) {
                    int nx = x + dx, ny = y + dy;
                    if (0 <= nx && nx < m && 0 <= ny && ny < n && classroom[nx][ny] != 'X') {
                        int new_e = classroom[nx][ny] == 'R' ? energy : e - 1;
                        int new_mask = mask | idx[nx][ny];
                        if (new_e > max_energy[nx][ny][new_mask]) {
                            max_energy[nx][ny][new_mask] = new_e;
                            q.emplace_back(nx, ny, new_e, new_mask);
                        }
                    }
                }
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minMoves(classroom []string, energy int) (ans int) {
	m, n := len(classroom), len(classroom[0])
	idx := make([][]int, m)
	for i := range idx {
		idx[i] = make([]int, n)
	}
	var cntL, sx, sy int
	for i, row := range classroom {
		for j, b := range row {
			if b == 'L' {
				idx[i][j] = 1 << cntL
				cntL++
			} else if b == 'S' {
				sx, sy = i, j
			}
		}
	}

	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	u := 1 << cntL
	maxEnergy := make([][][]int8, m)
	for i := range maxEnergy {
		maxEnergy[i] = make([][]int8, n)
		for j := range maxEnergy[i] {
			maxEnergy[i][j] = make([]int8, u)
			for k := range maxEnergy[i][j] {
				maxEnergy[i][j][k] = -1
			}
		}
	}

	maxEnergy[sx][sy][0] = int8(energy)
	type tuple struct{ x, y, e, mask int }
	q := []tuple{{sx, sy, energy, 0}}

	for ; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.mask == u-1 {
				return
			}
			if p.e == 0 {
				continue
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && classroom[x][y] != 'X' {
					newE := p.e - 1
					if classroom[x][y] == 'R' {
						newE = energy
					}
					newMask := p.mask | idx[x][y]
					if int8(newE) > maxEnergy[x][y][newMask] {
						maxEnergy[x][y][newMask] = int8(newE)
						q = append(q, tuple{x, y, newE, newMask})
					}
				}
			}
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\cdot n\cdot \textit{energy}\cdot 2^L)$，其中 $m$ 和 $n$ 分别为 $\textit{classroom}$ 的行数和列数，$L\le 10$ 为垃圾个数。每个状态至多访问一次。
- 空间复杂度：$\mathcal{O}(m\cdot n\cdot \textit{energy}\cdot 2^L)$。这是双列表（队列）需要的空间。

## 相似题目

- [864. 获取所有钥匙的最短路径](https://leetcode.cn/problems/shortest-path-to-get-all-keys/)
- [LCP 13. 寻宝](https://leetcode.cn/problems/xun-bao/)

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
