## 方法一：二维前缀和

请看[【图解】二维前缀和](https://leetcode.cn/problems/range-sum-query-2d-immutable/solution/tu-jie-yi-zhang-tu-miao-dong-er-wei-qian-84qp/) 以及 [视频讲解](https://www.bilibili.com/video/BV14r421W7oR/)。

代码实现时，可以取 X 和 Y 的 ASCII 值二进制的最低位，表示 $0$ 和 $1$。

```py [sol-Python3]
class Solution:
    def numberOfSubmatrices(self, grid: List[List[str]]) -> int:
        ans = 0
        m, n = len(grid), len(grid[0])
        s = [[[0, 0] for _ in range(n + 1)] for _ in range(m + 1)]
        for i, row in enumerate(grid):
            for j, c in enumerate(row):
                s[i + 1][j + 1][0] = s[i + 1][j][0] + s[i][j + 1][0] - s[i][j][0]
                s[i + 1][j + 1][1] = s[i + 1][j][1] + s[i][j + 1][1] - s[i][j][1]
                if c != '.':
                    s[i + 1][j + 1][ord(c) & 1] += 1
                if s[i + 1][j + 1][0] and s[i + 1][j + 1][0] == s[i + 1][j + 1][1]:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSubmatrices(char[][] grid) {
        int ans = 0;
        int m = grid.length;
        int n = grid[0].length;
        int[][][] sum = new int[m + 1][n + 1][2];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                sum[i + 1][j + 1][0] = sum[i + 1][j][0] + sum[i][j + 1][0] - sum[i][j][0];
                sum[i + 1][j + 1][1] = sum[i + 1][j][1] + sum[i][j + 1][1] - sum[i][j][1];
                if (grid[i][j] != '.') {
                    sum[i + 1][j + 1][grid[i][j] & 1]++;
                }
                if (sum[i + 1][j + 1][0] > 0 && sum[i + 1][j + 1][0] == sum[i + 1][j + 1][1]) {
                    ans++;
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
    int numberOfSubmatrices(vector<vector<char>>& grid) {
        int ans = 0, m = grid.size(), n = grid[0].size();
        vector<vector<array<int, 2>>> sum(m + 1, vector<array<int, 2>>(n + 1));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                sum[i + 1][j + 1][0] = sum[i + 1][j][0] + sum[i][j + 1][0] - sum[i][j][0];
                sum[i + 1][j + 1][1] = sum[i + 1][j][1] + sum[i][j + 1][1] - sum[i][j][1];
                if (grid[i][j] != '.') {
                    sum[i + 1][j + 1][grid[i][j] & 1]++;
                }
                if (sum[i + 1][j + 1][0] && sum[i + 1][j + 1][0] == sum[i + 1][j + 1][1]) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubmatrices(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	sum := make([][][2]int, m+1)
	for i := range sum {
		sum[i] = make([][2]int, n+1)
	}
	for i, row := range grid {
		for j, c := range row {
			sum[i+1][j+1][0] = sum[i+1][j][0] + sum[i][j+1][0] - sum[i][j][0]
			sum[i+1][j+1][1] = sum[i+1][j][1] + sum[i][j+1][1] - sum[i][j][1]
			if c != '.' {
				sum[i+1][j+1][c&1]++
			}
			if sum[i+1][j+1][0] > 0 && sum[i+1][j+1][0] == sum[i+1][j+1][1] {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

## 方法二：维护每列的字符个数

遍历每一行，同时用一个长为 $n\times 2$ 的数组 $\textit{colCnt}$ 维护每一列的 X 和 Y 的个数。

遍历当前行时，一边更新 $\textit{colCnt}[j]$，一边累加 $\textit{colCnt}[j][0]$ 到变量 $s_0$ 中，累加 $\textit{colCnt}[j][1]$ 到变量 $s_1$ 中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Ry411q71f/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfSubmatrices(self, grid: List[List[str]]) -> int:
        ans = 0
        col_cnt = [[0, 0] for _ in grid[0]]
        for row in grid:
            s0 = s1 = 0
            for j, c in enumerate(row):
                if c != '.':
                    col_cnt[j][ord(c) & 1] += 1
                s0 += col_cnt[j][0]
                s1 += col_cnt[j][1]
                if s0 and s0 == s1:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSubmatrices(char[][] grid) {
        int ans = 0;
        int[][] colCnt = new int[grid[0].length][2];
        for (char[] row : grid) {
            int s0 = 0, s1 = 0;
            for (int j = 0; j < row.length; j++) {
                if (row[j] != '.') {
                    colCnt[j][row[j] & 1]++;
                }
                s0 += colCnt[j][0];
                s1 += colCnt[j][1];
                if (s0 > 0 && s0 == s1) {
                    ans++;
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
    int numberOfSubmatrices(vector<vector<char>>& grid) {
        int ans = 0;
        vector<array<int, 2>> col_cnt(grid[0].size());
        for (auto& row : grid) {
            int s0 = 0, s1 = 0;
            for (int j = 0; j < row.size(); j++) {
                if (row[j] != '.') {
                    col_cnt[j][row[j] & 1]++;
                }
                s0 += col_cnt[j][0];
                s1 += col_cnt[j][1];
                if (s0 && s0 == s1) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubmatrices(grid [][]byte) (ans int) {
	colCnt := make([][2]int, len(grid[0]))
	for _, row := range grid {
		s0, s1 := 0, 0
		for j, c := range row {
			if c != '.' {
				colCnt[j][c&1]++
			}
			s0 += colCnt[j][0]
			s1 += colCnt[j][1]
			if s0 > 0 && s0 == s1 {
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [3070. 元素和小于等于 k 的子矩阵的数目](https://leetcode.cn/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/)

## 思考题

如果不要求包含左上角呢？

> 提示：枚举子矩形上下边界，转换成一维的情况。

可以参考 [363. 矩形区域不超过 K 的最大数值和](https://leetcode.cn/problems/max-sum-of-rectangle-no-larger-than-k/) 这题的思路。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
