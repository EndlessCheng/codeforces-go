题目的约束相当于：

- 子串不能同时包含 $s$ 的第一个字母和最后一个字母。

这意味着，如果第一个字母或者最后一个字母已经在正确的位置上，那么至多操作 $1$ 次。否则，至少操作 $2$ 次（$n\ge 3$）。

分类讨论（从上到下依次判断）：

- **情况一**：如果 $s$ 已经是升序，无需操作。
- **情况二**：如果 $n=2$，无法排序。
- **情况三**：第一个字母或者最后一个字母已经在正确的位置上。
    - 如果 $s[0]$ 是最小值，排序 $[1,n-1]$ 即可，操作 $1$ 次。
    - 如果 $s[n-1]$ 是最大值，排序 $[0,n-2]$ 即可，操作 $1$ 次。
- **情况四**：只需 $1$ 次操作，就可以变成情况三。
    - 如果 $[1,n-2]$ 中有最小值，那么先排序 $[0,n-2]$，把最小值排到最前面，变成情况三，再排序 $[1,n-1]$ 即可，一共操作 $2$ 次。
    - 如果 $[1,n-2]$ 中有最大值，那么先排序 $[1,n-1]$，把最大值排到最后面，变成情况三，再排序 $[0,n-2]$ 即可，一共操作 $2$ 次。
- **情况五**：需要 $2$ 次操作，才能变成情况三。
    - 现在只剩下一种情况，$s[0]$ 是最大值，$s[n-1]$ 是最小值，且 $[1,n-2]$ 不含最小值和最大值。
    - 先排序 $[0,n-2]$，把最大值排到下标 $n-2$。
    - 然后排序 $[1,n-1]$，把最大值排到最后面，且最小值排到下标 $1$。我们用了 $2$ 次操作，变成了情况三。
    - 最后排序 $[0,n-2]$，把最小值排到最前面。
    - 一共操作 $3$ 次。

[本题视频讲解](https://www.bilibili.com/video/BV1H6NMzdEbo/?t=29m6s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str) -> int:
        # s 已经是升序
        if all(x <= y for x, y in pairwise(s)):
            return 0

        # 长为 2，无法排序
        if len(s) == 2:
            return -1

        mn = min(s)
        mx = max(s)

        # 如果 s[0] 是最小值，排序 [1,n-1] 即可
        # 如果 s[n-1] 是最大值，排序 [0,n-2] 即可
        if s[0] == mn or s[-1] == mx:
            return 1

        # 如果 [1,n-2] 中有最小值，那么先排序 [0,n-2]，把最小值排在最前面，然后排序 [1,n-1] 即可
        # 如果 [1,n-2] 中有最大值，那么先排序 [1,n-1]，把最大值排在最后面，然后排序 [0,n-2] 即可
        if any(ch == mn or ch == mx for ch in s[1:-1]):
            return 2

        # 现在只剩下一种情况：s[0] 是最大值，s[n-1] 是最小值，且 [1,n-2] 不含最小值和最大值
        # 先排序 [0,n-2]，把最大值排到 n-2
        # 然后排序 [1,n-1]，把最大值排在最后面，且最小值排在 1
        # 最后排序 [0,n-2]，把最小值排在最前面
        return 3
```

```java [sol-Java]
class Solution {
    public int minOperations(String S) {
        char[] s = S.toCharArray();
        int n = s.length;

        boolean isSorted = true;
        for (int i = 1; i < n; i++) {
            if (s[i - 1] > s[i]) {
                isSorted = false;
                break;
            }
        }
        // s 已经是升序
        if (isSorted) {
            return 0;
        }

        // 长为 2，无法排序
        if (n == 2) {
            return -1;
        }

        char mn = s[0];
        char mx = s[0];
        for (char ch : s) {
            mn = (char) Math.min(mn, ch);
            mx = (char) Math.max(mx, ch);
        }

        // 如果 s[0] 是最小值，排序 [1,n-1] 即可
        // 如果 s[n-1] 是最大值，排序 [0,n-2] 即可
        if (s[0] == mn || s[n - 1] == mx) {
            return 1;
        }

        // 如果 [1,n-2] 中有最小值，那么先排序 [0,n-2]，把最小值排在最前面，然后排序 [1,n-1] 即可
        // 如果 [1,n-2] 中有最大值，那么先排序 [1,n-1]，把最大值排在最后面，然后排序 [0,n-2] 即可
        for (int i = 1; i < n - 1; i++) {
            if (s[i] == mn || s[i] == mx) {
                return 2;
            }
        }

        // 现在只剩下一种情况：s[0] 是最大值，s[n-1] 是最小值，且 [1,n-2] 不含最小值和最大值
        // 先排序 [0,n-2]，把最大值排到 n-2
        // 然后排序 [1,n-1]，把最大值排在最后面，且最小值排在 1
        // 最后排序 [0,n-2]，把最小值排在最前面
        return 3;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s) {
        // s 已经是升序
        if (ranges::is_sorted(s)) {
            return 0;
        }

        int n = s.size();
        // 长为 2，无法排序
        if (n == 2) {
            return -1;
        }

        auto [mn, mx] = ranges::minmax(s);

        // 如果 s[0] 是最小值，排序 [1,n-1] 即可
        // 如果 s[n-1] 是最大值，排序 [0,n-2] 即可
        if (s[0] == mn || s[n - 1] == mx) {
            return 1;
        }

        // 如果 [1,n-2] 中有最小值，那么先排序 [0,n-2]，把最小值排在最前面，然后排序 [1,n-1] 即可
        // 如果 [1,n-2] 中有最大值，那么先排序 [1,n-1]，把最大值排在最后面，然后排序 [0,n-2] 即可
        for (int i = 1; i < n - 1; i++) {
            if (s[i] == mn || s[i] == mx) {
                return 2;
            }
        }

        // 现在只剩下一种情况：s[0] 是最大值，s[n-1] 是最小值，且 [1,n-2] 不含最小值和最大值
        // 先排序 [0,n-2]，把最大值排到 n-2
        // 然后排序 [1,n-1]，把最大值排在最后面，且最小值排在 1
        // 最后排序 [0,n-2]，把最小值排在最前面
        return 3;
    }
};
```

```go [sol-Go]
func minOperations(s string) int {
	t := []byte(s)
	// s 已经是升序
	if slices.IsSorted(t) {
		return 0
	}

	n := len(t)
	// 长为 2，无法排序
	if n == 2 {
		return -1
	}

	mn := slices.Min(t)
	mx := slices.Max(t)
	// 如果 s[0] 是最小值，排序 [1,n-1] 即可
	// 如果 s[n-1] 是最大值，排序 [0,n-2] 即可
	if t[0] == mn || t[n-1] == mx {
		return 1
	}

	// 如果 [1,n-2] 中有最小值，那么先排序 [0,n-2]，把最小值排在最前面，然后排序 [1,n-1] 即可
	// 如果 [1,n-2] 中有最大值，那么先排序 [1,n-1]，把最大值排在最后面，然后排序 [0,n-2] 即可
	for _, ch := range t[1 : n-1] {
		if ch == mn || ch == mx {
			return 2
		}
	}

	// 现在只剩下一种情况：s[0] 是最大值，s[n-1] 是最小值，且 [1,n-2] 不含最小值和最大值
	// 先排序 [0,n-2]，把最大值排到 n-2
	// 然后排序 [1,n-1]，把最大值排在最后面，且最小值排在 1
	// 最后排序 [0,n-2]，把最小值排在最前面
	return 3
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。部分语言把字符串转成了列表，只是为了处理起来方便，可以不转成列表，从而做到 $\mathcal{O}(1)$ 空间。

## 专题训练

见下面思维题单的「**§5.7 分类讨论**」。

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
