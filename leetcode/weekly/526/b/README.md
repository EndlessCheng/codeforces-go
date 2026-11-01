首先，如果 $\textit{value} = 1$，题目的 $\gcd > 1$ 的要求无法满足，返回 $-1$。下面讨论 $\textit{value}>1$ 的情况。

我们需要让 $\textit{nums}$ 增长的尽量慢，这样一方面能让 $\textit{nums}[0] > 0$，从而使每个 $\textit{nums}[i]$ 都大于 $0$；另一方面能让 $\textit{nums}[n-1]$ 尽量小。

## 能否让 nums[0] > 0？

设 $a = \textit{nums}[i],\ b = \textit{nums}[i+1]$。题目要求 $a<b$，即 $b-a>0$。

设 $g = \gcd(a,b)$。根据 $\gcd$ 的定义，有 $a = k_1g,\ b=k_2g$，所以 $b-a = (k_2-k_1)g\ge g$。

所以从 $\textit{nums}[\textit{index}] = \textit{value}$ 开始，每往左走一步，数值至少减少 $g\ge 2$。

分类讨论：

- 如果 $\textit{value}$ 是偶数，那么可以构造一个往左的 $\textit{value},\textit{value}-2,\textit{value}-4,\ldots$ 全为偶数的等差序列。偶数之间的 $\gcd \ge 2$，满足题目要求。
- 如果 $\textit{value}$ 是奇数，设 $\textit{value}$ 的最小质因子为 $p$，那么 $\textit{nums}[\textit{index}-1]$ 至多为 $\textit{value}-p$。由于 $p$ 是奇数，所以 $\textit{value}-p$ 是偶数，那么同上，从 $\textit{nums}[\textit{index}-1] = \textit{value}-p$ 开始，构造一个往左的 $\textit{value}-p,\textit{value}-p-2,\textit{value}-p-4,\ldots$ 全为偶数的等差序列。

这两种情况可以合并为，构造一个从 $\textit{nums}[\textit{index}] = \textit{value}$ 开始往左的序列：

$$
\textit{value},\textit{value}-p,\textit{value}-p-2,\textit{value}-p-4,\ldots
$$

其中

$$
\textit{nums}[0] =
\begin{cases}
\textit{value}, & \textit{index}=0     \\
\textit{value} - p - 2(\textit{index}-1), & \textit{index} > 0     \\
\end{cases}
$$

如果 $\textit{nums}[0] \le 0$，返回 $-1$。

## 让 nums[n-1] 尽量小

同理，构造一个从 $\textit{nums}[\textit{index}] = \textit{value}$ 开始往右的序列：

$$
\textit{value},\textit{value}+p,\textit{value}+p+2,\textit{value}+p+4,\ldots
$$

所以有

$$
\textit{nums}[n-1] =
\begin{cases}
\textit{value}, & \textit{index}=n-1     \\
\textit{value}+p+2(n-2-\textit{index}), & \textit{index} < n-1     \\
\end{cases}
$$

[本题视频讲解](https://www.bilibili.com/video/BV1kfxW63EpG/?t=2m13s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    # 返回 value 的最小质因子
    def min_prime_divisor(self, value: int) -> int:
        for p in range(2, isqrt(value) + 1):
            if value % p == 0:
                return p
        return value  # value 是质数或者 1

    def minimumMaxValue(self, n: int, index: int, value: int) -> int:
        min_p = self.min_prime_divisor(value)
        if value == 1 or value - min_p - (index - 1) * 2 <= 0:  # nums[0] <= 0
            return -1
        if index == n - 1:
            return value
        return value + min_p + (n - 2 - index) * 2
```

```java [sol-Java]
class Solution {
    public long minimumMaxValue(int n, int index, int value) {
        int minP = minPrimeDivisor(value);
        if (value == 1 || value - minP - (index - 1) * 2 <= 0) { // nums[0] <= 0
            return -1;
        }
        if (index == n - 1) {
            return value;
        }
        return (long) value + minP + (n - 2 - index) * 2;
    }

    // 返回 value 的最小质因子
    private int minPrimeDivisor(int value) {
        for (int p = 2; p * p <= value; p++) {
            if (value % p == 0) {
                return p;
            }
        }
        return value; // value 是质数或者 1
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回 value 的最小质因子
    int min_prime_divisor(int value) {
        for (int p = 2; p * p <= value; p++) {
            if (value % p == 0) {
                return p;
            }
        }
        return value; // value 是质数或者 1
    }

public:
    long long minimumMaxValue(int n, int index, int value) {
        int min_p = min_prime_divisor(value);
        if (value == 1 || value - min_p - (index - 1) * 2 <= 0) { // nums[0] <= 0
            return -1;
        }
        if (index == n - 1) {
            return value;
        }
        return 1LL * value + min_p + (n - 2 - index) * 2;
    }
};
```

```go [sol-Go]
// 返回 value 的最小质因子
func minPrimeDivisor(value int) int {
	for p := 2; p*p <= value; p++ {
		if value%p == 0 {
			return p
		}
	}
	return value // value 是质数或者 1
}

func minimumMaxValue(n, index, value int) int64 {
	minP := minPrimeDivisor(value)
	if value == 1 || value-minP-(index-1)*2 <= 0 { // nums[0] <= 0
		return -1
	}
	if index == n-1 {
		return int64(value)
	}
	return int64(value + minP + (n-2-index)*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\sqrt{\textit{value}})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数学题单的「**§1.6 最大公约数（GCD）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
