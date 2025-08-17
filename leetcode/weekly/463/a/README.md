## 方法一：前缀和

**前置知识**：[前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

计算两个前缀和数组：

- $\textit{prices}[i]\cdot \textit{strategy}[i]$ 的前缀和，记作 $\textit{sum}$。
- $\textit{prices}[i]$ 的前缀和，记作 $\textit{sumSell}$。

如果不修改，答案为 $\textit{sum}[n]$。

如果修改，枚举修改 $[i-k,i-1]$ 这个子数组，答案由三部分组成：

1. $[0,i-k-1]$ 的 $\textit{prices}[i]\cdot \textit{strategy}[i]$ 之和，即 $\textit{sum}[i-k]$。
2. $[i,n-1]$ 的 $\textit{prices}[i]\cdot \textit{strategy}[i]$ 之和，即 $\textit{sum}[n] - \textit{sum}[i]$。
3. $[i-k/2,i-1]$ 的 $\textit{prices}[i]$ 之和，即 $\textit{sumSell}[i] - \textit{sumSell}[i-k/2]$。

取上述三者之和的最大值，即为修改一次时的答案。与 $\textit{sum}[n]$ 取最大值，即为最终答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

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

考虑修改元素时，在不修改的基础上，**最大增量**是多少。

我们有两个相邻的、长度均为 $k/2$ 的定长滑动窗口：

1. $[i-k,i-k/2-1]$，这里面的元素和增加了 $\textit{prices}[i]\cdot (-\textit{strategy}[i])$ 之和。
2. $[i-k/2,i-1]$，这里面的元素和增加了 $\textit{prices}[i]\cdot (1-\textit{strategy}[i])$ 之和。

当窗口向右滑动时：

1. $\textit{prices}[i]\cdot (1-\textit{strategy}[i])$ 进入第二个窗口。
2. 下标为 $i-k/2$ 的元素从第二个窗口移到一个窗口中，交易策略从 $1$ 变成 $0$，所以元素和减少了 $\textit{prices}[i-k/2]$。
3. $\textit{prices}[i-k]\cdot (-\textit{strategy}[i-k])$ 离开第一个窗口。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxProfit(self, prices: List[int], strategy: List[int], k: int) -> int:
        total = s = 0
        # 计算第一个窗口
        for p, st in zip(prices[:k // 2], strategy[:k // 2]):
            total += p * st
            s -= p * st
        for p, st in zip(prices[k // 2: k], strategy[k // 2: k]):
            total += p * st
            s += p * (1 - st)
        max_s = max(s, 0)

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
        // 计算第一个窗口
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
        // 计算第一个窗口
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
	// 计算第一个窗口
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

	for i := k; i < len(prices); i++ {
		p, s := prices[i], strategy[i]
		total += p * s
		sum += p*(1-s) - prices[i-k/2] + prices[i-k]*strategy[i-k]
		maxSum = max(maxSum, sum)
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
