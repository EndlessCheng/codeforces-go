根据题意，回文数的左半部分为

$$
10^{\Big\lfloor\dfrac{\textit{intLength}-1}{2}\Big\rfloor} + q - 1
$$

反转这个数，拼到左半部分之后即为第 $q$ 个回文数。

如果 $\textit{intLength}$ 为奇数则先去掉最低位再反转。

```python [sol1-Python3]
class Solution:
    def kthPalindrome(self, queries: List[int], intLength: int) -> List[int]:
        ans = [-1] * len(queries)
        base = 10 ** ((intLength - 1) // 2)
        for i, q in enumerate(queries):
            if q <= 9 * base:
                s = str(base + q - 1)  # 回文数左半部分
                s += s[::-1][intLength % 2:]
                ans[i] = int(s)
        return ans
```

```go [sol1-Go]
func kthPalindrome(queries []int, intLength int) []int64 {
	ans := make([]int64, len(queries))
	base := int(math.Pow10((intLength - 1) / 2))
	for i, q := range queries {
		if q > 9*base {
			ans[i] = -1
			continue
		}
		v := base + q - 1 // 回文数左半部分
		x := v
		if intLength%2 == 1 { x /= 10 } // 去掉最低位
		for ; x > 0; x /= 10 {
			v = v*10 + x%10 // 翻转 x 到 v 后
		}
		ans[i] = int64(v)
	}
	return ans
}
```
