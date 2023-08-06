题目意思是把 $\textit{purchaseAmount}$ 的个位数「四舍五入」，这可以用如下公式

$$
\left\lceil\dfrac{\textit{purchaseAmount}+5}{10}\right\rceil\cdot 10
$$

```py [sol-Python3]
class Solution:
    def accountBalanceAfterPurchase(self, purchaseAmount: int) -> int:
        return 100 - (purchaseAmount + 5) // 10 * 10
```

```go [sol-Go]
func accountBalanceAfterPurchase(purchaseAmount int) int {
	return 100 - (purchaseAmount+5)/10*10
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

---

下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！
