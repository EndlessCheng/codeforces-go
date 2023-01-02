[视频讲解](https://www.bilibili.com/video/BV1Dd4y1h72z/) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

对每个 $1$，向左向右找 $-1$，且中间的必须都是 $0$。

```py [sol1-Python3]
class Solution:
    def captureForts(self, forts: List[int]) -> int:
        ans = 0
        for i, x in enumerate(forts):
            if x != 1: continue
            j = i - 1
            while j >= 0 and forts[j] == 0: j -= 1
            if j >= 0 and forts[j] < 0: ans = max(ans, i - j - 1)
            j = i + 1
            while j < len(forts) and forts[j] == 0: j += 1
            if j < len(forts) and forts[j] < 0: ans = max(ans, j - i - 1)
        return ans
```

```go [sol1-Go]
func captureForts(forts []int) (ans int) {
	for i, x := range forts {
		if x != 1 {
			continue
		}
		j := i - 1
		for j >= 0 && forts[j] == 0 {
			j--
		}
		if j >= 0 && forts[j] < 0 {
			ans = max(ans, i-j-1)
		}
		j = i + 1
		for j < len(forts) && forts[j] == 0 {
			j++
		}
		if j < len(forts) && forts[j] < 0 {
			ans = max(ans, j-i-1)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{forts}$ 的长度。注意每个数至多被遍历两次。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
