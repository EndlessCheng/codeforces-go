## 引入

考虑以下两种情况：

- Alice 吃掉第一个兵，Bob 吃掉第二个兵，Alice 吃掉第三个兵，现在轮到 Bob 操作。
- Alice 吃掉第二个兵，Bob 吃掉第一个兵，Alice 吃掉第三个兵，现在轮到 Bob 操作。

这两种情况，Bob 面临的局面是完全一样的：前三个兵都被吃掉，马现在在第三个兵的位置。

有重复的子问题，就可以用记忆化搜索/递推解决。

## 状态和状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前马在第 $i$ 个兵的位置。特别地，把 $(\textit{kx}, \textit{ky})$ 当作第 $n+1$ 个兵的位置。
- $\textit{mask}$：剩余没有被吃掉的兵的集合。

因此，定义状态为 $\textit{dfs}(i,\textit{mask})$，表示当前马在第 $i$ 个兵的位置，且剩余没有被吃掉的兵的集合为 $\textit{mask}$ 的情况下，继续游戏，两名玩家的总移动次数的最大值。

> 注意题目要计算的是两名玩家的总移动次数，不是 Alice 一个人的总移动次数。

接下来，思考如何从一个状态转移到另一个状态。

设从 $(x,y)$ 移动到第 $i$ 个兵的最小步数为 $\textit{dis}[i][x][y]$，这可以用网格图 BFS 算出来：反向思考，计算从第 $i$ 个兵的位置出发，通过「马走日」移动到 $(x,y)$ 的最小步数。

设当前位置为 $(x,y) = \textit{positions}[i]$，考虑枚举吃掉第 $j$ 个兵：

- 如果第 $j$ 个兵在集合 $\textit{mask}$ 中，把马移动 $\textit{dis}[j][x][y]$ 步，吃掉第 $j$ 个兵。现在问题变成当前马在第 $j$ 个兵的位置，且剩余没有被吃掉的兵的集合为 $\textit{mask}\setminus \{j\}$ 的情况下，继续游戏，两名玩家的总移动次数的最大值，即 $\textit{dfs}(j,\textit{mask}\setminus \{j\})$。

如果当前是 Alice 操作，则有

$$
\textit{dfs}(i,\textit{mask}) = \max_{j\in \textit{mask}}  \textit{dfs}(j,\textit{mask}\setminus \{j\}) + \textit{dis}[j][x][y]
$$

如果当前是 Bob 操作，则有

$$
\textit{dfs}(i,\textit{mask}) = \min_{j\in \textit{mask}}  \textit{dfs}(j,\textit{mask}\setminus \{j\}) + \textit{dis}[j][x][y]
$$

如何判断当前是谁在操作？

可以添加一个状态 $\textit{curPlayer}$，也可以用已被吃掉的兵的集合 $\complement_U \textit{mask}$ 的大小的奇偶性来判断，其中全集 $\textit{U}={0,1,2,\cdots,n-1}$。

- 如果吃掉了 $0,2,4,\cdots$ 个兵，那么当前轮到 Alice 操作。
- 如果吃掉了 $1,3,5,\cdots$ 个兵，那么当前轮到 Bob 操作。

> 注：也可以写两个递归函数，直接从代码层面区分开谁在操作。

递归边界：$\textit{dfs}(i,\varnothing) = 0$。所有兵都被吃掉，游戏结束。

递归入口：$\textit{dfs}(n,U)$。即答案。

代码中用到了一些位运算技巧，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

具体请看 [视频讲解](https://www.bilibili.com/video/BV1z5pieUEkQ/) 第四题，欢迎点赞关注~

## 记忆化搜索

```py [sol-Python3]
DIRS = ((2, 1), (1, 2), (-1, 2), (-2, 1), (-2, -1), (-1, -2), (1, -2), (2, -1))

class Solution:
    def maxMoves(self, kx: int, ky: int, positions: List[List[int]]) -> int:
        n = len(positions)
        # 计算马到兵的步数，等价于计算兵到其余格子的步数
        dis = [[[-1] * 50 for _ in range(50)] for _ in range(n)]
        for d, (px, py) in zip(dis, positions):
            d[px][py] = 0
            q = [(px, py)]
            step = 1
            while q:
                tmp = q
                q = []
                for p in tmp:
                    for dx, dy in DIRS:
                        x, y = p[0] + dx, p[1] + dy
                        if 0 <= x < 50 and 0 <= y < 50 and d[x][y] < 0:
                            d[x][y] = step
                            q.append((x, y))
                step += 1

        positions.append((kx, ky))
        u = (1 << n) - 1
        @cache
        def dfs(i: int, mask: int) -> int:
            if mask == 0:
                return 0
            odd = (u ^ mask).bit_count() % 2
            res = inf if odd else 0
            op = min if odd else max
            x, y = positions[i]
            for j, d in enumerate(dis):
                if mask >> j & 1:
                    res = op(res, dfs(j, mask ^ (1 << j)) + d[x][y])
            return res
        return dfs(n, u)
```

```py [sol-Python3 写法二]
DIRS = ((2, 1), (1, 2), (-1, 2), (-2, 1), (-2, -1), (-1, -2), (1, -2), (2, -1))

class Solution:
    def maxMoves(self, kx: int, ky: int, positions: List[List[int]]) -> int:
        n = len(positions)
        # 计算马到兵的步数，等价于计算兵到其余格子的步数
        dis = [[[-1] * 50 for _ in range(50)] for _ in range(n)]
        for d, (px, py) in zip(dis, positions):
            d[px][py] = 0
            q = [(px, py)]
            step = 1
            while q:
                tmp = q
                q = []
                for p in tmp:
                    for dx, dy in DIRS:
                        x, y = p[0] + dx, p[1] + dy
                        if 0 <= x < 50 and 0 <= y < 50 and d[x][y] < 0:
                            d[x][y] = step
                            q.append((x, y))
                step += 1

        positions.append((kx, ky))

        @cache
        def alice(i: int, mask: int) -> int:
            if mask == 0:
                return 0
            res = 0
            x, y = positions[i]
            for j, d in enumerate(dis):
                if mask >> j & 1:
                    res = max(res, bob(j, mask ^ (1 << j)) + d[x][y])
            return res

        @cache
        def bob(i: int, mask: int) -> int:
            if mask == 0:
                return 0
            res = inf
            x, y = positions[i]
            for j, d in enumerate(dis):
                if mask >> j & 1:
                    res = min(res, alice(j, mask ^ (1 << j)) + d[x][y])
            return res

        return alice(n, (1 << n) - 1)
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}};

    public int maxMoves(int kx, int ky, int[][] positions) {
        int n = positions.length;
        // 计算马到兵的步数，等价于计算兵到其余格子的步数
        int[][][] dis = new int[n][50][50];
        for (int i = 0; i < n; i++) {
            int[][] d = dis[i];
            for (int j = 0; j < 50; j++) {
                Arrays.fill(d[j], -1);
            }
            int px = positions[i][0];
            int py = positions[i][1];
            d[px][py] = 0;
            List<int[]> q = List.of(new int[]{px, py});
            for (int step = 1; !q.isEmpty(); step++) {
                List<int[]> tmp = q;
                q = new ArrayList<>();
                for (int[] p : tmp) {
                    for (int[] dir : DIRS) {
                        int x = p[0] + dir[0];
                        int y = p[1] + dir[1];
                        if (0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0) {
                            d[x][y] = step;
                            q.add(new int[]{x, y});
                        }
                    }
                }
            }
        }

        int[][] memo = new int[n + 1][1 << n];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(n, (1 << n) - 1, kx, ky, positions, dis, memo);
    }

    private int dfs(int i, int mask, int kx, int ky, int[][] positions, int[][][] dis, int[][] memo) {
        if (mask == 0) {
            return 0;
        }
        if (memo[i][mask] != -1) { // 之前计算过
            return memo[i][mask];
        }
        int n = positions.length;
        int x = i < n ? positions[i][0] : kx;
        int y = i < n ? positions[i][1] : ky;

        int res = 0;
        int u = (1 << n) - 1;
        if (Integer.bitCount(u ^ mask) % 2 == 0) { // Alice 操作
            for (int j = 0; j < n; j++) {
                if ((mask >> j & 1) > 0) {
                    res = Math.max(res, dfs(j, mask ^ (1 << j), kx, ky, positions, dis, memo) + dis[j][x][y]);
                }
            }
        } else { // Bob 操作
            res = Integer.MAX_VALUE;
            for (int j = 0; j < n; j++) {
                if ((mask >> j & 1) > 0) {
                    res = Math.min(res, dfs(j, mask ^ (1 << j), kx, ky, positions, dis, memo) + dis[j][x][y]);
                }
            }
        }
        return memo[i][mask] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int dirs[8][2] = {{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}};
public:
    int maxMoves(int kx, int ky, vector<vector<int>>& positions) {
        int n = positions.size();
        int dis[n][50][50];
        memset(dis, -1, sizeof(dis));
        // 计算马到兵的步数，等价于计算兵到其余格子的步数
        for (int i = 0; i < n; i++) {
            int px = positions[i][0], py = positions[i][1];
            dis[i][px][py] = 0;
            vector<pair<int, int>> q = {{px, py}};
            for (int step = 1; !q.empty(); step++) {
                vector<pair<int, int>> tmp;
                for (auto& [qx, qy] : q) {
                    for (auto& [dx, dy] : dirs) {
                        int x = qx + dx, y = qy + dy;
                        if (0 <= x && x < 50 && 0 <= y && y < 50 && dis[i][x][y] < 0) {
                            dis[i][x][y] = step;
                            tmp.emplace_back(x, y);
                        }
                    }
                }
                q = move(tmp);
            }
        }

        positions.push_back({kx, ky});
        vector<vector<int>> memo(n + 1, vector<int>(1 << n, -1)); // -1 表示没有计算过
        int u = (1 << n) - 1;
        auto dfs = [&](auto&& dfs, int i, int mask) -> int {
            if (mask == 0) {
                return 0;
            }
            int& res = memo[i][mask]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            int x = positions[i][0], y = positions[i][1];
            if (__builtin_parity(u ^ mask) == 0) { // Alice 操作
                for (int j = 0; j < n; j++) {
                    if (mask >> j & 1) {
                        res = max(res, dfs(dfs, j, mask ^ (1 << j)) + dis[j][x][y]);
                    }
                }
            } else { // Bob 操作
                res = INT_MAX;
                for (int j = 0; j < n; j++) {
                    if (mask >> j & 1) {
                        res = min(res, dfs(dfs, j, mask ^ (1 << j)) + dis[j][x][y]);
                    }
                }
            }
            return res;
        };
        return dfs(dfs, n, u);
    }
};
```

```go [sol-Go]
func maxMoves(kx, ky int, positions [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	n := len(positions)
	// 计算马到兵的步数，等价于计算兵到其余格子的步数
	dis := make([][50][50]int, n)
	for i, pos := range positions {
		d := &dis[i]
		for j := range d {
			for k := range d[j] {
				d[j][k] = -1
			}
		}
		px, py := pos[0], pos[1]
		d[px][py] = 0
		q := []pair{{px, py}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y
					if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
						d[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}

	positions = append(positions, []int{kx, ky})
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, 1<<n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	u := 1<<n - 1
	var dfs func(int, int) int
	dfs = func(i, mask int) int {
		if mask == 0 {
			return 0
		}
		p := &memo[i][mask]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		x, y := positions[i][0], positions[i][1]
		if bits.OnesCount(uint(u^mask))%2 == 0 { // Alice 操作
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				res = max(res, dfs(j, mask^1<<j)+dis[j][x][y])
			}
		} else { // Bob 操作
			res = math.MaxInt
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				res = min(res, dfs(j, mask^1<<j)+dis[j][x][y])
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n, u)
}
```

## 递推

注意要先把 $\textit{mask}$ 小的状态算出来，再算 $\textit{mask}$ 大的。所以外层循环 $\textit{mask}$，内层循环 $i$。

```py [sol-Python3]
DIRS = ((2, 1), (1, 2), (-1, 2), (-2, 1), (-2, -1), (-1, -2), (1, -2), (2, -1))

class Solution:
    def maxMoves(self, kx: int, ky: int, positions: List[List[int]]) -> int:
        n = len(positions)
        # 计算马到兵的步数，等价于计算兵到其余格子的步数
        dis = [[[-1] * 50 for _ in range(50)] for _ in range(n)]
        for d, (px, py) in zip(dis, positions):
            d[px][py] = 0
            q = [(px, py)]
            step = 1
            while q:
                tmp = q
                q = []
                for p in tmp:
                    for dx, dy in DIRS:
                        x, y = p[0] + dx, p[1] + dy
                        if 0 <= x < 50 and 0 <= y < 50 and d[x][y] < 0:
                            d[x][y] = step
                            q.append((x, y))
                step += 1

        positions.append((kx, ky))
        u = (1 << n) - 1
        f = [[0] * (n + 1) for _ in range(1 << n)]
        for mask in range(1, 1 << n):
            for i, (x, y) in enumerate(positions):
                odd = (u ^ mask).bit_count() % 2
                res = inf if odd else 0
                op = min if odd else max
                for j, d in enumerate(dis):
                    if mask >> j & 1:
                        res = op(res, f[mask ^ (1 << j)][j] + d[x][y])
                f[mask][i] = res
        return f[-1][n]
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}};

    public int maxMoves(int kx, int ky, int[][] positions) {
        int n = positions.length;
        // 计算马到兵的步数，等价于计算兵到其余格子的步数
        int[][][] dis = new int[n][50][50];
        for (int i = 0; i < n; i++) {
            int[][] d = dis[i];
            for (int j = 0; j < 50; j++) {
                Arrays.fill(d[j], -1);
            }
            int px = positions[i][0];
            int py = positions[i][1];
            d[px][py] = 0;
            List<int[]> q = List.of(new int[]{px, py});
            for (int step = 1; !q.isEmpty(); step++) {
                List<int[]> tmp = q;
                q = new ArrayList<>();
                for (int[] p : tmp) {
                    for (int[] dir : DIRS) {
                        int x = p[0] + dir[0];
                        int y = p[1] + dir[1];
                        if (0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0) {
                            d[x][y] = step;
                            q.add(new int[]{x, y});
                        }
                    }
                }
            }
        }

        int u = (1 << n) - 1;
        int[][] f = new int[1 << n][n + 1];
        for (int mask = 1; mask < (1 << n); mask++) {
            for (int i = 0; i <= n; i++) {
                int x = i < n ? positions[i][0] : kx;
                int y = i < n ? positions[i][1] : ky;
                if (Integer.bitCount(u ^ mask) % 2 == 0) { // Alice 操作
                    for (int j = 0; j < n; j++) {
                        if ((mask >> j & 1) > 0) {
                            f[mask][i] = Math.max(f[mask][i], f[mask ^ (1 << j)][j] + dis[j][x][y]);
                        }
                    }
                } else { // Bob 操作
                    f[mask][i] = Integer.MAX_VALUE;
                    for (int j = 0; j < n; j++) {
                        if ((mask >> j & 1) > 0) {
                            f[mask][i] = Math.min(f[mask][i], f[mask ^ (1 << j)][j] + dis[j][x][y]);
                        }
                    }
                }
            }
        }
        return f[u][n];
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int dirs[8][2] = {{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}};
public:
    int maxMoves(int kx, int ky, vector<vector<int>>& positions) {
        int n = positions.size();
        int dis[n][50][50];
        memset(dis, -1, sizeof(dis));
        // 计算马到兵的步数，等价于计算兵到其余格子的步数
        for (int i = 0; i < n; i++) {
            int px = positions[i][0], py = positions[i][1];
            dis[i][px][py] = 0;
            vector<pair<int, int>> q = {{px, py}};
            for (int step = 1; !q.empty(); step++) {
                vector<pair<int, int>> tmp;
                for (auto& [qx, qy] : q) {
                    for (auto& [dx, dy] : dirs) {
                        int x = qx + dx, y = qy + dy;
                        if (0 <= x && x < 50 && 0 <= y && y < 50 && dis[i][x][y] < 0) {
                            dis[i][x][y] = step;
                            tmp.emplace_back(x, y);
                        }
                    }
                }
                q = move(tmp);
            }
        }

        positions.push_back({kx, ky});
        int u = (1 << n) - 1;
        vector<vector<int>> f(1 << n, vector<int>(n + 1));
        for (int mask = 1; mask < (1 << n); mask++) {
            for (int i = 0; i <= n; i++) {
                int x = positions[i][0], y = positions[i][1];
                if (__builtin_parity(u ^ mask) == 0) { // Alice 操作
                    for (int j = 0; j < n; j++) {
                        if (mask >> j & 1) {
                            f[mask][i] = max(f[mask][i], f[mask ^ (1 << j)][j] + dis[j][x][y]);
                        }
                    }
                } else { // Bob 操作
                    f[mask][i] = INT_MAX;
                    for (int j = 0; j < n; j++) {
                        if (mask >> j & 1) {
                            f[mask][i] = min(f[mask][i], f[mask ^ (1 << j)][j] + dis[j][x][y]);
                        }
                    }
                }
            }
        }
        return f[u][n];
    }
};
```

```go [sol-Go]
func maxMoves(kx, ky int, positions [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	n := len(positions)
	// 计算马到兵的步数，等价于计算兵到其余格子的步数
	dis := make([][50][50]int, n)
	for i, pos := range positions {
		d := &dis[i]
		for j := range d {
			for k := range d[j] {
				d[j][k] = -1
			}
		}
		px, py := pos[0], pos[1]
		d[px][py] = 0
		q := []pair{{px, py}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y
					if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
						d[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}

	positions = append(positions, []int{kx, ky})
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for mask := 1; mask < 1<<n; mask++ {
		for i, p := range positions {
			x, y := p[0], p[1]
			odd := bits.OnesCount(uint(u^mask))%2 > 0
			if odd {
				f[mask][i] = math.MaxInt
			}
			op := func(a, b int) int {
				if odd {
					return min(a, b)
				}
				return max(a, b)
			}
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				f[mask][i] = op(f[mask][i], f[mask^1<<j][j]+dis[j][x][y])
			}
		}
	}
	return f[u][n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nL^2 + n^2 2^n)$，其中 $n$ 是 $\textit{positions}$ 的长度，$L=50$。
- 空间复杂度：$\mathcal{O}(nL^2 + n2^n)$。

## 思考题

1. 如果要计算的是 Alice 一个人的移动次数之和呢？
2. 如果棋盘无限大呢？能否用数学公式 $\mathcal{O}(1)$ 算出马的最小移动步数？见 [1197. 进击的骑士](https://leetcode.cn/problems/minimum-knight-moves/)（会员题）

更多相似题目，见下面动态规划题单中的「**状压 DP**」。

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
