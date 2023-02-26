下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

模拟。按要求计算前缀和 $\textit{leftSum}$ 与后缀和 $\textit{rightSum}$，然后计算答案。

```py [sol1-Python3]
class Solution:
    def leftRigthDifference(self, nums: List[int]) -> List[int]:
        n = len(nums)
        left_sum, right_sum = [0] * n, [0] * n
        for i in range(n - 1):
            left_sum[i + 1] = left_sum[i] + nums[i]
            right_sum[-2 - i] = right_sum[-1 - i] + nums[-1 - i]
        return [abs(x - y) for x, y in zip(left_sum, right_sum)]
```

```go [sol1-Go]
func leftRigthDifference(nums []int) []int {
	n := len(nums)
	rightSum := make([]int, n)
	for i := n - 1; i > 0; i-- {
		rightSum[i-1] = rightSum[i] + nums[i]
	}
	
	ans := make([]int, n)
	leftSum := 0
	for i, x := range nums {
		ans[i] = abs(leftSum - rightSum[i])
		leftSum += x
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
