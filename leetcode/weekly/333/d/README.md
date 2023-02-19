构造方案如下：

1. 初始化长为 $n$ 的字符串 $s$，初始化所有 $s[i]$ 为 `\0`，初始化待填入字母 $c$ 为 $\text{`a'}$。
2. 找到第一个 `\0` 的下标 $i$，把所有 $j\ge i$ 且 $\textit{lcp}[i][j]>0$ 的 $s[j]$ 填入 $c$。
3. $c$ 加一，重复第 2 步，直到没有这样的 $i$，或者用完了 $26$ 个小写字母（此时不合法）。

构造完之后，计算 $s$ 的真实 LCP，检查是否和输入的 $\textit{lcp}$ 数组完全一致。

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
                if s[i] == s[j]:  # 相等
                    if i == n - 1 or j == n - 1:
                        if lcp[i][j] != 1: return ""
                    elif lcp[i][j] != lcp[i + 1][j + 1] + 1: return ""
                elif lcp[i][j]: return ""
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
            for (int j = n - 1; j >= 0; --j)
                if (s[i] == s[j]) {
                    if (i == n - 1 || j == n - 1) {
                        if (lcp[i][j] != 1)
                            return "";
                    } else if (lcp[i][j] != lcp[i + 1][j + 1] + 1)
                        return "";
                } else if (lcp[i][j] > 0)
                    return "";
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
            for (int j = n - 1; j >= 0; --j)
                if (s[i] == s[j]) {
                    if (i == n - 1 || j == n - 1) {
                        if (lcp[i][j] != 1)
                            return "";
                    } else if (lcp[i][j] != lcp[i + 1][j + 1] + 1)
                        return "";
                } else if (lcp[i][j])
                    return "";
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
			if s[i] == s[j] {
				if i == n-1 || j == n-1 {
					if lcp[i][j] != 1 {
						return ""
					}
				} else if lcp[i][j] != lcp[i+1][j+1]+1 {
					return ""
				}
			} else if lcp[i][j] > 0 {
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
