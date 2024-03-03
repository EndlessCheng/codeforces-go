请看 [视频讲解](https://www.bilibili.com/video/BV14r421W7oR/) 第三题，包含思考题（更大值域）的做法。

统计 Y 中的元素出现次数，记到一个长为 $3$ 的数组 $\textit{cnt}_1$ 中。统计不在 Y 中的元素出现次数，记到一个长为 $3$ 的数组 $\textit{cnt}_2$ 中。

计算最多可以保留多少个元素**不变**，设这个值为 $\textit{maxNotChange}$。

在 $0,1,2$ 中枚举 $i$ 和 $j$，其中 $i\ne j$。让 Y 中的元素都变成 $i$，不在 Y 中的元素都变成 $j$，那么 $\textit{maxNotChange}$ 就是 $\textit{cnt}_1[i]+\textit{cnt}_2[j]$ 的最大值。

最后返回 $n^2 - \textit{maxNotChange}$，即最少要修改的元素个数。

```py [sol-Python3]
class Solution:
    def minimumOperationsToWriteY(self, grid: List[List[int]]) -> int:
        cnt1 = [0] * 3
        cnt2 = [0] * 3
        n = len(grid)
        m = n // 2

        for i, row in enumerate(grid[:m]):
            cnt1[row[i]] += 1
            cnt1[row[-1 - i]] += 1
            for j, x in enumerate(row):
                if j != i and j != n - 1 - i:
                    cnt2[x] += 1
        for row in grid[m:]:
            cnt1[row[m]] += 1
            for j, x in enumerate(row):
                if j != m:
                    cnt2[x] += 1

        max_not_change = 0
        for i, c1 in enumerate(cnt1):
            for j, c2 in enumerate(cnt2):
                if i != j:
                    max_not_change = max(max_not_change, c1 + c2)
        return n * n - max_not_change
```

```java [sol-Java]
class Solution {
    public int minimumOperationsToWriteY(int[][] grid) {
        int[] cnt1 = new int[3];
        int[] cnt2 = new int[3];
        int n = grid.length;
        int m = n / 2;
        for (int i = 0; i < m; i++) {
            cnt1[grid[i][i]]++;
            cnt1[grid[i][n - 1 - i]]++;
            for (int j = 0; j < n; j++) {
                if (j != i && j != n - 1 - i) {
                    cnt2[grid[i][j]]++;
                }
            }
        }
        for (int i = m; i < n; i++) {
            cnt1[grid[i][m]]++;
            for (int j = 0; j < n; j++) {
                if (j != m) {
                    cnt2[grid[i][j]]++;
                }
            }
        }

        int maxNotChange = 0;
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                if (i != j) {
                    maxNotChange = Math.max(maxNotChange, cnt1[i] + cnt2[j]);
                }
            }
        }
        return n * n - maxNotChange;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOperationsToWriteY(vector<vector<int>> &grid) {
        int cnt1[3]{}, cnt2[3]{};
        int n = grid.size();
        int m = n / 2;
        for (int i = 0; i < m; i++) {
            cnt1[grid[i][i]]++;
            cnt1[grid[i][n - 1 - i]]++;
            for (int j = 0; j < n; j++) {
                if (j != i && j != n - 1 - i) {
                    cnt2[grid[i][j]]++;
                }
            }
        }
        for (int i = m; i < n; i++) {
            cnt1[grid[i][m]]++;
            for (int j = 0; j < n; j++) {
                if (j != m) {
                    cnt2[grid[i][j]]++;
                }
            }
        }

        int max_not_change = 0;
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                if (i != j) {
                    max_not_change = max(max_not_change, cnt1[i] + cnt2[j]);
                }
            }
        }
        return n * n - max_not_change;
    }
};
```

```go [sol-Go]
func minimumOperationsToWriteY(grid [][]int) int {
	var cnt1, cnt2 [3]int
	n := len(grid)
	m := n / 2
	for i, row := range grid[:m] {
		cnt1[row[i]]++
		cnt1[row[n-1-i]]++
		for j, x := range row {
			if j != i && j != n-1-i {
				cnt2[x]++
			}
		}
	}
	for _, row := range grid[m:] {
		cnt1[row[m]]++
		for j, x := range row {
			if j != m {
				cnt2[x]++
			}
		}
	}

	maxNotChange := 0
	for i, c1 := range cnt1 {
		for j, c2 := range cnt2 {
			if i != j {
				maxNotChange = max(maxNotChange, c1+c2)
			}
		}
	}
	return n*n - maxNotChange
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2 + k^2)$，其中 $n$ 为 $\textit{grid}$ 的长度，$k=3$。
- 空间复杂度：$\mathcal{O}(k)$。Python 忽略切片开销。

## 思考题（更快的做法）

如果元素不仅仅是 $0,1,2$ 要怎么做？

**解答**：$\textit{maxNotChange}$ 只有三种情况：

- $\textit{cnt}_1$ 中的最大值加上 $\textit{cnt}_2$ 中的最大值，前提是这两个最大值对应的元素不同。
- 如果对应的元素相同，我们可以取 $\textit{cnt}_1$ 中的最大值加上 $\textit{cnt}_2$ 中的次大值，或者 $\textit{cnt}_1$ 中的次大值加上 $\textit{cnt}_2$ 中的最大值。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
