下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

```py [sol1-Python3]
class Solution:
    def mostFrequentEven(self, nums: List[int]) -> int:
        cnt = Counter(x for x in nums if x % 2 == 0)
        if len(cnt) == 0: return -1
        max_cnt = max(cnt.values())
        return min(x for x, c in cnt.items() if c == max_cnt)
```

```go [sol1-Go]
func mostFrequentEven(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		if x%2 == 0 {
			cnt[x]++
		}
	}
	if len(cnt) == 0 {
		return -1
	}
	ans := -1
	for x, c := range cnt {
		if ans < 0 || c > cnt[ans] || c == cnt[ans] && x < ans {
			ans = x
		}
	}
	return ans
}
```
