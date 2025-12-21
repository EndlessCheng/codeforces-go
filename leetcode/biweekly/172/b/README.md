把每个 $x = \textit{nums}[i]$ 看成一个体积为 $(1,x)$，价值为 $x$ 的物品。其中 $(1,x)$ 表示物品的第一类体积为 $1$，第二类体积为 $x$。

我们需要选择一些物品装入背包，使得的第一类体积之和恰好是 $3$，第二类体积之和恰好是 $3$ 的倍数。所以本题是一个「恰好装满型」二维 0-1 背包。

请先完成二维 0-1 背包的经典题目 [474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/)，并阅读 [我的题解](https://leetcode.cn/problems/ones-and-zeroes/solutions/3038333/yi-bu-bu-si-kao-cong-ji-yi-hua-sou-suo-d-lqio/)。

类似 474 题，定义 $f[i+1][j][r]$ 表示在下标 $[0,i]$ 中选元素，选了恰好 $j$ 个数，元素和模 $3$ 为 $r$ 时，所选元素之和的最大值。

讨论 $x = \textit{nums}[i]$ 选或不选：

- 不选 $x$，问题变成在下标 $[0,i-1]$ 中选元素，选了恰好 $j$ 个数，元素和模 $3$ 为 $r$ 时，所选元素之和的最大值，即 $f[i][j][r]$。
- 选 $x$，问题变成在下标 $[0,i-1]$ 中选元素，选了恰好 $j-1$ 个数，元素和模 $3$ 为 $(r-x)\bmod 3$ 时，所选元素之和的最大值，即 $f[i][j-1][(r-x)\bmod 3]$，再加上 $x$（选了 $x$，获得了 $x$ 的价值）。

二者取最大值，得

$$
f[i+1][j][r] = \max(f[i][j][r], f[i][j-1][(r-x)\bmod 3] + x)
$$

初始值 $f[0][0][0] = 0$，其余 $f[0][j][r]=-\infty$。一开始没有选数字，只有 $j=r=0$ 是合法的。

答案为 $f[n][3][0]$，表示从 $\textit{nums}$ 中选了恰好 $3$ 个数，元素和模 $3$ 为 $0$ 时，所选元素之和的最大值。

如果 $f[n][3][0] < 0$，则无解，返回 $0$。

代码实现时，第一个维度可以优化掉，原理见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

本题涉及到中途取模，以及减法的取模，原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

## 写法一：查表法

```py [sol-Python3]
max = lambda a, b: b if b > a else a

class Solution:
    def maximumSum(self, nums: List[int]) -> int:
        K = 3
        MOD = 3
        f = [[-inf] * MOD for _ in range(K + 1)]
        f[0][0] = 0
        for x in nums:
            for j in range(K, 0, -1):
                for r in range(MOD):
                    f[j][r] = max(f[j][r], f[j - 1][(r - x) % MOD] + x)
        return max(f[K][0], 0)
```

```java [sol-Java]
class Solution {
    public int maximumSum(int[] nums) {
        final int K = 3;
        final int MOD = 3;
        int[][] f = new int[K + 1][MOD];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        f[0][0] = 0;
        for (int x : nums) {
            for (int j = K; j > 0; j--) {
                for (int r = 0; r < MOD; r++) {
                    f[j][r] = Math.max(f[j][r], f[j - 1][(r - x % MOD + MOD) % MOD] + x); // 保证取模结果非负
                }
            }
        }
        return Math.max(f[K][0], 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSum(vector<int>& nums) {
        constexpr int K = 3;
        constexpr int MOD = 3;
        vector f(K + 1, vector<int>(MOD, INT_MIN));
        f[0][0] = 0;
        for (int x : nums) {
            for (int j = K; j > 0; j--) {
                for (int r = 0; r < MOD; r++) {
                    f[j][r] = max(f[j][r], f[j - 1][(r - x % MOD + MOD) % MOD] + x); // 保证取模结果非负
                }
            }
        }
        return max(f[K][0], 0);
    }
};
```

```go [sol-Go]
func maximumSum(nums []int) int {
	const K = 3
	const MOD = 3
	f := [K + 1][MOD]int{}
	for i := range f {
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0
	for _, x := range nums {
		for j := K; j > 0; j-- {
			for r := range MOD {
				f[j][r] = max(f[j][r], f[j-1][(r-x%MOD+MOD)%MOD]+x) // 保证取模结果非负
			}
		}
	}
	return max(f[K][0], 0)
}
```

## 写法二：刷表法

用当前状态 $f[i][j][r]$ 去更新其他状态 $f[i+1][j+1][(r+x)\bmod 3]$，即为刷表法。

```py [sol-Python3]
max = lambda a, b: b if b > a else a

class Solution:
    def maximumSum(self, nums: List[int]) -> int:
        K = 3
        MOD = 3
        f = [[-inf] * MOD for _ in range(K + 1)]
        f[0][0] = 0
        for x in nums:
            for j in range(K - 1, -1, -1):
                for r in range(MOD):
                    f[j + 1][(r + x) % MOD] = max(f[j + 1][(r + x) % MOD], f[j][r] + x)
        return max(f[K][0], 0)
```

```java [sol-Java]
class Solution {
    public int maximumSum(int[] nums) {
        final int K = 3;
        final int MOD = 3;
        int[][] f = new int[K + 1][MOD];
        for (int[] row : f) {
            Arrays.fill(row, Integer.MIN_VALUE);
        }
        f[0][0] = 0;
        for (int x : nums) {
            for (int j = K - 1; j >= 0; j--) {
                for (int r = 0; r < MOD; r++) {
                    f[j + 1][(r + x) % MOD] = Math.max(f[j + 1][(r + x) % MOD], f[j][r] + x);
                }
            }
        }
        return Math.max(f[K][0], 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumSum(vector<int>& nums) {
        constexpr int K = 3;
        constexpr int MOD = 3;
        vector f(K + 1, vector<int>(MOD, INT_MIN));
        f[0][0] = 0;
        for (int x : nums) {
            for (int j = K - 1; j >= 0; j--) {
                for (int r = 0; r < MOD; r++) {
                    f[j + 1][(r + x) % MOD] = max(f[j + 1][(r + x) % MOD], f[j][r] + x);
                }
            }
        }
        return max(f[K][0], 0);
    }
};
```

```go [sol-Go]
func maximumSum(nums []int) int {
	const K = 3
	const MOD = 3
	f := [K + 1][MOD]int{}
	for i := range f {
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0
	for _, x := range nums {
		for j := K - 1; j >= 0; j-- {
			for r := range MOD {
				f[j+1][(r+x)%MOD] = max(f[j+1][(r+x)%MOD], f[j][r]+x)
			}
		}
	}
	return max(f[K][0], 0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nKM)$，其中 $n$ 是 $\textit{nums}$ 的长度，$K=3$ 是目标子序列的长度，$M=3$ 是模数（目标子序列之和的因子）。
- 空间复杂度：$\mathcal{O}(KM)$。

## 相似题目

- [474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/)
- [1262. 可被三整除的最大和](https://leetcode.cn/problems/greatest-sum-divisible-by-three/)

## 专题训练

见下面动态规划题单的「**§3.1 0-1 背包**」和「**§7.6 多维 DP**」。

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
