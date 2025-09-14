## 寻找子问题

考虑最后一个数 $x = \textit{nums}[n-1]$ 选或不选。

- 不选：问题变成在 $[0,n-2]$ 中能选出多少个稳定子序列。
- 选，那么有如下情况：
  - $x$ 单独组成一个长为 $1$ 的子序列。
  - 如果 $x$ 是偶数，那么 $x$ 可以添加到末尾为奇数的子序列的后面。问题变成在 $[0,n-2]$ 中能选出多少个末尾为奇数的稳定子序列。
  - 如果 $x$ 是偶数，那么 $x$ 可以添加到后两个数为奇数、偶数的子序列的后面。问题变成在 $[0,n-2]$ 中能选出多少个后两个数为奇数、偶数的稳定子序列。
  - 对于 $x$ 是奇数的情况，同理。

这些问题都是**和原问题相似的、规模更小的子问题**。

## 状态定义和状态转移方程

根据上面的讨论，定义 $f[i+1][x][j]$ 表示满足如下约束的稳定子序列的个数（$+1$ 是为了用 $f[0]$ 表示空前缀）：

- 元素下标在 $[0,i]$ 中。
- 子序列末尾元素的奇偶性为 $x$，其中 $x=0$ 或者 $1$。
- 子序列末尾恰好有连续 $j+1$ 个奇偶性都为 $x$ 的数。恰好的意思是，倒数第 $j+2$ 个数的奇偶性与 $x$ 相反，即 $x$ 异或 $1$，下文用 $x\oplus 1$ 表示。

设 $\textit{nums}[i]$ 的奇偶性为 $x$。考虑 $\textit{nums}[i]$ 选或不选。

- 不选：问题变成在 $[0,i-1]$ 中的满足 $x$ 和 $j$ 的稳定子序列的个数，即 $f[i][x][j]$。
- 选：有三种情况：
  - $\textit{nums}[i]$ 单独组成一个长为 $1$ 的稳定子序列，个数为 $1$。
  - 当 $j=0$ 时，$\textit{nums}[i]$ 只能添加到末尾元素奇偶性为 $x\oplus 1$ 的稳定子序列的后面，从 $f[i][x\oplus][0] + f[i][x\oplus][1]$ 转移过来。比如 $x$ 是偶数，那么可以添加到末尾恰好有 $1$ 个或者 $2$ 个奇数的稳定子序列的后面。
  - 当 $j=1$ 时，$\textit{nums}[i]$ 只能添加到末尾元素奇偶性为 $x$，且末尾恰好有一个奇偶性为 $x$ 的稳定子序列的后面，从 $f[i][x][0]$ 转移过来。比如 $x$ 是偶数，那么可以添加到末尾恰好有 $1$ 偶数的稳定子序列的后面。
 
累加得

$$
\begin{aligned}
f[i+1][x][0] &= f[i][x][0] + f[i][x\oplus 1][0] + f[i][x\oplus 1][1] + 1     \\
f[i+1][x][1] &= f[i][x][1] + f[i][x][0]     \\
\end{aligned}
$$

初始值：$f[0][x][j] = 0$。

答案：$f[n][0][0] + f[n][0][1] + f[n][1][0] + f[n][1][1]$。

代码实现时，$f$ 的第一个维度可以去掉。去掉后，由于 $f[x][1]$ 会用到 $f[x][0]$，所以要先更新 $f[x][1]$，再更新 $f[x][0]$。

注意取模。为什么可以在计算中途取模？见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1TBpczdE8P/?t=24m20s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countStableSubsequences(self, nums: list[int]) -> int:
        MOD = 1_000_000_007
        f = [[0, 0], [0, 0]]
        for x in nums:
            x %= 2
            f[x][1] = (f[x][1] + f[x][0]) % MOD
            f[x][0] = (f[x][0] + f[x ^ 1][0] + f[x ^ 1][1] + 1) % MOD
        return (f[0][0] + f[0][1] + f[1][0] + f[1][1]) % MOD
```

```java [sol-Java]
class Solution {
    public int countStableSubsequences(int[] nums) {
        final int MOD = 1_000_000_007;
        long[][] f = new long[2][2];
        for (int x : nums) {
            x %= 2;
            f[x][1] = (f[x][1] + f[x][0]) % MOD;
            f[x][0] = (f[x][0] + f[x ^ 1][0] + f[x ^ 1][1] + 1) % MOD;
        }
        return (int) ((f[0][0] + f[0][1] + f[1][0] + f[1][1]) % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countStableSubsequences(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        long long f[2][2]{};
        for (int x : nums) {
            x %= 2;
            f[x][1] = (f[x][1] + f[x][0]) % MOD;
            f[x][0] = (f[x][0] + f[x ^ 1][0] + f[x ^ 1][1] + 1) % MOD;
        }
        return (f[0][0] + f[0][1] + f[1][0] + f[1][1]) % MOD;
    }
};
```

```go [sol-Go]
func countStableSubsequences(nums []int) int {
	const mod = 1_000_000_007
	f := [2][2]int{}
	for _, x := range nums {
		x %= 2
		f[x][1] = (f[x][1] + f[x][0]) % mod
		f[x][0] = (f[x][0] + f[x^1][0] + f[x^1][1] + 1) % mod
	}
	return (f[0][0] + f[0][1] + f[1][0] + f[1][1]) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度，$k=2$ 表示最多 $k$ 个连续元素的奇偶性相同。
- 空间复杂度：$\mathcal{O}(k)$。

## 专题训练

见下面动态规划题单的「**§7.4 合法子序列 DP**」。

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
