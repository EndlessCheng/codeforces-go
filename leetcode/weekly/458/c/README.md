推荐先把本题的简单版本 [3307. 找出第 K 个字符 II](https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii/) 做了。

正着做可以吗？

生成的 $\textit{result}$ 字符串太长，超出内存限制，不行。

那么，倒着做可以吗？

为了方便下面计算 $k$ 的位置，首先，计算出每个 $s[i]$ 处理完后的 $\textit{result}$ 的长度，记作 $\textit{size}[i]$。最终 $\textit{result}$ 的长度为 $\textit{size}[n-1]$。

首先特判 $k\ge \textit{size}[n-1]$ 的情况，下标越界，返回 $\texttt{.}$ 号。

否则答案一定存在。我们从 $c = s[n-1]$ 开始倒着思考：

- 如果 $c$ 是字母：
    - 如果 $k = \textit{size}[n-1]-1$，那么答案就是 $\textit{result}$ 的最后一个字母，即 $c$。
    - 否则需要继续分析，问题变成 $s$ 的前 $n-1$ 个字母的**子问题**，可以**递归/迭代**解决。
- 如果 $c = \texttt{*}$：
    - 无影响（我们已经特判下标越界的情况），问题变成 $s$ 的前 $n-1$ 个字母的子问题。
- 如果 $c = \texttt{#}$：
    - 设 $m = \dfrac{\textit{size}[n-1]}{2}$。
    - 如果 $k < m$，无影响，问题变成 $s$ 的前 $n-1$ 个字母的子问题。
    - 如果 $k \ge m$，问题变成 $s$ 的前 $n-1$ 个字母的子问题中，下标为 $k - m$ 的字母（这个字母复制后的下标就是 $k$）。
- 如果 $c = \texttt{\%}$：
    - 问题变成 $s$ 的前 $n-1$ 个字母的子问题中，下标为 $\textit{size}[n-1]-1-k$ 的字母（这个字母反转后的下标就是 $k$）。

其余步骤和最后一步的处理逻辑是完全一样的，所以**想清楚最后一步怎么做，就想清楚了每一步怎么做**。

下面代码用迭代解决。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1xSuFzHEa1/?t=6m41s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def processStr(self, s: str, k: int) -> str:
        n = len(s)
        size = [0] * n
        sz = 0
        for i, c in enumerate(s):
            if c == '*':
                sz = max(sz - 1, 0)
            elif c == '#':
                sz *= 2
            elif c != '%':  # c 是字母
                sz += 1
            size[i] = sz

        if k >= size[-1]:  # 下标越界
            return '.'

        # 迭代
        for i in range(n - 1, -1, -1):
            c = s[i]
            sz = size[i]
            if c == '#':
                if k >= sz // 2:  # k 在复制后的右半边
                    k -= sz // 2
            elif c == '%':
                k = sz - 1 - k  # 反转前的下标为 sz-1-k 的字母就是答案
            elif c != '*' and k == sz - 1:  # 找到答案
                return c
```

```java [sol-Java]
class Solution {
    char processStr(String S, long k) {
        char[] s = S.toCharArray();
        int n = s.length;
        long[] size = new long[n];
        long sz = 0;
        for (int i = 0; i < n; i++) {
            char c = s[i];
            if (c == '*') {
                sz = Math.max(sz - 1, 0);
            } else if (c == '#') {
                sz *= 2; // 题目保证 sz <= 1e15
            } else if (c != '%') { // c 是字母
                sz++;
            }
            size[i] = sz;
        }

        if (k >= size[n - 1]) { // 下标越界
            return '.';
        }

        // 迭代
        for (int i = n - 1; ; i--) {
            char c = s[i];
            sz = size[i];
            if (c == '#') {
                if (k >= sz / 2) { // k 在复制后的右半边
                    k -= sz / 2;
                }
            } else if (c == '%') {
                k = sz - 1 - k; // 反转前的下标为 sz-1-k 的字母就是答案
            } else if (c != '*' && k == sz - 1) { // 找到答案
                return c;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    char processStr(string s, long long k) {
        int n = s.size();
        vector<long long> size(n);
        long long sz = 0;
        for (int i = 0; i < n; i++) {
            char c = s[i];
            if (c == '*') {
                sz = max(sz - 1, 0LL);
            } else if (c == '#') {
                sz *= 2; // 题目保证 sz <= 1e15
            } else if (c != '%') { // c 是字母
                sz++;
            }
            size[i] = sz;
        }

        if (k >= size[n - 1]) { // 下标越界
            return '.';
        }

        // 迭代
        for (int i = n - 1; ; i--) {
            char c = s[i];
            sz = size[i];
            if (c == '#') {
                if (k >= sz / 2) { // k 在复制后的右半边
                    k -= sz / 2;
                }
            } else if (c == '%') {
                k = sz - 1 - k; // 反转前的下标为 sz-1-k 的字母就是答案
            } else if (c != '*' && k == sz - 1) { // 找到答案
                return c;
            }
        }
    }
};
```

```go [sol-Go]
func processStr(s string, k int64) byte {
	n := len(s)
	size := make([]int64, n)
	sz := int64(0)
	for i, c := range s {
		if c == '*' {
			sz = max(sz-1, 0)
		} else if c == '#' {
			sz *= 2
		} else if c != '%' { // c 是字母
			sz++
		}
		size[i] = sz
	}

	if k >= size[n-1] { // 下标越界
		return '.'
	}

	// 迭代
	for i := n - 1; ; i-- {
		c := s[i]
		sz = size[i]
		if c == '#' {
			if k >= sz/2 { // k 在复制后的右半边
				k -= sz / 2
			}
		} else if c == '%' {
			k = sz - 1 - k // 反转前的下标为 sz-1-k 的字母就是答案
		} else if c != '*' && k == sz-1 { // 找到答案
			return c
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [3307. 找出第 K 个字符 II](https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii/) 2232
- [1545. 找出第 N 个二进制字符串中的第 K 位](https://leetcode.cn/problems/find-kth-bit-in-nth-binary-string/) 做到 $\mathcal{O}(n)$

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
