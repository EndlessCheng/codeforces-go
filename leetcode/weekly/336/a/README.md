下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

遍历 $[\textit{left},\textit{right}]$ 内的字符串，按照要求判断。

```py [sol1-Python3]
class Solution:
    def vowelStrings(self, words: List[str], left: int, right: int) -> int:
        return sum(s[0] in "aeiou" and s[-1] in "aeiou" for s in words[left:right+1])
```

```go [sol1-Go]
func vowelStrings(words []string, left, right int) (ans int) {
	for _, s := range words[left : right+1] {
		if strings.Contains("aeiou", s[:1]) && strings.Contains("aeiou", s[len(s)-1:]) {
			ans++
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{words}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
