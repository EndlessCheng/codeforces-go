### 提示 1-1

设要构造的字符串为 $s$。

既然要使 $s$ 字典序最小，可以先从 $s[0]$ 开始思考。

$s[0]$ 一定可以填入 $\text{`a'}$ 吗？

### 提示 1-2

证明如下：

1. 如果最终没有 $s[i]$ 等于 $\text{`a'}$，那么可以把所有 $s[i]$ 都变小。
2. 此时一定有某些 $s[i]$ 为 $\text{`a'}$，如果 $s[0]\ne\text{`a'}$，假设 $s[0]=\text{`b'}$，那么可以把所有 $\text{`b'}$ 替换成 $\text{`a'}$，所有 $\text{`a'}$ 替换成 $\text{`b'}$。
3. 由于 1 和 2 都不会影响 LCP，所以 $s[0]$ 一定可以填入 $\text{`a'}$。

### 提示 2-1

根据 $\textit{lcp}[0]$，还有哪些 $s[i]$ 一定是 $\text{`a'}$？哪些 $s[i]$ 一定不是 $\text{`a'}$？

### 提示 2-2

根据 LCP 的定义，$\textit{lcp}[0][i]>0$ 的一定是 $\text{`a'}$，$\textit{lcp}[0][i]=0$ 的一定不是 $\text{`a'}$。

把一定是 $\text{`a'}$ 的 $s[i]$ 填入 $\text{`a'}$ 后，接着想想，哪些 $s[i]$ 要填入 $\text{`b'}$？

### 提示 2-3

找到下一个没有填入字母的 $s[i]$，它可以填入 $\text{`b'}$，并且所有 $\textit{lcp}[i][j]>0$ 的 $s[j]$ 可以填入 $\text{`b'}$。

### 提示 3

构造方案如下：

1. 初始化长为 $n$ 的字符串 $s$，初始化所有 $s[i]$ 为 `\0`，初始化待填入字母 $c$ 为 $\text{`a'}$。
2. 找到第一个 `\0` 的下标 $i$，把所有 $j\ge i$ 且 $\textit{lcp}[i][j]>0$ 的 $s[j]$ 填入 $c$。
3. $c$ 加一，重复第 2 步，直到没有这样的 $i$，或者用完了 $26$ 个小写字母（此时不合法）。

### 提示 4

构造完之后，还需要计算 $s$ 的真实 LCP，检查是否和输入的 $\textit{lcp}$ 数组完全一致。

计算方式是一个 DP：

- 如果 $s[i]=s[j]$，那么 $\textit{lcp}[i][j] = \textit{lcp}[i+1][j+1]+1$；
- 如果 $s[i]\ne s[j]$，那么 $\textit{lcp}[i][j] = 0$。

附：[视频讲解](https://www.bilibili.com/video/BV1jM411J7y7/)

```py [sol1-Python3]
class Solution:
    def findTheString(self, lcp: List[List[int]]) -> str:
        i, n = 0, len(lcp)
        s = [''] * n
        for c in ascii_lowercase:
            while i < n and s[i]: i += 1
            if i == n: break  # 构造完毕
            for j in range(i, n):
                if lcp[i][j]:
                    s[j] = c
        if '' in s: return ""  # 没有构造完

        # 直接在原数组上验证
        for i in range(n - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                actual_lcp = 0 if s[i] != s[j] else 1 if i == n - 1 or j == n - 1 else lcp[i + 1][j + 1] + 1
                if lcp[i][j] != actual_lcp: return ""
        return "".join(s)
```

```java [sol1-Java]
class Solution {
    public String findTheString(int[][] lcp) {
        int i = 0, n = lcp.length;
        var s = new char[n];
        for (char c = 'a'; c <= 'z'; ++c) {
            while (i < n && s[i] > 0) ++i;
            if (i == n) break; // 构造完毕
            for (int j = i; j < n; ++j)
                if (lcp[i][j] > 0) s[j] = c;
        }
        while (i < n) if (s[i++] == 0) return ""; // 没有构造完

        // 直接在原数组上验证
        for (i = n - 1; i >= 0; --i)
            for (int j = n - 1; j >= 0; --j) {
                int actualLCP = s[i] != s[j] ? 0 : i == n - 1 || j == n - 1 ? 1 : lcp[i + 1][j + 1] + 1;
                if (lcp[i][j] != actualLCP) return "";
            }
        return new String(s);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    string findTheString(vector<vector<int>> &lcp) {
        int i = 0, n = lcp.size();
        string s(n, 0);
        for (char c = 'a'; c <= 'z'; ++c) {
            while (i < n && s[i]) ++i;
            if (i == n) break; // 构造完毕
            for (int j = i; j < n; ++j)
                if (lcp[i][j]) s[j] = c;
        }
        while (i < n) if (s[i++] == 0) return ""; // 没有构造完

        // 直接在原数组上验证
        for (int i = n - 1; i >= 0; --i)
            for (int j = n - 1; j >= 0; --j) {
                int actual_lcp = s[i] != s[j] ? 0 : i == n - 1 || j == n - 1 ? 1 : lcp[i + 1][j + 1] + 1;
                if (lcp[i][j] != actual_lcp) return "";
            }
        return s;
    }
};
```

```go [sol1-Go]
func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	for c := byte('a'); c <= 'z'; c++ {
		i := bytes.IndexByte(s, 0)
		if i < 0 { // 构造完毕
			break
		}
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				s[j] = c
			}
		}
	}
	if bytes.IndexByte(s, 0) >= 0 { // 没有构造完
		return ""
	}

	// 直接在原数组上验证
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			actualLCP := 0
			if s[i] == s[j] {
				actualLCP = 1
				if i < n-1 && j < n-1 {
					actualLCP += lcp[i+1][j+1]
				}
			}
			if lcp[i][j] != actualLCP {
				return ""
			}
		}
	}
	return string(s)
}
```

### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{lcp}$ 的长度。
- 空间复杂度：$O(1)$。不计入返回值，仅用到若干额外变量。
