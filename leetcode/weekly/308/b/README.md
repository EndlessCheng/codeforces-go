下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

用栈模拟即可，做法类似 [1047. 删除字符串中的所有相邻重复项](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/)。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def removeStars(self, s: str) -> str:
        st = []
        for c in s:
            if c == '*': st.pop()
            else: st.append(c)
        return ''.join(st)
```

```go [sol1-Go]
func removeStars(s string) string {
	st := []rune{}
	for _, c := range s {
		if c == '*' {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
```
