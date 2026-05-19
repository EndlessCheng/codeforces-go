先思考一个更简单的问题：

- 如果不能包含子序列 $\texttt{01}$ 或 $\texttt{10}$，最少修改次数是多少？

如果 $s$ 中 $\texttt{0}$ 和 $\texttt{1}$ 都有，那么 $\texttt{0}$ 在左 $\texttt{1}$ 在右，意味着 $s$ 包含子序列 $\texttt{01}$；$\texttt{1}$ 在左 $\texttt{0}$ 在右，意味着 $s$ 包含子序列 $\texttt{10}$。所以 $s$ 要么全为 $\texttt{0}$，要么全为 $\texttt{1}$。

思考这个更简单的问题，会启发我们往全为 $\texttt{0}$ 或者 $\texttt{1}$ 这个方向思考。

回到本题。$s$ 不含子序列 $\texttt{110}$ 或 $\texttt{011}$，意味着对于 $s$ 中的任意一个 $\texttt{0}$，其左侧不能有超过一个 $\texttt{1}$，右侧也不能有超过一个 $\texttt{1}$。

设 $n$ 是 $s$ 的长度，设 $c_0$ 是 $s$ 中的 $\texttt{0}$ 的个数，那么 $n-c_0$ 是 $s$ 中的 $\texttt{1}$ 的个数。分类讨论：

- 把 $s$ 中的 $\texttt{0}$ 全变成 $\texttt{1}$，没有 $\texttt{0}$，不就万事大吉了吗？操作 $c_0$ 次。
- 如果 $s[0] = \texttt{0}$，那么 $s$ 不能有超过一个 $\texttt{1}$，否则 $s$ 包含子序列 $\texttt{011}$；如果 $s[n-1] = \texttt{0}$，那么 $s$ 不能有超过一个 $\texttt{1}$，否则 $s$ 包含子序列 $\texttt{110}$。此时至多保留一个 $\texttt{1}$，其余 $\texttt{1}$ 全变成 $\texttt{0}$，操作 $\max(n-c_0-1, 0)$ 次。和 $0$ 取最大值是为了兼容 $s$ 没有 $\texttt{1}$ 的情况，此时 $n-c_0-1=-1$。
- 如果 $s[0] = s[n-1] = \texttt{1}$，那么可以保留 $s[0]$ 和 $s[n-1]$ 这两个 $\texttt{1}$（多保留一个 $\texttt{1}$），其余 $\texttt{1}$ 全变成 $\texttt{0}$，操作 $\max(n-c_0-2,0)$ 次。和 $0$ 取最大值是为了兼容 $n=1$ 的情况，此时 $n-c_0-2=-1$。

这些情况取最小值，即为答案。

```py [sol-Python3]
class Solution:
    def minFlips(self, s: str) -> int:
        c0 = s.count('0')
        c1 = len(s) - c0 - 1
        if s[0] == '1' and s[-1] == '1':
            c1 -= 1
        return min(c0, max(c1, 0))
```

```java [sol-Java]
class Solution {
    public int minFlips(String S) {
        char[] s = S.toCharArray();
        int n = s.length;

        int c0 = 0;
        for (char ch : s) {
            c0 += '1' - ch;
        }

        int c1 = n - c0 - 1;
        if (s[0] == '1' && s[n - 1] == '1') {
            c1--;
        }

        return Math.min(c0, Math.max(c1, 0));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minFlips(string s) {
        int n = s.size();
        int c0 = ranges::count(s, '0');
        int c1 = n - c0 - 1;
        if (s[0] == '1' && s[n - 1] == '1') {
            c1--;
        }
        return min(c0, max(c1, 0));
    }
};
```

```go [sol-Go]
func minFlips(s string) int {
	n := len(s)
	c0 := strings.Count(s, "0")
	c1 := n - c0 - 1
	if s[0] == '1' && s[n-1] == '1' {
		c1--
	}
	return min(c0, max(c1, 0))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 变形题

要让 $s$ 不包含子序列 $\texttt{010}$ 或 $\texttt{101}$，至少要操作多少次？

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面贪心与思维题单的「**§5.7 分类讨论**」。

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
