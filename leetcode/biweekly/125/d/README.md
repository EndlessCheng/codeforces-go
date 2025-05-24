## 性质一

由于一个数异或两次 $k$ 后保持不变，所以对于一条从 $x$ 到 $y$ 的简单路径，我们把路径上的所有边操作后，路径上除了 $x$ 和 $y$ 的其它节点都恰好操作两次，所以只有 $\textit{nums}[x]$ 和 $\textit{nums}[y]$ 都异或了 $k$，其余元素不变。

所以题目中的操作可以作用在**任意两个数**上。我们不需要建树，$\textit{edges}$ 是多余的！

## 性质二

注意到，无论操作多少次，总是有**偶数个**元素异或了 $k$。理由如下：

- 如果我们操作的两个数之前都没有异或过，那么操作后，异或 $k$ 的元素增加了 $2$。
- 如果我们操作的两个数之前都异或过，那么操作后，异或 $k$ 的元素减少了 $2$。
- 如果我们操作的两个数之前一个异或过，另一个没有异或过，那么操作后，异或 $k$ 的元素加一减一，不变。

## 思路

结合这两个性质，问题变成：

- 选择 $\textit{nums}$ 中的**偶数**个元素（长为偶数的子序列），把这些数都异或 $k$ 之后，整个 $\textit{nums}$ 的元素和最大是多少？

这可以用状态机 DP 解决。如果你没有做过状态机 DP 的题目，可以看[【基础算法精讲 21】](https://www.bilibili.com/video/BV1ho4y1W7QK/)。

- 定义 $f[i+1][0]$ 表示选择 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中的偶数个元素异或 $k$ 之后，$\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的元素和的最大值。这里写 $i+1$ 是为了用 $f[0]$ 表示初始值。
- 定义 $f[i+1][1]$ 表示选择 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中的奇数个元素异或 $k$，$\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的元素和的最大值。

设 $x=\textit{nums}[i]$。用**选或不选**思考：

- 不选 $x$ 操作，问题变成选择 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中的偶数/奇数个元素异或 $k$，得到的最大元素和，即 $f[i][0]$ 和 $f[i][1]$：
  - $f[i+1][0] = f[i][0] + x$。
  - $f[i+1][1] = f[i][1] + x$。
- 选 $x$ 操作，问题变成选择 $\textit{nums}[0]$ 到 $\textit{nums}[i-1]$ 中的奇数/偶数个元素异或 $k$，得到的最大元素和，即 $f[i][1]$ 和 $f[i][0]$：
  - $f[i+1][0] = f[i][1] + (x\oplus k)$。
  - $f[i+1][1] = f[i][0] + (x\oplus k)$。

两种情况取最大值，有

$$
\begin{aligned}
f[i+1][0] &= \max(f[i][0] + x, f[i][1] + (x\oplus k))  \\
f[i+1][1] &= \max(f[i][1] + x, f[i][0] + (x\oplus k))  \\
\end{aligned}
$$

初始值 $f[0][0] = 0,\ f[0][1] = -\infty$。用 $-\infty$ 表示不合法的状态，这样计算 $\max$ 的时候，合法状态一定会大于不合法的状态。$f[0][1] = -\infty$ 是因为没有数字，不可能选奇数个数。

答案为 $f[n][0]$。

代码实现时，$f$ 数组的第一个维度可以压缩掉。

[视频讲解](https://www.bilibili.com/video/BV1AU411F7Fp/) 第四题。

```py [sol-Python3]
class Solution:
    def maximumValueSum(self, nums: List[int], k: int, _: List[List[int]]) -> int:
        f0, f1 = 0, -inf
        for x in nums:
            f0, f1 = max(f0 + x, f1 + (x ^ k)), max(f1 + x, f0 + (x ^ k))
        return f0
```

```java [sol-Java]
class Solution {
    public long maximumValueSum(int[] nums, int k, int[][] edges) {
        long f0 = 0;
        long f1 = Long.MIN_VALUE;
        for (int x : nums) {
            long t = Math.max(f1 + x, f0 + (x ^ k));
            f0 = Math.max(f0 + x, f1 + (x ^ k));
            f1 = t;
        }
        return f0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumValueSum(vector<int> &nums, int k, vector<vector<int>> &edges) {
        long long f0 = 0, f1 = LLONG_MIN;
        for (int x : nums) {
            long long t = max(f1 + x, f0 + (x ^ k));
            f0 = max(f0 + x, f1 + (x ^ k));
            f1 = t;
        }
        return f0;
    }
};
```

```go [sol-Go]
func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	f0, f1 := 0, math.MinInt
	for _, x := range nums {
		f0, f1 = max(f0+x, f1+(x^k)), max(f1+x, f0+(x^k))
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
