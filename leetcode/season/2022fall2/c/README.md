[视频讲解](https://www.bilibili.com/video/BV1rT411P7NA) 已出炉，**包括本题滑窗的原理和时间复杂度分析**，欢迎点赞三连，在评论区分享你对这场力扣杯的看法~

```py [sol1-Python3]
class Solution:
    def beautifulBouquet(self, flowers: List[int], cnt: int) -> int:
        ans = left = 0
        c = defaultdict(int)
        for right, x in enumerate(flowers):
            c[x] += 1
            while c[x] > cnt:
                c[flowers[left]] -= 1
                left += 1
            ans += right - left + 1
        return ans % (10 ** 9 + 7)
```

```go [sol1-Go]
func beautifulBouquet(flowers []int, cnt int) (ans int) {
	c := map[int]int{}
	left := 0
	for right, x := range flowers {
		c[x]++
		for c[x] > cnt {
			c[flowers[left]]--
			left++
		}
		ans += right - left + 1
	}
	return ans % (1e9 + 7)
}
```