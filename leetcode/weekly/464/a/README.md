我们计算的是公差为 $2$ 的等差数列之和，根据等差数列求和公式，我们有

$$
\textit{sumOdd} = 1 + 3 + \dots + (2n-1) = \dfrac{(1+2n-1)n}{2} = n^2
$$

以及

$$
\textit{sumEven} = 2 + 4 + \dots + 2n = \dfrac{(2+2n)n}{2} = n(n+1)
$$

提取公约数 $n$ 后，$\textit{sumOdd}$ 剩下 $n$，$\textit{sumEven}$ 剩下 $n+1$，这两个数一定互质。**反证法**：如果 $n$ 和 $n+1$ 不互质，那么这两个数都是某个大于 $1$ 的整数 $x$ 的倍数。由于 $n\ne n+1$，两个不同的 $x$ 的倍数至少相差 $x$，而 $x > 1$，这与两数相差 $1$ 矛盾，所以 $n$ 和 $n+1$ 一定互质。

所以最大公约数为 $n$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1X9eJz2EWE/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def gcdOfOddEvenSums(self, n: int) -> int:
        return n
```

```java [sol-Java]
class Solution {
    public int gcdOfOddEvenSums(int n) {
        return n;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int gcdOfOddEvenSums(int n) {
        return n;
    }
};
```

```go [sol-Go]
func gcdOfOddEvenSums(n int) int {
	return n
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
