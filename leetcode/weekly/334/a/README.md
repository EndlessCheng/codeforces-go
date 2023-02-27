模拟。按要求计算前缀和 $\textit{leftSum}$ 与后缀和 $\textit{rightSum}$，然后计算答案。

代码实现时，可以把所有元素和计算出来，当作 $\textit{rightSum}$，然后一边遍历 $\textit{nums}$，一边更新 $\textit{leftSum}$ 和 $\textit{rightSum}$，同时把结果直接记到 $\textit{nums}[i]$ 中。

```py [sol1-Python3]
class Solution:
    def leftRigthDifference(self, nums: List[int]) -> List[int]:
        left_sum, right_sum = 0, sum(nums)
        for i, x in enumerate(nums):
            right_sum -= x
            nums[i] = abs(left_sum - right_sum)
            left_sum += x
        return nums
```

```go [sol1-Go]
func leftRigthDifference(nums []int) []int {
	rightSum := 0
	for _, x := range nums {
		rightSum += x
	}
	leftSum := 0
	for i, x := range nums {
		rightSum -= x
		nums[i] = abs(leftSum - rightSum)
		leftSum += x
	}
	return nums
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

---

欢迎关注【biIibiIi@灵茶山艾府】，高质量算法教学，持续更新中~
