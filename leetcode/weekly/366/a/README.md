[视频讲解](https://www.bilibili.com/video/BV1e84y117R9/)

设 $k = \left\lfloor\dfrac{n}{m}\right\rfloor$。

$\textit{num}_2$ 是 $[1,n]$ 内的 $m$ 的倍数之和，即

$$
\begin{aligned}
&m + 2m + \cdots + km\\
=\ & (1+2+\cdots+k)\cdot m\\
=\ & \dfrac{k(k+1)}{2}\cdot m
\end{aligned}
$$

$\textit{num}_1$ 相当于 $(1+2+\cdots+n) - \textit{num}_2$，所以

$$
\begin{aligned}
&\textit{num}_1 - \textit{num}_2\\
=\ & (1+2+\cdots+n) - \textit{num}_2 \cdot 2\\
=\ & \dfrac{n(n+1)}{2} - k(k+1)m
\end{aligned}
$$

```py [sol-Python3]
class Solution:
    def differenceOfSums(self, n: int, m: int) -> int:
        return n * (n + 1) // 2 - n // m * (n // m + 1) * m
```

```java [sol-Java]
class Solution {
    public int differenceOfSums(int n, int m) {
        return n * (n + 1) / 2 - n / m * (n / m + 1) * m;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int differenceOfSums(int n, int m) {
        return n * (n + 1) / 2 - n / m * (n / m + 1) * m;
    }
};
```

```go [sol-Go]
func differenceOfSums(n, m int) int {
	return n*(n+1)/2 - n/m*(n/m+1)*m
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
