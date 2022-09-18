下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

```py [sol1-Python3]
class Solution:
    def longestContinuousSubstring(self, s: str) -> int:
        ans = start = 0
        for i in range(1, len(s)):
            if ord(s[i]) != ord(s[i - 1]) + 1:
                ans = max(ans, i - start)
                start = i  # 新起点
        return max(ans, len(s) - start)
```

```go [sol1-Go]
func longestContinuousSubstring(s string) (ans int) {
	start := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1]+1 {
			ans = max(ans, i-start)
			start = i // 新起点
		}
	}
	return max(ans, len(s)-start)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
