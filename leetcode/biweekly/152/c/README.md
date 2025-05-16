## 不删除字符串

先想一想，如果不删除字符串，那么 $k$ 个字符串的最长公共前缀（LCP）的长度是多少？

把字符串排序后，有着相同前缀的字符串可以聚在一起。例如排序后，字符串是

$$
\begin{aligned}
& \texttt{aab} \\
& \texttt{ab}  \\
& \texttt{aba} \\
& \texttt{abb} \\
& \texttt{ac}  \\
& \texttt{b}   \\
\end{aligned}
$$

其中 $\texttt{a}$ 开头的字符串都聚在一起，$\texttt{ab}$ 开头的字符串都聚在一起。

如果要选 $k$ 个字符串，选择长为 $k$ 的连续子数组是更优的，这是排序带来的好处。

如何快速计算子数组的 LCP？

**定理 1**：有序子数组的 LCP，等于子数组第一个字符串和最后一个字符串的 LCP。

例如 $\texttt{ab},\texttt{aba},\texttt{abb}$ 的 LCP 等于 $\texttt{ab}$ 和 $\texttt{abb}$ 的 LCP，即 $\texttt{ab}$。

**证明**：把子数组第一个字符串记作 $s$，最后一个字符串记作 $t$，设 $s$ 和 $t$ 的 LCP 为 $m$。子数组的 LCP 不会比 $m$ 更大，如果更大，那么 $s$ 和 $t$ 的 LCP 会比 $m$ 还大，矛盾。子数组的 LCP 也不会比 $m$ 小，用上面的例子来说，这意味着中间的 $\texttt{aba}$ 的前两个字母不等于 $\texttt{ab}$，但我们已经把字符串排序了，在前后都是 $\texttt{ab}$ 开头的字符串的情况下，中间的字符串也必然以 $\texttt{ab}$ 开头，矛盾。综上所述，子数组的 LCP 恰好等于子数组第一个字符串和最后一个字符串的 LCP。

枚举所有长为 $k$ 的子数组，计算子数组第一个字符串和最后一个字符串的 LCP 长度，取最大值，即为不删除字符串时的答案。

## 删除一个字符串

**引理**：在有序数组中，两个字符串相隔越远，LCP 越短（不会变长）；两个字符串相隔越近，LCP 越长（不会变短）。

**证明**：设有三个字符串 $a,b,c$，其中 $b$ 排在 $a$ 和 $c$ 的中间，并且 $a$ 和 $c$ 的 LCP 是 $m$。根据定理 1 的证明，$a$ 和 $b$ 的 LCP 至少是 $m$，不可能小于 $m$，否则就破坏了数组的有序性。

在不删除的情况下，设所有长为 $k$ 的子数组中，LCP 的最大长度为 $\textit{mx}$，次大长度为 $\textit{mx}_2$。设 $\textit{mx}$ 对应的子数组为 $[\textit{mxI},\textit{mxI}+k-1]$。

**定理 2**：如果删除的字符串不在 $[\textit{mxI},\textit{mxI}+k-1]$ 中，答案是 $\textit{mx}$。

**证明**：需要证明新的 LCP 长度不会超过 $\textit{mx}$。删除一个字符串后，还需要添加一个在子数组外面的字符串，恢复成 $k$ 个字符串。设删除的字符串在某个子数组 $[i,i+k-1]$ 中，为了让 LCP 尽量长，应当选与子数组相邻的字符串，即下标为 $i-1$ 或者 $i+k$ 的字符串。如果删除的字符串下标不是 $i$ 也不是 $i+k-1$，根据定理 1 和引理，删除后的 LCP 来自两个更远的字符串，所以不可能变长；如果删除 $i+k-1$，添加 $i-1$，那么 LCP 来自一个以 $i-1$ 开头的子数组，但其 LCP 长度不会超过 $\textit{mx}$；删除 $i$ 添加 $i+k$ 的情况同理。

**定理 3**：如果删除的字符串在 $[\textit{mxI},\textit{mxI}+k-1]$ 中，答案是 $\textit{mx}_2$。

**证明**：需要证明以下两点。

删除一个字符串后，设新的 LCP 长为 $m$，那么 $m$ 不可能比 $\textit{mx}_2$ 还大。

- 反证，假设 $m > \textit{mx}_2$。如果删除字符串后，添加的是下标为 $\textit{mxI}-1$ 的字符串，说明以 $\textit{mxI}-1$ 开头的长为 $k+1$ 的子数组的 LCP 长为 $m$，那么根据定理 1 和引理，以 $\textit{mxI}-1$ 开头的长为 $k$ 的子数组的 LCP 至少是 $m$，所以我们先前算出来的 $\textit{mx}_2$ 也应当来自以 $\textit{mxI}-1$ 开头的子数组，所以 $\textit{mx}_2 = m$，矛盾。其他「删除-添加」字符串的情况，证明方法类似。

删除一个字符串不会导致 $\textit{mx}_2$ 变小。

- 如果次大子数组不包含删除的字符串，$\textit{mx}_2$ 不会变小。
- 如果次大子数组包含删除的字符串（记作 $s$），说明最大子数组和次大子数组都包含 $s$，这两个子数组是相交的（有公共部分）。在这种情况下，**次大子数组的 LCP，是最大子数组的 LCP 的前缀**。所以从次大子数组中删除 $s$，再增加一个在最大子数组中的字符串（恢复成 $k$ 个字符串），相当于增加了一个前缀包含次大子数组 LCP 的字符串，所以次大子数组的 LCP 是保持不变的，$\textit{mx}_2$ 不会变小。

代码实现时，由于答案和 $\textit{words}$ 的字符串顺序相关，不能直接对 $\textit{words}$ 排序，可以改为对下标排序。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1i6Q8YUEtN/?t=3m29s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    # 计算 s 和 t 的最长公共前缀（LCP）长度
    def calc_lcp(self, s: str, t: str) -> int:
        for i, (x, y) in enumerate(zip(s, t)):
            if x != y:
                return i
        return min(len(s), len(t))

    def longestCommonPrefix(self, words: List[str], k: int) -> List[int]:
        n = len(words)
        if k >= n:  # 移除一个字符串后，剩余字符串少于 k 个
            return [0] * n

        idx = list(range(n))
        idx.sort(key=lambda i: words[i])

        # 计算最大 LCP 长度和次大 LCP 长度，同时记录最大 LCP 来自哪里
        mx = mx2 = mx_i = -1
        for i in range(n - k + 1):
            # 排序后，[i, i+k-1] 的 LCP 等于两端点的 LCP
            lcp = self.calc_lcp(words[idx[i]], words[idx[i + k - 1]])
            if lcp > mx:
                mx, mx2, mx_i = lcp, mx, i
            elif lcp > mx2:
                mx2 = lcp

        ans = [mx] * n  # 先初始化成最大 LCP 长度
        # 移除下标在 idx[mx_i:mx_i+k] 中的字符串，会导致最大 LCP 变成次大 LCP
        for i in idx[mx_i: mx_i + k]:
            ans[i] = mx2  # 改成次大 LCP 长度
        return ans
```

```java [sol-Java]
class Solution {
    public int[] longestCommonPrefix(String[] words, int k) {
        int n = words.length;
        if (k >= n) { // 移除一个字符串后，剩余字符串少于 k 个
            return new int[n];
        }

        Integer[] idx = new Integer[n];
        Arrays.setAll(idx, i -> i);
        Arrays.sort(idx, Comparator.comparing(i -> words[i]));

        // 计算最大 LCP 长度和次大 LCP 长度，同时记录最大 LCP 来自哪里
        int mx = -1, mx2 = -1, mxI = -1;
        for (int i = 0; i <= n - k; i++) {
            // 排序后，[i, i+k-1] 的 LCP 等于两端点的 LCP
            int lcp = calcLCP(words[idx[i]], words[idx[i + k - 1]]);
            if (lcp > mx) {
                mx2 = mx;
                mx = lcp;
                mxI = i;
            } else if (lcp > mx2) {
                mx2 = lcp;
            }
        }

        int[] ans = new int[n];
        Arrays.fill(ans, mx); // 先初始化成最大 LCP 长度
        // 移除下标在 idx[mxI, mxI+k-1] 中的字符串，会导致最大 LCP 变成次大 LCP
        for (int i = mxI; i < mxI + k; i++) {
            ans[idx[i]] = mx2; // 改成次大 LCP 长度
        }
        return ans;
    }

    // 计算 s 和 t 的最长公共前缀（LCP）长度
    private int calcLCP(String s, String t) {
        int len = Math.min(s.length(), t.length());
        for (int i = 0; i < len; i++) {
            if (s.charAt(i) != t.charAt(i)) {
                return i;
            }
        }
        return len;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 计算 s 和 t 的最长公共前缀（LCP）长度
    int calc_lcp(string& s, string& t) {
        int len = min(s.size(), t.size());
        for (int i = 0; i < len; i++) {
            if (s[i] != t[i]) {
                return i;
            }
        }
        return len;
    }

public:
    vector<int> longestCommonPrefix(vector<string>& words, int k) {
        int n = words.size();
        if (k >= n) { // 移除一个字符串后，剩余字符串少于 k 个
            return vector<int>(n);
        }

        vector<int> idx(n);
        ranges::iota(idx, 0);
        ranges::sort(idx, {}, [&](int i) -> auto& { return words[i]; }); // 注意要返回引用，不然会 copy

        // 计算最大 LCP 长度和次大 LCP 长度，同时记录最大 LCP 来自哪里
        int mx = -1, mx2 = -1, mx_i = -1;
        for (int i = 0; i <= n - k; i++) {
            // 排序后，[i, i+k-1] 的 LCP 等于两端点的 LCP
            int lcp = calc_lcp(words[idx[i]], words[idx[i + k - 1]]);
            if (lcp > mx) {
                mx2 = mx;
                mx = lcp;
                mx_i = i;
            } else if (lcp > mx2) {
                mx2 = lcp;
            }
        }

        vector<int> ans(n, mx); // 先初始化成最大 LCP 长度
        // 移除下标在 idx[mx_i, mx_i+k-1] 中的字符串，会导致最大 LCP 变成次大 LCP
        for (int i : ranges::subrange(idx.begin() + mx_i, idx.begin() + mx_i + k)) {
            ans[i] = mx2; // 改成次大 LCP 长度
        }
        return ans;
    }
};
```

```go [sol-Go]
// 计算 s 和 t 的最长公共前缀（LCP）长度
func calcLCP(s, t string) int {
	n := min(len(s), len(t))
	for i := range n {
		if s[i] != t[i] {
			return i
		}
	}
	return n
}

func longestCommonPrefix(words []string, k int) []int {
	n := len(words)
	if k >= n { // 移除一个字符串后，剩余字符串少于 k 个
		return make([]int, n)
	}

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return cmp.Compare(words[i], words[j]) })

	// 计算最大 LCP 长度和次大 LCP 长度，同时记录最大 LCP 来自哪里
	mx, mx2, mxI := -1, -1, -1
	for i := range n - k + 1 {
		// 排序后，[i, i+k-1] 的 LCP 等于两端点的 LCP
		lcp := calcLCP(words[idx[i]], words[idx[i+k-1]])
		if lcp > mx {
			mx, mx2, mxI = lcp, mx, i
		} else if lcp > mx2 {
			mx2 = lcp
		}
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = mx // 先初始化成最大 LCP 长度
	}
	// 移除下标在 idx[mxI:mxI+k] 中的字符串，会导致最大 LCP 变成次大 LCP
	for _, i := range idx[mxI : mxI+k] {
		ans[i] = mx2 // 改成次大 LCP 长度
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L\log n)$，其中 $n$ 是 $\textit{words}$ 的长度，$L$ 是所有 $\textit{words}[i]$ 的长度之和。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
