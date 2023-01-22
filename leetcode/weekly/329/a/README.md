下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

从最低位开始模拟。

如果最后到了最高位，发现它取的是负号，则把答案取反。

```py [sol1-Python3]
class Solution:
    def alternateDigitSum(self, n: int) -> int:
        ans, sign = 0, 1
        while n:
            ans += n % 10 * sign
            sign = -sign
            n //= 10
        return ans * -sign
```

```go [sol1-Go]
func alternateDigitSum(n int) (ans int) {
	sign := 1
	for ; n > 0; n /= 10 {
		ans += n % 10 * sign
		sign = -sign
	}
	return ans * -sign
}
```

### 复杂度分析

- 时间复杂度：$O(\log n)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
