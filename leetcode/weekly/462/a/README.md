根据题意，交换的范围是行号 $[x,x+k-1]$，列号 $[y,y+k-1]$。

类似 [344. 反转字符串](https://leetcode.cn/problems/reverse-string/)，用**双指针**实现：

- 初始化 $l=x$，$r=x+k-1$。
- 循环直到 $l\ge r$。
- 每次循环，对于 $[y,y+k-1]$ 中的每个整数 $j$，交换 $\textit{grid}[l][j]$ 和 $\textit{grid}[r][j]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1QNbNzxEtZ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def reverseSubmatrix(self, grid: List[List[int]], x: int, y: int, k: int) -> List[List[int]]:
        l, r = x, x + k - 1
        while l < r:
            for j in range(y, y + k):
                grid[l][j], grid[r][j] = grid[r][j], grid[l][j]
            l += 1
            r -= 1
        return grid
```

```py [sol-Python3 整体交换]
class Solution:
    def reverseSubmatrix(self, grid: List[List[int]], x: int, y: int, k: int) -> List[List[int]]:
        l, r = x, x + k - 1
        while l < r:
            grid[l][y: y + k], grid[r][y: y + k] = grid[r][y: y + k], grid[l][y: y + k]
            l += 1
            r -= 1
        return grid
```

```java [sol-Java]
class Solution {
    public int[][] reverseSubmatrix(int[][] grid, int x, int y, int k) {
        int l = x;
        int r = x + k - 1;
        while (l < r) {
            for (int j = y; j < y + k; j++) {
                int tmp = grid[l][j];
                grid[l][j] = grid[r][j];
                grid[r][j] = tmp;
            }
            l++;
            r--;
        }
        return grid;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> reverseSubmatrix(vector<vector<int>>& grid, int x, int y, int k) {
        int l = x, r = x + k - 1;
        while (l < r) {
            for (int j = y; j < y + k; j++) {
                swap(grid[l][j], grid[r][j]);
            }
            l++;
            r--;
        }
        return grid;
    }
};
```

```go [sol-Go]
func reverseSubmatrix(grid [][]int, x, y, k int) [][]int {
	l, r := x, x+k-1
	for l < r {
		for j := y; j < y+k; j++ {
			grid[l][j], grid[r][j] = grid[r][j], grid[l][j]
		}
		l++
		r--
	}
	return grid
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k^2)$。
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
