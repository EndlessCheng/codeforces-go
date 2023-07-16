下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲这场周赛的题目，欢迎关注！

---

按照题目要求计算即可。

```py [sol-Python3]
class Solution:
    def sumOfSquares(self, nums: List[int]) -> int:
        return sum(x * x for i, x in enumerate(nums, 1)
                         if len(nums) % i == 0)
```

```go [sol-Go]
func sumOfSquares(nums []int) (ans int) {
	for i, x := range nums {
		if len(nums)%(i+1) == 0 {
			ans += x * x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
