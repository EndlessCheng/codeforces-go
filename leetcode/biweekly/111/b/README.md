请看 [视频讲解](https://www.bilibili.com/video/BV1Yu4y1v7H6/) 第二题。

下文将 $\textit{str}_1$ 简记为 $s$，将 $\textit{str}_2$ 简记为 $t$。

想一想，如果 $s[0]$ 可以匹配 $t[0]$，那么此时一定要匹配，不然后面可能没有机会匹配 $t[0]$。

因此，双指针遍历 $s[i]$ 和 $t[j]$，如果 $s[i]$ 可以匹配 $t[j]$，那么 $i$ 和 $j$ 都加一，否则只有 $i$ 加一。

如果 $j$ 等于 $t$ 的长度，则返回 `true`，否则返回 `false`。

```py [sol-Python3]
class Solution:
    def canMakeSubsequence(self, s: str, t: str) -> bool:
        if len(s) < len(t):
            return False
        j = 0
        for b in s:
            c = chr(ord(b) + 1) if b != 'z' else 'a'
            if b == t[j] or c == t[j]:
                j += 1
                if j == len(t):
                    return True
        return False
```

```go [sol-Go]
func canMakeSubsequence(s, t string) bool {
	if len(s) < len(t) {
		return false
	}
	j := 0
	for _, b := range s {
		c := byte(b) + 1
		if b == 'z' {
			c = 'a'
		}
		if byte(b) == t[j] || c == t[j] {
			j++
			if j == len(t) {
				return true
			}
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
