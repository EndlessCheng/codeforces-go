下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

```py [sol1-Python3]
class Solution:
    def equalFrequency(self, word: str) -> bool:
        for i in range(len(word)):
            cnt = Counter(word[:i] + word[i + 1:])
            same = cnt.popitem()[1]
            if all(c == same for c in cnt.values()):
                return True
        return False
```

```go [sol1-Go]
func equalFrequency(word string) bool {
next:
	for i := range word {
		cnt := map[rune]int{}
		for _, c := range word[:i] + word[i+1:] {
			cnt[c]++
		}
		same := 0
		for _, c := range cnt {
			if same == 0 {
				same = c
			} else if c != same {
				continue next
			}
		}
		return true
	}
	return false
}
```

$O(n)$ 做法见 [1224. 最大相等频率](https://leetcode.cn/problems/maximum-equal-frequency/)。
