这题和 [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/) 其实是一样的，把那道题的 $1$ 和 $2$ 替换成 $\textit{zero}$ 和 $\textit{one}$ 你就认识了。

具体请看[【双周赛 91】](https://www.bilibili.com/video/BV1gd4y1b7qj/?t=1m49s)第二题。

```py [sol-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = 10 ** 9 + 7
        f = [1] + [0] * high  # f[i] 表示构造长为 i 的字符串的方案数，其中构造空串的方案数为 1
        for i in range(1, high + 1):
            if i >= one:  f[i] = (f[i] + f[i - one]) % MOD
            if i >= zero: f[i] = (f[i] + f[i - zero]) % MOD
        return sum(f[low:]) % MOD
```

```java [sol-Java]
class Solution {
    public int countGoodStrings(int low, int high, int zero, int one) {
        final int MOD = (int) 1e9 + 7;
        int ans = 0;
        var f = new int[high + 1]; // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for (int i = 1; i <= high; i++) {
            if (i >= one) f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
            if (i >= low) ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countGoodStrings(int low, int high, int zero, int one) {
        const int MOD = 1e9 + 7;
        int ans = 0;
        vector<int> f(high + 1); // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for (int i = 1; i <= high; i++) {
            if (i >= one) f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
            if (i >= low) ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countGoodStrings(low, high, zero, one int) (ans int) {
	const mod = 1_000_000_007
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

```js [sol-JavaScript]
var countGoodStrings = function (low, high, zero, one) {
    const mod = 1e9 + 7;
    let ans = 0;
    let f = Array(high + 1).fill(0); // f[i] 表示构造长为 i 的字符串的方案数
    f[0] = 1; // 构造空串的方案数为 1
    for (let i = 1; i <= high; i++) {
        if (i >= one) f[i] = (f[i] + f[i - one]) % mod;
        if (i >= zero) f[i] = (f[i] + f[i - zero]) % mod;
        if (i >= low) ans = (ans + f[i]) % mod;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{high})$。
- 空间复杂度：$\mathcal{O}(\textit{high})$。
