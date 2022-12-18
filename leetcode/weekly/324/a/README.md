[视频讲解](https://www.bilibili.com/video/BV1LW4y1T7if/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

由于只有小写字母，可以用一个整数 $\textit{mask}$ 表示字符串中出现过的字母。

遍历 $\textit{words}$ 的同时，用一个哈希表 $\textit{cnt}$ 维护 $\textit{mask}$ 的出现次数。先把 $\textit{cnt}[\textit{mask}]$ 加到答案中，然后把 $\textit{mask}$ 的出现次数加一。

```py [sol1-Python3]
class Solution:
    def similarPairs(self, words: List[str]) -> int:
        ans, cnt = 0, Counter()
        for s in words:
            mask = 0
            for c in s:
                mask |= 1 << (ord(c) - ord('a'))
            ans += cnt[mask]
            cnt[mask] += 1
        return ans
```

```go [sol1-Go]
func similarPairs(words []string) (ans int) {
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, c := range s {
			mask |= 1 << (c - 'a')
		}
		ans += cnt[mask]
		cnt[mask]++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(L)$，其中 $L$ 为 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$O(n)$，其中 $n$ 为 $\textit{words}$ 的长度。
