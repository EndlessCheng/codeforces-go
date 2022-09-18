下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

根据题意，当 $n$ 为奇数时，答案为 $2n$，当 $n$ 为偶数时，答案为 $n$。

因此答案为

$$
(n\bmod 2 + 1) \cdot n
$$

```py [sol1-Python3]
class Solution:
    def smallestEvenMultiple(self, n: int) -> int:
        return (n % 2 + 1) * n
```

```go [sol1-Go]
func smallestEvenMultiple(n int) int {
	return (n%2 + 1) * n
}
```

#### 复杂度分析

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。
