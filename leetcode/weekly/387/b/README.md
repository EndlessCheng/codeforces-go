## 方法一：二维前缀和

请看[【图解】二维前缀和](https://leetcode.cn/problems/range-sum-query-2d-immutable/solution/tu-jie-yi-zhang-tu-miao-dong-er-wei-qian-84qp/) 以及 [视频讲解](https://www.bilibili.com/video/BV14r421W7oR/)。

```py [sol-Python3]
class Solution:
    def countSubmatrices(self, grid: List[List[int]], k: int) -> int:
        ans = 0
        m, n = len(grid), len(grid[0])
        s = [[0] * (n + 1) for _ in range(m + 1)]
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                s[i + 1][j + 1] = s[i + 1][j] + s[i][j + 1] - s[i][j] + x
                if s[i + 1][j + 1] <= k:
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countSubmatrices(int[][] grid, int k) {
        int ans = 0;
        int m = grid.length;
        int n = grid[0].length;
        int[][] sum = new int[m + 1][n + 1];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                sum[i + 1][j + 1] = sum[i + 1][j] + sum[i][j + 1] - sum[i][j] + grid[i][j];
                if (sum[i + 1][j + 1] <= k) {
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
    int countSubmatrices(vector<vector<int>> &grid, int k) {
        int ans = 0, m = grid.size(), n = grid[0].size();
        vector<vector<int>> sum(m + 1, vector<int>(n + 1));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                sum[i + 1][j + 1] = sum[i + 1][j] + sum[i][j + 1] - sum[i][j] + grid[i][j];
                ans += sum[i + 1][j + 1] <= k;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubmatrices(grid [][]int, k int) (ans int) {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range grid {
		sum[i+1] = make([]int, n+1)
		for j, x := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
			if sum[i+1][j+1] <= k {
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

注：如果原地计算二维前缀和，可以做到 $\mathcal{O}(1)$ 额外空间。

## 方法二：维护每列的元素和

遍历每一行，同时用一个长为 $n$ 的数组 $\textit{colSum}$ 维护每一列的元素和。

遍历当前行时，一边更新 $\textit{colSum}[j]$，一边累加 $\textit{colSum}[j]$ 到变量 $s$ 中。

如果 $s\le k$ 则把答案加一，否则可以退出循环（因为矩阵元素都非负）。

```py [sol-Python3]
class Solution:
    def countSubmatrices(self, grid: List[List[int]], k: int) -> int:
        ans = 0
        col_sum = [0] * len(grid[0])
        for row in grid:
            s = 0
            for j, x in enumerate(row):
                col_sum[j] += x
                s += col_sum[j]
                if s > k:
                    break
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countSubmatrices(int[][] grid, int k) {
        int ans = 0;
        int n = grid[0].length;
        int[] colSum = new int[n];
        for (int[] row : grid) {
            int s = 0;
            for (int j = 0; j < n; j++) {
                colSum[j] += row[j];
                s += colSum[j];
                if (s > k) break;
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSubmatrices(vector<vector<int>> &grid, int k) {
        int ans = 0, n = grid[0].size();
        vector<int> col_sum(n);
        for (auto &row : grid) {
            int s = 0;
            for (int j = 0; j < n; j++) {
                col_sum[j] += row[j];
                s += col_sum[j];
                if (s > k) break;
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubmatrices(grid [][]int, k int) (ans int) {
	colSum := make([]int, len(grid[0]))
	for _, row := range grid {
		s := 0
		for j, x := range row {
			colSum[j] += x
			s += colSum[j]
			if s > k {
				break
			}
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

注：如果把每列元素和保存到 $\textit{grid}$ 的第一行，可以做到 $\mathcal{O}(1)$ 额外空间。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
