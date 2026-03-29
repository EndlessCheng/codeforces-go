其实这题改成往上下左右四个方向移动，也能做。

写一个网格图 DFS，除了有参数 $(i,j)$ 表示当前位置外，还需要参数 $\textit{xor}$ 表示走过的格子的异或和。

每访问一个格子，就把格子的值异或到 $\textit{xor}$ 中。

移动到终点时，用 $\textit{xor}$ 更新答案的最小值。

为避免重复访问相同的状态，用 $\textit{vis}$ 数组（哈希集合）标记访问过的状态 $(i,j,\textit{xor})$。

**最优性剪枝**：如果答案已经最小（等于 $0$），那么不再搜索。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minCost(self, grid: List[List[int]]) -> int:
        vis = set()
        ans = inf

        def dfs(i: int, j: int, xor: int) -> None:
            nonlocal ans
            # 最优性剪枝：如果答案已经最小（等于 0），那么不再搜索
            if ans == 0 or i < 0 or j < 0 or (i, j, xor) in vis:
                return
            vis.add((i, j, xor))
            xor ^= grid[i][j]
            if i == 0 and j == 0:
                ans = min(ans, xor)
                return
            dfs(i - 1, j, xor)
            dfs(i, j - 1, xor)

        dfs(len(grid) - 1, len(grid[0]) - 1, 0)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans = Integer.MAX_VALUE;

    public int minCost(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;

        // 异或和不会超过所有元素的 OR
        int orAll = 0;
        for (int[] row : grid) {
            for (int x : row) {
                orAll |= x;
            }
        }

        boolean[][][] vis = new boolean[m][n][orAll + 1];
        dfs(m - 1, n - 1, 0, grid, vis);
        return ans;
    }

    private void dfs(int i, int j, int xor, int[][] grid, boolean[][][] vis) {
        // 最优性剪枝：如果答案已经最小（等于 0），那么不再搜索
        if (ans == 0 || i < 0 || j < 0 || vis[i][j][xor]) {
            return;
        }
        vis[i][j][xor] = true;
        xor ^= grid[i][j];
        if (i == 0 && j == 0) {
            ans = Math.min(ans, xor);
            return;
        }
        dfs(i - 1, j, xor, grid, vis);
        dfs(i, j - 1, xor, grid, vis);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        // 异或和不会超过所有元素的 OR
        int or_all = 0;
        for (auto& row : grid) {
            for (int x : row) {
                or_all |= x;
            }
        }

        vector vis(m, vector(n, vector<int8_t>(or_all + 1)));
        int ans = INT_MAX;

        auto dfs = [&](this auto&& dfs, int i, int j, int xor_val) -> void {
            // 最优性剪枝：如果答案已经最小（等于 0），那么不再搜索
            if (ans == 0 || i < 0 || j < 0 || vis[i][j][xor_val]) {
                return;
            }
            vis[i][j][xor_val] = true;
            xor_val ^= grid[i][j];
            if (i == 0 && j == 0) {
                ans = min(ans, xor_val);
                return;
            }
            dfs(i - 1, j, xor_val);
            dfs(i, j - 1, xor_val);
        };

        dfs(m - 1, n - 1, 0);
        return ans;
    }
};
```

```go [sol-Go]
func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 异或和不会超过所有元素的 OR
	orAll := 0
	for _, row := range grid {
		for _, x := range row {
			orAll |= x
		}
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, orAll+1)
		}
	}
	ans := math.MaxInt

	var dfs func(int, int, int)
	dfs = func(i, j, xor int) {
		// 最优性剪枝：如果答案已经最小（等于 0），那么不再搜索
		if ans == 0 || i < 0 || j < 0 || vis[i][j][xor] {
			return
		}
		vis[i][j][xor] = true
		xor ^= grid[i][j]
		if i == 0 && j == 0 {
			ans = min(ans, xor)
			return
		}
		dfs(i-1, j, xor)
		dfs(i, j-1, xor)
	}

	dfs(m-1, n-1, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数，$U=\max(\textit{grid})$。
- 空间复杂度：$\mathcal{O}(mnU)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
