两根木材，三辆车，每辆车只能装一根木材，所以**至多切一次**。

设被切割的木材长为 $n$，那么 $n\le k$ 时无需切割，否则设分成两段 $x$ 和 $n-x$，其中 $n-k\le x\le k$。

题目要求最小化

$$
x(n-x) = -x^2 + nx
$$

这是开口向下的抛物线，最小值在 $x=k$（或者对称的 $x=n-k$）处取到，代入得

$$
k(n-k)
$$

整合一下，无论 $n$ 是否大于 $k$，把答案增加

$$
\max(k(n-k), 0)
$$

由于题目保证存在能被运输的方案，$n$ 和 $m$ 一定有一个 $\le k$，我们只需考虑二者的较大值 $\max(n,m)$。所以最终答案为

$$
\max(k\cdot(\max(n,m)-k), 0)
$$

具体请看 [视频讲解](https://www.bilibili.com/video/BV1o1jgzJE51/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minCuttingCost(self, n: int, m: int, k: int) -> int:
        return max(k * (max(n, m) - k), 0)
```

```java [sol-Java]
class Solution {
    public long minCuttingCost(int n, int m, int k) {
        return Math.max((long) k * (Math.max(n, m) - k), 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCuttingCost(int n, int m, int k) {
        return max(1LL * k * (max(n, m) - k), 0LL);
    }
};
```

```go [sol-Go]
func minCuttingCost(n, m, k int) int64 {
	return int64(max(k*(max(n, m)-k), 0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
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
