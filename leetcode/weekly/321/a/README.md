$1$ 到 $x$ 的元素和为 $\dfrac{x(x+1)}{2}$，$x$ 到 $n$ 的元素和为 $1$ 到 $n$ 的元素和减去 $1$ 到 $x-1$ 的元素和，即 $\dfrac{n(n+1)-x(x-1)}{2}$。

两式相等，简化后即

$$
x = \sqrt{\dfrac{n(n+1)}{2}}
$$

如果 $x$ 不是整数则返回 $-1$。

```py [sol1-Python3]
class Solution:
    def pivotInteger(self, n: int) -> int:
        m = n * (n + 1) // 2
        x = int(m ** 0.5)
        return x if x * x == m else -1
```

```go [sol1-Go]
func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。计算平方根有专门的 CPU 指令，可以视作是 $O(1)$ 时间。
- 空间复杂度：$O(1)$，仅用到若干变量。
