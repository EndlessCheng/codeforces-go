先计算所有行变成回文最少需要翻转多少次。

也就是对于每一行 $\textit{row}$，计算这一行变成回文最少需要翻转多少次。

也就是累加 $\textit{row}[j]\ne \textit{row}[n-1-j]$ 的个数，其中 $0\le j \le \lfloor n/2 \rfloor$。

对于列，统计方式同理。

两种情况取最小值，即为答案。

```py [sol-Python3]
class Solution:
    def minFlips(self, grid: List[List[int]]) -> int:
        diff_row = 0
        for row in grid:
            for j in range(len(row) // 2):
                if row[j] != row[-1 - j]:
                    diff_row += 1

        diff_col = 0
        for col in zip(*grid):
            for i in range(len(grid) // 2):
                if col[i] != col[-1 - i]:
                    diff_col += 1

        return min(diff_row, diff_col)
```

```java [sol-Java]
class Solution {
    public int minFlips(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;

        int diffRow = 0;
        for (int[] row : grid) {
            for (int j = 0; j < n / 2; j++) {
                if (row[j] != row[n - 1 - j]) {
                    diffRow++;
                }
            }
        }

        int diffCol = 0;
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m / 2; i++) {
                if (grid[i][j] != grid[m - 1 - i][j]) {
                    diffCol++;
                }
            }
        }

        return Math.min(diffRow, diffCol);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minFlips(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();

        int diff_row = 0;
        for (auto& row : grid) {
            for (int j = 0; j < n / 2; j++) {
                diff_row += row[j] != row[n - 1 - j];
            }
        }

        int diff_col = 0;
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m / 2; i++) {
                diff_col += grid[i][j] != grid[m - 1 - i][j];
            }
        }

        return min(diff_row, diff_col);
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int minFlips(int** grid, int gridSize, int* gridColSize) {
    int m = gridSize, n = gridColSize[0];

    int diff_row = 0;
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n / 2; j++) {
            if (grid[i][j] != grid[i][n - 1 - j]) {
                diff_row++;
            }
        }
    }

    int diff_col = 0;
    for (int j = 0; j < n; j++) {
        for (int i = 0; i < m / 2; i++) {
            if (grid[i][j] != grid[m - 1 - i][j]) {
                diff_col++;
            }
        }
    }

    return MIN(diff_row, diff_col);
}
```

```go [sol-Go]
func minFlips(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    diffRow := 0
    for _, row := range grid {
        for j := 0; j < n/2; j++ {
            if row[j] != row[n-1-j] {
                diffRow++
            }
        }
    }

    diffCol := 0
    for j := 0; j < n; j++ {
        for i, row := range grid[:m/2] {
            if row[j] != grid[m-1-i][j] {
                diffCol++
            }
        }
    }

    return min(diffRow, diffCol)
}
```

```js [sol-JavaScript]
var minFlips = function(grid) {
    const m = grid.length, n = grid[0].length;

    let diffRow = 0;
    for (const row of grid) {
        for (let j = 0; j < n / 2; j++) {
            if (row[j] !== row[n - 1 - j]) {
                diffRow++;
            }
        }
    }

    let diffCol = 0;
    for (let j = 0; j < n; j++) {
        for (let i = 0; i < m / 2; i++) {
            if (grid[i][j] !== grid[m - 1 - i][j]) {
                diffCol++;
            }
        }
    }

    return Math.min(diffRow, diffCol);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_flips(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();

        let mut diff_row = 0;
        for row in &grid {
            for j in 0..n / 2 {
                if row[j] != row[n - 1 - j] {
                    diff_row += 1;
                }
            }
        }

        let mut diff_col = 0;
        for j in 0..n {
            for i in 0..m / 2 {
                if grid[i][j] != grid[m - 1 - i][j] {
                    diff_col += 1;
                }
            }
        }

        diff_row.min(diff_col)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略 `zip` 的开销。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
