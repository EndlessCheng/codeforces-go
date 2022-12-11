[视频讲解](https://www.bilibili.com/video/BV1kR4y1r7Df/) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

按题意模拟即可。

```py [sol1-Python3]
class Solution:
    def maximumValue(self, strs: List[str]) -> int:
        ans = 0
        for s in strs:
            try: x = int(s)
            except: x = len(s)
            ans = max(ans, x)
        return ans
```

```go [sol1-Go]
func maximumValue(strs []string) (ans int) {
	for _, s := range strs {
		x, err := strconv.Atoi(s)
		if err != nil {
			x = len(s)
		}
		ans = max(ans, x)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为所有字符串的长度之和。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
