下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

```go
func partitionString(s string) int {
	ans, vis := 1, 0
	for _, c := range s {
		if vis>>(c&31)&1 > 0 {
			vis = 0
			ans++
		}
		vis |= 1 << (c & 31)
	}
	return ans
}
```
