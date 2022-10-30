基本思路是，不断 $+1$ 直到产生进位，就可能让数位和变小。

代码实现时，可以直接计算每个数位进位后的结果。

比如 $467$，十位数进位为 $470$，百位数进位为 $500$，千位数进位为 $1000$（这一点在 $\textit{target}=1$ 时尤为重要）。

```py [sol1-Python3]
class Solution:
    def makeIntegerBeautiful(self, n: int, target: int) -> int:
        tail = 1
        while True:
            m = x = n + (tail - n % tail) % tail  # 进位后的数字
            s = 0
            while x:
                s += x % 10
                x //= 10
            if s <= target: return m - n
            tail *= 10
```

```go [sol1-Go]
func makeIntegerBeautiful(n int64, target int) int64 {
	for tail := int64(1); ; tail *= 10 {
		m := n + (tail-n%tail)%tail // 进位后的数字
		sum := 0
		for x := m; x > 0; x /= 10 {
			sum += int(x % 10)
		}
		if sum <= target {
			return m - n
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$O(\log^2 n)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
