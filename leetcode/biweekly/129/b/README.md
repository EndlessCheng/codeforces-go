套路：有三个顶点，枚举「**中间**」的直角顶点更容易计算。

想一想，直角顶点为 $(i,j)$ 的「直角三角形」有多少个？

设第 $i$ 行有 $\textit{rowSum}$ 个 $1$，第 $j$ 列有 $\textit{colSum}$ 个 $1$。根据**乘法原理**，直角顶点为 $(i,j)$ 的「直角三角形」有

$$
(\textit{rowSum} - 1)\cdot(\textit{colSum} - 1)
$$

个，加到答案中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16t421c7GB/)，欢迎点赞关注！

```py [sol-Py3]
class Solution:
    def numberOfRightTriangles(self, grid: List[List[int]]) -> int:
        col_sum = [sum(col) - 1 for col in zip(*grid)]  # 提前减一
        ans = 0
        for row in grid:
            row_sum = sum(row) - 1  # 提前减一
            ans += row_sum * sum(cs for x, cs in zip(row, col_sum) if x)
        return ans
```

```py [sol-Py3 写法二]
class Solution:
    def numberOfRightTriangles(self, grid: List[List[int]]) -> int:
        col_sum = [sum(col) - 1 for col in zip(*grid)]  # 提前减一
        return sum((sum(row) - 1) * sum(cs for x, cs in zip(row, col_sum) if x) for row in grid)
```

```java [sol-Java]
class Solution {
    public long numberOfRightTriangles(int[][] grid) {
        int n = grid[0].length;
        int[] colSum = new int[n];
        Arrays.fill(colSum, -1); // 提前减一
        for (int[] row : grid) {
            for (int j = 0; j < n; j++) {
                colSum[j] += row[j];
            }
        }

        long ans = 0;
        for (int[] row : grid) {
            int rowSum = -1; // 提前减一
            for (int x : row) {
                rowSum += x;
            }
            for (int j = 0; j < row.length; j++) {
                if (row[j] == 1) {
                    ans += rowSum * colSum[j];
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfRightTriangles(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> col_sum(n, -1); // 提前减一
        for (auto& row : grid) {
            for (int j = 0; j < n; j++) {
                col_sum[j] += row[j];
            }
        }

        long long ans = 0;
        for (auto& row : grid) {
            int row_sum = reduce(row.begin(), row.end()) - 1; // 提前减一
            for (int j = 0; j < row.size(); j++) {
                if (row[j] == 1) {
                    ans += row_sum * col_sum[j];
                }
            }
        }
        return ans;
    }
};
```

```c [sol-C]
long long numberOfRightTriangles(int** grid, int gridSize, int* gridColSize) {
    int m = gridSize, n = gridColSize[0];
    int* col_sum = calloc(n, sizeof(int));
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            col_sum[j] += grid[i][j];
        }
    }

    long long ans = 0;
    for (int i = 0; i < m; i++) {
        int row_sum = -1; // 提前减一
        for (int j = 0; j < n; j++) {
            row_sum += grid[i][j];
        }
        for (int j = 0; j < n; j++) {
            if (grid[i][j] == 1) {
                ans += row_sum * (col_sum[j] - 1);
            }
        }
    }

    free(col_sum);
    return ans;
}
```

```go [sol-Go]
func numberOfRightTriangles(grid [][]int) (ans int64) {
    n := len(grid[0])
    colSum := make([]int, n)
    for _, row := range grid {
        for j, x := range row {
            colSum[j] += x
        }
    }

    for _, row := range grid {
        rowSum := -1 // 提前减一
        for _, x := range row {
            rowSum += x
        }
        for j, x := range row {
            if x == 1 {
                ans += int64(rowSum * (colSum[j] - 1))
            }
        }
    }
    return
}
```

```js [sol-JS]
var numberOfRightTriangles = function(grid) {
    const n = grid[0].length;
    const colSum = Array(n).fill(-1); // 提前减一
    for (const row of grid) {
        for (let j = 0; j < n; j++) {
            colSum[j] += row[j];
        }
    }

    let ans = 0;
    for (const row of grid) {
        const rowSum = _.sum(row) - 1; // 提前减一
        for (let j = 0; j < row.length; j++) {
            if (row[j] === 1) {
                ans += rowSum * colSum[j];
            }
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn number_of_right_triangles(grid: Vec<Vec<i32>>) -> i64 {
        let n = grid[0].len();
        let col_sum = (0..n).map(|j| grid.iter().map(|row| row[j]).sum::<i32>() - 1).collect::<Vec<_>>();
        grid.iter()
            .map(|row| (row.iter().sum::<i32>() - 1) as i64 * // 把 rowSum-1 提出来
                row.iter()
                    .zip(col_sum.iter())
                    .filter_map(|(&x, &c)| (x != 0).then_some(c))
                    .sum::<i32>() as i64)
            .sum()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 套路：枚举中间

- [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/)
- [2867. 统计树中的合法路径数目](https://leetcode.cn/problems/count-valid-paths-in-a-tree/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
