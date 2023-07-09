下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

注意 $[3,4,3,4,5,4,5]$ 这样的数组，第一组交替子数组为 $[3,4,3,4]$，第二组交替子数组为 $[4,5,4,5]$，这两组有重叠部分，所以下面代码循环末尾要把 $i$ 减一。

```py [sol-Python3]
class Solution:
    def alternatingSubarray(self, nums: List[int]) -> int:
        ans = -1
        i, n = 0, len(nums)
        while i < n - 1:
            if nums[i + 1] - nums[i] != 1:
                i += 1
                continue
            i0 = i
            i += 1
            while i < n and nums[i] == nums[i0] + (i - i0) % 2:
                i += 1
            ans = max(ans, i - i0)
            i -= 1
        return ans
```

```go [sol-Go]
func alternatingSubarray(nums []int) int {
	ans := -1
	for i, n := 0, len(nums); i < n-1; {
		if nums[i+1]-nums[i] != 1 {
			i++
			continue
		}
		st := i
		for i++; i < n && nums[i] == nums[st]+(i-st)%2; i++ {}
		ans = max(ans, i-st)
		i--
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。虽然写了个二重循环，但是内层循环中对 $i$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
