由于 $n-1$ 一定满足要求，不断循环后，$[2,n]$ 都会在桌面上，答案为 $n-1$。

注意特判 $n=1$ 的情况，此时答案为 $1$。

附：[视频讲解](https://www.bilibili.com/video/BV1mD4y1E7QK/)

```py [sol1-Python3]
class Solution:
    def distinctIntegers(self, n: int) -> int:
        return max(n - 1, 1)
```

```go [sol1-Go]
func distinctIntegers(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}
```

### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。
