示例 1 的操作流程如下：

- 序列一开始是 $[1, 2, 3, 4, 5, 6, 7, 8]$，这是一个首项为 $1$，公差为 $1$ 的等差数列。
- 删除所有奇数下标数字，得到序列 $[1, 3, 5, 7]$。由于下次操作要从右侧开始删，我们可以把序列**反转**，得到 $[7,5,3,1]$，那么下次操作相当于还是从左侧开始删。序列 $[7,5,3,1]$ 是一个首项为 $7$，公差为 $-2$ 的等差数列。
- 删除所有奇数下标数字，得到序列 $[7,3]$。由于下次操作要从右侧开始删，我们可以把序列**反转**，得到 $[3,7]$，那么下次操作相当于还是从左侧开始删。序列 $[3,7]$ 是一个首项为 $3$，公差为 $4$ 的等差数列。
- 删除所有奇数下标数字，得到序列 $[3]$。

考察上述流程：

- 每次操作，删除所有奇数下标数字。当 $n$ 是偶数时，删除一半，当 $n$ 是奇数时，删除 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个（比如 $n=5$ 时删除 $2$ 个）。所以每次都会删除 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个元素，序列元素个数从 $n$ 变成 $\left\lceil\dfrac{n}{2}\right\rceil$。
- 公差初始为 $d=1$，每次操作把 $d$ 乘以 $-2$。
- 如果 $n$ 是奇数，那么序列的最后一个数一定保留。操作后，首项 $\textit{start}$ 变成序列的最后一个数 $\textit{start} + (n-1)\cdot d$，其中 $n$ 和 $d$ 都是操作之前的值。
- 如果 $n$ 是偶数，那么序列的最后一个数被删除，倒数第二个数保留。操作后，首项 $\textit{start}$ 变成序列的倒数第二个数 $\textit{start} + (n-2)\cdot d$，其中 $n$ 和 $d$ 都是操作之前的值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def lastInteger(self, n: int) -> int:
        start = d = 1  # 等差数列首项，公差
        while n > 1:
            start += (n - 2 + n % 2) * d
            d *= -2
            n = (n + 1) // 2
        return start
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
        long long start = 1, d = 1; // 等差数列首项，公差
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
	start, d := int64(1), int64(1) // 等差数列首项，公差
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

## 附

Python3 可以直接用 $\texttt{range}$ 模拟。

```py
class Solution:
    def lastInteger(self, n: int) -> int:
        r = range(1, n + 1)
        while len(r) > 1:
            r = r[::2][::-1]
        return r[0]
```

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
