下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

按题意模拟即可。

```py [sol1-Python3]
class Solution:
    def edgeScore(self, edges: List[int]) -> int:
        ans, score = 0, [0] * len(edges)
        for i, to in enumerate(edges):
            score[to] += i
            if score[to] > score[ans] or score[to] == score[ans] and to < ans:
                ans = to
        return ans
```

```go [sol1-Go]
func edgeScore(edges []int) (ans int) {
	score := make([]int, len(edges))
	for i, to := range edges {
		score[to] += i
		if score[to] > score[ans] || score[to] == score[ans] && to < ans {
			ans = to
		}
	}
	return
}
```
