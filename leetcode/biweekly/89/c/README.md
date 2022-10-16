# 方法一：二分答案

「最小化最大值」就是二分答案的代名词。

二分答案 $\textit{limit}$，那么我们可以从后往前模拟：如果 $\textit{nums}[i]>\textit{limit}$，那么应当去掉多余的 $\textit{nums}[i]-\textit{limit}$ 加到 $\textit{nums}[i-1]$ 上，最后如果 $\textit{nums}[0]\le\textit{limit}$，则二分判定成功。

```py [sol1-Python3]
class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        def check(limit: int) -> bool:
            extra = 0
            for i in range(len(nums) - 1, 0, -1):
                extra = max(nums[i] + extra - limit, 0)
            return nums[0] + extra <= limit
        return bisect_left(range(max(nums)), True, key=check)
```

```go [sol1-Go]
func minimizeArrayValue(nums []int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	return sort.Search(mx, func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干变量。

# 方法二：分类讨论

从 $\textit{nums}[0]$ 开始讨论：

- 如果数组中只有 $\textit{nums}[0]$，那么最大值为 $\textit{nums}[0]$。
- 再考虑 $\textit{nums}[1]$，如果 $\textit{nums}[1]\le\textit{nums}[0]$，最大值还是 $\textit{nums}[0]$；否则可以平均这两个数，平均后的最大值为平均值的上取整，即 $\left\lceil\dfrac{\textit{nums}[0]+\textit{nums}[1]}{2}\right\rceil$。
- 再考虑 $\textit{nums}[2]$，如果 $\textit{nums}[2]\le$ 前面算出的最大值，那么最大值不变；否则可以平均这三个数，做法同上。
- 以此类推直到最后一个数。
- 过程中的最大值为答案。

```py [sol2-Python3]
class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        return max((s + i - 1) // i for i, s in enumerate(accumulate(nums), 1))
```

```go [sol2-Go]
func minimizeArrayValue(nums []int) (ans int) {
	s := 0
	for i, x := range nums {
		s += x
		ans = max(ans, (s+i)/(i+1))
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
