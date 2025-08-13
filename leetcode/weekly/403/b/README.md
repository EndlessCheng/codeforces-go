遍历 $\textit{grid}$：

- 找到最左最右的 $1$ 的列号，分别记作 $\textit{left},\textit{right}$，则矩形底边长至少为 $\textit{right}-\textit{left}+1$。
- 找到最上最下的 $1$ 的行号，分别记作 $\textit{top},\textit{bottom}$，则矩形高至少为 $\textit{bottom} - \textit{top} + 1$。

矩形面积至少为

$$
(\textit{right} - \textit{left} + 1) \cdot (\textit{bottom} - \textit{top} + 1)
$$

```py [sol-Python3]
class Solution:
    def minimumArea(self, grid: List[List[int]]) -> int:
        left = top = inf
        right = bottom = 0
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if x:
                    left = min(left, j)
                    right = max(right, j)
                    top = min(top, i)
                    bottom = i
        return (right - left + 1) * (bottom - top + 1)
```

```java [sol-Java]
class Solution {
    public int minimumArea(int[][] grid) {
        int left = Integer.MAX_VALUE;
        int right = 0;
        int top = Integer.MAX_VALUE;
        int bottom = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] == 1) {
                    left = Math.min(left, j);
                    right = Math.max(right, j);
                    top = Math.min(top, i);
                    bottom = i;
                }
            }
        }
        return (right - left + 1) * (bottom - top + 1);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumArea(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        int left = m, right = 0, top = n, bottom = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j]) {
                    left = min(left, j);
                    right = max(right, j);
                    top = min(top, i);
                    bottom = i;
                }
            }
        }
        return (right - left + 1) * (bottom - top + 1);
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int minimumArea(int** grid, int gridSize, int* gridColSize) {
    int m = gridSize, n = gridColSize[0];
    int left = m, right = 0, top = n, bottom = 0;
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            if (grid[i][j]) {
                left = MIN(left, j);
                right = MAX(right, j);
                top = MIN(top, i);
                bottom = i;
            }
        }
    }
    return (right - left + 1) * (bottom - top + 1);
}
```

```go [sol-Go]
func minimumArea(grid [][]int) int {
	left, right := len(grid[0]), 0
	top, bottom := len(grid), 0
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				left = min(left, j)
				right = max(right, j)
				top = min(top, i)
				bottom = i
			}
		}
	}
	return (right - left + 1) * (bottom - top + 1)
}
```

```js [sol-JavaScript]
var minimumArea = function(grid) {
    const m = grid.length, n = grid[0].length;
    let left = m, right = 0, top = n, bottom = 0;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j]) {
                left = Math.min(left, j);
                right = Math.max(right, j);
                top = Math.min(top, i);
                bottom = i;
            }
        }
    }
    return (right - left + 1) * (bottom - top + 1);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_area(grid: Vec<Vec<i32>>) -> i32 {
        let mut left = usize::MAX;
        let mut right = 0;
        let mut top = usize::MAX;
        let mut bottom = 0;
        for (i, row) in grid.into_iter().enumerate() {
            for (j, x) in row.into_iter().enumerate() {
                if x != 0 {
                    left = left.min(j);
                    right = right.max(j);
                    top = top.min(i);
                    bottom = i;
                }
            }
        }
        ((right - left + 1) * (bottom - top + 1)) as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
