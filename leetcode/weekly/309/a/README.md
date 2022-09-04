下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

遍历 $s$，记录 $s[i]$ 上一次出现的位置，如果再次遇到 $s[i]$，通过位置之差可以得到两个字母之间的字母个数，与 $\textit{distance}$ 比较即可。

```py [sol1-Python3]
class Solution:
    def checkDistances(self, s: str, distance: List[int]) -> bool:
        last = [0] * 26
        for i, c in enumerate(s):
            c = ord(c) - ord('a')
            if last[c] and i - last[c] != distance[c]:
                return False
            last[c] = i + 1
        return True
```

```go [sol1-Go]
func checkDistances(s string, distance []int) bool {
	last := [26]int{}
	for i, c := range s {
		c -= 'a'
		if last[c] > 0 && i-last[c] != distance[c] {
			return false
		}
		last[c] = i + 1
	}
	return true
}
```
