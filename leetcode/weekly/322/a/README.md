1. 首尾字符相同；
2. 每个空格左右的字符相同。

```py [sol1-Python3]
class Solution:
    def isCircularSentence(self, s: str) -> bool:
        return s[0] == s[-1] and all(c != ' ' or s[i - 1] == s[i + 1] for i, c in enumerate(s))
```

```go [sol1-Go]
func isCircularSentence(s string) bool {
	if s[0] != s[len(s)-1] {
		return false
	}
	for i, c := range s {
		if c == ' ' && s[i-1] != s[i+1] {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{sentence}$ 的长度。
- 空间复杂度：$O(1)$，只用到若干额外变量。
