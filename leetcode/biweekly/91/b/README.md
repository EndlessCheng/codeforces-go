下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

这题和 [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/) 有什么区别呢？不就是把那道题的 $1$ 和 $2$ 替换成了 $\textit{zero}$ 和 $\textit{one}$ 嘛！

```py [sol1-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = 10 ** 9 + 7
        f = [1] + [0] * high  # f[i] 表示构造长为 i 的字符串的方案数，其中构造空串的方案数为 1
        for i in range(1, high + 1):
            if i >= one:  f[i] = (f[i] + f[i - one]) % MOD
            if i >= zero: f[i] = (f[i] + f[i - zero]) % MOD
        return sum(f[low:]) % MOD
```

```go [sol1-Go]
func countGoodStrings(low, high, zero, one int) (ans int) {
	const mod int = 1e9 + 7
	f := make([]int, high+1) // f[i] 表示构造长为 i 的字符串的方案数
	f[0] = 1 // 构造空串的方案数为 1
	for i := 1; i <= high; i++ {
		if i >= one  { f[i] = (f[i] + f[i-one]) % mod }
		if i >= zero { f[i] = (f[i] + f[i-zero]) % mod }
		if i >= low  { ans = (ans + f[i]) % mod }
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(\textit{high})$。
- 空间复杂度：$O(\textit{high})$。
