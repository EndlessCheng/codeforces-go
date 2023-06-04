下午两点直播讲题，记得关注哦~（见个人主页）

---

对于本题，把问题转换成：

- 计算 $\le \textit{num}_2$ 的合法数字个数 $a$。
- 计算 $\le \textit{num}_1-1$ 的合法数字个数 $b$。

那么答案就是 $a-b$。

考虑到 $\textit{num}_1$ 是个字符串，可以直接计算 $\le \textit{num}_1$ 的合法数字个数，再单独判断 $\textit{num}_1$ 这个数是否合法。

然后就可以套 [数位 DP 通用模板](https://leetcode.cn/problems/numbers-with-repeated-digits/solutions/1748539/by-endlesscheng-c5vg/) 了。下午直播也会再次教大家这个模板。

把模板中的 $\textit{mask}$ 换成 $\textit{sum}$，表示数位和。在递归中，如果 $\textit{sum}>\textit{maxSum}$ 则直接返回 $0$（因为 $\textit{sum}$ 不可能变小）。递归到终点时，如果 $\textit{sum}\ge \textit{minSum}$，说明找到了**一个**合法的数字，返回 $1$，否则返回 $0$。

此外，由于前导零对数位和无影响（数位和加上 $0$ 不变），$\textit{isNum}$ 可以省略。

> 注：如果你不知道为什么要在计算中途取模，可以看文末的讲解。

```py [sol-Python3]
class Solution:
    def count(self, num1: str, num2: str, min_sum: int, max_sum: int) -> int:
        MOD = 10 ** 9 + 7
        def f(s: string) -> int:
            @cache  # 记忆化搜索
            def f(i: int, sum: int, is_limit: bool) -> int:
                if sum > max_sum:  # 非法
                    return 0
                if i == len(s):
                    return sum >= min_sum
                res = 0
                up = int(s[i]) if is_limit else 9
                for d in range(up + 1):  # 枚举要填入的数字 d
                    res += f(i + 1, sum + d, is_limit and d == up)
                return res % MOD
            return f(0, 0, True)
        ans = f(num2) - f(num1) + (min_sum <= sum(map(int, num1)) <= max_sum)
        return ans % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;
    private int minSum, maxSum;

    public int count(String num1, String num2, int minSum, int maxSum) {
        this.minSum = minSum;
        this.maxSum = maxSum;
        int ans = count(num2) - count(num1) + MOD; // 避免负数
        int sum = 0;
        for (char c : num1.toCharArray()) sum += c - '0';
        if (minSum <= sum && sum <= maxSum) ans++; // x=num1 是合法的，补回来
        return ans % MOD;
    }

    private int count(String S) {
        var s = S.toCharArray();
        int n = s.length;
        var memo = new int[n][Math.min(9 * n, maxSum) + 1];
        for (int i = 0; i < n; i++)
            Arrays.fill(memo[i], -1); // -1 表示没有计算过
        return f(s, memo, 0, 0, true);
    }

    private int f(char[] s, int[][] memo, int i, int sum, boolean isLimit) {
        if (sum > maxSum) return 0; // 非法数字
        if (i == s.length) return sum >= minSum ? 1 : 0;
        if (!isLimit && memo[i][sum] != -1) return memo[i][sum];
        int res = 0;
        int up = isLimit ? s[i] - '0' : 9;
        for (int d = 0; d <= up; ++d) // 枚举要填入的数字 d
            res = (res + f(s, memo, i + 1, sum + d, isLimit && d == up)) % MOD;
        if (!isLimit) memo[i][sum] = res;
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1e9 + 7;

    int f(string s, int min_sum, int max_sum) {
        int n = s.length(), memo[n][min(9 * n, max_sum) + 1];
        memset(memo, -1, sizeof(memo)); // -1 表示没有计算过
        function<int(int, int, bool)> f = [&](int i, int sum, bool is_limit) -> int {
            if (sum > max_sum) return 0; // 非法数字
            if (i == n) return sum >= min_sum;
            if (!is_limit && memo[i][sum] != -1) return memo[i][sum];
            int res = 0;
            int up = is_limit ? s[i] - '0' : 9;
            for (int d = 0; d <= up; ++d) // 枚举要填入的数字 d
                res = (res + f(i + 1, sum + d, is_limit && d == up)) % MOD;
            if (!is_limit) memo[i][sum] = res;
            return res;
        };
        return f(0, 0, true);
    }

public:
    int count(string num1, string num2, int min_sum, int max_sum) {
        int ans = f(num2, min_sum, max_sum) - f(num1, min_sum, max_sum) + MOD; // 避免负数
        int sum = 0;
        for (char c: num1) sum += c - '0';
        ans += min_sum <= sum && sum <= max_sum; // x=num1 是合法的，补回来
        return ans % MOD;
    }
};
```

```go [sol-Go]
func count(num1, num2 string, minSum, maxSum int) int {
	const mod int = 1e9 + 7
	f := func(s string) int {
		memo := make([][]int, len(s))
		for i := range memo {
			memo[i] = make([]int, min(9*len(s), maxSum)+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(p, sum int, limitUp bool) int
		dfs = func(p, sum int, limitUp bool) (res int) {
			if sum > maxSum { // 非法
				return
			}
			if p == len(s) {
				if sum >= minSum { // 合法
					return 1
				}
				return
			}
			if !limitUp {
				ptr := &memo[p][sum]
				if *ptr >= 0 {
					return *ptr
				}
				defer func() { *ptr = res }()
			}
			up := 9
			if limitUp {
				up = int(s[p] - '0')
			}
			for d := 0; d <= up; d++ { // 枚举要填入的数字 d
				res = (res + dfs(p+1, sum+d, limitUp && d == up)) % mod
			}
			return
		}
		return dfs(0, 0, true)
	}
	ans := f(num2) - f(num1) + mod // 避免负数
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}
	if minSum <= sum && sum <= maxSum { // x=num1 是合法的，补回来
		ans++
	}
	return ans % mod
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(10nm)$，其中 $n$ 为 $\textit{nums}_2$ 的长度，$m=\min\{9n, \textit{maxSum}\}$。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(10)$，因此时间复杂度为 $\mathcal{O}(10nm)$。
- 空间复杂度：$\mathcal{O}(nm)$。

## 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模。
