下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

```py [sol-Python3]
class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        ans = 0
        for i in range(low, high + 1):
            s = str(i)
            n = len(s)
            ans += n % 2 == 0 and sum(map(int, s[:n // 2])) == sum(map(int, s[n // 2:]))
        return ans
```

```go [sol-Go]
func countSymmetricIntegers(low int, high int) (ans int) {
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		n := len(s)
		if n%2 > 0 {
			continue
		}
		sum := 0
		for _, c := range s[:n/2] {
			sum += int(c)
		}
		for _, c := range s[n/2:] {
			sum -= int(c)
		}
		if sum == 0 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((\textit{high} - \textit{low})\log \textit{high})$。
- 空间复杂度：$\mathcal{O}(\log \textit{high})$。

## 思考题

你能用 [数位 DP](https://www.bilibili.com/video/BV1rS4y1s721/?t=20m05s) 解决本题吗？

时间复杂度可以做到 $\mathcal{O}(\log^2 \textit{high})$。
