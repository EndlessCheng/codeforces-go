想象有一个人在网格图上移动，按照题目要求，移动经过的值必须为 $1,2,0,2,0,\dots$

由于至多转弯一次，所以移动路径不会形成环，可以写一个记忆化搜索，模拟人在网格图上的移动。

定义 $\textit{dfs}(i,j,k,\textit{canTurn},\textit{target})$ 表示在如下约束下的最长移动步数。

- **上一步**的位置在 $(i,j)$。（定义成上一步，方便编程实现）
- 移动方向为 $\textit{DIRS}[k]$，其中 $\textit{DIRS}$ 是一个长为 $4$ 的方向数组。
- 是否可以右转，用布尔值 $\textit{canTurn}$ 表示。
- 当前位置的目标值必须等于 $\textit{target}$。

**转移**：

- 设 $(i',j')$ 是从 $(i,j)$ 向 $\textit{DIRS}[k]$ 方向移动一步后的位置。
- 直行：递归到 $\textit{dfs}(i',j',k,\textit{canTurn}, 2-\textit{target})$。这里用 $2-\textit{target}$ 来实现 $2$ 和 $0$ 的切换。也可以写成 $\textit{target}\oplus 2$。 
- 右转：如果 $\textit{canTurn} = \texttt{true}$，那么递归到 $\textit{dfs}(i',j',(k+1)\bmod 4, \texttt{false}, 2-\textit{target})$。其中 $(k+1)\bmod 4$ 表示环形数组 $\textit{DIRS}$ 的下一个元素的下标。如果 $k<3$，那么 $k$ 更新为 $k+1$；否则 $k$ 更新为 $0$。
- 返回二者加一后的最大值。其中加一表示走了一步。

**递归边界**：

- 出界，返回 $0$。
- 如果 $\textit{grid}[i'][j']\ne \textit{target}$，返回 $0$。

**递归入口**：

- 如果 $\textit{grid}[i][j]=1$，那么枚举 $k=0,1,2,3$，递归 $\textit{dfs}(i,j,k,\texttt{true},2)$。其中 $2$ 是因为下一步的值必须是 $2$。

计算所有 $\textit{dfs}(i,j,k,\texttt{true},2)+1$ 的最大值，即为答案。

⚠**注意**：$\textit{target}$ 无需记忆化，因为知道 $(i,j)$ 就间接知道 $\textit{target}$ 是多少，代码只是为了方便实现，额外传入了 $\textit{target}$。

关于记忆化搜索的原理，请看 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

[本题视频讲解](https://www.bilibili.com/video/BV1pmAGegEcw/?t=35m09s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def lenOfVDiagonal(self, grid: List[List[int]]) -> int:
        DIRS = (1, 1), (1, -1), (-1, -1), (-1, 1)
        m, n = len(grid), len(grid[0])

        # 上一步在 (i,j)，移动方向为 DIRS[k]，是否可以右转，当前位置目标值
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（一行代码实现记忆化）
        def dfs(i: int, j: int, k: int, can_turn: bool, target: int) -> int:
            i += DIRS[k][0]
            j += DIRS[k][1]
            if not (0 <= i < m and 0 <= j < n) or grid[i][j] != target:
                return 0
            res = dfs(i, j, k, can_turn, 2 - target) + 1  # 直行
            if can_turn:
                res = max(res, dfs(i, j, (k + 1) % 4, False, 2 - target) + 1)  # 右转
            return res

        ans = 0
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x == 1:
                    # 枚举起始方向
                    for k in range(4):
                        ans = max(ans, dfs(i, j, k, True, 2) + 1)
        return ans
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{1, 1}, {1, -1}, {-1, -1}, {-1, 1}};

    public int lenOfVDiagonal(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        // 数组维度太多影响效率，这里把 k 和 canTurn 压缩成一个 int
        int[][][] memo = new int[m][n][1 << 3];
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    for (int k = 0; k < 4; k++) { // 枚举起始方向
                        ans = Math.max(ans, dfs(i, j, k, 1, 2, grid, memo) + 1);
                    }
                }
            }
        }
        return ans;
    }

    // 上一步在 (i,j)，移动方向为 DIRS[k]，是否可以右转，当前位置目标值
    private int dfs(int i, int j, int k, int canTurn, int target, int[][] grid, int[][][] memo) {
        i += DIRS[k][0];
        j += DIRS[k][1];
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[i].length || grid[i][j] != target) {
            return 0;
        }
        int mask = k << 1 | canTurn;
        if (memo[i][j][mask] > 0) { // 之前计算过
            return memo[i][j][mask];
        }
        int res = dfs(i, j, k, canTurn, 2 - target, grid, memo) + 1; // 直行
        if (canTurn == 1) {
            res = Math.max(res, dfs(i, j, (k + 1) % 4, 0, 2 - target, grid, memo) + 1); // 右转
        }
        return memo[i][j][mask] = res; // 记忆化
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{1, 1}, {1, -1}, {-1, -1}, {-1, 1}};

public:
    int lenOfVDiagonal(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector<array<array<int, 2>, 4>>(n));

        // 上一步在 (i,j)，移动方向为 DIRS[k]，是否可以右转，当前位置目标值
        auto dfs = [&](this auto&& dfs, int i, int j, int k, bool can_turn, int target) -> int {
            i += DIRS[k][0];
            j += DIRS[k][1];
            if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target) {
                return 0;
            }
            int& res = memo[i][j][k][can_turn]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            res = dfs(i, j, k, can_turn, 2 - target) + 1; // 直行
            if (can_turn) {
                res = max(res, dfs(i, j, (k + 1) % 4, false, 2 - target) + 1); // 右转
            }
            return res;
        };

        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    for (int k = 0; k < 4; k++) { // 枚举起始方向
                        ans = max(ans, dfs(i, j, k, true, 2) + 1);
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
var DIRS = [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func lenOfVDiagonal(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][][4][2]int, m)
	for i := range memo {
		memo[i] = make([][4][2]int, n)
	}

	// 上一步在 (i,j)，移动方向为 DIRS[k]，是否可以右转，当前位置目标值
	var dfs func(int, int, int, int, int) int
	dfs = func(i, j, k, canTurn, target int) int {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return 0
		}
		p := &memo[i][j][k][canTurn]
		if *p > 0 { // 之前计算过
			return *p
		}
		res := dfs(i, j, k, canTurn, 2-target) + 1 // 直行
		if canTurn == 1 {
			res = max(res, dfs(i, j, (k+1)%4, 0, 2-target)+1) // 右转
		}
		*p = res // 记忆化
		return res
	}

	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				for k := range 4 { // 枚举起始方向
					ans = max(ans, dfs(i, j, k, 1, 2)+1)
				}
			}
		}
	}
	return
}
```

## 最优性剪枝

#### 优化一

在递归之前，如果发现即使走到底，能访问的格子数也不会比目前算出的 $\textit{ans}$ 更大，那么不递归。

![lc3459.jpg](https://pic.leetcode.cn/1739840631-DeGuFJ-lc3459.jpg)

看图，设当前在 $(i,j)$：

- 如果一开始往右下走，无论是否右转，**每一步都在往下走**，所以至多访问 $m-i$ 个格子。
- 如果一开始往左下走，无论是否右转，**每一步都在往左走**，所以至多访问 $j+1$ 个格子。
- 如果一开始往左上走，无论是否右转，**每一步都在往上走**，所以至多访问 $i+1$ 个格子。
- 如果一开始往右上走，无论是否右转，**每一步都在往右走**，所以至多访问 $n-j$ 个格子。

如果理论最大值 $\textit{mx}$ 大于 $\textit{ans}$，那么递归，否则不递归。

#### 优化二

同理，在递归中，如果要右转，可以先判断右转后继续走的理论最大值是否大于 $\textit{res}$，如果大于则递归，否则不递归。

以右下走为例，如果 $(i,j)$ 离网格图下边缘更近，那么至多访问 $m-i$ 个格子；如果 $(i,j)$ 离网格图右边缘更近，那么至多访问 $n-j$ 个格子。二者的最小值 $\min(m-i,n-j)$ 就是还能访问的格子数的理论最大值。

```py [sol-Python3]
class Solution:
    def lenOfVDiagonal(self, grid: List[List[int]]) -> int:
        DIRS = (1, 1), (1, -1), (-1, -1), (-1, 1)
        m, n = len(grid), len(grid[0])

        @cache
        def dfs(i: int, j: int, k: int, can_turn: bool, target: int) -> int:
            i += DIRS[k][0]
            j += DIRS[k][1]
            if not (0 <= i < m and 0 <= j < n) or grid[i][j] != target:
                return 0
            res = dfs(i, j, k, can_turn, 2 - target) + 1
            if can_turn:
                maxs = (m - i, j + 1, i + 1, n - j)  # 理论最大值（走到底）
                k = (k + 1) % 4
                # 优化二：如果理论最大值没有超过 res，那么不递归
                if min(maxs[k], maxs[k - 1]) > res:
                    res = max(res, dfs(i, j, k, False, 2 - target) + 1)
            return res

        ans = 0
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x != 1:
                    continue
                maxs = (m - i, j + 1, i + 1, n - j)  # 理论最大值（走到底）
                for k, mx in enumerate(maxs):  # 枚举起始方向
                    # 优化一：如果理论最大值没有超过 ans，那么不递归
                    if mx > ans:
                        ans = max(ans, dfs(i, j, k, True, 2) + 1)
        return ans
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{1, 1}, {1, -1}, {-1, -1}, {-1, 1}};

    public int lenOfVDiagonal(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        // 开太多维度影响效率，这里把 k 和 canTurn 压缩成一个 int
        int[][][] memo = new int[m][n][1 << 3];
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] != 1) {
                    continue;
                }
                int[] maxs = {m - i, j + 1, i + 1, n - j}; // 理论最大值（走到底）
                for (int k = 0; k < 4; k++) { // 枚举起始方向
                    // 优化一：如果理论最大值没有超过 ans，那么不递归
                    if (maxs[k] > ans) {
                        ans = Math.max(ans, dfs(i, j, k, 1, 2, grid, memo) + 1);
                    }
                }
            }
        }
        return ans;
    }

    private int dfs(int i, int j, int k, int canTurn, int target, int[][] grid, int[][][] memo) {
        i += DIRS[k][0];
        j += DIRS[k][1];
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[i].length || grid[i][j] != target) {
            return 0;
        }
        int mask = k << 1 | canTurn;
        if (memo[i][j][mask] > 0) {
            return memo[i][j][mask];
        }
        int res = dfs(i, j, k, canTurn, 2 - target, grid, memo) + 1;
        if (canTurn == 1) {
            int[] maxs = {grid.length - i, j + 1, i + 1, grid[i].length - j}; // 理论最大值（走到底）
            k = (k + 1) % 4;
            // 优化二：如果理论最大值没有超过 res，那么不递归
            if (Math.min(maxs[k], maxs[(k + 3) % 4]) > res) {
                res = Math.max(res, dfs(i, j, k, 0, 2 - target, grid, memo) + 1);
            }
        }
        return memo[i][j][mask] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{1, 1}, {1, -1}, {-1, -1}, {-1, 1}};

public:
    int lenOfVDiagonal(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector<array<array<int, 2>, 4>>(n));

        auto dfs = [&](this auto&& dfs, int i, int j, int k, bool can_turn, int target) -> int {
            i += DIRS[k][0];
            j += DIRS[k][1];
            if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target) {
                return 0;
            }
            int& res = memo[i][j][k][can_turn];
            if (res) {
                return res;
            }
            res = dfs(i, j, k, can_turn, 2 - target) + 1;
            if (can_turn) {
                int maxs[4] = {m - i, j + 1, i + 1, n - j}; // 理论最大值（走到底）
                k = (k + 1) % 4;
                // 优化二：如果理论最大值没有超过 res，那么不递归
                if (min(maxs[k], maxs[(k + 3) % 4]) > res) {
                    res = max(res, dfs(i, j, k, false, 2 - target) + 1);
                }
            }
            return res;
        };

        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] != 1) {
                    continue;
                }
                int maxs[4] = {m - i, j + 1, i + 1, n - j}; // 理论最大值（走到底）
                for (int k = 0; k < 4; k++) { // 枚举起始方向
                    // 优化一：如果理论最大值没有超过 ans，那么不递归
                    if (maxs[k] > ans) {
                        ans = max(ans, dfs(i, j, k, true, 2) + 1);
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
var DIRS = [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func lenOfVDiagonal(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][][4][2]int, m)
	for i := range memo {
		memo[i] = make([][4][2]int, n)
	}

	var dfs func(int, int, int, int, int) int
	dfs = func(i, j, k, canTurn, target int) int {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return 0
		}
		p := &memo[i][j][k][canTurn]
		if *p > 0 {
			return *p
		}
		res := dfs(i, j, k, canTurn, 2-target) + 1
		if canTurn == 1 {
			maxs := [4]int{m - i, j + 1, i + 1, n - j} // 理论最大值（走到底）
			k = (k + 1) % 4
			// 优化二：如果理论最大值没有超过 res，那么不递归
			if min(maxs[k], maxs[(k+3)%4]) > res {
				res = max(res, dfs(i, j, k, 0, 2-target)+1)
			}
		}
		*p = res
		return res
	}

	for i, row := range grid {
		for j, x := range row {
			if x != 1 {
				continue
			}
			maxs := [4]int{m - i, j + 1, i + 1, n - j} // 理论最大值（走到底）
			for k, mx := range maxs { // 枚举起始方向
				// 优化一：如果理论最大值没有超过 ans，那么不递归
				if mx > ans {
					ans = max(ans, dfs(i, j, k, 1, 2)+1)
				}
			}
		}
	}
	return
}
```

## 空间优化

从一个 $1$ 出发，在不拐弯的情况下，我们不可能与另一条从 $1$ 出发且没有拐弯的路径重合（否则路径上就有两个 $1$ 了）。

所以 $\textit{canTurn} = \texttt{true}$ 的状态不会重复访问到，也就无需记忆化了。换言之，类似数位 DP，只需在 $\textit{canTurn} = \texttt{false}$ 时记忆化。

> Python 用户可以忽略这个优化，仍然用方便的 `@cache` 装饰器。

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{1, 1}, {1, -1}, {-1, -1}, {-1, 1}};

    public int lenOfVDiagonal(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][] memo = new int[m][n][4];
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] != 1) {
                    continue;
                }
                int[] maxs = {m - i, j + 1, i + 1, n - j}; // 理论最大值（走到底）
                for (int k = 0; k < 4; k++) { // 枚举起始方向
                    // 优化一：如果理论最大值没有超过 ans，那么不递归
                    if (maxs[k] > ans) {
                        ans = Math.max(ans, dfs(i, j, k, true, 2, grid, memo) + 1);
                    }
                }
            }
        }
        return ans;
    }

    private int dfs(int i, int j, int k, boolean canTurn, int target, int[][] grid, int[][][] memo) {
        i += DIRS[k][0];
        j += DIRS[k][1];
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[i].length || grid[i][j] != target) {
            return 0;
        }
        // 只在 canTurn=false 时读取和写入 memo
        if (!canTurn && memo[i][j][k] > 0) {
            return memo[i][j][k];
        }
        int res = dfs(i, j, k, canTurn, 2 - target, grid, memo) + 1;
        if (!canTurn) {
            return memo[i][j][k] = res;
        }
        int[] maxs = {grid.length - i, j + 1, i + 1, grid[i].length - j}; // 理论最大值（走到底）
        k = (k + 1) % 4;
        // 优化二：如果理论最大值没有超过 res，那么不递归
        if (Math.min(maxs[k], maxs[(k + 3) % 4]) > res) {
            res = Math.max(res, dfs(i, j, k, false, 2 - target, grid, memo) + 1);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{1, 1}, {1, -1}, {-1, -1}, {-1, 1}};

public:
    int lenOfVDiagonal(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector memo(m, vector<array<int, 4>>(n));

        auto dfs = [&](this auto&& dfs, int i, int j, int k, bool can_turn, int target) -> int {
            i += DIRS[k][0];
            j += DIRS[k][1];
            if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target) {
                return 0;
            }
            // 只在 can_turn=false 时读取和写入 memo
            if (!can_turn && memo[i][j][k]) {
                return memo[i][j][k];
            }
            int res = dfs(i, j, k, can_turn, 2 - target) + 1;
            if (!can_turn) {
                return memo[i][j][k] = res;
            }
            int maxs[4] = {m - i, j + 1, i + 1, n - j}; // 理论最大值（走到底）
            k = (k + 1) % 4;
            // 优化二：如果理论最大值没有超过 res，那么不递归
            if (min(maxs[k], maxs[(k + 3) % 4]) > res) {
                res = max(res, dfs(i, j, k, false, 2 - target) + 1);
            }
            return res;
        };

        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] != 1) {
                    continue;
                }
                int maxs[4] = {m - i, j + 1, i + 1, n - j}; // 理论最大值（走到底）
                for (int k = 0; k < 4; k++) { // 枚举起始方向
                    // 优化一：如果理论最大值没有超过 ans，那么不递归
                    if (maxs[k] > ans) {
                        ans = max(ans, dfs(i, j, k, true, 2) + 1);
                    }
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
var DIRS = [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func lenOfVDiagonal(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][][4]int, m)
	for i := range memo {
		memo[i] = make([][4]int, n)
	}

	var dfs func(int, int, int, bool, int) int
	dfs = func(i, j, k int, canTurn bool, target int) int {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return 0
		}
		// 只在 canTurn=false 时读取和写入 memo
		if !canTurn && memo[i][j][k] > 0 {
			return memo[i][j][k]
		}
		res := dfs(i, j, k, canTurn, 2-target) + 1
		if !canTurn {
			memo[i][j][k] = res
			return res
		}
		maxs := [4]int{m - i, j + 1, i + 1, n - j} // 理论最大值（走到底）
		k = (k + 1) % 4
		// 优化二：如果理论最大值没有超过 res，那么不递归
		if min(maxs[k], maxs[(k+3)%4]) > res {
			res = max(res, dfs(i, j, k, false, 2-target)+1)
		}
		return res
	}

	for i, row := range grid {
		for j, x := range row {
			if x != 1 {
				continue
			}
			maxs := [4]int{m - i, j + 1, i + 1, n - j} // 理论最大值（走到底）
			for k, mx := range maxs { // 枚举起始方向
				// 优化一：如果理论最大值没有超过 ans，那么不递归
				if mx > ans {
					ans = max(ans, dfs(i, j, k, true, 2)+1)
				}
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。保存多少状态，就需要多少空间。

## 相似题目

- [329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)
- [2328. 网格图中递增路径的数目](https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/)

更多相似题目，见下面动态规划题单的「**二、网格图 DP**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
