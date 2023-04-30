下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

由于最大值加一后还是最大值，那么反复利用最大值即可。

设数组的最大值为 $m$，答案就是

$$
m+(m+1)+(m+2)+\cdots + (m+k-1) = \dfrac{(2m+k-1)\cdot k}{2}
$$

```py [sol1-Python3]
class Solution:
    def maximizeSum(self, nums: List[int], k: int) -> int:
        return (max(nums) * 2 + k - 1) * k // 2
```

```go [sol1-Go]
func maximizeSum(nums []int, k int) int {
	max := nums[0]
	for _, x := range nums[1:] {
		if x > max {
			max = x
		}
	}
	return (max*2 + k - 1) * k / 2
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
