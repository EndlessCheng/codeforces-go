由于每回合玩家恰好摘掉一朵花，所以当且仅当 $x+y$ 是奇数，Alice 必胜。

- 如果 $x$ 是奇数，那么 $y$ 必须是偶数。
- 如果 $x$ 是偶数，那么 $y$ 必须是奇数。

$[1,n]$ 中有 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个偶数，$\left\lceil\dfrac{n}{2}\right\rceil$ 个奇数。对于 $[1,m]$ 也同理。

所以答案为

$$
\left\lfloor\dfrac{n}{2}\right\rfloor\cdot \left\lceil\dfrac{m}{2}\right\rceil + \left\lfloor\dfrac{m}{2}\right\rfloor\cdot \left\lceil\dfrac{n}{2}\right\rceil
$$

进一步优化：

- 当 $y=1,3,5,\cdots$ 时，可以选 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个 $x$。
- 当 $y=2,4,6,\cdots$ 时，可以选 $\left\lceil\dfrac{n}{2}\right\rceil$ 个 $x$。

两个两个一组，把 $y=1$ 和 $y=2$ 加起来，可以选 $\left\lfloor\dfrac{n}{2}\right\rfloor + \left\lceil\dfrac{n}{2}\right\rceil = n$ 个 $x$。依此类推。

- 如果 $m$ 是偶数，那么答案就是 $\dfrac{nm}{2}$。
- 如果 $m$ 是奇数，当 $y=m$ 时，可以选 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个 $x$，所以答案是 $\left\lfloor\dfrac{nm}{2}\right\rfloor$。

综合这两种情况，答案是

$$
\left\lfloor\dfrac{nm}{2}\right\rfloor
$$

[视频讲解](https://www.bilibili.com/video/BV1we411J7Y8/) 中讲了另外一种推导方法（利用国际象棋棋盘），欢迎收看！

```py [sol-Python3]
class Solution:
    def flowerGame(self, n: int, m: int) -> int:
        return n * m // 2
```

```java [sol-Java]
class Solution {
    public long flowerGame(int n, int m) {
        return (long) n * m / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long flowerGame(int n, int m) {
        return (long long) n * m / 2;
    }
};
```

```go [sol-Go]
func flowerGame(n, m int) int64 {
	return int64(n) * int64(m) / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
