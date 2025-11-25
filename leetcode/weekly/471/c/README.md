## 问题拆解

分成如下三类问题，依次解答：

1. 子串只包含一种字母。
2. 子串只包含两种字母。
3. 子串包含三种字母。

## 子串只包含一种字母

即 $s$ 中的最长连续相同子串的长度。这题是 [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/)。

这可以用**分组循环**解决。

**适用场景**：按照题目要求，序列会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：

- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

[分组循环详细讲解](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)。

## 子串只包含两种字母

同样地，用分组循环分组，每组只包含两种字母。

对于每一组，计算含有相同数量的两种字母的最长子串。这题是 [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)。

做法见 [我的题解](https://leetcode.cn/problems/contiguous-array/solutions/3805089/shi-zi-bian-xing-mei-ju-you-wei-hu-zuo-p-x9q2/)。

## 子串包含三种字母

仿照 525 题的做法，设 $\texttt{a}$ 在这个组的**个数前缀和**数组为 $S_a$，$\texttt{b}$ 在这个组的个数前缀和数组为 $S_b$，$\texttt{c}$ 在这个组的个数前缀和数组为 $S_c$。

子串 $[l,r)$ 中的字母 $\texttt{a},\texttt{b},\texttt{c}$ 的出现次数相等，可以拆分为如下两个约束：

1. 子串 $[l,r)$ 中的字母 $\texttt{a}$ 和 $\texttt{b}$ 的出现次数相等。
2. 子串 $[l,r)$ 中的字母 $\texttt{b}$ 和 $\texttt{c}$ 的出现次数相等。

只要满足这两个约束，由等号的传递性可知，子串 $[l,r)$ 中的字母 $\texttt{a}$ 和 $\texttt{c}$ 的出现次数相等，即三个字母的出现次数都相等。

两个约束即如下两个等式

$$
\begin{aligned}
S_a[r] - S_b[r] &= S_a[l] - S_b[l]      \\
S_b[r] - S_c[r] &= S_b[l] - S_c[l]      \\
\end{aligned}
$$

定义数组 $a[i] = (S_a[i] - S_b[i], S_b[i] - S_c[i])$，问题变成：

- 计算数组 $a$ 中的一对相等元素的最远距离。

做法同上。

[本题视频讲解](https://www.bilibili.com/video/BV1FJ4uz1EkN/?t=5m42s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestBalanced(self, s: str) -> int:
        n = len(s)

        # 一种字母
        ans = i = 0
        while i < n:
            start = i
            i += 1
            while i < n and s[i] == s[i - 1]:
                i += 1
            ans = max(ans, i - start)

        # 两种字母
        def f(x: str, y: str) -> None:
            nonlocal ans
            i = 0
            while i < n:
                pos = {0: i - 1}  # 前缀和数组的首项是 0，位置相当于在 i-1
                d = 0  # x 的个数减去 y 的个数
                while i < n and (s[i] == x or s[i] == y):
                    d += 1 if s[i] == x else -1
                    if d in pos:
                        ans = max(ans, i - pos[d])
                    else:
                        pos[d] = i
                    i += 1
                i += 1

        f('a', 'b')
        f('a', 'c')
        f('b', 'c')

        # 三种字母
        # 前缀和数组的首项是 0，位置相当于在 -1
        pos = {(0, 0): -1}
        cnt = defaultdict(int)
        for i, b in enumerate(s):
            cnt[b] += 1
            p = (cnt['a'] - cnt['b'], cnt['b'] - cnt['c'])
            if p in pos:
                ans = max(ans, i - pos[p])
            else:
                pos[p] = i
        return ans
```

```java [sol-Java]
class Solution {
    public int longestBalanced(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;

        // 一种字母
        for (int i = 0; i < n; ) {
            int start = i;
            for (i++; i < n && s[i] == s[i - 1]; i++) ;
            ans = Math.max(ans, i - start);
        }

        // 两种字母
        ans = Math.max(ans, f(s, 'a', 'b'));
        ans = Math.max(ans, f(s, 'a', 'c'));
        ans = Math.max(ans, f(s, 'b', 'c'));

        // 三种字母
        // 把 (x, y) 压缩成一个 long，方便保存至哈希表
        // (x, y) 变成 (x + n) << 20 | (y + n)，其中 +n 避免出现负数
        Map<Long, Integer> pos = new HashMap<>();
        pos.put((long) n << 20 | n, -1); // 前缀和数组的首项是 0，位置相当于在 -1
        int[] cnt = new int[3];
        for (int i = 0; i < n; i++) {
            cnt[s[i] - 'a']++;
            long p = (long) (cnt[0] - cnt[1] + n) << 20 | (cnt[1] - cnt[2] + n);
            if (pos.containsKey(p)) {
                ans = Math.max(ans, i - pos.get(p));
            } else {
                pos.put(p, i);
            }
        }
        return ans;
    }

    private int f(char[] s, char x, char y) {
        int n = s.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            Map<Integer, Integer> pos = new HashMap<>();
            pos.put(0, i - 1); // 前缀和数组的首项是 0，位置相当于在 i-1
            int d = 0; // x 的个数减去 y 的个数
            for (; i < n && (s[i] == x || s[i] == y); i++) {
                d += s[i] == x ? 1 : -1;
                if (pos.containsKey(d)) {
                    ans = Math.max(ans, i - pos.get(d));
                } else {
                    pos.put(d, i);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int ans = 0;
        
        // 一种字母
        for (int i = 0; i < n;) {
            int start = i;
            for (i++; i < n && s[i] == s[i - 1]; i++);
            ans = max(ans, i - start);
        }

        // 两种字母
        auto f = [&](char x, char y) -> void {
            for (int i = 0; i < n; i++) {
                unordered_map<int, int> pos = {{0, i - 1}}; // 前缀和数组的首项是 0，位置相当于在 i-1
                int d = 0; // x 的个数减去 y 的个数
                for (; i < n && (s[i] == x || s[i] == y); i++) {
                    d += s[i] == x ? 1 : -1;
                    if (pos.contains(d)) {
                        ans = max(ans, i - pos[d]);
                    } else {
                        pos[d] = i;
                    }
                }
            }
        };
        f('a', 'b');
        f('a', 'c');
        f('b', 'c');

        // 三种字母
        // 把 (x, y) 压缩成一个 long long，方便保存至哈希表
        // (x, y) 变成 (x + n) << 32 | (y + n)，其中 +n 避免出现负数
        unordered_map<long long, int> pos = {{1LL * n << 32 | n, -1}}; // 前缀和数组的首项是 0，位置相当于在 -1
        int cnt[3]{};
        for (int i = 0; i < n; i++) {
            cnt[s[i] - 'a']++;
            long long p = 1LL * (cnt[0] - cnt[1] + n) << 32 | (cnt[1] - cnt[2] + n);
            if (pos.contains(p)) {
                ans = max(ans, i - pos[p]);
            } else {
                pos[p] = i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	n := len(s)

	// 一种字母
	for i := 0; i < n; {
		start := i
		for i++; i < n && s[i] == s[i-1]; i++ {
		}
		ans = max(ans, i-start)
	}

	// 两种字母
	f := func(x, y byte) {
		for i := 0; i < n; i++ {
			pos := map[int]int{0: i - 1} // 前缀和数组的首项是 0，位置相当于在 i-1
			d := 0 // x 的个数减去 y 的个数
			for ; i < n && (s[i] == x || s[i] == y); i++ {
				if s[i] == x {
					d++
				} else {
					d--
				}
				if j, ok := pos[d]; ok {
					ans = max(ans, i-j)
				} else {
					pos[d] = i
				}
			}
		}
	}
	f('a', 'b')
	f('a', 'c')
	f('b', 'c')

	// 三种字母
	type pair struct{ diffAB, diffBC int }
	pos := map[pair]int{{}: -1} // 前缀和数组的首项是 0，位置相当于在 -1
	cnt := [3]int{}
	for i, b := range s {
		cnt[b-'a']++
		p := pair{cnt[0] - cnt[1], cnt[1] - cnt[2]}
		if j, ok := pos[p]; ok {
			ans = max(ans, i-j)
		} else {
			pos[p] = i
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 下面双指针题单的「**六、分组循环**」。
2. 下面数据结构题单的「**§1.2 前缀和与哈希表**」。

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
