从左到右，遍历每一列。先计算出列的最大值 $\textit{mx}$，再更新列中的 $-1$ 为 $\textit{mx}$。

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
    vector<vector<int>> modifiedMatrix(vector<vector<int>>& matrix) {
        for (int j = 0; j < matrix[0].size(); j++) {
            int mx = 0;
            for (auto& row: matrix) {
                mx = max(mx, row[j]);
            }
            for (auto& row: matrix) {
                if (row[j] == -1) {
                    row[j] = mx;
                }
            }
        }
        return matrix;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int** modifiedMatrix(int** matrix, int matrixSize, int* matrixColSize, int* returnSize, int** returnColumnSizes) {
    int n = matrixColSize[0];
    for (int j = 0; j < n; j++) {
        int mx = 0;
        for (int i = 0; i < matrixSize; i++) {
            mx = MAX(mx, matrix[i][j]);
        }
        for (int i = 0; i < matrixSize; i++) {
            if (matrix[i][j] == -1) {
                matrix[i][j] = mx;
            }
        }
    }
    *returnSize = matrixSize;
    *returnColumnSizes = matrixColSize;
    return matrix;
}
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

```js [sol-JavaScript]
var modifiedMatrix = function(matrix) {
    for (let j = 0; j < matrix[0].length; j++) {
        let mx = 0;
        for (const row of matrix) {
            mx = Math.max(mx, row[j]);
        }
        for (const row of matrix) {
            if (row[j] === -1) {
                row[j] = mx;
            }
        }
    }
    return matrix;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn modified_matrix(mut matrix: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        for j in 0..matrix[0].len() {
            let mx = matrix.iter().map(|row| row[j]).max().unwrap();
            for row in matrix.iter_mut() {
                if row[j] == -1 {
                    row[j] = mx;
                }
            }
        }
        matrix
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{matrix}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
