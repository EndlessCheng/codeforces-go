## 分析

本题 $m$ 很大，如果用最大堆暴力模拟，时间复杂度为 $\mathcal{O}(m\log n)$，太慢了。

对于第 $m$ 小/大问题，有如下转化套路：

- 第 $m$ 小等价于：求**最小**的 $x$，满足 $\le x$ 的数**至少**有 $m$ 个。
- 第 $m$ 大等价于：求**最大**的 $x$，满足 $\ge x$ 的数**至少**有 $m$ 个。

本题由于 $x$ 越小，能选的次数越多；$x$ 越大，能选的次数越少。有这样的性质，可以**二分**。关于二分的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

## 核心思路

1. 二分求出第 $m$ 大价值 $\textit{low}$（所选价值的最大下界）。
2. 用等差数列求和公式，算出价值严格大于 $\textit{low}$ 的价值之和，以及这些价值的个数 $\textit{cnt}$。剩余的 $m-\textit{cnt}$ 个价值都恰好等于 $\textit{low}$，加到价值之和中。

## 二分第 m 大价值

二分所选价值的下界 $\textit{low}$。我们只考虑 $\textit{value}[i] \ge \textit{low}$ 的下标 $i$。

对于下标 $i$，设我们选了 $k$ 次，则有

$$
\textit{value}[i] - \textit{decay}[i]\cdot(k-1)\ge \textit{low}
$$

解得

$$
k\le \left\lfloor\dfrac{\textit{value}[i]-\textit{low}}{\textit{decay}[i]}\right\rfloor + 1
$$

累加 $k$ 的上界，如果比 $m$ 小，说明 $\textit{low}$ 取大了，更新二分右边界，否则更新二分左边界。

## 计算价值总和

先求出价值严格大于 $\textit{low}$ 的价值之和，以及这些价值的个数 $\textit{cnt}$。

对于 $\textit{value}[i] > \textit{low}$ 的下标 $i$，我们一共选了

$$
k = \left\lfloor\dfrac{\textit{value}[i]-(\textit{low}+1)}{\textit{decay}[i]}\right\rfloor + 1
$$

次。

由等差数列求和公式可得

$$
\begin{aligned}
    & \textit{value}[i] + (\textit{value}[i] - \textit{decay}[i]) + \cdots +  (\textit{value}[i] - \textit{decay}[i]\cdot(k-1))    \\
={} & \dfrac{(\textit{value}[i]\cdot 2 - \textit{decay}[i]\cdot (k-1)) \cdot k}{2}        \\
\end{aligned}
$$

累加 $k$，得到 $\textit{cnt}$。剩余的 $m-\textit{cnt}$ 个价值都恰好等于 $\textit{low}$，加到价值之和中。

> **注**：也可以先计算价值 $\ge \textit{low}$ 的价值之和，最后减去超出 $m$ 次的价值和。

## 细节

下面代码采用开区间二分。使用闭区间或者半闭半开区间也是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$0$。由于 $0$ 不影响价值之和，所以可以认为有无数个等于 $0$ 的价值，所以下界 $0$ 一定满足要求。**注**：如果实际上 $\ge 0$ 的价值不足 $m$ 个，二分结果是 $0$。此时代码中的 `m * low` 等于 $0$，不影响答案。
- 开区间右端点初始值：$\max(\textit{value})+1$。无法选出任何价值。

> **注 1**：对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。
> 
> **注 2**：由于本题 $m\le 10^9$，且我们每次选的价值不超过 $10^9$，所以价值总和不超过 $10^{18}$，可以在最后返回时取模。

## 答疑

**问**：有没有可能，不存在恰好等于二分结果 $\textit{low}$ 的价值？

**答**：不可能。如果没有等于 $\textit{low}$ 的价值，说明「$\ge \textit{low}$ 的价值个数」等于「$\ge \textit{low}+1$ 的价值个数」，但根据循环不变量，$\ge \textit{low}+1$ 的价值个数不足 $m$ 个，矛盾。所以一定存在恰好等于 $\textit{low}$ 的价值。

[本题视频讲解](https://www.bilibili.com/video/BV1uqjt6zEMT/?t=20m10s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxTotalValue(self, value: list[int], decay: list[int], m: int) -> int:
        def check(low: int) -> bool:
            left_m = m
            for v, d in zip(value, decay):
                if v >= low:
                    left_m -= (v - low) // d + 1
                    if left_m < 0:  # 提前跳出循环
                        return True
            return False

        left, right = 0, max(value) + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        low = left

        ans = 0
        # 计算价值严格大于 low 的价值和，以及这些价值的个数
        for v, d in zip(value, decay):
            if v > low:
                k = (v - low - 1) // d + 1
                m -= k
                ans += (v * 2 - d * (k - 1)) * k
        ans //= 2  # 把除以 2 提到循环外面
        ans += m * low  # 剩余 m 次选的价值都是 low
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int maxTotalValue(int[] value, int[] decay, int m) {
        int mx = 0;
        for (int v : value) {
            mx = Math.max(mx, v);
        }

        int left = 0;
        int right = mx + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(mid, value, decay, m)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        int low = left;

        long ans = 0;
        // 计算价值严格大于 low 的价值和，以及这些价值的个数
        for (int i = 0; i < value.length; i++) {
            int v = value[i];
            if (v > low) {
                int d = decay[i];
                int k = (v - low - 1) / d + 1;
                m -= k;
                ans += (v * 2 - (long) d * (k - 1)) * k;
            }
        }
        ans /= 2; // 把除以 2 提到循环外面
        ans += (long) m * low; // 剩余 m 次选的价值都是 low
        return (int) (ans % 1_000_000_007);
    }

    private boolean check(int low, int[] value, int[] decay, int m) {
        int leftM = m;
        for (int i = 0; i < value.length; i++) {
            int v = value[i];
            if (v >= low) {
                leftM -= (v - low) / decay[i] + 1;
                if (leftM < 0) { // 提前跳出循环
                    return true;
                }
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxTotalValue(vector<int>& value, vector<int>& decay, int m) {
        auto check = [&](int low) -> bool {
            int left_m = m;
            for (int i = 0; i < value.size(); i++) {
                int v = value[i];
                if (v >= low) {
                    left_m -= (v - low) / decay[i] + 1;
                    if (left_m < 0) { // 提前跳出循环
                        return true;
                    }
                }
            }
            return false;
        };

        int left = 0, right = ranges::max(value) + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        int low = left;

        long long ans = 0;
        // 计算价值严格大于 low 的价值和，以及这些价值的个数
        for (int i = 0; i < value.size(); i++) {
            int v = value[i];
            if (v > low) {
                int d = decay[i];
                int k = (v - low - 1) / d + 1;
                m -= k;
                ans += (v * 2 - 1LL * d * (k - 1)) * k;
            }
        }
        ans /= 2; // 把除以 2 提到循环外面
        ans += 1LL * m * low; // 剩余 m 次选的价值都是 low
        return ans % 1'000'000'007;
    }
};
```

```go [sol-Go]
func maxTotalValue(value, decay []int, m int) (ans int) {
	low := sort.Search(slices.Max(value), func(low int) bool {
		low++
		leftM := m
		for i, v := range value {
			if v >= low {
				leftM -= (v-low)/decay[i] + 1
				if leftM < 0 { // 提前跳出循环
					return false
				}
			}
		}
		return true
	})

	// 计算价值严格大于 low 的价值和，以及这些价值的个数
	for i, v := range value {
		if v > low {
			dec := decay[i]
			k := (v-low-1)/dec + 1
			m -= k
			ans += (v*2 - dec*(k-1)) * k
		}
	}
	ans /= 2 // 把除以 2 提到循环外面
	ans += m * low // 剩余 m 次选的价值都是 low
	return ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{value}$ 的长度，$U=\max(\textit{value})$。二分 $\mathcal{O}(\log U)$ 次，每次 $\mathcal{O}(n)$ 时间。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面二分题单的「**§2.6 第 K 小/大**」。

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
