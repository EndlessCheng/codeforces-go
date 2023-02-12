按题意模拟。

附：[视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)

```py [sol1-Python3]
class Solution:
    def findTheArrayConcVal(self, nums: List[int]) -> int:
        ans, i, j = 0, 0, len(nums) - 1
        while i < j:
            x, y = nums[i], nums[j]
            while y:
                x *= 10
                y //= 10
            ans += x + nums[j]
            i += 1
            j -= 1
        if i == j: ans += nums[i]
        return ans
```

```go [sol1-Go]
func findTheArrayConcVal(a []int) (ans int64) {
	for len(a) > 1 {
		x := a[0]
		for y := a[len(a)-1]; y > 0; y /= 10 {
			x *= 10
		}
		ans += int64(x + a[len(a)-1])
		a = a[1 : len(a)-1]
	}
	if len(a) > 0 {
		ans += int64(a[0])
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
