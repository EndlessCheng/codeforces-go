做法同 [200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/)，[我的题解](https://leetcode.cn/problems/number-of-islands/solutions/2965773/ba-fang-wen-guo-de-ge-zi-cha-shang-qi-zi-9gs0/)，在 DFS 岛屿的同时统计岛屿格子的元素和。

如果岛屿元素和是 $k$ 的倍数，那么把答案加一。

如何理解网格图 DFS，请看 [视频讲解](https://www.bilibili.com/video/BV1R5g8zDEGY/?t=4m46s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])

        def dfs(i: int, j: int) -> int:
            res = grid[i][j]
            grid[i][j] = 0  # 标记 (i,j) 访问过
            for x, y in (i, j + 1), (i, j - 1), (i + 1, j), (i - 1, j):
                if 0 <= x < m and 0 <= y < n and grid[x][y]:
                    res += dfs(x, y)
            return res

        ans = 0
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x and dfs(i, j) % k == 0:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{0, -1}, {0, 1}, {-1, 0}, {1, 0}};

    public int countIslands(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] > 0 && dfs(i, j, grid) % k == 0) {
                    ans++;
                }
            }
        }
        return ans;
    }

    private long dfs(int i, int j, int[][] grid) {
        long res = grid[i][j];
        grid[i][j] = 0; // 标记为访问过
        for (int[] d : DIRS) {
            int x = i + d[0], y = j + d[1];
            if (0 <= x && x < grid.length && 0 <= y && y < grid[x].length && grid[x][y] > 0) {
                res += dfs(x, y, grid);
            }
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int DIRS[4][2] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
public:
    int countIslands(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();

        auto dfs = [&](this auto&& dfs, int i, int j) -> long long {
            long long res = grid[i][j];
            grid[i][j] = 0; // 标记为访问过
            for (auto [dx, dy] : DIRS) {
                int x = i + dx, y = j + dy;
                if (0 <= x && x < m && 0 <= y && y < n && grid[x][y]) {
                    res += dfs(x, y);
                }
            }
            return res;
        };

        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] && dfs(i, j) % k == 0) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countIslands(grid [][]int, k int) (ans int) {
	dirs := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		res := grid[i][j]
		grid[i][j] = 0 // 标记 (i,j) 访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
				res += dfs(x, y)
			}
		}
		return res
	}

	for i, row := range grid {
		for j, x := range row {
			if x > 0 && dfs(i, j)%k == 0 {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。最坏情况下，对于之字形岛屿，递归深度为 $\mathcal{O}(mn)$，需要 $\mathcal{O}(mn)$ 的栈空间。

## 专题训练

见下面网格图题单的「**一、网格图 DFS**」。

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
