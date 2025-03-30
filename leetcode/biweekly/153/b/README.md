⚠**注意**：题目求的是 $\texttt{1}$ 的个数，并没有要求这些 $\texttt{1}$ 是连续的。

根据题意，答案为 $s$ 中 $\texttt{1}$ 的个数，加上一个 $\texttt{010}$ 子串中的 $\texttt{0}$ 的个数。这里 $\texttt{010}$ 子串是指一段连续的 $\texttt{0}$，紧跟着一段连续的 $\texttt{1}$，再紧跟着一段连续的 $\texttt{0}$。

我们需要找一个 $\texttt{0}$ 最多的 $\texttt{010}$ 子串。

遍历 $s$ 的过程中，记录连续相同段的长度 $\textit{cnt}$，以及上一段连续 $\texttt{0}$ 的个数 $\textit{pre}_0$。如果当前这段是 $\texttt{0}$，那么用 $\textit{pre}_0+\textit{cnt}$ 更新 $\textit{mx}$。

最终答案为 $s$ 中 $\texttt{1}$ 的个数加上 $\textit{mx}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JrZzYhEHt/?t=1m33s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxActiveSectionsAfterTrade(self, s: str) -> int:
        total1 = mx = cnt = 0
        pre0 = -inf
        for i, b in enumerate(s):
            cnt += 1
            if i == len(s) - 1 or b != s[i + 1]:  # i 是这一段的末尾
                if b == '1':
                    total1 += cnt
                else:
                    mx = max(mx, pre0 + cnt)
                    pre0 = cnt
                cnt = 0
        return total1 + mx
```

```py [sol-Python3 groupby]
class Solution:
    def maxActiveSectionsAfterTrade(self, s: str) -> int:
        total1 = mx = 0
        pre0 = -inf
        for b, group in groupby(s):
            cnt = len(list(group))
            if b == '1':
                total1 += cnt
            else:
                mx = max(mx, pre0 + cnt)
                pre0 = cnt
        return total1 + mx
```

```java [sol-Java]
class Solution {
    public int maxActiveSectionsAfterTrade(String S) {
        char[] s = S.toCharArray();
        int total1 = 0;
        int mx = 0;
        int pre0 = Integer.MIN_VALUE;
        int cnt = 0;
        for (int i = 0; i < s.length; i++) {
            cnt++;
            if (i == s.length - 1 || s[i] != s[i + 1]) { // i 是这一段的末尾
                if (s[i] == '1') {
                    total1 += cnt;
                } else {
                    mx = Math.max(mx, pre0 + cnt);
                    pre0 = cnt;
                }
                cnt = 0;
            }
        }
        return total1 + mx;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxActiveSectionsAfterTrade(string s) {
        int total1 = 0, mx = 0, pre0 = INT_MIN, cnt = 0;
        for (int i = 0; i < s.size(); i++) {
            cnt++;
            if (i == s.size() - 1 || s[i] != s[i + 1]) { // i 是这一段的末尾
                if (s[i] == '1') {
                    total1 += cnt;
                } else {
                    mx = max(mx, pre0 + cnt);
                    pre0 = cnt;
                }
                cnt = 0;
            }
        }
        return total1 + mx;
    }
};
```

```go [sol-Go]
func maxActiveSectionsAfterTrade(s string) (ans int) {
	mx := 0
	pre0 := math.MinInt
	cnt := 0
	for i := range len(s) {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] { // i 是这一段的末尾
			if s[i] == '1' {
				ans += cnt
			} else {
				mx = max(mx, pre0+cnt)
				pre0 = cnt
			}
			cnt = 0
		}
	}
	return ans + mx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面双指针题单中的「**六、分组循环**」。

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
