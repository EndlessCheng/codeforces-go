遍历每一列：先计算出列的最大值 $\textit{mx}$，再更新列中的 $-1$ 为 $\textit{mx}$。

```py [sol-Python3]
class Solution:
    def modifiedMatrix(self, matrix: List[List[int]]) -> List[List[int]]:
        for j in range(len(matrix[0])):
            mx = max(row[j] for row in matrix)
            for row in matrix:
                if row[j] == -1:
                    row[j] = mx
        return matrix
```

```java [sol-Java]
class Solution {
    public int[][] modifiedMatrix(int[][] matrix) {
        for (int j = 0; j < matrix[0].length; j++) {
            int mx = 0;
            for (int[] row : matrix) {
                mx = Math.max(mx, row[j]);
            }
            for (int[] row : matrix) {
                if (row[j] == -1) {
                    row[j] = mx;
                }
            }
        }
        return matrix;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> modifiedMatrix(vector<vector<int>> &matrix) {
        for (int j = 0; j < matrix[0].size(); j++) {
            int mx = 0;
            for (auto &row: matrix) {
                mx = max(mx, row[j]);
            }
            for (auto &row: matrix) {
                if (row[j] == -1) {
                    row[j] = mx;
                }
            }
        }
        return matrix;
    }
};
```

```go [sol-Go]
func modifiedMatrix(matrix [][]int) [][]int {
	for j := range matrix[0] {
		mx := 0
		for _, row := range matrix {
			mx = max(mx, row[j])
		}
		for _, row := range matrix {
			if row[j] == -1 {
				row[j] = mx
			}
		}
	}
	return matrix
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{matrix}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
