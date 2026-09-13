## 写法一：按题意模拟

```py [sol-Python3]
class Solution:
    def cyclicShift(self, n: int, grid: list[list[int]], rowShift: list[int], colShift: list[int]) -> list[list[int]]:
        for i, row in enumerate(grid):
            shift = rowShift[i]
            grid[i] = row[shift:] + row[:shift]

        for j, col in enumerate(zip(*grid)):
            shift = colShift[j]
            new_col = col[shift:] + col[:shift]
            for row, x in zip(grid, new_col):
                row[j] = x

        return grid
```

```java [sol-Java]
class Solution {
    public int[][] cyclicShift(int n, int[][] grid, int[] rowShift, int[] colShift) {
        for (int i = 0; i < grid.length; i++) {
            int shift = rowShift[i];
            int[] row = grid[i];
            int[] newRow = new int[n];
            System.arraycopy(row, shift, newRow, 0, n - shift);
            System.arraycopy(row, 0, newRow, n - shift, shift);
            grid[i] = newRow;
        }

        int[] col = new int[n];
        for (int j = 0; j < colShift.length; j++) {
            int shift = colShift[j];
            // 收集列元素
            int k = 0;
            for (int i = shift; i < n; i++) {
                col[k++] = grid[i][j];
            }
            for (int i = 0; i < shift; i++) {
                col[k++] = grid[i][j];
            }
            // 填入列
            for (int i = 0; i < n; i++) {
                grid[i][j] = col[i];
            }
        }

        return grid;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> cyclicShift(int n, vector<vector<int>>& grid, vector<int>& rowShift, vector<int>& colShift) {
        for (int i = 0; i < n; i++) {
            ranges::rotate(grid[i], grid[i].begin() + rowShift[i]);
        }

        vector<int> col(n);
        for (int j = 0; j < n; j++) {
            int shift = colShift[j];
            // 收集列元素
            int k = 0;
            for (int i = shift; i < n; i++) {
                col[k++] = grid[i][j];
            }
            for (int i = 0; i < shift; i++) {
                col[k++] = grid[i][j];
            }
            // 填入列
            for (int i = 0; i < n; i++) {
                grid[i][j] = col[i];
            }
        }

        return grid;
    }
};
```

```go [sol-Go]
func cyclicShift(n int, grid [][]int, rowShift, colShift []int) [][]int {
	for i, row := range grid {
		shift := rowShift[i]
		grid[i] = append(row[shift:], row[:shift]...)
	}

	column := make([]int, n)
	for j, shift := range colShift {
		// 收集列元素
		col := column[:0]
		for _, row := range grid[shift:] {
			col = append(col, row[j])
		}
		for _, row := range grid[:shift] {
			col = append(col, row[j])
		}
		// 填入列
		for i, row := range grid {
			row[j] = col[i]
		}
	}

	return grid
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：直接计算

答案 $\textit{ans}[i][j]$ 来自哪个格子？

来自哪一行？列移位后，$\textit{row} = (i+\textit{colShift}[j])\bmod n$ 行移到了 $i$ 行。

来自哪一列？行移位后，$\textit{row}$ 行的 $\textit{col} = (j + \textit{rowShift}[\textit{row}])\bmod n$ 列移到了 $j$ 列。

所以 $\textit{ans}[i][j] = \textit{grid}[\textit{row}][\textit{col}]$。

```py [sol-Python3]
class Solution:
    def cyclicShift(self, n: int, grid: list[list[int]], rowShift: list[int], colShift: list[int]) -> list[list[int]]:
        ans = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(n):
                row = (i + colShift[j]) % n
                col = (j + rowShift[row]) % n
                ans[i][j] = grid[row][col]
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] cyclicShift(int n, int[][] grid, int[] rowShift, int[] colShift) {
        int[][] ans = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                int row = (i + colShift[j]) % n;
                int col = (j + rowShift[row]) % n;
                ans[i][j] = grid[row][col];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> cyclicShift(int n, vector<vector<int>>& grid, vector<int>& row_shift, vector<int>& col_shift) {
        vector ans(n, vector<int>(n));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                int row = (i + col_shift[j]) % n;
                int col = (j + row_shift[row]) % n;
                ans[i][j] = grid[row][col];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func cyclicShift(n int, grid [][]int, rowShift, colShift []int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			row := (i + colShift[j]) % n
			col := (j + rowShift[row]) % n
			ans[i][j] = grid[row][col]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

**注**：利用 [189. 轮转数组](https://leetcode.cn/problems/rotate-array/) 的技巧，可以做到完美的**原地修改**。详见 [我的题解](https://leetcode.cn/problems/rotate-array/solutions/2784427/tu-jie-yuan-di-zuo-fa-yi-tu-miao-dong-py-ryfv/)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
