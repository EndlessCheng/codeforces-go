### 本题视频讲解

见[【周赛 342】](https://www.bilibili.com/video/BV1Bs4y1A7Wa/)第二题。

### 思路

在 $[1,n]$ 中 $m$ 的倍数有 $c = \left\lfloor\dfrac{n}{m}\right\rfloor$ 个，结合等差数列求和公式，这些数的和为

$$
\dfrac{(1+c)\cdot c}{2} \cdot m
$$

再结合容斥原理，可以算出 $3$ 或 $5$ 或 $7$ 的倍数之和。具体请看视频讲解。

```py [sol1-Python3]
class Solution:
    def sumOfMultiples(self, n: int) -> int:
        def s(m: int) -> int:
            return (1 + n // m) * (n // m) // 2 * m
        return s(3) + s(5) + s(7) - s(15) - s(21) - s(35) + s(105)
```

```java [sol1-Java]
class Solution {
    public int sumOfMultiples(int n) {
        return s(n, 3) + s(n, 5) + s(n, 7) - s(n, 15) - s(n, 21) - s(n, 35) + s(n, 105);
    }

    private int s(int n, int m) {
        return (1 + n / m) * (n / m) / 2 * m;
    }
}
```

```cpp [sol1-C++]
class Solution {
    int s(int n, int m) {
        return (1 + n / m) * (n / m) / 2 * m;
    }
public:
    int sumOfMultiples(int n) {
        return s(n, 3) + s(n, 5) + s(n, 7) - s(n, 15) - s(n, 21) - s(n, 35) + s(n, 105);
    }
};
```

```go [sol1-Go]
func s(n, m int) int {
	return (1 + n/m) * (n / m) / 2 * m
}

func sumOfMultiples(n int) int {
	return s(n, 3) + s(n, 5) + s(n, 7) - s(n, 15) - s(n, 21) - s(n, 35) + s(n, 105)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$，
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

### 相似题目

- [1201. 丑数 III](https://leetcode.cn/problems/ugly-number-iii/)
