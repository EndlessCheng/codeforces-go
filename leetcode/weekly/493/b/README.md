题目让我们统计

$$
\begin{array}{r}
1\\
2\\
3\\
\vdots\\
999\\
1,000\\
1,001\\
1,002\\
\vdots\\
999,999\\
1,000,000\\
1,000,001\\
1,000,002\\
\vdots \\
999,999,999\\
1,000,000,000\\
1,000,000,001\\
1,000,000,002\\
\vdots \\
n \\
\end{array}
$$

这些数一共有多少个逗号。

横看成岭侧成峰，从逗号的视角看，它出现在多少个数中？

- 最右边的逗号，出现在 $[10^3,n]$ 的每个整数中。这有 $n-10^3+1$ 个。
- 倒数第二个逗号，出现在 $[10^6,n]$ 的每个整数中。这有 $n-10^6+1$ 个。
- 倒数第三个逗号，出现在 $[10^9,n]$ 的每个整数中。这有 $n-10^9+1$ 个。
- 依此类推。

累加这些逗号的个数，即为答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countCommas(self, n: int) -> int:
        ans = 0
        # 从低到高，枚举逗号的位置
        low = 1000
        while low <= n:
            # [low, n] 中的每个数都在这个位置上有一个逗号
            ans += n - low + 1
            low *= 1000
        return ans
```

```java [sol-Java]
class Solution {
    public long countCommas(long n) {
        long ans = 0;
        // 从低到高，枚举逗号的位置
        for (long low = 1000; low <= n; low *= 1000) {
            // [low, n] 中的每个数都在这个位置上有一个逗号
            ans += n - low + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countCommas(long long n) {
        long long ans = 0;
        // 从低到高，枚举逗号的位置
        for (long long low = 1000; low <= n; low *= 1000) {
            // [low, n] 中的每个数都在这个位置上有一个逗号
            ans += n - low + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countCommas(n int64) (ans int64) {
	// 从低到高，枚举逗号的位置
	for low := int64(1000); low <= n; low *= 1000 {
		// [low, n] 中的每个数都在这个位置上有一个逗号
		ans += n - low + 1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 附：数学公式

设 $k$ 是满足 $10^{3k}\le n$ 的最大整数。

根据等比数列求和公式，答案为

$$
\begin{aligned}
    & (n-10^3+1) + (n-10^6+1) + \cdots + (n-10^{3k}+1)      \\
={} & k(n+1) - \dfrac{10^{3k+3}-1000}{999}        \\
\end{aligned}
$$

```go
func countCommas(n int64) (ans int64) {
	k := 5 // n == 1e15 时 Log10(n) 有误差，需要特判
	if n < 1e15 {
		k = int(math.Log10(float64(n))) / 3
	}
	return int64(k)*(n+1) - (int64(math.Pow10(k*3+3))-1000)/999
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面思维题单的「**§5.5 贡献法**」。

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
