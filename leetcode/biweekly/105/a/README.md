## 视频讲解

见[【双周赛 105】](https://www.bilibili.com/video/BV1YV4y1z74w/) 第一题，欢迎点赞投币！

## 思路

遍历 $\textit{prices}$，维护最小值 $\textit{mn}_1$ 和次小值 $\textit{mn}_2$。

如果 $\textit{mn}_1+\textit{mn}_2>\textit{money}$，那么剩余的钱会变成负数，返回 $\textit{money}$，否则返回 $\textit{money}-\textit{mn}_1-\textit{mn}_2$。

```py [sol-Python3]
class Solution:
    def buyChoco(self, prices: List[int], money: int) -> int:
        mn1 = mn2 = inf
        for p in prices:
            if p < mn1:
                mn2 = mn1
                mn1 = p
            elif p < mn2:
                mn2 = p
        return money if mn1 + mn2 > money else money - mn1 - mn2
```

```go [sol-Go]
func buyChoco(prices []int, money int) int {
	mn1, mn2 := math.MaxInt, math.MinInt
	for _, p := range prices {
		if p < mn1 {
			mn2 = mn1
			mn1 = p
		} else if p < mn2 {
			mn2 = p
		}
	}
	if mn1+mn2 <= money {
		return money - mn1 - mn2
	}
	return money
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
