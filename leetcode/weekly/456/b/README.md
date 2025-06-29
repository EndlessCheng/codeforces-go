## 方法一：前后缀分解

下文把最长公共前缀记作 LCP。

类似 [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)，移除 $\textit{words}[i]$ 后，问题变成：

- 计算前缀 $[0,i-1]$ 中的相邻 LCP 长度的最大值。
- 计算 $\textit{words}[i-1]$ 和 $\textit{words}[i+1]$ 的 LCP 长度。
- 计算后缀 $[i+1,n-1]$ 中的相邻 LCP 长度的最大值。

三者取最大值，即为 $\textit{ans}[i]$。

前缀相邻 LCP 长度的最大值可以正着递推计算（类比 238 题计算前缀乘积）。

后缀相邻 LCP 长度的最大值可以倒着递推计算。

代码实现时，计算前缀的循环可以和计算答案的循环合并。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1j6gZzqEdc/?t=10m6s)，欢迎点赞关注~

```py [sol-Python3]
@cache  # 避免重复计算
def lcp(s: str, t: str) -> int:
    cnt = 0
    for x, y in zip(s, t):
        if x != y:
            break
        cnt += 1
    return cnt

class Solution:
    def longestCommonPrefix(self, words: List[str]) -> List[int]:
        n = len(words)
        ans = [0] * n
        if n == 1:  # 不存在相邻对
            return ans

        # 后缀 [i,n-1] 中的相邻 LCP 长度的最大值
        suf_max = [0] * n
        for i in range(n - 2, 0, -1):
            suf_max[i] = max(suf_max[i + 1], lcp(words[i], words[i + 1]))

        ans[0] = suf_max[1]
        pre_max = 0  # 前缀 [0,i-1] 中的相邻 LCP 长度的最大值
        for i in range(1, n - 1):
            ans[i] = max(pre_max, lcp(words[i - 1], words[i + 1]), suf_max[i + 1])
            pre_max = max(pre_max, lcp(words[i - 1], words[i]))  # 为下一轮循环做准备
        ans[-1] = pre_max
        return ans
```

```java [sol-Java]
class Solution {
    public int[] longestCommonPrefix(String[] words) {
        int n = words.length;
        int[] ans = new int[n];
        if (n == 1) { // 不存在相邻对
            return ans;
        }

        // 后缀 [i,n-1] 中的相邻 LCP 长度的最大值
        int[] sufMax = new int[n];
        for (int i = n - 2; i > 0; i--) {
            sufMax[i] = Math.max(sufMax[i + 1], lcp(words[i], words[i + 1]));
        }

        ans[0] = sufMax[1];
        int preMax = 0; // 前缀 [0,i-1] 中的相邻 LCP 长度的最大值
        for (int i = 1; i < n - 1; i++) {
            ans[i] = Math.max(Math.max(preMax, lcp(words[i - 1], words[i + 1])), sufMax[i + 1]);
            preMax = Math.max(preMax, lcp(words[i - 1], words[i])); // 为下一轮循环做准备
        }
        ans[n - 1] = preMax;
        return ans;
    }

    private int lcp(String s, String t) {
        int n = Math.min(s.length(), t.length());
        int cnt = 0;
        for (int i = 0; i < n && s.charAt(i) == t.charAt(i); i++) {
            cnt++;
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    int lcp(const string& s, const string& t) {
        int n = min(s.size(), t.size());
        int cnt = 0;
        for (int i = 0; i < n && s[i] == t[i]; i++) {
            cnt++;
        }
        return cnt;
    }

public:
    vector<int> longestCommonPrefix(vector<string>& words) {
        int n = words.size();
        vector<int> ans(n);
        if (n == 1) { // 不存在相邻对
            return ans;
        }

        // 后缀 [i,n-1] 中的相邻 LCP 长度的最大值
        vector<int> suf_max(n);
        for (int i = n - 2; i > 0; i--) {
            suf_max[i] = max(suf_max[i + 1], lcp(words[i], words[i + 1]));
        }

        ans[0] = suf_max[1];
        int pre_max = 0; // 前缀 [0,i-1] 中的相邻 LCP 长度的最大值
        for (int i = 1; i < n - 1; i++) {
            ans[i] = max({pre_max, lcp(words[i - 1], words[i + 1]), suf_max[i + 1]});
            pre_max = max(pre_max, lcp(words[i - 1], words[i])); // 为下一轮循环做准备
        }
        ans[n - 1] = pre_max;
        return ans;
    }
};
```

```go [sol-Go]
func lcp(s, t string) (cnt int) {
	n := min(len(s), len(t))
	for i := 0; i < n && s[i] == t[i]; i++ {
		cnt++
	}
	return
}

func longestCommonPrefix(words []string) []int {
	n := len(words)
	ans := make([]int, n)
	if n == 1 { // 不存在相邻对
		return ans
	}

	// 后缀 [i,n-1] 中的相邻 LCP 长度的最大值
	sufMax := make([]int, n)
	for i := n - 2; i > 0; i-- {
		sufMax[i] = max(sufMax[i+1], lcp(words[i], words[i+1]))
	}

	ans[0] = sufMax[1]
	preMax := 0 // 前缀 [0,i-1] 中的相邻 LCP 长度的最大值
	for i := 1; i < n-1; i++ {
		ans[i] = max(preMax, lcp(words[i-1], words[i+1]), sufMax[i+1])
		preMax = max(preMax, lcp(words[i-1], words[i])) // 为下一轮循环做准备
	}
	ans[n-1] = preMax
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L\le 10^5$ 是 $\textit{words}[i]$ 的长度之和。每个 $\textit{words}[i]$ 至多参与 $6$ 次（线性时间的）LCP 的计算。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：维护前三大 LCP

遍历所有相邻字符串对，计算 LCP，维护前三大 LCP 长度以及下标（下标取相邻字符串的左边那个）。

然后计算 $\textit{ans}[i]$。设 $\ell$ 为 $\textit{words}[i-1]$ 与 $\textit{words}[i+1]$ 的 LCP 长度。

分类讨论：

- 设最大 LCP 的长度为 $\textit{mx}_1$，下标为 $(i_1,i_1+1)$。如果 $i\ne i_1$ 且 $i\ne i_1+1$，说明最大 LCP 没被破坏，答案为 $\max(\textit{mx}_1,\ell)$。
- 设次大 LCP 的长度为 $\textit{mx}_2$，下标为 $(i_2,i_2+1)$。如果 $i\ne i_2$ 且 $i\ne i_2+1$，说明次大 LCP 没被破坏，答案为 $\max(\textit{mx}_2,\ell)$。
- 否则第三大 LCP 一定没被破坏，答案为 $\max(\textit{mx}_3,\ell)$，其中 $\textit{mx}_3$ 为第三大的 LCP 的长度。

```py [sol-Python3]
@cache  # 不加就是 O(1) 空间，但测试发现，加这个更快（可能有些测试数据有很多重复的字符串）
def lcp(s: str, t: str) -> int:
    cnt = 0
    for x, y in zip(s, t):
        if x != y:
            break
        cnt += 1
    return cnt

class Solution:
    def longestCommonPrefix(self, words: List[str]) -> List[int]:
        n = len(words)
        ans = [0] * n
        if n == 1:
            return ans

        mx1 = mx2 = mx3 = 0
        i1 = i2 = -2
        for i in range(n - 1):
            l = lcp(words[i], words[i + 1])
            if l > mx1:
                mx3 = mx2
                mx2, i2 = mx1, i1
                mx1, i1 = l, i
            elif l > mx2:
                mx3 = mx2
                mx2, i2 = l, i
            elif l > mx3:
                mx3 = l

        for i in range(n):
            l = lcp(words[i - 1], words[i + 1]) if 0 < i < n - 1 else 0
            if i != i1 and i != i1 + 1:  # 最大 LCP 没被破坏
                ans[i] = max(mx1, l)
            elif i != i2 and i != i2 + 1:  # 次大 LCP 没被破坏
                ans[i] = max(mx2, l)
            else:  # 只能是第三大 LCP
                ans[i] = max(mx3, l)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] longestCommonPrefix(String[] words) {
        int n = words.length;
        int[] ans = new int[n];
        if (n == 1) {
            return ans;
        }

        int mx1 = 0, mx2 = 0, mx3 = 0;
        int i1 = -2, i2 = -2;
        for (int i = 0; i < n - 1; i++) {
            int l = lcp(words[i], words[i + 1]);
            if (l > mx1) {
                mx3 = mx2;
                mx2 = mx1;
                mx1 = l;
                i2 = i1;
                i1 = i;
            } else if (l > mx2) {
                mx3 = mx2;
                mx2 = l;
                i2 = i;
            } else if (l > mx3) {
                mx3 = l;
            }
        }

        for (int i = 0; i < n; i++) {
            int l = 0 < i && i < n - 1 ? lcp(words[i - 1], words[i + 1]) : 0;
            if (i != i1 && i != i1 + 1) { // 最大 LCP 没被破坏
                ans[i] = Math.max(mx1, l);
            } else if (i != i2 && i != i2 + 1) { // 次大 LCP 没被破坏
                ans[i] = Math.max(mx2, l);
            } else { // 只能是第三大 LCP
                ans[i] = Math.max(mx3, l);
            }
        }
        return ans;
    }

    private int lcp(String s, String t) {
        int n = Math.min(s.length(), t.length());
        int cnt = 0;
        for (int i = 0; i < n && s.charAt(i) == t.charAt(i); i++) {
            cnt++;
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    int lcp(const string& s, const string& t) {
        int n = min(s.size(), t.size());
        int cnt = 0;
        for (int i = 0; i < n && s[i] == t[i]; i++) {
            cnt++;
        }
        return cnt;
    }

public:
    vector<int> longestCommonPrefix(vector<string>& words) {
        int n = words.size();
        vector<int> ans(n);
        if (n == 1) {
            return ans;
        }

        int mx1 = 0, mx2 = 0, mx3 = 0;
        int i1 = -2, i2 = -2;
        for (int i = 0; i < n - 1; i++) {
            int l = lcp(words[i], words[i + 1]);
            if (l > mx1) {
                mx3 = mx2;
                mx2 = mx1;
                mx1 = l;
                i2 = i1;
                i1 = i;
            } else if (l > mx2) {
                mx3 = mx2;
                mx2 = l;
                i2 = i;
            } else if (l > mx3) {
                mx3 = l;
            }
        }

        for (int i = 0; i < n; i++) {
            int l = 0 < i && i < n - 1 ? lcp(words[i - 1], words[i + 1]) : 0;
            if (i != i1 && i != i1 + 1) { // 最大 LCP 没被破坏
                ans[i] = max(mx1, l);
            } else if (i != i2 && i != i2 + 1) { // 次大 LCP 没被破坏
                ans[i] = max(mx2, l);
            } else { // 只能是第三大 LCP
                ans[i] = max(mx3, l);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func lcp(s, t string) (cnt int) {
	n := min(len(s), len(t))
	for i := 0; i < n && s[i] == t[i]; i++ {
		cnt++
	}
	return
}

func longestCommonPrefix(words []string) []int {
	n := len(words)
	ans := make([]int, n)
	if n == 1 {
		return ans
	}

	mx1, mx2, mx3 := 0, 0, 0
	i1, i2 := -2, -2
	for i := range n - 1 {
		l := lcp(words[i], words[i+1])
		if l > mx1 {
			mx3 = mx2
			mx2, i2 = mx1, i1
			mx1, i1 = l, i
		} else if l > mx2 {
			mx3 = mx2
			mx2, i2 = l, i
		} else if l > mx3 {
			mx3 = l
		}
	}

	for i := range n {
		l := 0
		if 0 < i && i < n-1 {
			l = lcp(words[i-1], words[i+1])
		}
		if i != i1 && i != i1+1 { // 最大 LCP 没被破坏
			ans[i] = max(mx1, l)
		} else if i != i2 && i != i2+1 { // 次大 LCP 没被破坏
			ans[i] = max(mx2, l)
		} else { // 只能是第三大 LCP
			ans[i] = max(mx3, l)
		}
	}
	return ans
}
```

## 方法三：维护前二大 LCP

实际上，我们不需要维护第三大 LCP。

如果上面程序进入计算 $\max(\textit{mx}_3,\ell)$ 的分支，只可能是如下情况：

- 最大 LCP 和次大 LCP 是重叠的，即 $i_1+1 = i_2$ 或者 $i_2+1 = i_1$，且此时 $i$ 正好处在构成最大 LCP 和次大 LCP 的三个字符串的中间。

设这三个字符串分别为 $A,B,C$。删除 $B$ 后，我们计算的是 $\ell = \text{LCP}(A,C)$。

**定理**：$\text{LCP}(A,C)\ge \textit{mx}_2$。

**证明**：不失一般性，假设 $\textit{mx}_1 = \text{LCP}(A,B)$，$\textit{mx}_2 = \text{LCP}(B,C)$。这意味着 $B$ 和 $C$ 的前 $\textit{mx}_2$ 个字母是一样的，$A$ 和 $B$ 的前 $\textit{mx}_1$ 个字母是一样的，由于 $\textit{mx}_1\ge \textit{mx}_2$，所以 $A$ 和 $B$ 的前 $\textit{mx}_2$ 个字母也是一样的。所以 $A$ 和 $C$ 的前 $\textit{mx}_2$ 个字母是一样的，也就是说，$\text{LCP}(A,C)$ 至少是 $\textit{mx}_2$。证毕。

当程序进入计算 $\max(\textit{mx}_3,\ell)$ 的分支时，由于 $\ell = \text{LCP}(A,C) \ge \textit{mx}_2 \ge \textit{mx}_3$，所以 $\max(\textit{mx}_3,\ell) = \ell$，这意味着我们无需维护第三大 LCP 的长度 $\textit{mx}_3$。

```py [sol-Python3]
@cache  # 不加就是 O(1) 空间，但测试发现，加这个更快（可能有些测试数据有很多重复的字符串）
def lcp(s: str, t: str) -> int:
    cnt = 0
    for x, y in zip(s, t):
        if x != y:
            break
        cnt += 1
    return cnt

class Solution:
    def longestCommonPrefix(self, words: List[str]) -> List[int]:
        n = len(words)
        ans = [0] * n
        if n == 1:
            return ans

        mx1 = mx2 = 0
        i1 = i2 = -2
        for i in range(n - 1):
            l = lcp(words[i], words[i + 1])
            if l > mx1:
                mx2, i2 = mx1, i1
                mx1, i1 = l, i
            elif l > mx2:
                mx2, i2 = l, i

        for i in range(n):
            l = lcp(words[i - 1], words[i + 1]) if 0 < i < n - 1 else 0
            if i != i1 and i != i1 + 1:  # 最大 LCP 没被破坏
                ans[i] = max(mx1, l)
            elif i != i2 and i != i2 + 1:  # 次大 LCP 没被破坏
                ans[i] = max(mx2, l)
            else:
                ans[i] = l
        return ans
```

```java [sol-Java]
class Solution {
    public int[] longestCommonPrefix(String[] words) {
        int n = words.length;
        int[] ans = new int[n];
        if (n == 1) {
            return ans;
        }

        int mx1 = 0, mx2 = 0;
        int i1 = -2, i2 = -2;
        for (int i = 0; i < n - 1; i++) {
            int l = lcp(words[i], words[i + 1]);
            if (l > mx1) {
                mx2 = mx1;
                mx1 = l;
                i2 = i1;
                i1 = i;
            } else if (l > mx2) {
                mx2 = l;
                i2 = i;
            }
        }

        for (int i = 0; i < n; i++) {
            int l = 0 < i && i < n - 1 ? lcp(words[i - 1], words[i + 1]) : 0;
            if (i != i1 && i != i1 + 1) { // 最大 LCP 没被破坏
                ans[i] = Math.max(mx1, l);
            } else if (i != i2 && i != i2 + 1) { // 次大 LCP 没被破坏
                ans[i] = Math.max(mx2, l);
            } else {
                ans[i] = l;
            }
        }
        return ans;
    }

    private int lcp(String s, String t) {
        int n = Math.min(s.length(), t.length());
        int cnt = 0;
        for (int i = 0; i < n && s.charAt(i) == t.charAt(i); i++) {
            cnt++;
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    int lcp(const string& s, const string& t) {
        int n = min(s.size(), t.size());
        int cnt = 0;
        for (int i = 0; i < n && s[i] == t[i]; i++) {
            cnt++;
        }
        return cnt;
    }

public:
    vector<int> longestCommonPrefix(vector<string>& words) {
        int n = words.size();
        vector<int> ans(n);
        if (n == 1) {
            return ans;
        }

        int mx1 = 0, mx2 = 0;
        int i1 = -2, i2 = -2;
        for (int i = 0; i < n - 1; i++) {
            int l = lcp(words[i], words[i + 1]);
            if (l > mx1) {
                mx2 = mx1;
                mx1 = l;
                i2 = i1;
                i1 = i;
            } else if (l > mx2) {
                mx2 = l;
                i2 = i;
            }
        }

        for (int i = 0; i < n; i++) {
            int l = 0 < i && i < n - 1 ? lcp(words[i - 1], words[i + 1]) : 0;
            if (i != i1 && i != i1 + 1) { // 最大 LCP 没被破坏
                ans[i] = max(mx1, l);
            } else if (i != i2 && i != i2 + 1) { // 次大 LCP 没被破坏
                ans[i] = max(mx2, l);
            } else {
                ans[i] = l;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func lcp(s, t string) (cnt int) {
	n := min(len(s), len(t))
	for i := 0; i < n && s[i] == t[i]; i++ {
		cnt++
	}
	return
}

func longestCommonPrefix(words []string) []int {
	n := len(words)
	ans := make([]int, n)
	if n == 1 {
		return ans
	}

	mx1, mx2 := 0, 0
	i1, i2 := -2, -2
	for i := range n - 1 {
		l := lcp(words[i], words[i+1])
		if l > mx1 {
			mx2, i2 = mx1, i1
			mx1, i1 = l, i
		} else if l > mx2 {
			mx2, i2 = l, i
		}
	}

	for i := range n {
		l := 0
		if 0 < i && i < n-1 {
			l = lcp(words[i-1], words[i+1])
		}
		if i != i1 && i != i1+1 { // 最大 LCP 没被破坏
			ans[i] = max(mx1, l)
		} else if i != i2 && i != i2+1 { // 次大 LCP 没被破坏
			ans[i] = max(mx2, l)
		} else {
			ans[i] = l
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L\le 10^5$ 是 $\textit{words}[i]$ 的长度之和。每个 $\textit{words}[i]$ 至多参与 $4$ 次（线性时间的）LCP 的计算。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
