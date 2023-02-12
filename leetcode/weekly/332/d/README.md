### 提示 1

在 $[\textit{left}, \textit{right}]$ 之间的字符，删除是不影响得分的，且删除后更有机会让剩余部分是 $s$ 的子序列。

因此只需考虑删除的是 $t$ 的子串，而不是子序列。

### 提示 2

删除子串后，剩余部分是 $t$ 的一个前缀和一个后缀。

假设前缀匹配的是 $s$ 的一个前缀 $s[:i]$，后缀匹配的是 $s$ 的一个后缀 $s[i:]$。这里匹配指子序列匹配。

那么枚举 $i$，分别计算能够与 $s[:i]$ 和 $s[i:]$ 匹配的 $t$ 的最长前缀和最长后缀，就知道要删除的子串的最小值了。这个技巧叫做「前后缀分解」。

### 提示 3

具体来说：

定义 $\textit{pre}[i]$ 为 $s[:i]$ 对应的 $t$ 的最长前缀的结束下标。

定义 $\textit{suf}[i]$ 为 $s[i:]$ 对应的 $t$ 的最长后缀的开始下标。

那么删除的子串就是从 $\textit{pre}[i]+1$ 到 $\textit{suf}[i]-1$ 这段，答案就是 $\textit{suf}[i]-\textit{pre}[i]-1$ 的最小值。

代码实现时，可以先计算 $\textit{suf}$，然后一边计算 $\textit{pre}$，一边更新最小值，所以 $\textit{pre}$ 可以省略。

附：[视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)

```py [sol1-Python3]
class Solution:
    def minimumScore(self, s: str, t: str) -> int:
        n, m = len(s), len(t)
        suf = [m] * (n + 1)
        j = m - 1
        for i in range(n - 1, -1, -1):
            if j >= 0 and s[i] == t[j]:
                j -= 1
            suf[i] = j + 1
        ans = suf[0]  # 删除 t[:suf[0]]
        if ans == 0: return 0

        j = 0
        for i, c in enumerate(s):
            if c == t[j]:  # 注意 j 不会等于 m，因为上面 suf[0]>0 表示 t 不是 s 的子序列
                j += 1
                ans = min(ans, suf[i + 1] - j)  # 删除 t[j:suf[i+1]]
        return ans
```

```java [sol1-Java]
class Solution {
    public int minimumScore(String S, String T) {
        char[] s = S.toCharArray(), t = T.toCharArray();
        int n = s.length, m = t.length;
        var suf = new int[n + 1];
        suf[n] = m;
        for (int i = n - 1, j = m - 1; i >= 0; --i) {
            if (j >= 0 && s[i] == t[j]) --j;
            suf[i] = j + 1;
        }
        int ans = suf[0]; // 删除 t[:suf[0]]
        if (ans == 0) return 0;
        for (int i = 0, j = 0; i < n; ++i)
            if (s[i] == t[j]) // 注意 j 不会等于 m，因为上面 suf[0]>0 表示 t 不是 s 的子序列
                ans = Math.min(ans, suf[i + 1] - ++j); // ++j 后，删除 t[j:suf[i+1]]
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minimumScore(string s, string t) {
        int n = s.length(), m = t.length(), suf[n + 1];
        suf[n] = m;
        for (int i = n - 1, j = m - 1; i >= 0; --i) {
            if (j >= 0 && s[i] == t[j]) --j;
            suf[i] = j + 1;
        }
        int ans = suf[0]; // 删除 t[:suf[0]]
        if (ans == 0) return 0;
        for (int i = 0, j = 0; i < n; ++i)
            if (s[i] == t[j]) // 注意 j 不会等于 m，因为上面 suf[0]>0 表示 t 不是 s 的子序列
                ans = min(ans, suf[i + 1] - ++j); // ++j 后，删除 t[j:suf[i+1]]
        return ans;
    }
};
```

```go [sol1-Go]
func minimumScore(s, t string) int {
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = j + 1
	}
	ans := suf[0] // 删除 t[:suf[0]]
	if ans == 0 {
		return 0
	}
	for i, j := 0, 0; i < n; i++ {
		if s[i] == t[j] { // 注意 j 不会等于 m，因为上面 suf[0]>0 表示 t 不是 s 的子序列
			j++
			ans = min(ans, suf[i+1]-j) // 删除 t[j:suf[i+1]]
		}
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。注意时间复杂度与 $t$ 的长度无关。
- 空间复杂度：$O(n)$。
