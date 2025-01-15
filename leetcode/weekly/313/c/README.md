设 $n_1$ 为 $\textit{num}_1$ 的二进制表示的长度，$c_1$ 为 $\textit{num}_1$ 的置位数，$c_2$ 为 $\textit{num}_2$ 的置位数。

基本思路：

$x$ 的置位数和 $\textit{num}_2$ 相同，意味着 $x$ 的二进制表示中有 $c_2$ 个 $1$，我们需要合理地分配这 $c_2$ 个 $1$。

为了让异或和尽量小，这些 $1$ 应当从高位到低位匹配 $\textit{num}_1$ 中的 $1$；如果匹配完了还有多余的 $1$，那么就从低位到高位把 $0$ 改成 $1$。

分类讨论：

- 如果 $c_2\ge n_1$，$x$ 只能是 $2^{c_2}-1$，任何其他方案都会使异或和变大；
- 如果 $c_2=c_1$，那么 $x=\textit{num}_1$；
- 如果 $c_2<c_1$，那么将 $\textit{num}_1$ 的最低的 $c_1-c_2$ 个 $1$ 变成 $0$，其结果就是 $x$；
- 如果 $c_2>c_1$，那么将 $\textit{num}_1$ 的最低的 $c_2-c_1$ 个 $0$ 变成 $1$，其结果就是 $x$；

下面代码用到了一些位运算技巧，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

[本题视频讲解](https://www.bilibili.com/video/BV1kd4y1q7fC) 第三题。

```py [sol-Python3]
class Solution:
    def minimizeXor(self, num1: int, num2: int) -> int:
        c1 = num1.bit_count()
        c2 = num2.bit_count()
        while c2 < c1:
            num1 &= num1 - 1  # 最低的 1 变成 0
            c2 += 1
        while c2 > c1:
            num1 |= num1 + 1  # 最低的 0 变成 1
            c2 -= 1
        return num1
```

```java [sol-Java]
class Solution {
    public int minimizeXor(int num1, int num2) {
        int c1 = Integer.bitCount(num1);
        int c2 = Integer.bitCount(num2);
        for (; c2 < c1; c2++) num1 &= num1 - 1; // 最低的 1 变成 0
        for (; c2 > c1; c2--) num1 |= num1 + 1; // 最低的 0 变成 1
        return num1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeXor(int num1, int num2) {
        int c1 = popcount((unsigned) num1);
        int c2 = popcount((unsigned) num2);
        for (; c2 < c1; c2++) num1 &= num1 - 1; // 最低的 1 变成 0
        for (; c2 > c1; c2--) num1 |= num1 + 1; // 最低的 0 变成 1
        return num1;
    }
};
```

```go [sol-Go]
func minimizeXor(num1, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))
	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 最低的 1 变成 0
	}
	for ; c2 > c1; c2-- {
		num1 |= num1 + 1 // 最低的 0 变成 1
	}
	return num1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(|\log\textit{num}_1 - \log\textit{num}_2|)$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
