下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

将题目中的式子变形得

$$
\textit{nums}[i]-i \ne \textit{nums}[j]-i
$$

为了方便计算，我们可以算出满足

$$
\textit{nums}[i]-i = \textit{nums}[j]-i
$$

的下标对数，这可以一边遍历，一边统计。然后用所有对数 $\dfrac{n(n-1)}{2}$ 减去，即为答案。

```py [sol1-Python3]
class Solution:
    def countBadPairs(self, nums: List[int]) -> int:
        n, cnt = len(nums), Counter()
        ans = n * (n - 1) // 2
        for i, num in enumerate(nums):
            ans -= cnt[num - i]
            cnt[num - i] += 1
        return ans
```

```go [sol1-Go]
func countBadPairs(nums []int) int64 {
	n := len(nums)
	ans := n * (n - 1) / 2
	cnt := map[int]int{}
	for i, num := range nums {
		ans -= cnt[num-i]
		cnt[num-i]++
	}
	return int64(ans)
}
```
