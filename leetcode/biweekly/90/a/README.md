下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

计算出每个字符串的 $\textit{difference}$ 数组，作为哈希表的 key。把相同 key 的字符串存到同一组中，答案为组内只有一个字符串的那个字符串。

```py [sol1-Python3]
class Solution:
    def oddString(self, words: List[str]) -> str:
        d = defaultdict(list)
        for s in words:
            d[tuple(ord(x) - ord(y) for x, y in pairwise(s))].append(s)
        x, y = d.values()
        return x[0] if len(x) == 1 else y[0]
```

```go [sol1-Go]
func oddString(words []string) string {
	m := map[string][]string{}
	d := make([]byte, len(words[0])-1)
	for _, s := range words {
		for i := 0; i < len(s)-1; i++ {
			d[i] = s[i] - s[i+1]
		}
		t := string(d)
		m[t] = append(m[t], s)
	}
	for _, g := range m {
		if len(g) == 1 {
			return g[0]
		}
	}
	return ""
}
```

#### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 为 $\textit{words}$ 的长度，$n$ 为 $\textit{words}[i]$ 的长度。
- 空间复杂度：$O(m+n)$。哈希表的空间消耗为 $O(m+n)$。
