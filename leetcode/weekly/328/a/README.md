下午两点【B站@灵茶山艾府】直播讲题，关注UP不迷路~

---

由于元素值一定不小于其数位和，所以答案就是元素和减去数位和。

代码实现时可以用同一个变量。

```py [sol1-Python3]
class Solution:
    def differenceOfSum(self, nums: List[int]) -> int:
        ans = 0
        for x in nums:
            ans += x
            while x:
                ans -= x % 10
                x //= 10
        return ans
```

```go [sol1-Go]
func differenceOfSum(nums []int) (ans int) {
	for _, x := range nums {
		for ans += x; x > 0; x /= 10 {
			ans -= x % 10
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
