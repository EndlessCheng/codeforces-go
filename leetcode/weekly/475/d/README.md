## 在何处断环？

设 $M = \textit{nums}[i]$ 是 $\textit{nums}$ 的最大值。

假设我们得到了一个最优划分方案。考察其中包含 $M$ 的子数组，设该子数组的最小值在下标 $j$。

分类讨论：

- 如果 $i\le j$，我们可以把子数组在 $i$ **左边**的元素分给上一个子数组，这不会得到更差的结果。此时 $M$ 在子数组的最左边。
- 如果 $i\ge j$，我们可以把子数组在 $i$ **右边**的元素分给下一个子数组，这不会得到更差的结果。此时 $M$ 在子数组的最右边。

所以一定存在一个最优划分方案，其中 $M$ 在某个子数组的最左边或者最右边。

换句话说，我们可以把环形数组在 $(i-1,i)$ 处断开，或者在 $(i,i+1)$ 处断开，转化成非环形的问题。两种断环方式的计算结果取最大值。

## 非环形数组

对于划分出的一段子数组：

- 如果子数组最小值在最大值左边，可以视作**低买高卖**。
- 如果子数组最大值在最小值左边，可以视作**高借低还**（高位做空，低位平空）。

题目限定至多分出 $k$ 个子数组，这对应着至多 $k$ 次交易。

这和 [3573. 买卖股票的最佳时机 V](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/) **完全一样**，请看 [我的题解](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/solutions/3695611/zhuang-tai-ji-dpzai-188-ti-de-ji-chu-sha-aozb/)。

> **注**：换句话说，3573 题可以更抽象的描述为，划分成至多 $k$ 段子数组的绝对差之和的最大值。

[本题视频讲解](https://www.bilibili.com/video/BV1oskQBLEsY/?t=13m38s)，欢迎点赞关注~

```py [sol-Python3]
# 手写 max 更快
fmax = lambda a, b: b if b > a else a

class Solution:
    # 3573. 买卖股票的最佳时机 V
    def maximumProfit(self, prices: List[int], k: int) -> int:
        f = [[-inf] * 3 for _ in range(k + 2)]
        for j in range(1, k + 2):
            f[j][0] = 0
        for p in prices:
            for j in range(k + 1, 0, -1):
                f[j][0] = fmax(f[j][0], fmax(f[j][1] + p, f[j][2] - p))
                f[j][1] = fmax(f[j][1], f[j - 1][0] - p)
                f[j][2] = fmax(f[j][2], f[j - 1][0] + p)
        return f[-1][0]

    def maximumScore(self, nums: List[int], k: int) -> int:
        max_i = nums.index(max(nums))
        ans1 = self.maximumProfit(nums[max_i:] + nums[:max_i], k)  # nums[max_i] 是第一个数
        ans2 = self.maximumProfit(nums[max_i + 1:] + nums[:max_i + 1], k)  # nums[max_i] 是最后一个数
        return fmax(ans1, ans2)
```

```java [sol-Java]
class Solution {
    public long maximumScore(int[] nums, int k) {
        int n = nums.length;
        int maxI = 0;
        for (int i = 1; i < n; i++) {
            if (nums[i] > nums[maxI]) {
                maxI = i;
            }
        }

        long ans1 = maximumProfit(nums, maxI, maxI + n, k); // nums[maxI] 是第一个数
        long ans2 = maximumProfit(nums, maxI + 1, maxI + 1 + n, k); // nums[maxI] 是最后一个数
        return Math.max(ans1, ans2);
    }

    // 3573. 买卖股票的最佳时机 V
    private long maximumProfit(int[] prices, int l, int r, int k) {
        int n = prices.length;
        long[][] f = new long[k + 2][3];
        for (int j = 1; j <= k + 1; j++) {
            f[j][1] = Long.MIN_VALUE / 2; // 防止溢出
        }
        f[0][0] = Long.MIN_VALUE / 2;
        for (int i = l; i < r; i++) {
            int p = prices[i % n];
            for (int j = k + 1; j > 0; j--) {
                f[j][0] = Math.max(f[j][0], Math.max(f[j][1] + p, f[j][2] - p));
                f[j][1] = Math.max(f[j][1], f[j - 1][0] - p);
                f[j][2] = Math.max(f[j][2], f[j - 1][0] + p);
            }
        }
        return f[k + 1][0];
    }
}
```

```cpp [sol-C++]
class Solution {
    // 3573. 买卖股票的最佳时机 V
    long long maximumProfit(vector<int>& prices, int l, int r, int k) {
        int n = prices.size();
        vector<array<long long, 3>> f(k + 2, {LLONG_MIN / 2, LLONG_MIN / 2, LLONG_MIN / 2});
        for (int j = 1; j <= k + 1; j++) {
            f[j][0] = 0;
        }
        for (int i = l; i < r; i++) {
            int p = prices[i % n];
            for (int j = k + 1; j > 0; j--) {
                f[j][0] = max(f[j][0], max(f[j][1] + p, f[j][2] - p));
                f[j][1] = max(f[j][1], f[j - 1][0] - p);
                f[j][2] = max(f[j][2], f[j - 1][0] + p);
            }
        }
        return f[k + 1][0];
    }

public:
    long long maximumScore(vector<int>& nums, int k) {
        int n = nums.size();
        int max_i = ranges::max_element(nums) - nums.begin();
        long long ans1 = maximumProfit(nums, max_i, max_i + n, k); // nums[max_i] 是第一个数
        long long ans2 = maximumProfit(nums, max_i + 1, max_i + 1 + n, k); // nums[max_i] 是最后一个数
        return max(ans1, ans2);
    }
};
```

```go [sol-Go]
// 3573. 买卖股票的最佳时机 V
func maximumProfit(prices []int, l, r, k int) int64 {
	n := len(prices)
	f := make([][3]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
		f[j][2] = math.MinInt / 2
	}
	f[0][0] = math.MinInt / 2
	for i := l; i < r; i++ {
		p := prices[i%n]
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p, f[j][2]-p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][2] = max(f[j][2], f[j-1][0]+p)
		}
	}
	return int64(f[k+1][0])
}

func maximumScore(nums []int, k int) int64 {
	n := len(nums)
	maxI := 0
	for i, x := range nums {
		if x > nums[maxI] {
			maxI = i
		}
	}

	ans1 := maximumProfit(nums, maxI, maxI+n, k) // nums[maxI] 是第一个数
	ans2 := maximumProfit(nums, maxI+1, maxI+1+n, k) // nums[maxI] 是最后一个数
	return max(ans1, ans2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 专题训练

见下面动态规划题单的「**六、状态机 DP**」。

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
