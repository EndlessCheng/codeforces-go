为方便计算，考虑从上往下移动。

在移动的过程中，我们需要知道：

- 当前位置 $(i,j)$。
- 上一步的状态：是在 $i-1$ 行，还是在 $i$ 行？

定义 $f[i][j]$ 表示从 $0$ 行移动到 $(i,j)$，且上一步在 $i-1$ 行的方案数。

定义 $g[i][j]$ 表示从 $0$ 行移动到 $(i,j)$，且上一步也在 $i$ 行的方案数。

从 $0$ 行移动到 $(i,j)$ 的方案数为

$$
f[i][j] + g[i][j]
$$

如果 $(i,j)$ 被阻塞，那么上式为 $0$。

设上一步在 $(i-1,k)$，那么必须满足

$$
(k-j)^2 + 1 \le d^2
$$

即 

$$
|k-j|\le \left\lfloor \sqrt {d^2-1} \right\rfloor = d-1
$$

> 注：题目保证 $d\ge 1$，所以 $(d-1)^2 = d^2-2d+1\le d^2-1$。

对于 $f[i][j]$，枚举上一步在 $(i-1,k)$，从 $0$ 行移动到 $(i-1,k)$ 的方案数为

$$
f[i-1][k] + g[i-1][k]
$$

累加得

$$
f[i][j] = \sum_{k=j-d+1}^{j+d-1} f[i-1][k] + g[i-1][k]
$$

上式可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

对于 $g[i][j]$，枚举上一步在 $(i,k)$，那么有

$$
g[i][j] = \left(\sum_{k=j-d}^{j+d} f[i][k]\right) - f[i][j]
$$

注意不能原地不动，所以要减去 $f[i][j]$。

上式也可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

初始值：$f[0][j] = 1$。

答案：$\displaystyle\sum_{j=0}^{m-1} f[n-1][j] + g[n-1][j]$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def numberOfRoutes(self, grid: List[str], d: int) -> int:
        MOD = 1_000_000_007
        m = len(grid[0])
        sum_f = [0] * (m + 1)
        sum_fg = [0] * (m + 1)

        for i, row in enumerate(grid):
            # 从 i-1 行移动到 i 行的方案数
            f = [0] * m
            for j, ch in enumerate(row):
                if ch == '#':
                    continue
                if i == 0:  # 第一行（起点）
                    f[j] = 1  # DP 初始值
                else:
                    f[j] = sum_fg[min(j + d, m)] - sum_fg[max(j - d + 1, 0)]

            # f 的前缀和
            for j, v in enumerate(f):
                sum_f[j + 1] = (sum_f[j] + v) % MOD

            # 从 i 行移动到 i 行的方案数
            g = [0] * m
            for j, ch in enumerate(row):
                if ch == '#':
                    continue
                # 不能原地不动，减去 f[j]
                g[j] = sum_f[min(j + d + 1, m)] - sum_f[max(j - d, 0)] - f[j]

            # f[j] + g[j] 的前缀和
            for j, (fj, gj) in enumerate(zip(f, g)):
                sum_fg[j + 1] = (sum_fg[j] + fj + gj) % MOD

        return sum_fg[m]
```

```java [sol-Java]
class Solution {
    public int numberOfRoutes(String[] grid, int d) {
        final int MOD = 1_000_000_007;
        int m = grid[0].length();
        long[] sumF = new long[m + 1];
        long[] sum = new long[m + 1];

        for (int i = 0; i < grid.length; i++) {
            char[] row = grid[i].toCharArray();
            // 从 i-1 行移动到 i 行的方案数
            long[] f = new long[m];
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    continue;
                }
                if (i == 0) { // 第一行（起点）
                    f[j] = 1; // DP 初始值
                } else {
                    f[j] = sum[Math.min(j + d, m)] - sum[Math.max(j - d + 1, 0)];
                }
            }

            // f 的前缀和
            for (int j = 0; j < m; j++) {
                sumF[j + 1] = (sumF[j] + f[j]) % MOD;
            }

            // 从 i 行移动到 i 行的方案数
            long[] g = new long[m];
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    continue;
                }
                // 不能原地不动，减去 f[j]
                g[j] = sumF[Math.min(j + d + 1, m)] - sumF[Math.max(j - d, 0)] - f[j];
            }

            // f[j] + g[j] 的前缀和
            for (int j = 0; j < m; j++) {
                sum[j + 1] = (sum[j] + f[j] + g[j]) % MOD;
            }
        }

        return (int) ((sum[m] + MOD) % MOD); // +MOD 保证结果非负
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfRoutes(vector<string>& grid, int d) {
        constexpr int MOD = 1'000'000'007;
        int m = grid[0].size();
        vector<long long> sum_f(m + 1);
        vector<long long> sum(m + 1);

        for (int i = 0; i < grid.size(); i++) {
            auto& row = grid[i];
            // 从 i-1 行移动到 i 行的方案数
            vector<long long> f(m);
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    continue;
                }
                if (i == 0) { // 第一行（起点）
                    f[j] = 1; // DP 初始值
                } else {
                    f[j] = sum[min(j + d, m)] - sum[max(j - d + 1, 0)];
                }
            }

            // f 的前缀和
            for (int j = 0; j < m; j++) {
                sum_f[j + 1] = (sum_f[j] + f[j]) % MOD;
            }

            // 从 i 行移动到 i 行的方案数
            vector<long long> g(m);
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    continue;
                }
                // 不能原地不动，减去 f[j]
                g[j] = sum_f[min(j + d + 1, m)] - sum_f[max(j - d, 0)] - f[j];
            }

            // f[j] + g[j] 的前缀和
            for (int j = 0; j < m; j++) {
                sum[j + 1] = (sum[j] + f[j] + g[j]) % MOD;
            }
        }

        return (sum[m] + MOD) % MOD; // +MOD 保证结果非负
    }
};
```

```go [sol-Go]
func numberOfRoutes(grid []string, d int) int {
	const mod = 1_000_000_007
	m := len(grid[0])
	sumF := make([]int, m+1)
	sum := make([]int, m+1)

	for i, row := range grid {
		// 从 i-1 行移动到 i 行的方案数
		f := make([]int, m)
		for j, ch := range row {
			if ch == '#' {
				continue
			}
			if i == 0 { // 第一行（起点）
				f[j] = 1 // DP 初始值
			} else {
				f[j] = sum[min(j+d, m)] - sum[max(j-d+1, 0)]
			}
		}

		// f 的前缀和
		for j, v := range f {
			sumF[j+1] = (sumF[j] + v) % mod
		}

		// 从 i 行移动到 i 行的方案数
		g := make([]int, m)
		for j, ch := range row {
			if ch == '#' {
				continue
			}
			// 不能原地不动，减去 f[j]
			g[j] = sumF[min(j+d+1, m)] - sumF[max(j-d, 0)] - f[j]
		}

		// f[j] + g[j] 的前缀和
		for j, fj := range f {
			sum[j+1] = (sum[j] + fj + g[j]) % mod
		}
	}

	return (sum[m] + mod) % mod // +mod 保证结果非负
}
```

## 优化

计算 $f$ 的同时，计算 $f$ 的前缀和。

计算 $g$ 的同时，计算 $f[j]+g[j]$ 的前缀和。

```py [sol-Python3]
# 手写 min max 更快
min = lambda a, b: b if b < a else a
max = lambda a, b: b if b > a else a

class Solution:
    def numberOfRoutes(self, grid: List[str], d: int) -> int:
        MOD = 1_000_000_007
        m = len(grid[0])
        sum_f = [0] * (m + 1)
        sum_fg = [0] * (m + 1)

        for i, row in enumerate(grid):
            # f 的前缀和
            for j, ch in enumerate(row):
                if ch == '#':
                    sum_f[j + 1] = sum_f[j]
                elif i == 0:  # 第一行（起点）
                    sum_f[j + 1] = sum_f[j] + 1  # DP 初始值
                else:
                    sum_f[j + 1] = (sum_f[j] + sum_fg[min(j + d, m)] - sum_fg[max(j - d + 1, 0)]) % MOD

            # f[j] + g[j] 的前缀和
            for j, ch in enumerate(row):
                if ch == '#':
                    sum_fg[j + 1] = sum_fg[j]
                else:
                    # -f[j] 和 +f[j] 抵消了
                    sum_fg[j + 1] = (sum_fg[j] + sum_f[min(j + d + 1, m)] - sum_f[max(j - d, 0)]) % MOD

        return sum_fg[m]
```

```java [sol-Java]
class Solution {
    public int numberOfRoutes(String[] grid, int d) {
        final int MOD = 1_000_000_007;
        int m = grid[0].length();
        long[] sumF = new long[m + 1];
        long[] sum = new long[m + 1];

        for (int i = 0; i < grid.length; i++) {
            char[] row = grid[i].toCharArray();
            // f 的前缀和
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    sumF[j + 1] = sumF[j];
                } else if (i == 0) { // 第一行（起点）
                    sumF[j + 1] = sumF[j] + 1; // DP 初始值
                } else {
                    sumF[j + 1] = (sumF[j] + sum[Math.min(j + d, m)] - sum[Math.max(j - d + 1, 0)]) % MOD;
                }
            }

            // f[j] + g[j] 的前缀和
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    sum[j + 1] = sum[j];
                } else {
                    // -f[j] 和 +f[j] 抵消了
                    sum[j + 1] = (sum[j] + sumF[Math.min(j + d + 1, m)] - sumF[Math.max(j - d, 0)]) % MOD;
                }
            }
        }

        return (int) ((sum[m] + MOD) % MOD); // +MOD 保证结果非负
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfRoutes(vector<string>& grid, int d) {
        constexpr int MOD = 1'000'000'007;
        int m = grid[0].size();
        vector<long long> sum_f(m + 1);
        vector<long long> sum(m + 1);

        for (int i = 0; i < grid.size(); i++) {
            auto& row = grid[i];
            // f 的前缀和
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    sum_f[j + 1] = sum_f[j];
                } else if (i == 0) { // 第一行（起点）
                    sum_f[j + 1] = sum_f[j] + 1; // DP 初始值
                } else {
                    sum_f[j + 1] = (sum_f[j] + sum[min(j + d, m)] - sum[max(j - d + 1, 0)]) % MOD;
                }
            }

            // f[j] + g[j] 的前缀和
            for (int j = 0; j < m; j++) {
                if (row[j] == '#') {
                    sum[j + 1] = sum[j];
                } else {
                    // -f[j] 和 +f[j] 抵消了
                    sum[j + 1] = (sum[j] + sum_f[min(j + d + 1, m)] - sum_f[max(j - d, 0)]) % MOD;
                }
            }
        }

        return (sum[m] + MOD) % MOD; // +MOD 保证结果非负
    }
};
```

```go [sol-Go]
func numberOfRoutes(grid []string, d int) int {
	const mod = 1_000_000_007
	m := len(grid[0])
	sumF := make([]int, m+1)
	sum := make([]int, m+1)

	for i, row := range grid {
		// f 的前缀和
		for j, ch := range row {
			if ch == '#' {
				sumF[j+1] = sumF[j]
			} else if i == 0 { // 第一行（起点）
				sumF[j+1] = sumF[j] + 1 // DP 初始值
			} else {
				sumF[j+1] = (sumF[j] + sum[min(j+d, m)] - sum[max(j-d+1, 0)]) % mod
			}
		}

		// f[j] + g[j] 的前缀和
		for j, ch := range row {
			if ch == '#' {
				sum[j+1] = sum[j]
			} else {
				// -f[j] 和 +f[j] 抵消了
				sum[j+1] = (sum[j] + sumF[min(j+d+1, m)] - sumF[max(j-d, 0)]) % mod
			}
		}
	}

	return (sum[m] + mod) % mod // +mod 保证结果非负
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 和 $m$ 分别是 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$。

## 专题训练

见下面动态规划题单的「**§11.1 前缀和优化 DP**」。

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
