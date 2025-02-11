对于每条对角线，行号 $i$ 减列号 $j$ 是一个定值。（回想一下 [51. N 皇后](https://leetcode.cn/problems/n-queens/) 的写法）

设 $k=i-j+n$，那么右上角那条对角线的 $k=1$，左下角那条对角线的 $k=m+n-1$。（本题 $m=n$）

于是枚举 $k=1,2,3,\ldots,m+n-1$，就相当于在从右上到左下，一条一条地枚举对角线。

由于 $i = k+j-n$，知道 $j$ 就知道 $i$，所以我们只需要计算出每条对角线的 $j$ 的最小值和最大值，就可以开始遍历对角线了。

- 由于 $j=i-k+n$，当 $i=0$ 的时候 $j$ 取到最小值 $n-k$，但这个数不能是负数，所以最小的 $j$ 是 $\max(n-k,0)$。
- 由于 $j=i-k+n$，当 $i=m-1$ 的时候 $j$ 取到最大值 $m + n - 1 - k$，但这个数不能超过 $n-1$，所以最大的 $j$ 是 $\min(m + n - 1 - k, n - 1)$。

然后就可以模拟了：

1. 把对角线的元素存入列表 $a$ 中。
2. 如果最小的 $j$ 大于 $0$，说明我们在主对角线右上，升序排序；否则降序排序。
3. 把 $a$ 按顺序原样放回对角线中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ekN2ebEHx/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def sortMatrix(self, grid: List[List[int]]) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        # 第一排在右上，最后一排在左下
        # 每排从左上到右下
        # 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
        for k in range(1, m + n):
            # 核心：计算 j 的最小值和最大值
            min_j = max(n - k, 0)  # i=0 的时候，j=n-k，但不能是负数
            max_j = min(m + n - 1 - k, n - 1)  # i=m-1 的时候，j=m+n-1-k，但不能超过 n-1
            a = [grid[k + j - n][j] for j in range(min_j, max_j + 1)]  # 根据 k 的定义得 i=k+j-n
            a.sort(reverse=min_j == 0)
            for j, val in zip(range(min_j, max_j + 1), a):
                grid[k + j - n][j] = val
        return grid
```

```java [sol-Java]
class Solution {
    public int[][] sortMatrix(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        // 第一排在右上，最后一排在左下
        // 每排从左上到右下
        // 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
        for (int k = 1; k < m + n; k++) {
            // 核心：计算 j 的最小值和最大值
            int minJ = Math.max(n - k, 0); // i=0 的时候，j=n-k，但不能是负数
            int maxJ = Math.min(m + n - 1 - k, n - 1); // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1
            List<Integer> a = new ArrayList<>(maxJ - minJ + 1); // 预分配空间
            for (int j = minJ; j <= maxJ; j++) {
                a.add(grid[k + j - n][j]); // 根据 k 的定义得 i=k+j-n
            }
            a.sort(minJ > 0 ? null : Comparator.reverseOrder());
            for (int j = minJ; j <= maxJ; j++) {
                grid[k + j - n][j] = a.get(j - minJ);
            }
        }
        return grid;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> sortMatrix(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        // 第一排在右上，最后一排在左下
        // 每排从左上到右下
        // 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
        for (int k = 1; k < m + n; k++) {
            // 核心：计算 j 的最小值和最大值
            int min_j = max(n - k, 0); // i=0 的时候，j=n-k，但不能是负数
            int max_j = min(m + n - 1 - k, n - 1); // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1
            vector<int> a;
            for (int j = min_j; j <= max_j; j++) {
                a.push_back(grid[k + j - n][j]); // 根据 k 的定义得 i=k+j-n
            }
            if (min_j > 0) { // 右上角三角形
                ranges::sort(a);
            } else { // 左下角三角形（包括中间对角线）
                ranges::sort(a, greater<int>());
            }
            for (int j = min_j; j <= max_j; j++) {
                grid[k + j - n][j] = a[j - min_j];
            }
        }
        return grid;
    }
};
```

```go [sol-Go]
func sortMatrix(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	// 第一排在右上，最后一排在左下
	// 每排从左上到右下
	// 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
	for k := 1; k < m+n; k++ {
		// 核心：计算 j 的最小值和最大值
		minJ := max(n-k, 0)       // i=0 的时候，j=n-k，但不能是负数
		maxJ := min(m+n-1-k, n-1) // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1
		a := []int{}
		for j := minJ; j <= maxJ; j++ {
			a = append(a, grid[k+j-n][j]) // 根据 k 的定义得 i=k+j-n
		}
		if minJ > 0 { // 右上角三角形
			slices.Sort(a)
		} else { // 左下角三角形（包括中间对角线）
			slices.SortFunc(a, func(a, b int) int { return b - a })
		}
		for j := minJ; j <= maxJ; j++ {
			grid[k+j-n][j] = a[j-minJ]
		}
	}
	return grid
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$，其中 $n$ 为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [1329. 将矩阵按对角线排序](https://leetcode.cn/problems/sort-the-matrix-diagonally/)
- [2711. 对角线上不同值的数量差](https://leetcode.cn/problems/difference-of-number-of-distinct-values-on-diagonals/)
- [498. 对角线遍历](https://leetcode.cn/problems/diagonal-traverse/) 副对角线
- [562. 矩阵中最长的连续 1 线段](https://leetcode.cn/problems/longest-line-of-consecutive-one-in-matrix/)（会员题）

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
