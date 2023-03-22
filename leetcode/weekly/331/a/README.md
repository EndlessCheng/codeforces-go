用最大堆模拟。原地堆化可以做到 $O(1)$ 额外空间。

附：[视频讲解](https://www.bilibili.com/video/BV1sG4y1T7oc/)

```py [sol1-Python3]
class Solution:
    def pickGifts(self, gifts: List[int], k: int) -> int:
        for i in range(len(gifts)):
            gifts[i] *= -1  # 最大堆
        heapify(gifts)
        while k and -gifts[0] > 1:
            heapreplace(gifts, -isqrt(-gifts[0]))
            k -= 1
        return -sum(gifts)
```

```go [sol1-Go]
func pickGifts(gifts []int, k int) (ans int64) {
	h := &hp{gifts}
	heap.Init(h)
	for ; k > 0 && gifts[0] > 1; k-- {
		gifts[0] = int(math.Sqrt(float64(gifts[0])))
		heap.Fix(h, 0)
	}
	for _, x := range gifts {
		ans += int64(x)
	}
	return
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}
```

### 复杂度分析

- 时间复杂度：$O(k\log n)$，其中 $n$ 为 $\textit{gifts}$ 的长度。计算平方根有专门的 CPU 指令，可以视作是 $O(1)$ 时间。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
