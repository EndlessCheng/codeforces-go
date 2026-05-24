两人的移动路径，可以分别视作一个减函数和一个增函数（非严格）。所以交集要么是水平的一段，要么是垂直的一段。

所以问题等价于每行每列的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，[我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。

需要特别注意子数组长度为 $1$ 的情况，两人在交点处的移动路径只能形如佛教符号（如下图）。这意味着，交点不能在 $\textit{grid}$ 的边界上。所以边界上的子数组长度至少为 $2$（由示例 2 可知，子数组长度为 $2$ 的情况是存在的）。

![lc3938.png](https://pic.leetcode.cn/1779581072-bPmSGA-lc3938.png){:width=50px}

我们可以先计算不在边界上的 $\textit{grid}[i][j]$ 的最大值，然后就只需考虑子数组长度至少为 $2$ 的情况了，无需特判边界。

长度至少为 $2$ 的最大子数组和，可以用**前缀和**或者 **DP** 解决。

我们可以在 [DP 做法](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/) 上略作修改，先计算 `ans = max(ans, f + nums[i])`，再更新 `f = max(f, 0) + nums[i]`，这样 `f + nums[i]` 就可以保证子数组至少有两个数了。

[本题视频讲解](https://www.bilibili.com/video/BV1iuG76VEXy/?t=15m42s)，欢迎点赞关注~

```py [sol-Python3]
# 手写 max 更快
fmax = lambda a, b: b if b > a else a

class Solution:
    # 53. 最大子数组和（子数组长度 >= 2）
    def maxSubArray(self, nums: list[int]) -> int:
        ans = -inf  # 注意答案可以是负数，不能初始化成 0
        f = nums[0]
        for x in nums[1:]:
            ans = fmax(ans, f + x)  # f+x 保证子数组至少有两个数
            f = fmax(f, 0) + x
        return ans

    def maxScore(self, grid: list[list[int]]) -> int:
        ans = -inf

        # 单独计算子数组长为 1 的情况，此时子数组不能在 grid 的边界上
        if len(grid) > 2 and len(grid[0]) > 2:
            ans = max(max(row[1: -1]) for row in grid[1: -1])

        # 每行的最大子数组和（子数组长度 >= 2）
        for row in grid:
            ans = fmax(ans, self.maxSubArray(row))

        # 每列的最大子数组和（子数组长度 >= 2）
        for col in zip(*grid):
            ans = fmax(ans, self.maxSubArray(list(col)))

        return ans
```

```java [sol-Java]
class Solution {
    public int maxScore(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int ans = Integer.MIN_VALUE;

        // 单独计算子数组长为 1 的情况，此时子数组不能在 grid 的边界上
        for (int i = 1; i < m - 1; i++) {
            for (int j = 1; j < n - 1; j++) {
                ans = Math.max(ans, grid[i][j]);
            }
        }

        // 每行的最大子数组和（子数组长度 >= 2）
        for (int[] row : grid) {
            ans = Math.max(ans, maxSubArray(row));
        }

        // 每列的最大子数组和（子数组长度 >= 2）
        int[] col = new int[m];
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m; i++) {
                col[i] = grid[i][j];
            }
            ans = Math.max(ans, maxSubArray(col));
        }

        return ans;
    }

    // 53. 最大子数组和（子数组长度 >= 2）
    private int maxSubArray(int[] nums) {
        int ans = Integer.MIN_VALUE; // 注意答案可以是负数，不能初始化成 0
        int f = nums[0];
        for (int i = 1; i < nums.length; i++) {
            int x = nums[i];
            ans = Math.max(ans, f + x); // f+x 保证子数组至少有两个数
            f = Math.max(f, 0) + x;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 53. 最大子数组和（子数组长度 >= 2）
    int maxSubArray(vector<int>& nums) {
        int ans = INT_MIN; // 注意答案可以是负数，不能初始化成 0
        int f = nums[0];
        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i];
            ans = max(ans, f + x); // f+x 保证子数组至少有两个数
            f = max(f, 0) + x;
        }
        return ans;
    }

public:
    int maxScore(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        int ans = INT_MIN;

        // 单独计算子数组长为 1 的情况，此时子数组不能在 grid 的边界上
        for (int i = 1; i < m - 1; i++) {
            for (int j = 1; j < n - 1; j++) {
                ans = max(ans, grid[i][j]);
            }
        }

        // 每行的最大子数组和（子数组长度 >= 2）
        for (auto& row : grid) {
            ans = max(ans, maxSubArray(row));
        }

        // 每列的最大子数组和（子数组长度 >= 2）
        vector<int> col(m);
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m; i++) {
                col[i] = grid[i][j];
            }
            ans = max(ans, maxSubArray(col));
        }

        return ans;
    }
};
```

```go [sol-Go]
// 53. 最大子数组和（子数组长度 >= 2）
func maxSubArray(nums []int) int {
	ans := math.MinInt // 注意答案可以是负数，不能初始化成 0
	f := nums[0]
	for _, x := range nums[1:] {
		ans = max(ans, f+x) // f+x 保证子数组至少有两个数
		f = max(f, 0) + x
	}
	return ans
}

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := math.MinInt

	// 单独计算子数组长为 1 的情况，此时子数组不能在 grid 的边界上
	if m > 2 && n > 2 {
		for _, row := range grid[1 : m-1] {
			ans = max(ans, slices.Max(row[1:n-1]))
		}
	}

	// 每行的最大子数组和（子数组长度 >= 2）
	for _, row := range grid {
		ans = max(ans, maxSubArray(row))
	}

	// 每列的最大子数组和（子数组长度 >= 2）
	col := make([]int, m)
	for j := range n {
		for i, row := range grid {
			col[i] = row[j]
		}
		ans = max(ans, maxSubArray(col))
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$ 或 $\mathcal{O}(1)$，取决于是否使用辅助数组。

## 专题训练

见下面动态规划题单的「**§1.3 最大子数组和**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
