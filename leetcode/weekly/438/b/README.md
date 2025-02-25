本题没有负数，所以每排都可以取恰好 $\textit{limits}[i]$ 个数。

为了最大化元素和，每排取最大的 $\textit{limits}[i]$ 个数。

然后再取这些数中最大的 $k$ 个数，求和即为答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hiAUeWEUG/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxSum(self, grid: List[List[int]], limits: List[int], k: int) -> int:
        a = []
        for row, limit in zip(grid, limits):
            row.sort(reverse=True)
            a.extend(row[:limit])
        a.sort(reverse=True)
        return sum(a[:k])
```

```java [sol-Java]
// 更快的写法见【数组】
class Solution {
    public long maxSum(int[][] grid, int[] limits, int k) {
        List<Integer> a = new ArrayList<>();
        for (int i = 0; i < grid.length; i++) {
            int[] row = grid[i];
            Arrays.sort(row);
            for (int j = row.length - limits[i]; j < row.length; j++) {
                a.add(row[j]);
            }
        }
        a.sort(Collections.reverseOrder());
        long ans = 0;
        for (int i = 0; i < k; i++) {
            ans += a.get(i);
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public long maxSum(int[][] grid, int[] limits, int k) {
        int m = grid.length;
        int n = grid[0].length;
        int[] a = new int[m * n];
        int size = 0;
        for (int i = 0; i < m; i++) {
            int[] row = grid[i];
            Arrays.sort(row);
            for (int j = n - limits[i]; j < n; j++) {
                a[size++] = row[j];
            }
        }

        Arrays.sort(a, 0, size);
        long ans = 0;
        for (int i = size - k; i < size; i++) {
            ans += a[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSum(vector<vector<int>>& grid, vector<int>& limits, int k) {
        vector<int> a;
        for (int i = 0; i < grid.size(); i++) {
            auto& row = grid[i];
            ranges::sort(row, greater());
            a.insert(a.end(), row.begin(), row.begin() + limits[i]);
        }
        ranges::sort(a, greater());
        return reduce(a.begin(), a.begin() + k, 0LL);
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    long long maxSum(vector<vector<int>>& grid, vector<int>& limits, int k) {
        vector<int> a;
        for (int i = 0; i < grid.size(); i++) {
            auto& row = grid[i];
            ranges::nth_element(row, row.end() - limits[i]);
            a.insert(a.end(), row.end() - limits[i], row.end());
        }
        ranges::nth_element(a, a.end() - k);
        return reduce(a.end() - k, a.end(), 0LL);
    }
};
```

```go [sol-Go]
func maxSum(grid [][]int, limits []int, k int) (ans int64) {
	a := []int{}
	cmp := func(a, b int) int { return b - a }
	for i, row := range grid {
		slices.SortFunc(row, cmp)
		a = append(a, row[:limits[i]]...)
	}
	slices.SortFunc(a, cmp)
	for _, x := range a[:k] {
		ans += int64(x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log (mn))$ 或 $\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。用快速选择可以做到 $\mathcal{O}(mn)$，见 C++ 代码。
- 空间复杂度：$\mathcal{O}(mn)$，或者 $\textit{limits}[i]$ 之和。

更多相似题目，见下面贪心题单中的「**§1.1 从最小/最大开始贪心**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
