套路：有三个顶点，枚举「**中间**」的直角顶点更容易计算。

设第 $i$ 行有 $\textit{rowSum}$ 个 $1$，第 $j$ 列有 $\textit{colSum}$ 个 $1$。根据**乘法原理**，直角顶点为 $(i,j)$ 的「直角三角形」有

$$
(\textit{rowSum} - 1)\cdot(\textit{colSum} - 1)
$$

个，加到答案中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16t421c7GB/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfRightTriangles(self, grid: List[List[int]]) -> int:
        col_sum = [sum(col) - 1 for col in zip(*grid)]  # 提前减一
        ans = 0
        for row in grid:
            row_sum = sum(row) - 1
            ans += row_sum * sum(c for x, c in zip(row, col_sum) if x)
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def numberOfRightTriangles(self, grid: List[List[int]]) -> int:
        col_sum = [sum(col) - 1 for col in zip(*grid)]  # 提前减一
        return sum((sum(row) - 1) * sum(c for x, c in zip(row, col_sum) if x) for row in grid)
```

```java [sol-Java]
class Solution {
    public long numberOfRightTriangles(int[][] grid) {
        int n = grid[0].length;
        int[] colSum = new int[n];
        Arrays.fill(colSum, -1); // 提前减一
        for (int j = 0; j < n; j++) {
            for (int[] row : grid) {
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
        for (int j = 0; j < n; j++) {
            for (auto& row : grid) {
                col_sum[j] += row[j];
            }
        }

        long long ans = 0;
        for (auto& row : grid) {
            int row_sum = accumulate(row.begin(), row.end(), 0) - 1; // 提前减一
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

```go [sol-Go]
func numberOfRightTriangles(grid [][]int) (ans int64) {
	n := len(grid[0])
	colSum := make([]int, n)
	for j := 0; j < n; j++ {
		for _, row := range grid {
			colSum[j] += row[j]
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 套路：枚举中间

- [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/)
- [2867. 统计树中的合法路径数目](https://leetcode.cn/problems/count-valid-paths-in-a-tree/)

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
