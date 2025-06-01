暴力枚举所有子矩形，把子矩形中的所有元素添加到一个列表 $a$ 中。

把 $a$ 排序后，不同元素之差的最小值一定在相邻元素中，计算相邻不同元素之差的最小值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Dz76zfEdi/?t=10m54s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minAbsDiff(self, grid: List[List[int]], k: int) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        ans = [[0] * (n - k + 1) for _ in range(m - k + 1)]
        for i in range(m - k + 1):
            sub_grid = grid[i: i + k]
            for j in range(n - k + 1):
                a = []
                for row in sub_grid:
                    a.extend(row[j: j + k])
                a.sort()

                res = inf
                for x, y in pairwise(a):
                    if x < y:
                        res = min(res, y - x)
                if res < inf:
                    ans[i][j] = res
        return ans
```

```java [sol-Java]
class Solution {
    public int[][] minAbsDiff(int[][] grid, int k) {
        int m = grid.length;
        int n = grid[0].length;
        int[][] ans = new int[m - k + 1][n - k + 1];
        int[] a = new int[k * k];
        for (int i = 0; i <= m - k; i++) {
            for (int j = 0; j <= n - k; j++) {
                int idx = 0;
                for (int x = 0; x < k; x++) {
                    for (int y = 0; y < k; y++) {
                        a[idx++] = grid[i + x][j + y];
                    }
                }
                Arrays.sort(a);

                int res = Integer.MAX_VALUE;
                for (int p = 1; p < a.length; p++) {
                    if (a[p] > a[p - 1]) {
                        res = Math.min(res, a[p] - a[p - 1]);
                    }
                }
                if (res < Integer.MAX_VALUE) {
                    ans[i][j] = res;
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
    vector<vector<int>> minAbsDiff(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        vector ans(m - k + 1, vector<int>(n - k + 1));
        for (int i = 0; i <= m - k; i++) {
            for (int j = 0; j <= n - k; j++) {
                vector<int> a;
                for (int x = 0; x < k; x++) {
                    for (int y = 0; y < k; y++) {
                        a.push_back(grid[i + x][j + y]);
                    }
                }
                ranges::sort(a);

                int res = INT_MAX;
                for (int p = 1; p < a.size(); p++) {
                    if (a[p] > a[p - 1]) {
                        res = min(res, a[p] - a[p - 1]);
                    }
                }
                if (res < INT_MAX) {
                    ans[i][j] = res;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m-k+1)
	arr := make([]int, k*k)
	for i := range ans {
		ans[i] = make([]int, n-k+1)
		for j := range ans[i] {
			a := arr[:0] // 避免反复 make
			for _, row := range grid[i : i+k] {
				a = append(a, row[j:j+k]...)
			}
			slices.Sort(a)

			res := math.MaxInt
			for p := 1; p < len(a); p++ {
				if a[p] > a[p-1] {
					res = min(res, a[p]-a[p-1])
				}
			}
			if res < math.MaxInt {
				ans[i][j] = res
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((m-k)(n-k)k^2\log k)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(k^2)$。返回值不计入。

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
