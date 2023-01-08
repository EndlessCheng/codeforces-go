用一个最大堆模拟，每次循环累加堆顶，同时修改堆顶。

原地修改可以做到 $O(1)$ 空间复杂度。

```py [sol1-Python3]
class Solution:
    def maxKelements(self, nums: List[int], k: int) -> int:
        for i in range(len(nums)):
            nums[i] = -nums[i]  # 最大堆
        heapify(nums)
        ans = 0
        for _ in range(k):
            ans += -heapreplace(nums, nums[0] // 3)
        return ans
```

```go [sol1-Go]
func maxKelements(nums []int, k int) (ans int64) {
	h := hp{nums}
	heap.Init(&h)
	for ; k > 0; k-- {
		ans += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
```

#### 复杂度分析

- 时间复杂度：$O(k\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
