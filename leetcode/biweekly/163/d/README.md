如果没有传送，本题就是 [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)。注意本题不计入起点的值。

接着 [64 题我的题解](https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/) 继续讲。

在有传送的情况下，可以用一个额外的维度表示传送次数。定义 $f[t][i+1][j+1]$ 表示在使用**恰好** $t$ 次传送的情况下，从左上角 $(0,0)$ 到 $(i,j)$ 的最小总成本。

考虑转移来源，即我们是从哪个格子移动到 $(i,j)$ 的。

- 普通移动：从 $(i,j-1)$ 和 $(i-1,j)$ 移动到 $(i,j)$。转移来源分别为 $f[t][i+1][j]$ 和 $f[t][i][j+1]$。
- 传送：设 $x = \textit{grid}[i][j]$，我们可以从格子值 $\ge x$ 的任意格子传送到 $(i,j)$。转移来源为 $f[t-1][i'+1][j'+1]$，满足 $\textit{grid}[i'][j']\ge x$。如何快速得到这些 $f[t-1][i'+1][j'+1]$ 的最小值？
   - 定义 $\textit{sufMinF}_{t-1}[x]$ 表示满足 $\textit{grid}[i][j]\ge x$ 的 $f[t-1][i+1][j+1]$ 的最小值。
   - 在计算完 $f[t-1][i+1][j+1]$ 后，把格子值 $x=\textit{grid}[i][j]$ 及其对应的状态值 $f[t-1][i+1][j+1]$ 保存到一个数组 $\textit{minF}$ 中，其中 $\textit{minF}[x]$ 表示格子值为 $x$ 的最小状态值（如果不存在则为 $\infty$）。然后倒序遍历 $\textit{minF}$，计算后缀最小值，即为 $\textit{sufMinF}_{t-1}$。

状态转移方程为

$$
f[t][i+1][j+1] = \min(f[t][i+1][j]  + x, f[t][i][j+1]  + x, \textit{sufMinF}_{t-1}[x])
$$

其中 $x = \textit{grid}[i][j]$。

初始值同 64 题。

答案为 $f[k][m-1][n-1]$。虽然题目要求使用「至多」$k$ 次传送，但由于我们可以原地传送，所以传送的次数越多，总成本是不会增大的。所以「至多」$k$ 次传送等于「恰好」$k$ 次传送。

代码实现时，$f$ 数组的前两个维度可以优化掉。

具体请看 [视频讲解](https://www.bilibili.com/video/BV191YCzjEvc/?t=15m04s)，欢迎点赞关注~

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minCost(self, grid: List[List[int]], k: int) -> int:
        n = len(grid[0])
        mx = max(map(max, grid))

        suf_min_f = [inf] * (mx + 2)
        for _ in range(k + 1):
            min_f = [inf] * (mx + 1)

            # 64. 最小路径和（空间优化写法）
            f = [inf] * (n + 1)
            f[1] = -grid[0][0]  # 起点的成本不算
            for row in grid:
                for j, x in enumerate(row):
                    f[j + 1] = min(min(f[j], f[j + 1]) + x, suf_min_f[x])
                    min_f[x] = min(min_f[x], f[j + 1])
   
            # 计算 min_f 的后缀最小值
            for i in range(mx, -1, -1):
                suf_min_f[i] = min(suf_min_f[i + 1], min_f[i])

        return f[n]
```

```java [sol-Java]
class Solution {
    public int minCost(int[][] grid, int k) {
        int n = grid[0].length;
        int mx = 0;
        for (int[] row : grid) {
            for (int x : row) {
                mx = Math.max(mx, x);
            }
        }

        int[] sufMinF = new int[mx + 2];
        Arrays.fill(sufMinF, Integer.MAX_VALUE);
        int[] minF = new int[mx + 1];
        int[] f = new int[n + 1];

        for (int t = 0; t <= k; t++) {
            Arrays.fill(minF, Integer.MAX_VALUE);

            // 64. 最小路径和（空间优化写法）
            Arrays.fill(f, Integer.MAX_VALUE / 2);
            f[1] = -grid[0][0]; // 起点的成本不算
            for (int[] row : grid) {
                for (int j = 0; j < row.length; j++) {
                    int x = row[j];
                    f[j + 1] = Math.min(Math.min(f[j], f[j + 1]) + x, sufMinF[x]);
                    minF[x] = Math.min(minF[x], f[j + 1]);
                }
            }

            // 计算 minF 的后缀最小值
            for (int i = mx; i >= 0; i--) {
                sufMinF[i] = Math.min(sufMinF[i + 1], minF[i]);
            }
        }

        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<vector<int>>& grid, int k) {
        int n = grid[0].size();
        int mx = 0;
        for (auto& row : grid) {
            mx = max(mx, ranges::max(row));
        }

        vector<int> suf_min_f(mx + 2, INT_MAX);
        vector<int> min_f(mx + 1);
        vector<int> f(n + 1);

        for (int t = 0; t <= k; t++) {
            ranges::fill(min_f, INT_MAX);

            // 64. 最小路径和（空间优化写法）
            ranges::fill(f, INT_MAX / 2);
            f[1] = -grid[0][0]; // 起点的成本不算
            for (auto& row : grid) {
                for (int j = 0; j < row.size(); j++) {
                    int x = row[j];
                    f[j + 1] = min(min(f[j], f[j + 1]) + x, suf_min_f[x]);
                    min_f[x] = min(min_f[x], f[j + 1]);
                }
            }

            // 计算 min_f 的后缀最小值
            for (int i = mx; i >= 0; i--) {
                suf_min_f[i] = min(suf_min_f[i + 1], min_f[i]);
            }
        }

        return f[n];
    }
};
```

```go [sol-Go]
func minCost(grid [][]int, k int) int {
	n := len(grid[0])
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}

	sufMinF := make([]int, mx+2)
	for i := range sufMinF {
		sufMinF[i] = math.MaxInt
	}
	minF := make([]int, mx+1)
	f := make([]int, n+1)

	for range k + 1 {
		for i := range minF {
			minF[i] = math.MaxInt
		}

		// 64. 最小路径和（空间优化写法）
		for i := range f {
			f[i] = math.MaxInt / 2
		}
		f[1] = -grid[0][0] // 起点的成本不算
		for _, row := range grid {
			for j, x := range row {
				f[j+1] = min(f[j]+x, f[j+1]+x, sufMinF[x])
				minF[x] = min(minF[x], f[j+1])
			}
		}

		// 计算 minF 的后缀最小值
		for i := mx; i >= 0; i-- {
			sufMinF[i] = min(sufMinF[i+1], minF[i])
		}
	}

	return f[n]
}
```

## 优化

每次循环我们会计算一遍 $\textit{sufMinF}$。如果发现某次循环没有改变 $\textit{sufMinF}$，那么无论再传送多少次，都不会再改变 $\textit{sufMinF}$ 了，此时我们已经找到了答案。

力扣喜欢出随机数据。测试发现，$m=n=80$，值域 $10^4$ 的随机数据，迭代约 $2.2$ 次就收敛了，然后再循环一次发现收敛，即 $\textit{sufMinF}$ 在循环前后是相同的。所以平均外层循环约 $3.2$ 次就可以退出循环了，而不是循环 $k+1$ 次。

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minCost(self, grid: List[List[int]], k: int) -> int:
        n = len(grid[0])
        mx = max(map(max, grid))

        suf_min_f = [inf] * (mx + 2)
        for _ in range(k + 1):
            min_f = [inf] * (mx + 1)

            # 64. 最小路径和（空间优化写法）
            f = [inf] * (n + 1)
            f[1] = -grid[0][0]  # 起点的成本不算
            for row in grid:
                for j, x in enumerate(row):
                    f[j + 1] = min(min(f[j], f[j + 1]) + x, suf_min_f[x])
                    min_f[x] = min(min_f[x], f[j + 1])
   
            tmp = suf_min_f.copy()
            # 计算 min_f 的后缀最小值
            for i in range(mx, -1, -1):
                suf_min_f[i] = min(suf_min_f[i + 1], min_f[i])
            if suf_min_f == tmp:
                # 收敛了：传送不改变 suf_min_f，那么无论再传送多少次都不会改变 suf_min_f
                break

        return f[n]
```

```java [sol-Java]
class Solution {
    public int minCost(int[][] grid, int k) {
        int n = grid[0].length;
        int mx = 0;
        for (int[] row : grid) {
            for (int x : row) {
                mx = Math.max(mx, x);
            }
        }

        int[] sufMinF = new int[mx + 2];
        Arrays.fill(sufMinF, Integer.MAX_VALUE);
        int[] minF = new int[mx + 1];
        int[] f = new int[n + 1];

        for (int t = 0; t <= k; t++) {
            Arrays.fill(minF, Integer.MAX_VALUE);

            // 64. 最小路径和（空间优化写法）
            Arrays.fill(f, Integer.MAX_VALUE / 2);
            f[1] = -grid[0][0]; // 起点的成本不算
            for (int[] row : grid) {
                for (int j = 0; j < row.length; j++) {
                    int x = row[j];
                    f[j + 1] = Math.min(Math.min(f[j], f[j + 1]) + x, sufMinF[x]);
                    minF[x] = Math.min(minF[x], f[j + 1]);
                }
            }

            int[] tmp = sufMinF.clone();
            // 计算 minF 的后缀最小值
            for (int i = mx; i >= 0; i--) {
                sufMinF[i] = Math.min(sufMinF[i + 1], minF[i]);
            }
            if (Arrays.equals(sufMinF, tmp)) {
                // 收敛了：传送不改变 sufMinF，那么无论再传送多少次都不会改变 sufMinF
                break;
            }
        }

        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minCost(vector<vector<int>>& grid, int k) {
        int n = grid[0].size();
        int mx = 0;
        for (auto& row : grid) {
            mx = max(mx, ranges::max(row));
        }

        vector<int> suf_min_f(mx + 2, INT_MAX);
        vector<int> min_f(mx + 1);
        vector<int> f(n + 1);

        for (int t = 0; t <= k; t++) {
            ranges::fill(min_f, INT_MAX);

            // 64. 最小路径和（空间优化写法）
            ranges::fill(f, INT_MAX / 2);
            f[1] = -grid[0][0]; // 起点的成本不算
            for (auto& row : grid) {
                for (int j = 0; j < row.size(); j++) {
                    int x = row[j];
                    f[j + 1] = min(min(f[j], f[j + 1]) + x, suf_min_f[x]);
                    min_f[x] = min(min_f[x], f[j + 1]);
                }
            }

            auto tmp = suf_min_f;
            // 计算 min_f 的后缀最小值
            for (int i = mx; i >= 0; i--) {
                suf_min_f[i] = min(suf_min_f[i + 1], min_f[i]);
            }
            if (suf_min_f == tmp) {
                // 收敛了：传送不改变 suf_min_f，那么无论再传送多少次都不会改变 suf_min_f
                break;
            }
        }

        return f[n];
    }
};
```

```go [sol-Go]
func minCost(grid [][]int, k int) int {
	n := len(grid[0])
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}

	sufMinF := make([]int, mx+2)
	for i := range sufMinF {
		sufMinF[i] = math.MaxInt
	}
	minF := make([]int, mx+1)
	f := make([]int, n+1)

	for range k + 1 {
		for i := range minF {
			minF[i] = math.MaxInt
		}

		// 64. 最小路径和（空间优化写法）
		for i := range f {
			f[i] = math.MaxInt / 2
		}
		f[1] = -grid[0][0] // 起点的成本不算
		for _, row := range grid {
			for j, x := range row {
				f[j+1] = min(f[j]+x, f[j+1]+x, sufMinF[x])
				minF[x] = min(minF[x], f[j+1])
			}
		}

		tmp := slices.Clone(sufMinF)
		// 计算 minF 的后缀最小值
		for i := mx; i >= 0; i-- {
			sufMinF[i] = min(sufMinF[i+1], minF[i])
		}
		if slices.Equal(sufMinF, tmp) {
			// 收敛了：传送不改变 sufMinF，那么无论再传送多少次都不会改变 sufMinF
			break
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((mn+U)k)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U$ 为 $\textit{grid}[i][j]$ 的最大值。
- 空间复杂度：$\mathcal{O}(n+U)$。

## 专题训练

见下面动态规划题单的「**二、网格图 DP**」和「**§7.6 多维 DP**」。

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
