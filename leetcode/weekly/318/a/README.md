不需要额外创建数组，直接在原数组上操作。

```py [sol1-Python3]
class Solution:
    def applyOperations(self, a: List[int]) -> List[int]:
        j, n = 0, len(a)
        for i in range(n - 1):
            if a[i]:
                if a[i] == a[i + 1]:
                    a[i] *= 2
                    a[i + 1] = 0
                a[j] = a[i]  # 非零数字排在前面
                j += 1
        if a[-1]:
            a[j] = a[-1]
            j += 1
        for i in range(j, n):
            a[i] = 0
        return a
```

```go [sol1-Go]
func applyOperations(a []int) []int {
	n := len(a)
	b := a[:0]
	for i := 0; i < n-1; i++ {
		if a[i] > 0 {
			if a[i] == a[i+1] {
				a[i] *= 2
				a[i+1] = 0
			}
			b = append(b, a[i]) // 非零数字排在前面
		}
	}
	if a[n-1] > 0 {
		b = append(b, a[n-1])
	}
	for i := len(b); i < n; i++ {
		a[i] = 0
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。
