下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

由于选的是子序列，且子序列的元素都相等，所以**元素顺序对答案没有影响**，可以先对数组**排序**。

由于替换操作替换的是一个连续范围内的数，所以排序后，选出的子序列必然也是一段**连续子数组**。

那么问题变成：「找最长的连续子数组，其最大值减最小值不超过 $2k$」，只要子数组满足这个要求，其中的元素都可以变成同一个数。

这个问题可以用 [同向双指针](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 解决。枚举 $\textit{nums}[\textit{right}]$ 作为子数组的最后一个数，一旦 $\textit{nums}[\textit{right}]-\textit{nums}[\textit{left}]>2k$，就移动左端点。

$\textit{right}-\textit{left}+1$ 是子数组的长度，取所有长度最大值，即为答案。

```py [sol-Python3]
class Solution:
    def maximumBeauty(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = left = 0
        for right, x in enumerate(nums):
            while x - nums[left] > k * 2:
                left += 1
            ans = max(ans, right - left + 1)
        return ans
```

```go [sol-Go]
func maximumBeauty(nums []int, k int) (ans int) {
	sort.Ints(nums)
	left := 0
	for right, x := range nums {
		for x-nums[left] > k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。
