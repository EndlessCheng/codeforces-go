## 方法一：115 题 + 前后缀分解

先计算不插入字母时的 $\texttt{LCT}$ 子序列个数，然后加上插入字母**额外**产生的 $\texttt{LCT}$ 子序列个数。

不插入字母，问题就是 [115. 不同的子序列](https://leetcode.cn/problems/distinct-subsequences/)，请看 [我的题解](https://leetcode.cn/problems/distinct-subsequences/solutions/3060706/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-9va6/)。

插入字母，分类讨论：

- 插入 $\texttt{L}$，插在 $s$ 的最左边最优，额外产生的 $\texttt{LCT}$ 子序列个数就是 $s$ 中的 $\texttt{CT}$ 子序列个数。复用 115 题代码即可。
- 插入 $\texttt{T}$，插在 $s$ 的最右边最优，额外产生的 $\texttt{LCT}$ 子序列个数就是 $s$ 中的 $\texttt{LC}$ 子序列个数。复用 115 题代码即可。
- 插入 $\texttt{C}$，我们枚举插在 $s[i]$ 和 $s[i+1]$ 之间，根据乘法原理，产生的额外 $\texttt{LCT}$ 子序列个数，等于 $s[0]$ 到 $s[i]$ 中的 $\texttt{L}$ 的个数，乘以 $s[i+1]$ 到 $s[n-1]$ 中的 $\texttt{T}$ 的个数。取所有插入位置的最大值。

三种情况取最大值，即为插入字母所产生的额外 $\texttt{LCT}$ 子序列个数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1pm8vzAEXx/?t=54m38s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    # 115. 不同的子序列
    def numDistinct(self, s: str, t: str) -> int:
        n, m = len(s), len(t)
        if n < m:
            return 0

        f = [1] + [0] * m
        for i, x in enumerate(s):
            for j in range(min(i, m - 1), max(m - n + i, 0) - 1, -1):
                if x == t[j]:
                    f[j + 1] += f[j]
        return f[m]

    # 计算插入 C 额外产生的 LCT 子序列个数的最大值
    def calcInsertC(self, s: str) -> int:
        cnt_t = s.count('T')  # s[i+1] 到 s[n-1] 的 'T' 的个数
        cnt_l = 0  # s[0] 到 s[i] 的 'L' 的个数
        res = 0
        for c in s:
            if c == 'T':
                cnt_t -= 1
            if c == 'L':
                cnt_l += 1
            res = max(res, cnt_l * cnt_t)
        return res

    def numOfSubsequences(self, s: str) -> int:
        extra = max(self.numDistinct(s, "CT"), self.numDistinct(s, "LC"), self.calcInsertC(s))
        return self.numDistinct(s, "LCT") + extra
```

```java [sol-Java]
class Solution {
    public long numOfSubsequences(String S) {
        char[] s = S.toCharArray();
        long extra = Math.max(Math.max(numDistinct(s, "CT"), numDistinct(s, "LC")), calcInsertC(s));
        return numDistinct(s, "LCT") + extra;
    }

    // 115. 不同的子序列
    private long numDistinct(char[] s, String T) {
        int n = s.length;
        int m = T.length();
        if (n < m) {
            return 0;
        }

        char[] t = T.toCharArray();
        long[] f = new long[m + 1];
        f[0] = 1;
        for (int i = 0; i < n; i++) {
            for (int j = Math.min(i, m - 1); j >= Math.max(m - n + i, 0); j--) {
                if (s[i] == t[j]) {
                    f[j + 1] += f[j];
                }
            }
        }
        return f[m];
    }

    // 计算插入 C 额外产生的 LCT 子序列个数的最大值
    private long calcInsertC(char[] s) {
        int cntT = 0;
        for (char c : s) {
            if (c == 'T') {
                cntT++;
            }
        }

        long res = 0;
        int cntL = 0;
        for (char c : s) {
            if (c == 'T') {
                cntT--;
            }
            if (c == 'L') {
                cntL++;
            }
            res = Math.max(res, (long) cntL * cntT);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 115. 不同的子序列
    long long numDistinct(const string& s, const string& t) {
        int n = s.size(), m = t.size();
        if (n < m) {
            return 0;
        }

        vector<long long> f(m + 1);
        f[0] = 1;
        for (int i = 0; i < n; i++) {
            for (int j = min(i, m - 1); j >= max(m - n + i, 0); j--) {
                if (s[i] == t[j]) {
                    f[j + 1] += f[j];
                }
            }
        }
        return f[m];
    }

    // 计算插入 C 额外产生的 LCT 子序列个数的最大值
    long long calcInsertC(string s) {
        int cnt_t = ranges::count(s, 'T'); // s[i+1] 到 s[n-1] 的 'T' 的个数
        int cnt_l = 0; // s[0] 到 s[i] 的 'L' 的个数
        long long res = 0;
        for (char c : s) {
            if (c == 'T') {
                cnt_t--;
            }
            if (c == 'L') {
                cnt_l++;
            }
            res = max(res, 1LL * cnt_l * cnt_t);
        }
        return res;
    }

public:
    long long numOfSubsequences(string s) {
        long long extra = max({numDistinct(s, "CT"), numDistinct(s, "LC"), calcInsertC(s)});
        return numDistinct(s, "LCT") + extra;
    }
};
```

```go [sol-Go]
// 115. 不同的子序列
func numDistinct(s, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	f := make([]int, m+1)
	f[0] = 1
	for i, x := range s {
		for j := min(i, m-1); j >= max(m-n+i, 0); j-- {
			if byte(x) == t[j] {
				f[j+1] += f[j]
			}
		}
	}
	return f[m]
}

// 计算插入 C 额外产生的 LCT 子序列个数的最大值
func calcInsertC(s string) (res int) {
	cntT := strings.Count(s, "T") // s[i+1] 到 s[n-1] 的 'T' 的个数
	cntL := 0 // s[0] 到 s[i] 的 'L' 的个数
	for _, c := range s {
		if c == 'T' {
			cntT--
		}
		if c == 'L' {
			cntL++
		}
		res = max(res, cntL*cntT)
	}
	return
}

func numOfSubsequences(s string) int64 {
	extra := max(numDistinct(s, "CT"), numDistinct(s, "LC"), calcInsertC(s))
	return int64(numDistinct(s, "LCT") + extra)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：状态机 DP + 前后缀分解

方法一计算子序列个数的那三个 DP，可以合并到同一个循环中。

对于 $\texttt{LCT}$ 子序列的个数，计算方式如下：

- 定义 $\textit{l}$ 表示遍历过的 $\texttt{L}$ 的个数。每次遍历到 $\texttt{L}$ 的时候，就把 $\textit{l}$ 增加 $1$。
- 定义 $\textit{lc}$ 表示遍历过的 $\texttt{LC}$ 子序列的个数。每次遍历到 $\texttt{C}$ 的时候，就把 $\textit{lc}$ 增加 $\textit{l}$。
- 定义 $\textit{lct}$ 表示遍历过的 $\texttt{LCT}$ 子序列的个数。每次遍历到 $\texttt{T}$ 的时候，就把 $\textit{lct}$ 增加 $\textit{lc}$。

对于 $\texttt{LC}$ 和 $\texttt{CT}$ 的计算方法同理。

```py [sol-Python3]
class Solution:
    def numOfSubsequences(self, s: str) -> int:
        t = s.count('T')
        l = lc = lct = c = ct = lt = 0
        for b in s:
            if b == 'L':
                l += 1
            elif b == 'C':
                lc += l
                c += 1
            elif b == 'T':
                lct += lc
                ct += c
                t -= 1
            lt = max(lt, l * t)
        return lct + max(ct, lc, lt)
```

```py [sol-Python3 手写 max]
class Solution:
    def numOfSubsequences(self, s: str) -> int:
        t = s.count('T')
        l = lc = lct = c = ct = lt = 0
        for b in s:
            if b == 'L':
                l += 1
            elif b == 'C':
                lc += l
                c += 1
            elif b == 'T':
                lct += lc
                ct += c
                t -= 1
            v = l * t
            if v > lt:
                lt = v
        return lct + max(ct, lc, lt)
```

```java [sol-Java]
class Solution {
    public long numOfSubsequences(String S) {
        char[] s = S.toCharArray();
        int t = 0;
        for (char c : s) {
            if (c == 'T') {
                t++;
            }
        }

        long l = 0, lc = 0, lct = 0, c = 0, ct = 0, lt = 0;
        for (char b : s) {
            if (b == 'L') {
                l++;
            } else if (b == 'C') {
                lc += l;
                c++;
            } else if (b == 'T') {
                lct += lc;
                ct += c;
                t--;
            }
            lt = Math.max(lt, l * t);
        }
        return lct + Math.max(Math.max(ct, lc), lt);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numOfSubsequences(string s) {
        int t = ranges::count(s, 'T');
        long long l = 0, lc = 0, lct = 0, c = 0, ct = 0, lt = 0;
        for (char b : s) {
            if (b == 'L') {
                l++;
            } else if (b == 'C') {
                lc += l;
                c++;
            } else if (b == 'T') {
                lct += lc;
                ct += c;
                t--;
            }
            lt = max(lt, l * t);
        }
        return lct + max({ct, lc, lt});
    }
};
```

```go [sol-Go]
func numOfSubsequences(s string) int64 {
	t := strings.Count(s, "T")
	var l, lc, lct, c, ct, lt int
	for _, b := range s {
		if b == 'L' {
			l++
		} else if b == 'C' {
			lc += l
			c++
		} else if b == 'T' {
			lct += lc
			ct += c
			t--
		}
		lt = max(lt, l*t)
	}
	return int64(lct + max(ct, lc, lt))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[2222. 选择建筑的方案数](https://leetcode.cn/problems/number-of-ways-to-select-buildings/)

## 专题训练

见下面动态规划题单的「**§4.1 最长公共子序列（LCS）**」「**六、状态机 DP**」和「**专题：前后缀分解**」。

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
