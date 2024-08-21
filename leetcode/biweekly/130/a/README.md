遍历矩阵，挨个判断 $\textit{grid}[i][j]$：

- 如果 $j>0$ 且 $\textit{grid}[i][j]=\textit{grid}[i][j-1]$，返回 $\texttt{false}$。
- 如果 $i>0$ 且 $\textit{grid}[i][j]\ne\textit{grid}[i-1][j]$，返回 $\texttt{false}$。
- 如果遍历中没有返回 $\texttt{false}$，最后返回 $\texttt{true}$。

```py [sol-Python3]
class Solution:
    def satisfiesConditions(self, grid: List[List[int]]) -> bool:
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                if j and x == row[j - 1] or i and x != grid[i - 1][j]:
                    return False
        return True
```

```java [sol-Java]
class Solution {
    public boolean satisfiesConditions(int[][] grid) {
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (j > 0 && grid[i][j] == grid[i][j - 1] ||
                    i > 0 && grid[i][j] != grid[i - 1][j]) {
                    return false;
                }
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool satisfiesConditions(vector<vector<int>>& grid) {
        for (int i = 0; i < grid.size(); i++) {
            for (int j = 0; j < grid[i].size(); j++) {
                if (j && grid[i][j] == grid[i][j - 1] ||
                    i && grid[i][j] != grid[i - 1][j]) {
                    return false;
                }
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func satisfiesConditions(grid [][]int) bool {
	for i, row := range grid {
		for j, x := range row {
			if j > 0 && x == row[j-1] || i > 0 && x != grid[i-1][j] {
				return false
			}
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

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
