只要 $s$ 有不同字母，那么 $s$ 一定存在相邻不同的字母，我们就可以把这两个字母（作为一个子串）移除，直到 $s$ 只剩下一种字母为止。

比如有 $3$ 个 $\texttt{a}$ 和 $2$ 个 $\texttt{b}$，每次消除 $1$ 个 $\texttt{a}$ 和 $1$ 个 $\texttt{b}$，最终剩下 $3-2=1$ 个 $\texttt{a}$。

设 $s$ 有 $k$ 个 $\texttt{a}$，那么有 $n-k$ 个 $\texttt{b}$。

消除后，剩余字母个数为 $|k-(n-k)| = |2k-n|$。

[本题视频讲解](https://www.bilibili.com/video/BV1ZuCQBJEjD/?t=1m45s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minLengthAfterRemovals(self, s: str) -> int:
        return abs(s.count('a') * 2 - len(s))
```

```java [sol-Java]
class Solution {
    public int minLengthAfterRemovals(String s) {
        int k = 0;
        for (char ch : s.toCharArray()) {
            k += ch - 'a'; // 也可以统计 b 的个数
        }
        return Math.abs(k * 2 - s.length());
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLengthAfterRemovals(string s) {
        int k = ranges::count(s, 'a');
        return abs(k * 2 - (int) s.size());
    }
};
```

```go [sol-Go]
func minLengthAfterRemovals(s string) int {
	k := strings.Count(s, "a")
	return abs(k*2 - len(s))
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面思维题单的「**§5.2 脑筋急转弯**」。

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
