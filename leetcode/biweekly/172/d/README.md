## 方法一：用等差数列模拟

示例 1 的操作流程如下：

- 序列一开始是 $[1, 2, 3, 4, 5, 6, 7, 8]$，这是一个首项为 $1$，公差为 $1$ 的等差数列。
- 删除所有奇数下标数字，得到序列 $[1, 3, 5, 7]$。由于下次操作要从右侧开始删，我们可以把序列**反转**，得到 $[7,5,3,1]$，那么下次操作相当于还是从左侧开始删。序列 $[7,5,3,1]$ 是一个首项为 $7$，公差为 $-2$ 的等差数列。
- 删除所有奇数下标数字，得到序列 $[7,3]$。由于下次操作要从右侧开始删，我们可以把序列**反转**，得到 $[3,7]$，那么下次操作相当于还是从左侧开始删。序列 $[3,7]$ 是一个首项为 $3$，公差为 $4$ 的等差数列。
- 删除所有奇数下标数字，得到序列 $[3]$。

考察上述流程：

- 每次操作，删除所有奇数下标数字。当 $n$ 是偶数时，删除一半，当 $n$ 是奇数时，删除 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个（比如 $n=5$ 时删除 $2$ 个）。所以每次都会删除 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个元素，序列元素个数从 $n$ 变成 $\left\lceil\dfrac{n}{2}\right\rceil$。
- 间隔地删除一个等差数列中的元素，得到的仍然是一个等差数列。
- 公差初始为 $d=1$，每次操作把 $d$ 乘以 $-2$。
- 如果 $n$ 是奇数，那么序列的最后一个数一定保留。操作后，首项 $\textit{start}$ 变成序列的最后一个数 $\textit{start} + (n-1)\cdot d$，其中 $n$ 和 $d$ 都是操作之前的值。
- 如果 $n$ 是偶数，那么序列的最后一个数被删除，倒数第二个数保留。操作后，首项 $\textit{start}$ 变成序列的倒数第二个数 $\textit{start} + (n-2)\cdot d$，其中 $n$ 和 $d$ 都是操作之前的值。

[本题视频讲解](https://www.bilibili.com/video/BV14LqmBMECK/?t=16m13s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def lastInteger(self, n: int) -> int:
        start = d = 1  # 等差数列的首项和公差
        while n > 1:
            start += (n - 2 + n % 2) * d
            d *= -2
            n = (n + 1) // 2
        return start
```

```py [sol-Python3 range]
class Solution:
    def lastInteger(self, n: int) -> int:
        # range 是一个等差数列对象，不是 list
        # 计算 len、计算切片都可以用数学公式做到 O(1) 时间复杂度
        r = range(1, n + 1)
        while len(r) > 1:
            r = r[::2][::-1]
        return r[0]
```

```java [sol-Java]
class Solution {
    public long lastInteger(long n) {
        long start = 1; // 等差数列首项
        long d = 1; // 等差数列公差
        for (; n > 1; n = (n + 1) / 2) {
            start += (n - 2 + n % 2) * d;
            d *= -2;
        }
        return start;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long lastInteger(long long n) {
        long long start = 1, d = 1; // 等差数列的首项和公差
        for (; n > 1; n = (n + 1) / 2) {
            start += (n - 2 + n % 2) * d;
            d *= -2;
        }
        return start;
    }
};
```

```go [sol-Go]
func lastInteger(n int64) int64 {
	start, d := int64(1), int64(1) // 等差数列的首项和公差
	for ; n > 1; n = (n + 1) / 2 {
		start += (n - 2 + n%2) * d
		d *= -2
	}
	return start
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。根据 [下取整恒等式及其应用](https://zhuanlan.zhihu.com/p/1893240318645732760)，$k$ 次操作后序列长度为 $\left\lceil\dfrac{n}{2^k}\right\rceil$，所以操作次数为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：位运算

为方便计算，把初始序列改成从 $0$ 开始，即 $0,1,2,\ldots,n-1$。

第一次操作，我们删除了所有的奇数，剩余的都是偶数。这意味着，（从 $0$ 开始的）最终答案，二进制最低位一定是 $0$。

把剩余元素 $0,2,4,\ldots$ 全部右移一位，我们又得到了序列 $0,1,2,\ldots$

在此基础上，执行第二次操作。

在第二次操作中，我们要从右往左删除：

- 如果序列最后一个数是偶数，例如 $0,1,2,3,4$，那么我们会删除所有的奇数，剩余的都是偶数。
- 如果序列最后一个数是奇数，例如 $0,1,2,3,4,5$，那么我们会删除所有的偶数，剩余的都是奇数。

这意味着，（从 $0$ 开始的）最终答案，二进制从低到高第二位一定等于 $n-1$ 从低到高第二位。

依此类推。

一般地，（从 $0$ 开始的）最终答案，二进制从低到高第 $1,3,5,\ldots$ 位一定是 $0$；第 $2,4,6,\ldots$ 位一定和 $n-1$ 的第 $2,4,6,\ldots$ 位相同。

算出答案后，把答案加一（因为原题的序列是从 $1$ 开始的）。

```py [sol-Python3]
class Solution:
    def lastInteger(self, n: int) -> int:
        MASK = 0xAAAAAAAAAAAAAAA  # ...1010
        return ((n - 1) & MASK) + 1  # 取出 n-1 的从低到高第 2,4,6,... 位，最后再加一（从 1 开始）
```

```java [sol-Java]
class Solution {
    public long lastInteger(long n) {
        final long MASK = 0xAAAAAAAAAAAAAAAL; // ...1010
        return ((n - 1) & MASK) + 1; // 取出 n-1 的从低到高第 2,4,6,... 位，最后再加一（从 1 开始）
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long lastInteger(long long n) {
        constexpr long long MASK = 0xAAAAAAAAAAAAAAALL; // ...1010
        return ((n - 1) & MASK) + 1; // 取出 n-1 的从低到高第 2,4,6,... 位，最后再加一（从 1 开始）
    }
};
```

```go [sol-Go]
func lastInteger(n int64) int64 {
	const mask = 0xAAAAAAAAAAAAAAA // ...1010
	return (n-1)&mask + 1 // 取出 n-1 的从低到高第 2,4,6,... 位，最后再加一（从 1 开始）
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[390. 消除游戏](https://leetcode.cn/problems/elimination-game/)

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
