## 方法一：前缀和

计算两个前缀和数组：

- 定义数组 $c$，其中 $c[i] = \textit{prices}[i]\cdot \textit{strategy}[i]$。计算 $c$ 的前缀和，记作 $\textit{sum}$。
- 计算 $\textit{prices}$ 的前缀和，记作 $\textit{sumSell}$。
- 关于前缀和数组的详细定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

如果不修改，答案为 $\textit{sum}[n]$。

如果修改，枚举修改子数组 $[i-k,i-1]$。修改后的利润由三部分组成：

1. $[0,i-k-1]$ 的 $\textit{prices}[i]\cdot \textit{strategy}[i]$ 之和，即 $\textit{sum}[i-k]$。
2. $[i,n-1]$ 的 $\textit{prices}[i]\cdot \textit{strategy}[i]$ 之和，即 $\textit{sum}[n] - \textit{sum}[i]$。
3. $[i-k/2,i-1]$ 的 $\textit{prices}[i]$ 之和，即 $\textit{sumSell}[i] - \textit{sumSell}[i-k/2]$。

总和为

$$
\textit{sum}[i-k] + \textit{sum}[n] - \textit{sum}[i] + \textit{sumSell}[i] - \textit{sumSell}[i-k/2]
$$

用上式更新答案的最大值。

[本题视频讲解](https://www.bilibili.com/video/BV1kTYyzwEDD/?t=29m23s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxProfit(self, prices: List[int], strategy: List[int], k: int) -> int:
        n = len(prices)
        s = list(accumulate((p * s for p, s in zip(prices, strategy)), initial=0))
        s_sell = list(accumulate(prices, initial=0))

        # 修改一次
        ans = max(s[i - k] + s[n] - s[i] + s_sell[i] - s_sell[i - k // 2] for i in range(k, n + 1))
        return max(ans, s[n])  # 不修改
```

```java [sol-Java]
class Solution {
    public long maxProfit(int[] prices, int[] strategy, int k) {
        int n = prices.length;
        long[] sum = new long[n + 1];
        long[] sumSell = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + prices[i] * strategy[i];
            sumSell[i + 1] = sumSell[i] + prices[i];
        }

        long ans = sum[n]; // 不修改
        for (int i = k; i <= n; i++) {
            long res = sum[i - k] + sum[n] - sum[i] + sumSell[i] - sumSell[i - k / 2];
            ans = Math.max(ans, res);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxProfit(vector<int>& prices, vector<int>& strategy, int k) {
        int n = prices.size();
        vector<long long> sum(n + 1), sum_sell(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + prices[i] * strategy[i];
            sum_sell[i + 1] = sum_sell[i] + prices[i];
        }

        long long ans = sum[n]; // 不修改
        for (int i = k; i <= n; i++) {
            long long res = sum[i - k] + sum[n] - sum[i] + sum_sell[i] - sum_sell[i - k / 2];
            ans = max(ans, res);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	sum := make([]int, n+1)
	sumSell := make([]int, n+1)
	for i, p := range prices {
		sum[i+1] = sum[i] + p*strategy[i]
		sumSell[i+1] = sumSell[i] + p
	}

	ans := sum[n] // 不修改
	for i := k; i <= n; i++ {
		res := sum[i-k] + sum[n] - sum[i] + sumSell[i] - sumSell[i-k/2]
		ans = max(ans, res)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：定长滑动窗口

**前置知识**：[定长滑动窗口](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)。

设不修改时的利润为 $\textit{total}$。修改后，利润（相比不修改）增加了 $\textit{sum}$。所有窗口的 $\textit{sum}$ 的最大值为 $\textit{maxSum}$。那么答案为 $\textit{total} + \max(\textit{maxSum},0)$。这里可能出现 $\textit{maxSum} < 0$ 的情况，此时不修改更好，也就是与 $0$ 取最大值。

对于价格 $\textit{p}$，如果修改前策略是 $x$，修改后策略是 $y$，那么利润增加了 $p\cdot(y-x)$。比如原来买入，现在持有（不买入），那么利润增加了 $p\cdot (0 - (-1)) = p$。又比如原来买入，现在卖出，那么利润增加了 $p\cdot (1 - (-1)) = 2p$。

下面来计算每个窗口的 $\textit{sum}$。考察从 $[i-k,i-1]$ 向右滑到 $[i-k+1,i]$，$\textit{sum}$ 如何变化。

先看窗口 $[i-k,i-1]$ 的 $\textit{sum}$，分为左右两部分：

1. 左半为 $[i-k,i-k/2-1]$。修改前的策略为 $\textit{strategy}[j]$，修改后的策略为 $0$，所以利润增加了 $\textit{prices}[j]\cdot (-\textit{strategy}[j])$ 之和，其中 $j$ 在左半中。
2. 右半为 $[i-k/2,i-1]$。修改前的策略为 $\textit{strategy}[j]$，修改后的策略为 $1$，所以利润增加了 $\textit{prices}[j]\cdot (1-\textit{strategy}[j])$ 之和，其中 $j$ 在右半中。

当窗口向右滑动时，有三个位置的元素发生了变化：

1. $i$ 进入窗口（在右半），$\textit{sum}$ 增加了 $\textit{prices}[i]\cdot (1-\textit{strategy}[i])$。
2. 下标为 $i-k/2$ 的元素从右半移到左半，交易策略从 $1$ 变成 $0$，所以 $\textit{sum}$ 减少了 $\textit{prices}[i-k/2]$。
3. $i-k$ 离开窗口（离开前在左半），$\textit{sum}$ 减少了 $\textit{prices}[i-k]\cdot (-\textit{strategy}[i-k])$。

## 写法一

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxProfit(self, prices: List[int], strategy: List[int], k: int) -> int:
        total = s = 0
        # 计算第一个窗口的 s
        for p, st in zip(prices[:k // 2], strategy[:k // 2]):
            total += p * st
            s -= p * st
        for p, st in zip(prices[k // 2: k], strategy[k // 2: k]):
            total += p * st
            s += p * (1 - st)

        max_s = max(s, 0)
        # 向右滑动，计算后续窗口的 s
        for i in range(k, len(prices)):
            p, st = prices[i], strategy[i]
            total += p * st
            s += p * (1 - st) - prices[i - k // 2] + prices[i - k] * strategy[i - k]
            max_s = max(max_s, s)
        return total + max_s
```

```java [sol-Java]
class Solution {
    public long maxProfit(int[] prices, int[] strategy, int k) {
        long total = 0, sum = 0;
        // 计算第一个窗口的 sum
        for (int i = 0; i < k / 2; i++) {
            int p = prices[i], s = strategy[i];
            total += p * s;
            sum -= p * s;
        }
        for (int i = k / 2; i < k; i++) {
            int p = prices[i], s = strategy[i];
            total += p * s;
            sum += p * (1 - s);
        }

        long maxSum = Math.max(sum, 0);
        // 向右滑动，计算后续窗口的 sum
        for (int i = k; i < prices.length; i++) {
            int p = prices[i], s = strategy[i];
            total += p * s;
            sum += p * (1 - s) - prices[i - k / 2] + prices[i - k] * strategy[i - k];
            maxSum = Math.max(maxSum, sum);
        }
        return total + maxSum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxProfit(vector<int>& prices, vector<int>& strategy, int k) {
        long long total = 0, sum = 0;
        // 计算第一个窗口的 sum
        for (int i = 0; i < k / 2; i++) {
            int p = prices[i], s = strategy[i];
            total += p * s;
            sum -= p * s;
        }
        for (int i = k / 2; i < k; i++) {
            int p = prices[i], s = strategy[i];
            total += p * s;
            sum += p * (1 - s);
        }

        long long max_sum = max(sum, 0LL);
        // 向右滑动，计算后续窗口的 sum
        for (int i = k; i < prices.size(); i++) {
            int p = prices[i], s = strategy[i];
            total += p * s;
            sum += p * (1 - s) - prices[i - k / 2] + prices[i - k] * strategy[i - k];
            max_sum = max(max_sum, sum);
        }
        return total + max_sum;
    }
};
```

```go [sol-Go]
func maxProfit(prices, strategy []int, k int) int64 {
	total, sum := 0, 0
	// 计算第一个窗口的 sum
	for i := range k / 2 {
		p, s := prices[i], strategy[i]
		total += p * s
		sum -= p * s
	}
	for i := k / 2; i < k; i++ {
		p, s := prices[i], strategy[i]
		total += p * s
		sum += p * (1 - s)
	}

	maxSum := max(sum, 0)
	// 向右滑动，计算后续窗口的 sum
	for i := k; i < len(prices); i++ {
		p, s := prices[i], strategy[i]
		total += p * s
		sum += p*(1-s) - prices[i-k/2] + prices[i-k]*strategy[i-k]
		maxSum = max(maxSum, sum)
	}
	return int64(total + maxSum)
}
```

## 写法二

```py [sol-Python3]
class Solution:
    def maxProfit(self, prices: List[int], strategy: List[int], k: int) -> int:
        total = max_s = s = 0
        for i, (p, st) in enumerate(zip(prices, strategy)):
            total += p * st  # 不修改时的最大利润

            # 1. 下标为 i 的元素入右半，交易策略从 st 变成 1
            s += p * (1 - st)  # 修改带来的额外利润

            if i < k - 1:  # 尚未形成第一个窗口
                # 在下一轮循环中，下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
                if i >= k // 2 - 1:  
                    s -= prices[i - k // 2 + 1]
                continue

            # 2. 更新
            max_s = max(max_s, s)  # 修改带来的最大额外利润

            # 3. 出，为下一个窗口做准备
            # 下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
            # 下标为 i-k+1 的元素从左半离开窗口，交易策略从 0 恢复为 strategy[i-k+1]
            s -= prices[i - k // 2 + 1] - prices[i - k + 1] * strategy[i - k + 1]

        return total + max_s
```

```java [sol-Java]
class Solution {
    public long maxProfit(int[] prices, int[] strategy, int k) {
        long total = 0, maxSum = 0, sum = 0;
        for (int i = 0; i < prices.length; i++) {
            int p = prices[i], s = strategy[i];
            total += p * s; // 不修改时的最大利润

            // 1. 下标为 i 的元素入右半，交易策略从 s 变成 1
            sum += p * (1 - s);

            if (i < k - 1) { // 尚未形成第一个窗口
                // 在下一轮循环中，下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
                if (i >= k / 2 - 1) {
                    sum -= prices[i - k / 2 + 1];
                }
                continue;
            }

            // 2. 更新
            maxSum = Math.max(maxSum, sum); // 修改带来的最大额外利润

            // 3. 出，为下一个窗口做准备
            // 下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
            // 下标为 i-k+1 的元素从左半离开窗口，交易策略从 0 恢复为 strategy[i-k+1]
            sum -= prices[i - k / 2 + 1] - prices[i - k + 1] * strategy[i - k + 1];
        }

        return total + maxSum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxProfit(vector<int>& prices, vector<int>& strategy, int k) {
        long long total = 0, max_sum = 0, sum = 0;
        for (int i = 0; i < prices.size(); i++) {
            int p = prices[i], s = strategy[i];
            total += p * s; // 不修改时的最大利润

            // 1. 下标为 i 的元素入右半，交易策略从 s 变成 1
            sum += p * (1 - s);

            if (i < k - 1) { // 尚未形成第一个窗口
                // 在下一轮循环中，下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
                if (i >= k / 2 - 1) {
                    sum -= prices[i - k / 2 + 1];
                }
                continue;
            }

            // 2. 更新
            max_sum = max(max_sum, sum); // 修改带来的最大额外利润

            // 3. 出，为下一个窗口做准备
            // 下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
            // 下标为 i-k+1 的元素从左半离开窗口，交易策略从 0 恢复为 strategy[i-k+1]
            sum -= prices[i - k / 2 + 1] - prices[i - k + 1] * strategy[i - k + 1];
        }

        return total + max_sum;
    }
};
```

```go [sol-Go]
func maxProfit(prices, strategy []int, k int) int64 {
	var total, maxSum, sum int
	for i, p := range prices {
		s := strategy[i]
		total += p * s // 不修改时的最大利润

		// 1. 下标为 i 的元素入右半，交易策略从 s 变成 1
		sum += p * (1 - s)

		if i < k-1 { // 尚未形成第一个窗口
			// 在下一轮循环中，下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
			if i >= k/2-1 {
				sum -= prices[i-k/2+1]
			}
			continue
		}

		// 2. 更新
		maxSum = max(maxSum, sum) // 修改带来的最大额外利润

		// 3. 出，为下一个窗口做准备
		// 下标为 i-k/2+1 的元素从右半移到左半，交易策略从 1 变成 0
		// 下标为 i-k+1 的元素从左半离开窗口，交易策略从 0 恢复为 strategy[i-k+1]
		sum -= prices[i-k/2+1] - prices[i-k+1]*strategy[i-k+1]
	}

	return int64(total + maxSum)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 数据结构题单的「**一、前缀和**」。
2. 滑动窗口题单的「**一、定长滑动窗口**」。

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
