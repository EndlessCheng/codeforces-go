想象有一个人在网格图上移动，按照题目要求，移动经过的值必须为 $1,2,0,2,0,\cdots$。

由于移动路径不会形成环（否则路径会有多个 $1$，或者说拐弯了不止一次），我们可以写一个递归，来模拟人在网格图上的移动。

定义 $\textit{dfs}(i,j,k,\textit{canTurn},\textit{target})$ 表示在如下约束下的最长移动步数。

- **上一步**在 $(i,j)$。
- 移动方向为 $\textit{DIRS}[k]$，其中 $\textit{DIRS}$ 是一个长为 $4$ 的方向数组。
- 是否可以右转，用布尔值 $\textit{canTurn}$ 表示。
- 当前位置的目标值必须等于 $\textit{target}$。

**递归边界**：

- 出界，返回 $0$。
- 如果 $\textit{grid}[i'][j']\ne \textit{target}$，返回 $0$。其中 $(i',j')$ 是从 $(i,j)$ 向 $\textit{DIRS}[k]$ 方向移动一步后的位置。

**递归入口**：

- 如果 $\textit{grid}[i][j]=1$，那么枚举 $k=0,1,2,3$，递归 $\textit{dfs}(i,j,k,\texttt{true},2)$。其中 $2$ 是因为下一步的值必须是 $2$。

计算所有 $\textit{dfs}(i,j,k,\texttt{true},2)+1$ 的最大值，即为答案。

⚠**注意**：$\textit{target}$ 无需记忆化，因为知道 $(i,j)$ 就间接知道 $\textit{target}$ 是多少，代码只是为了方便实现，额外传入了 $\textit{target}$。

关于记忆化搜索的原理，请看 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

[本题视频讲解](https://www.bilibili.com/video/BV1pmAGegEcw/?t=35m09s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def lenOfVDiagonal(self, grid: List[List[int]]) -> int:
        DIRS = (1, 1), (1, -1), (-1, -1), (-1, 1)
        m, n = len(grid), len(grid[0])

        # 上一步在 (i,j)，移动方向为 DIRS[k]，是否可以右转，当前位置目标值
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（一行代码实现记忆化）
        def dfs(i: int, j: int, k: int, can_turn: bool, target: int):
            i += DIRS[k][0]
            j += DIRS[k][1]
            if not (0 <= i < m and 0 <= j < n) or grid[i][j] != target:
                return 0
            res = dfs(i, j, k, can_turn, 2 - target)  # 直行
            if can_turn:
                res = max(res, dfs(i, j, (k + 1) % 4, False, 2 - target))  # 右转
            return res + 1  # 算上当前位置

        ans = 0
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x == 1:
                    # 枚举出发方向
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
        // 开太多维度影响效率，这里把 k 和 canTurn 压缩成一个 int
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
        int res = dfs(i, j, k, canTurn, 2 - target, grid, memo); // 直行
        if (canTurn == 1) {
            res = Math.max(res, dfs(i, j, (k + 1) % 4, 0, 2 - target, grid, memo)); // 右转
        }
        return memo[i][j][mask] = res + 1; // 算上当前位置
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
            int& res = memo[i][j][k][can_turn]; // 注意这里是引用
            if (res) { // 之前计算过
                return res;
            }
            res = dfs(i, j, k, can_turn, 2 - target); // 直行
            if (can_turn) {
                res = max(res, dfs(i, j, (k + 1) % 4, false, 2 - target)); // 右转
            }
            return ++res; // 算上当前位置
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

	var dfs func(int, int, int, int, int) int
	dfs = func(i, j, k, canTurn, target int) (res int) {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return
		}
		p := &memo[i][j][k][canTurn]
		if *p > 0 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		res = dfs(i, j, k, canTurn, 2-target) // 直行
		if canTurn == 1 {
			res = max(res, dfs(i, j, (k+1)%4, 0, 2-target)) // 右转
		}
		return res + 1 // 算上当前位置
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。保存多少状态，就需要多少空间。

## 相似题目

- [329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)
- [2328. 网格图中递增路径的数目](https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/)

更多相似题目，见下面动态规划题单中的「**二、网格图 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
