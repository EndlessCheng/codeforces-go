## 方法一：前后缀分解

### 写法一

设修改的字母与 $t[i]$ 匹配。枚举 $i=0,1,2\ldots,|t|-1$，我们需要知道：

- 看左边，设 $t$ 的前缀 $[0,i-1]$ 中的最长匹配子序列为 $s$ 的前缀 $[0,\textit{pre}[i-1]]$。
- 看右边，设 $t$ 的后缀 $[i+1,|t|-1]$ 中的最长匹配子序列为 $s$ 的后缀 $[\textit{suf}[i+1],|s|-1]$。
- 如果 $s$ 的前缀 $[0,\textit{pre}[i-1]]$ 和后缀 $[\textit{suf}[i+1],|s|-1]$ 之间至多有一个字母，那么把这个字母改成 $t[i]$，即可让 $s$ 是 $t$ 的子序列。也就是 $(\textit{pre}[i-1]+1) + 1 + (|s| - \textit{suf}[i+1]) \ge |s|$，即 $\textit{pre}[i-1]+2\ge \textit{suf}[i+1]$。

> 如果 $s$ 不是 $t$ 的子序列，上式中的 $\ge $ 可以写成 $=$。

如何计算 $\textit{pre}$ 和 $\textit{suf}$？见 [392. 判断子序列](https://leetcode.cn/problems/is-subsequence/)，[我的题解](https://leetcode.cn/problems/is-subsequence/solution/jian-ji-xie-fa-pythonjavaccgojsrust-by-e-mz22/)。  

代码实现时，可以先计算 $\textit{suf}$，然后在枚举 $i$ 的同时计算 $\textit{pre}$。

```py [sol-Python3]
class Solution:
    def canMakeSubsequence(self, s: str, t: str) -> bool:
        n, m = len(s), len(t)

        # s[suf[i]:] 是 t[i:] 的子序列
        suf = [0] * (m + 1)
        suf[m] = n
        j = n
        for i in range(m - 1, -1, -1):
            if s[j - 1] == t[i]:
                j -= 1
                if j == 0:
                    # s 已是 t 的子序列
                    return True
            suf[i] = j

        pre = -1
        for i, ch in enumerate(t):
            # 此时 s[:pre+1] 是 t[:i] 的子序列
            if pre + 2 == suf[i + 1]:  # 公式推导见题解
                return True
            if s[pre + 1] == ch:
                pre += 1

        return False
```

```java [sol-Java]
class Solution {
    public boolean canMakeSubsequence(String S, String T) {
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;
        int m = t.length;

        // s[suf[i]:] 是 t[i:] 的子序列
        int[] suf = new int[m + 1];
        suf[m] = n;
        int j = n;
        for (int i = m - 1; i >= 0; i--) {
            if (s[j - 1] == t[i]) {
                j--;
                if (j == 0) {
                    // s 已是 t 的子序列
                    return true;
                }
            }
            suf[i] = j;
        }

        int pre = -1;
        for (int i = 0; i < m; i++) {
            // 此时 s[:pre+1] 是 t[:i] 的子序列
            if (pre + 2 == suf[i + 1]) { // 公式推导见题解
                return true;
            }
            if (s[pre + 1] == t[i]) {
                pre++;
            }
        }

        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canMakeSubsequence(string s, string t) {
        int n = s.size(), m = t.size();

        // s[suf[i]:] 是 t[i:] 的子序列
        vector<int> suf(m + 1);
        suf[m] = n;
        int j = n;
        for (int i = m - 1; i >= 0; i--) {
            if (s[j - 1] == t[i]) {
                j--;
                if (j == 0) {
                    // s 已是 t 的子序列
                    return true;
                }
            }
            suf[i] = j;
        }

        int pre = -1;
        for (int i = 0; i < m; i++) {
            // 此时 s[:pre+1] 是 t[:i] 的子序列
            if (pre + 2 == suf[i + 1]) { // 公式推导见题解
                return true;
            }
            if (s[pre + 1] == t[i]) {
                pre++;
            }
        }

        return false;
    }
};
```

```go [sol-Go]
func canMakeSubsequence(s, t string) bool {
	n, m := len(s), len(t)

	// s[suf[i]:] 是 t[i:] 的子序列
	suf := make([]int, m+1)
	suf[m] = n
	j := n
	for i := m - 1; i >= 0; i-- {
		if s[j-1] == t[i] {
			j--
			if j == 0 {
				// s 已是 t 的子序列
				return true
			}
		}
		suf[i] = j
	}

	pre := -1
	for i, ch := range t {
		// 此时 s[:pre+1] 是 t[:i] 的子序列
		if pre+2 == suf[i+1] { // 公式推导见题解
			return true
		}
		if s[pre+1] == byte(ch) {
			pre++
		}
	}

	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

### 写法二

枚举我们修改了哪个 $s[i]$，其中 $i=0,1,2\ldots,|s|-1$，我们需要知道：

- 看左边，设 $s$ 的前缀 $[0,i-1]$ 是 $t$ 的前缀 $[0,\textit{pre}[i-1]]$ 的子序列。
- 看右边，设 $s$ 的后缀 $[i+1,|s|-1]$ 是 $t$ 的后缀 $[\textit{suf}[i+1],|t|-1]$ 的子序列。
- 如果 $\textit{pre}[i-1]$ 和 $\textit{suf}[i+1]$ 之间至少有一个下标 $j$，也就是 $\textit{suf}[i+1] - \textit{pre}[i-1] > 1$，那么就可以把 $s[i]$ 改成 $t[j]$，使 $s$ 是 $t$ 的子序列。
- 所以 $\textit{pre}[i-1]$ 越小越好，$\textit{suf}[i+1]$ 越大越好。

```py [sol-Python3]
class Solution:
    def canMakeSubsequence(self, s: str, t: str) -> bool:
        n, m = len(s), len(t)

        # s[i:] 是 t[suf[i]:] 的子序列（如果 suf[i]=-1 则不是子序列）
        suf = [0] * (n + 1)
        suf[n] = m
        j = m
        for i in range(n - 1, -1, -1):
            # 上一轮循环 s[i+1] 匹配了 t[j]，减一后继续匹配 s[i]
            j -= 1
            while j >= 0 and t[j] != s[i]:
                j -= 1
            suf[i] = j

        if suf[0] >= 0:
            # s 已是 t 的子序列
            return True

        pre = -1
        for i, ch in enumerate(s):
            # 此时 s[:i] 是 t[:pre+1] 的子序列（如果 pre=m 则不是子序列）
            # 修改 s[i]，那么在 pre 和 suf[i+1] 之间，至少要有一个字母
            if suf[i + 1] - pre > 1:
                return True

            # 上一轮循环 s[i-1] 匹配了 t[pre]，加一后继续匹配 s[i]
            pre += 1
            while pre < m and t[pre] != ch:
                pre += 1

        return False
```

```java [sol-Java]
class Solution {
    public boolean canMakeSubsequence(String S, String T) {
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;
        int m = t.length;

        // s[i,n-1] 是 t[suf[i]m-1] 的子序列（如果 suf[i]=-1 则不是子序列）
        int[] suf = new int[n + 1];
        suf[n] = m;
        int j = m;
        for (int i = n - 1; i >= 0; i--) {
            // 上一轮循环 s[i+1] 匹配了 t[j]，减一后继续匹配 s[i]
            j--;
            while (j >= 0 && t[j] != s[i]) {
                j--;
            }
            suf[i] = j;
        }

        if (suf[0] >= 0) {
            // s 已是 t 的子序列
            return true;
        }

        int pre = -1;
        for (int i = 0; i < n; i++) {
            // 此时 s[0,i-1] 是 t[0,pre] 的子序列（如果 pre=m 则不是子序列）
            // 修改 s[i]，那么在 pre 和 suf[i+1] 之间，至少要有一个字母
            if (suf[i + 1] - pre > 1) {
                return true;
            }

            // 上一轮循环 s[i-1] 匹配了 t[pre]，加一后继续匹配 s[i]
            pre++;
            while (pre < m && t[pre] != s[i]) {
                pre++;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canMakeSubsequence(string s, string t) {
        int n = s.size(), m = t.size();

        // s[i,n-1] 是 t[suf[i],m-1] 的子序列（如果 suf[i]=-1 则不是子序列）
        vector<int> suf(n + 1);
        suf[n] = m;
        int j = m;
        for (int i = n - 1; i >= 0; i--) {
            // 上一轮循环 s[i+1] 匹配了 t[j]，减一后继续匹配 s[i]
            j--;
            while (j >= 0 && t[j] != s[i]) {
                j--;
            }
            suf[i] = j;
        }

        if (suf[0] >= 0) {
            // s 已是 t 的子序列
            return true;
        }

        int pre = -1;
        for (int i = 0; i < n; i++) {
            // 此时 s[0,i-1] 是 t[0,pre] 的子序列（如果 pre=m 则不是子序列）
            // 修改 s[i]，那么在 pre 和 suf[i+1] 之间，至少要有一个字母
            if (suf[i + 1] - pre > 1) {
                return true;
            }

            // 上一轮循环 s[i-1] 匹配了 t[pre]，加一后继续匹配 s[i]
            pre++;
            while (pre < m && t[pre] != s[i]) {
                pre++;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func canMakeSubsequence(s, t string) bool {
	n, m := len(s), len(t)

	// s[i:] 是 t[suf[i]:] 的子序列（如果 suf[i]=-1 则不是子序列）
	suf := make([]int, n+1)
	suf[n] = m
	j := m
	for i := n - 1; i >= 0; i-- {
		// 上一轮循环 s[i+1] 匹配了 t[j]，减一后继续匹配 s[i]
		j--
		for j >= 0 && t[j] != s[i] {
			j--
		}
		suf[i] = j
	}

	if suf[0] >= 0 {
		// s 已是 t 的子序列
		return true
	}

	pre := -1
	for i, ch := range s {
		// 此时 s[:i] 是 t[:pre+1] 的子序列（如果 pre=m 则不是子序列）
		// 修改 s[i]，那么在 pre 和 suf[i+1] 之间，至少要有一个字母
		if suf[i+1]-pre > 1 {
			return true
		}

		// 上一轮循环 s[i-1] 匹配了 t[pre]，加一后继续匹配 s[i]
		pre++
		for pre < m && t[pre] != byte(ch) {
			pre++
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。虽然写了个二重循环，但内层循环中的下标只会减小（或增大），所以二重循环的总循环次数是 $\mathcal{O}(n+m)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：状态机 DP

在 392 题中，我们定义 $j_0$ 表示在不修改的情况下，$s$ 的前缀 $[0, j_0-1]$ 是 $t$ 的（当前正在遍历的）前缀的子序列，且 $j_0$ 尽量大。

对于本题，我们定义 $j_1$ 表示在改过一次的情况下，$s$ 的前缀 $[0, j_1-1]$ 是 $t$ 的（当前正在遍历的）前缀的子序列，且 $j_1$ 尽量大。

对于 $j_1$，有两种情况：

- 普通匹配：如果 $s[j_1] = t[i]$，那么 $j_1$ 增加一。
- 修改：把 $s[j_0]$ 改成 $t[i]$，那么 $j_1$ 在 $j_0$ 的基础上加一。

两种情况取最大值。

对于 $j_0$，只能普通匹配：如果 $s[j_0] = t[i]$，那么 $j_0$ 增加一。

任意时刻，只要 $j_1=|s|$，则说明 $s$ 是 $t$ 的子序列。注：由于 $j_0\le j_1$，我们只需比较 $j_1$ 和 $|s|$。

```py [sol-Python3]
class Solution:
    def canMakeSubsequence(self, s: str, t: str) -> bool:
        n = len(s)
        if n > len(t):
            return False

        j0 = 0  # 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
        j1 = 0  # 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列

        for ch in t:
            # j1 普通匹配
            if s[j1] == ch:
                j1 += 1

            # 也可以修改 s[j0] 为 ch，强行匹配
            j1 = max(j1, j0 + 1)

            # j0 普通匹配
            if s[j0] == ch:
                j0 += 1

            if j1 == n:
                # s 是 t 的子序列
                return True

        return False
```

```java [sol-Java]
class Solution {
    public boolean canMakeSubsequence(String S, String t) {
        if (S.length() > t.length()) {
            return false;
        }

        char[] s = S.toCharArray();
        int n = s.length;
        int j0 = 0; // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
        int j1 = 0; // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
        for (char ch : t.toCharArray()) {
            // j1 普通匹配
            if (s[j1] == ch) {
                j1++;
            }

            // 也可以修改 s[j0] 为 ch，强行匹配
            j1 = Math.max(j1, j0 + 1);

            // j0 普通匹配
            if (s[j0] == ch) {
                j0++;
            }

            if (j1 == n) {
                // s 是 t 的子序列
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canMakeSubsequence(string s, string t) {
        int n = s.size();
        if (n > t.size()) {
            return false;
        }

        int j0 = 0; // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
        int j1 = 0; // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
        for (char ch : t) {
            // j1 普通匹配
            if (s[j1] == ch) {
                j1++;
            }

            // 也可以修改 s[j0] 为 ch，强行匹配
            j1 = max(j1, j0 + 1);

            // j0 普通匹配
            if (s[j0] == ch) {
                j0++;
            }

            if (j1 == n) {
                // s 是 t 的子序列
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func canMakeSubsequence(s, t string) bool {
	n := len(s)
	if n > len(t) {
		return false
	}

	j0 := 0 // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
	j1 := 0 // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
	for _, ch := range t {
		// j1 普通匹配
		if s[j1] == byte(ch) {
			j1++
		}

		// 也可以修改 s[j0] 为 ch，强行匹配
		j1 = max(j1, j0+1)

		// j0 普通匹配
		if s[j0] == byte(ch) {
			j0++
		}

		if j1 == n {
			// s 是 t 的子序列
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m)$，其中 $m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 双指针题单的「**§4.2 判断子序列**」。
2. 动态规划题单的「**六、状态机 DP**」。
3. 动态规划题单的「**专题：前后缀分解**」。

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
