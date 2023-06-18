下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，不仅讲做法，还会教你如何一步步思考，记得关注哦~

---

答案的理论最小值，是任意两个元素的差的绝对值的最小值。能否取到这个最小值呢？

是可以的，把数组排序后，最小值必然对应两个相邻元素，设其为 $\textit{nums}[i-1]$ 和 $\textit{nums}[i]$。

把不超过 $\textit{nums}[i-1]$ 的数分到第一个数组中，把不低于 $\textit{nums}[i]$ 的数分到第二个数组中，即可满足题目要求。

```py [sol-Python3]
class Solution:
    def findValueOfPartition(self, nums: List[int]) -> int:
        nums.sort()
        return min(y - x for x, y in pairwise(nums))
```

```go [sol-Go]
func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
