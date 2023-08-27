请看 [视频讲解](https://www.bilibili.com/video/BV1Rx4y1f75Y/) 第二题。

把 $\textit{target}$ 记作 $k$。

对于 $[1,k-1]$ 内的数字：

- $1$ 和 $k-1$ 只能选其中一个；
- $2$ 和 $k-2$ 只能选其中一个；
- $3$ 和 $k-3$ 只能选其中一个；
- ……
- 一直到 $\left\lfloor\dfrac{k}{2}\right\rfloor$，无论 $k$ 是奇数还是偶数，它都可以选。

设 $m=\min\left(\left\lfloor\dfrac{k}{2}\right\rfloor, n\right)$，那么答案的第一段是从 $1$ 到 $m$，元素和为

$$
\dfrac{m(m+1)}{2}
$$

此时还剩下 $n-m$ 个数，只能从 $k$ 开始往后选，那么答案的第二段是从 $k$ 到 $k+n-m-1$，元素和为

$$
\dfrac{(2k+n-m-1)(n-m)}{2}
$$

所以答案为

$$
\dfrac{m(m+1) + (2k+n-m-1)(n-m)}{2}
$$

```py [sol-Python3]
class Solution:
    def minimumPossibleSum(self, n: int, k: int) -> int:
        m = min(k // 2, n)
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) // 2
```

```java [sol-Java]
class Solution {
    public long minimumPossibleSum(int n, int k) {
        long m = Math.min(k / 2, n);
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumPossibleSum(int n, int k) {
        long long m = min(k / 2, n);
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2;
    }
};
```

```go [sol-Go]
func minimumPossibleSum(n, k int) int64 {
    m := min(k/2, n)
    return int64((m*(m+1) + (k*2+n-m-1)*(n-m)) / 2)
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
