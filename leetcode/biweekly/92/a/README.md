根据对称性：

- $n$ 为偶数时，每一块扇形都有对称的扇形，所以是切割 $\dfrac{n}{2}$ 次；
- $n$ 为奇数时，不存在对称性，所以是切割 $n$ 次。

注意 $n=1$ 时无需切割。

```py [sol1-Python3]
class Solution:
    def numberOfCuts(self, n: int) -> int:
        return n if n > 1 and n % 2 else n // 2
```

```go [sol1-Go]
func numberOfCuts(n int) int {
	if n == 1 || n%2 == 0 {
		return n / 2
	}
	return n
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
