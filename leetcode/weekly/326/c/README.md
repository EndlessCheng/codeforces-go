[视频讲解](https://www.bilibili.com/video/BV1H8411E7hn) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

```py [sol1-Python3]
class Solution:
    def minimumPartition(self, s: str, k: int) -> int:
        ans, x = 1, 0
        for v in map(int, s):
            if v > k: return -1
            x = x * 10 + v
            if x > k:
                ans += 1
                x = v
        return ans
```

```go [sol1-Go]
func minimumPartition(s string, k int) int {
	ans, x := 1, 0
	for _, c := range s {
		v := int(c - '0')
		if v > k {
			return -1
		}
		x = x*10 + v
		if x > k {
			ans++
			x = v
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
