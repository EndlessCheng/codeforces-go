非负整数 $x$ 二进制的所有位都相同，意味着 $x = 0$ 或者 $x$ 的二进制全为 $1$，即 $x = 2^k - 1$，其中 $k$ 是非负整数。

问题变成求 $[0,n]$ 中有多少个 $2^k - 1$，或者说有多少个不同的 $k$。

$0\le 2^k -1 \le n$ 即 $1\le 2^k\le n+1$。

解得 $0\le k\le w-1$，其中 $w$ 是 $n+1$ 的二进制长度。

比如 $n+1 = 10110$，二进制长度为 $5$，那么 $2^k\le 10000$，所以 $k\le 4$。

所以有 $w$ 个不同的 $k$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countMonobit(self, n: int) -> int:
        return (n + 1).bit_length()
```

```java [sol-Java]
class Solution {
    public int countMonobit(int n) {
        return 32 - Integer.numberOfLeadingZeros(n + 1);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countMonobit(int n) {
        return bit_width((uint32_t) n + 1);
    }
};
```

```go [sol-Go]
func countMonobit(n int) int {
	return bits.Len(uint(n + 1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。读者可以看看库函数的源码。
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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
