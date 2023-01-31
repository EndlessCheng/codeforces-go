**请记住：有序是一个非常好的性质。**

把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 组合起来，按照 $\textit{nums}_2[i]$ 从大到小排序。枚举 $\textit{nums}_2[i]$ 作为序列的最小值，那么 $\textit{nums}_1$ 就只能在下标 $\le i$ 的数中选了。要选最大的 $k$ 个数。

根据 [703. 数据流中的第 K 大元素](https://leetcode.cn/problems/kth-largest-element-in-a-stream/)，这可以用一个大小固定为 $k$ 的最小堆来做，如果当前元素大于堆顶，就替换堆顶，这样可以让堆中元素之和变大。

附：[视频讲解](https://www.bilibili.com/video/BV1jG4y197qD/)

```py [sol1-Python3]
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        a = sorted(zip(nums1, nums2), key=lambda p: -p[1])
        h = [x for x, _ in a[:k]]
        heapify(h)
        s = sum(h)
        ans = s * a[k - 1][1]
        for x, y in a[k:]:
            if x > h[0]:
                s += x - heapreplace(h, x)
                ans = max(ans, s * y)
        return ans
```

```go [sol1-Go]
func maxScore(nums1, nums2 []int, k int) int64 {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].y > a[j].y })

	h := hp{nums2[:k]} // 复用内存
	for i, p := range a[:k] {
		sum += p.x
		h.IntSlice[i] = p.x
	}
	ans := sum * a[k-1].y
	heap.Init(&h)
	for _, p := range a[k:] {
		if p.x > h.IntSlice[0] {
			sum += p.x - h.replace(p.x)
			ans = max(ans, sum*p.y)
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }
func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}
func (h hp) replace(v int) int  { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。瓶颈在排序上。
- 空间复杂度：$O(n)$。
