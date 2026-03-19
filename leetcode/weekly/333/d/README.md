本文把 $\textit{word}$ 简称为 $s$。

想象有 $n$ 个空位，我们要填入 $n$ 个字母，得到一个字符串 $s$。要求 $s$ 的字典序尽量小。

### 一、确定 s[0] 填什么字母

要让 $s$ 的字典序尽量小，先从 $s[0]$ 开始思考，$s[0]$ 越小越好。

$s[0]$ 可以填 $\texttt{a}$ 吗？

可以，理由如下：

1. 如果最终所有 $s[i]$ 都大于 $\texttt{a}$，那么把所有 $s[i]$ 同时减小，直到有 $s[i] = \texttt{a}$，LCP 会变吗？比如 $\texttt{cbcd}$ 和 $\texttt{cbbc}$ 的 LCP 等于 $2$，所有字母都减小 $1$ 后，得到 $\texttt{babc}$ 和 $\texttt{baab}$，LCP 仍然是 $2$。一般地，当且仅当 $x=y$ 成立时，$x-1=y-1$ 成立。所以这个操作不会影响 LCP。
2. 如果减小后，$s[0] = \texttt{a}$，那么目标达成。如果 $s[0] \ne \texttt{a}$，比如 $s[0]=\texttt{b}$，那么把所有的 $\texttt{b}$ 都替换成 $\texttt{a}$，把所有的 $\texttt{a}$ 都替换成 $\texttt{b}$，LCP 会变吗？接着上面的例子，$\texttt{babc}$ 和 $\texttt{baab}$ 替换后，得到 $\texttt{abac}$ 和 $\texttt{abba}$，LCP 仍然是 $2$。一般地，替换前相等的字母，替换后仍然相等；替换前不相等的字母，替换后一定不相等。所以这个操作不会影响 LCP。

所以在有解的情况下，$s[0] = \texttt{a}$。

### 二、确定其他位置上的字母

根据 $\textit{lcp}[0]$，还有哪些位置一定填 $\texttt{a}$？

根据 LCP 的定义：

- 如果 $\textit{lcp}[0][j]>0$，那么 $s[j] = s[0] = \texttt{a}$。
- 如果 $\textit{lcp}[0][j]=0$，那么 $s[j] \ne s[0] = \texttt{a}$。

所以当且仅当 $\textit{lcp}[0][j]>0$ 时，$s[j] = \texttt{a}$。

现在我们确定了所有填 $\texttt{a}$ 的位置。

下一步，哪些位置要填 $\texttt{b}$？

找到第一个没有填入字母的位置 $i$，为了让 $s$ 的字典序尽量小，$s[i]$ 填 $\texttt{b}$ 是最优的。同理，所有满足 $\textit{lcp}[i][j]>0$ 的 $s[j]$ 都一定是 $\texttt{b}$。

依此类推，用这个方法依次填入字母 $\texttt{a},\texttt{b},\texttt{c},\ldots,\texttt{z}$。

如果 $26$ 个字母都填完后，还有空位，那么无解，返回空串。

### 三、验证 s 是否符合 lcp 数组

在上面的计算过程中，我们只访问了 $\textit{lcp}$ 矩阵中的至多 $26$ 行，其余行完全没有访问过。如果这里面有不匹配的数据呢？

所以构造完 $s$ 后，还需要计算 $s$ 的**实际 LCP 矩阵**，看看是否和输入的 $\textit{lcp}$ 矩阵完全一致。

如何计算后缀 $[i,n-1]$ 和后缀 $[j,n-1]$ 的 LCP？

这个计算过程是一个 DP。分类讨论：

- 如果 $s[i]\ne s[j]$，那么 $\textit{lcp}[i][j]$ 必须是 $0$。
- 如果 $s[i]=s[j]$，那么问题变成计算后缀 $[i+1,n-1]$ 和后缀 $[j+1,n-1]$ 的 LCP，所以必须满足 $\textit{lcp}[i][j] = \textit{lcp}[i+1][j+1]+1$。特别地，如果 $i=n-1$ 或者 $j=n-1$，那么必须满足 $\textit{lcp}[i][j] = 1$。

```py [sol-Python3]
class Solution:
    def findTheString(self, lcp: List[List[int]]) -> str:
        n = len(lcp)
        s = [''] * n
        i = 0  # s[i] 没有填字母
        for c in ascii_lowercase:
            for j in range(i, n):
                if lcp[i][j]:  # s[j] == s[i]
                    s[j] = c
            # 找下一个空位
            while i < n and s[i]:
                i += 1
            if i == n:  # 没有空位
                break

        if i < n:  # 还有空位
            return ""

        # 验证 s 是否符合 lcp 矩阵
        for i in range(n - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                # 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
                actual_lcp = 0 if s[i] != s[j] else (1 if i == n - 1 or j == n - 1 else lcp[i + 1][j + 1] + 1)
                if lcp[i][j] != actual_lcp:  # 矛盾
                    return ""
        return ''.join(s)
```

```java [sol-Java]
class Solution {
    public String findTheString(int[][] lcp) {
        int n = lcp.length;
        char[] s = new char[n];
        int i = 0; // s[i] 没有填字母
        for (char c = 'a'; c <= 'z'; c++) {
            for (int j = i; j < n; j++) {
                if (lcp[i][j] > 0) { // s[j] == s[i]
                    s[j] = c;
                }
            }
            // 找下一个空位
            while (i < n && s[i] > 0) {
                i++;
            }
            if (i == n) { // 没有空位
                break;
            }
        }

        if (i < n) { // 还有空位
            return "";
        }

        // 验证 s 是否符合 lcp 矩阵
        for (i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                // 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
                int actualLcp = s[i] != s[j] ? 0 : (i == n - 1 || j == n - 1 ? 1 : lcp[i + 1][j + 1] + 1);
                if (lcp[i][j] != actualLcp) { // 矛盾
                    return "";
                }
            }
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string findTheString(vector<vector<int>>& lcp) {
        int n = lcp.size();
        string s(n, 0);
        int i = 0; // s[i] 没有填字母
        for (char c = 'a'; c <= 'z'; c++) {
            for (int j = i; j < n; j++) {
                if (lcp[i][j]) { // s[j] == s[i]
                    s[j] = c;
                }
            }
            // 找下一个空位
            while (i < n && s[i]) {
                i++;
            }
            if (i == n) { // 没有空位
                break;
            }
        }

        if (i < n) { // 还有空位
            return "";
        }

        // 验证 s 是否符合 lcp 矩阵
        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                // 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
                int actual_lcp = s[i] != s[j] ? 0 : (i == n - 1 || j == n - 1 ? 1 : lcp[i + 1][j + 1] + 1);
                if (lcp[i][j] != actual_lcp) { // 矛盾
                    return "";
                }
            }
        }
        return s;
    }
};
```

```c [sol-C]
char* findTheString(int** lcp, int lcpSize, int* lcpColSize) {
    int n = lcpSize;
    char* s = calloc(n + 1, sizeof(char));
    int i = 0; // s[i] 没有填字母
    for (char c = 'a'; c <= 'z'; c++) {
        for (int j = i; j < n; j++) {
            if (lcp[i][j]) { // s[j] == s[i]
                s[j] = c;
            }
        }
        // 找下一个空位
        while (i < n && s[i]) {
            i++;
        }
        if (i == n) { // 没有空位
            break;
        }
    }

    if (i < n) { // 还有空位
        s[0] = '\0';
        return s;
    }

    // 验证 s 是否符合 lcp 矩阵
    for (int i = n - 1; i >= 0; i--) {
        for (int j = n - 1; j >= 0; j--) {
            // 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
            int actual_lcp = s[i] != s[j] ? 0 : (i == n - 1 || j == n - 1 ? 1 : lcp[i + 1][j + 1] + 1);
            if (lcp[i][j] != actual_lcp) { // 矛盾
                s[0] = '\0';
                return s;
            }
        }
    }
    return s;
}
```

```go [sol-Go]
func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	i := 0 // s[i] 没有填字母
	for c := byte('a'); c <= 'z'; c++ {
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 { // s[j] == s[i]
				s[j] = c
			}
		}
		// 找下一个空位
		for i < n && s[i] > 0 {
			i++
		}
		if i == n { // 没有空位
			break
		}
	}

	if i < n { // 还有空位
		return ""
	}

	// 验证 s 是否符合 lcp 矩阵
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
			actualLcp := 0
			if s[i] == s[j] {
				if i == n-1 || j == n-1 {
					actualLcp = 1
				} else {
					actualLcp = lcp[i+1][j+1] + 1
				}
			}
			if lcp[i][j] != actualLcp { // 矛盾
				return ""
			}
		}
	}
	return string(s)
}
```

```js [sol-JavaScript]
var findTheString = function(lcp) {
    const n = lcp.length;
    const s = Array(n).fill(null);
    let i = 0; // s[i] 没有填字母
    for (let c = 'a'.charCodeAt(0); c <= 'z'.charCodeAt(0); c++) {
        for (let j = i; j < n; j++) {
            if (lcp[i][j] > 0) { // s[j] == s[i]
                s[j] = c;
            }
        }
        // 找下一个空位
        while (i < n && s[i] !== null) {
            i++;
        }
        if (i === n) { // 没有空位
            break;
        }
    }

    if (i < n) { // 还有空位
        return "";
    }

    // 验证 s 是否符合 lcp 矩阵
    for (i = n - 1; i >= 0; i--) {
        for (let j = n - 1; j >= 0; j--) {
            // 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
            const actualLcp = s[i] !== s[j] ? 0 : (i === n - 1 || j === n - 1 ? 1 : lcp[i + 1][j + 1] + 1);
            if (lcp[i][j] !== actualLcp) { // 矛盾
                return "";
            }
        }
    }
    return String.fromCharCode(...s);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_the_string(lcp: Vec<Vec<i32>>) -> String {
        let n = lcp.len();
        let mut s = vec![0; n];
        let mut i = 0; // s[i] 没有填字母
        for c in b'a'..=b'z' {
            for j in i..n {
                if lcp[i][j] > 0 { // s[j] == s[i]
                    s[j] = c;
                }
            }
            // 找下一个空位
            while i < n && s[i] > 0 {
                i += 1;
            }
            if i == n { // 没有空位
                break;
            }
        }

        if i < n { // 还有空位
            return String::new();
        }

        // 验证 s 是否符合 lcp 矩阵
        for i in (0..n).rev() {
            for j in (0..n).rev() {
                // 计算后缀 [i,n-1] 和后缀 [j,n-1] 的实际 LCP
                let actual_lcp = if s[i] != s[j] {
                    0
                } else if i == n - 1 || j == n - 1 {
                    1
                } else {
                    lcp[i + 1][j + 1] + 1
                };
                if lcp[i][j] != actual_lcp { // 矛盾
                    return String::new();
                }
            }
        }
        unsafe { String::from_utf8_unchecked(s) }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{lcp}$ 矩阵的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 专题训练

1. 字符串题单的「**二、Z 函数**」中的「**LCP 数组**」。
2. 贪心题单的「**§3.1 字典序最小/最大**」。

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
