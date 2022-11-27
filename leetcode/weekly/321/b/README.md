贪心，双指针遍历 $s$ 和 $t$，$t[j]$ 应匹配 $i$ 尽量小（但大于上一个的匹配的位置）的 $s[i]$。

```py [sol1-Python3]
class Solution:
    def appendCharacters(self, s: str, t: str) -> int:
        i, n = 0, len(s)
        for j, c in enumerate(t):
            while i < n and s[i] != t[j]: i += 1
            if i == n: return len(t) - j
            i += 1
        return 0
```

```go [sol1-Go]
func appendCharacters(s, t string) int {
	i, n := 0, len(s)
	for j := range t {
		for i < n && s[i] != t[j] {
			i++
		}
		if i == n {
			return len(t) - j
		}
		i++
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $s$ 的长度，$m$ 为 $t$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
