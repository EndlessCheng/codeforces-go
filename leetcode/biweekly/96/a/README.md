先来个一行写法：把并集求出来，然后求最小值。

```py [sol1-Python3]
class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        return min(set(nums1) & set(nums2), default=-1)
```

双指针可以把空间复杂度优化到 $O(1)$。

```py [sol2-Python3]
class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        j, m = 0, len(nums2)
        for x in nums1:
            while j < m and nums2[j] < x:  # 找下一个 nums2[j] >= x
                j += 1
            if j < m and nums2[j] == x:
                return x
        return -1
```

```go [sol2-Go]
func getCommon(nums1, nums2 []int) int {
	j, m := 0, len(nums2)
	for _, x := range nums1 {
		for j < m && nums2[j] < x { // 找下一个 nums2[j] >= x
			j++
		}
		if j < m && nums2[j] == x {
			return x
		}
	}
	return -1
}
```

附：[视频讲解](https://www.bilibili.com/video/BV1jG4y197qD/)

#### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
