下午 2 点在 B 站直播讲周赛（和双周赛）的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

答案可以分为两部分：

- 所有垃圾的数目，即 $\textit{garbage}$ 中所有字符串的长度之和。
- 每一种字符在 $\textit{garbage}$ 中最后一次出现的下标，即每辆垃圾车需要向右开到的房子的最小值。

遍历 $\textit{garbage}$ 可以求出。

```py [sol1-Python3]
class Solution:
    def garbageCollection(self, garbage: List[str], travel: List[int]) -> int:
        ans = 0
        right = [0] * 3
        for i, s in enumerate(garbage):
            ans += len(s)
            for j, c in enumerate("MPG"):
                if c in s:
                    right[j] = i
        return ans + sum(sum(travel[:r]) for r in right)
```

```go [sol1-Go]
func garbageCollection(garbage []string, travel []int) (ans int) {
	right := [3]int{}
	for i, s := range garbage {
		ans += len(s)
		for j, c := range "MPG" {
			if strings.ContainsRune(s, c) {
				right[j] = i
			}
		}
	}
	for _, r := range right {
		for _, t := range travel[:r] {
			ans += t
		}
	}
	return
}
```
