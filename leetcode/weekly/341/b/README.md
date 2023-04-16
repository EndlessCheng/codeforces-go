下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

```py [sol1-Python3]
class Solution:
    def maxDivScore(self, nums: List[int], divisors: List[int]) -> int:
        max_cnt, ans = -1, 0
        for d in divisors:
            cnt = sum(x % d == 0 for x in nums)
            if cnt > max_cnt or cnt == max_cnt and d < ans:
                max_cnt, ans = cnt, d
        return ans
```

```go [sol1-Go]
func maxDivScore(nums, divisors []int) (ans int) {
	maxCnt := -1
	for _, d := range divisors {
		cnt := 0
		for _, x := range nums {
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && d < ans {
			maxCnt, ans = cnt, d
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(nm)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{divisors}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
