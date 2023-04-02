### 本题视频讲解

见[【周赛 339】](https://www.bilibili.com/video/BV1va4y1M7Fr/)。

### 思路

记录连续 $0$ 的个数 $\textit{pre}$ 和连续 $1$ 的个数 $\textit{cur}$，那么这个 $01$ 串的长度就是 $2\cdot \min(\textit{pre}, \textit{cur})$。

取所有长度的最大值，即为答案。

```py [sol1-Python3]
class Solution:
    def findTheLongestBalancedSubstring(self, s: str) -> int:
        ans = pre = cur = 0
        for i, c in enumerate(s):
            cur += 1
            if i == len(s) - 1 or c != s[i + 1]:
                if c == '1':
                    ans = max(ans, min(pre, cur) * 2)
                pre = cur
                cur = 0
        return ans
```

```java [sol1-Java]
class Solution {
    public int findTheLongestBalancedSubstring(String S) {
        var s = S.toCharArray();
        int ans = 0, pre = 0, cur = 0, n = s.length;
        for (int i = 0; i < n; ++i) {
            ++cur;
            if (i == s.length - 1 || s[i] != s[i + 1]) {
                if (s[i] == '1')
                    ans = Math.max(ans, Math.min(pre, cur) * 2);
                pre = cur;
                cur = 0;
            }
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int findTheLongestBalancedSubstring(string s) {
        int ans = 0, pre = 0, cur = 0, n = s.length();
        for (int i = 0; i < n; ++i) {
            ++cur;
            if (i == s.length() - 1 || s[i] != s[i + 1]) {
                if (s[i] == '1')
                    ans = max(ans, min(pre, cur) * 2);
                pre = cur;
                cur = 0;
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func findTheLongestBalancedSubstring(s string) (ans int) {
	pre, cur := 0, 0
	for i, c := range s {
		cur++
		if i == len(s)-1 || byte(c) != s[i+1] {
			if c == '1' {
				ans = max(ans, min(pre, cur)*2)
			}
			pre = cur
			cur = 0
		}
	}
	return
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
