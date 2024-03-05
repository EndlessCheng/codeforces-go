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

最后，别忘了对 $10^9+7$ 取模。

```py [sol-Python3]
class Solution:
    def minimumPossibleSum(self, n: int, k: int) -> int:
        m = min(k // 2, n)
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) // 2 % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int minimumPossibleSum(int n, int k) {
        long m = Math.min(k / 2, n);
        return (int) ((m * (m + 1) + (n - m - 1 + k * 2) * (n - m)) / 2 % 1_000_000_007);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumPossibleSum(int n, int k) {
        long long m = min(k / 2, n);
        return (m * (m + 1) + (n - m - 1 + k * 2) * (n - m)) / 2 % 1'000'000'007;
    }
};
```

```go [sol-Go]
func minimumPossibleSum(n, k int) int {
	m := min(k/2, n)
	return (m*(m+1) + (k*2+n-m-1)*(n-m)) / 2 % 1_000_000_007
}
```

```js [sol-JavaScript]
var minimumPossibleSum = function(n, k) {
    const m = Math.min(k >> 1, n);
    return ((BigInt(m) * BigInt(m + 1) + BigInt(k * 2 + n - m - 1) * BigInt(n - m)) / 2n) % 1_000_000_007n;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_possible_sum(n: i32, target: i32) -> i32 {
        let n = n as i64;
        let k = target as i64;
        let m = n.min(k / 2);
        ((m * (m + 1) + (n - m - 1 + k * 2) * (n - m)) / 2 % 1_000_000_007) as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
