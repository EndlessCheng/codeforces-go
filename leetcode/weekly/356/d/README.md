视频讲解：[【周赛 356】](https://www.bilibili.com/video/BV1BM4y1W7AQ/) 第四题。

### 前置知识：记忆化搜索

见[【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

### 思路

定义 $\text{calc}(n)$ 表示不超过 $n$ 的步进数字数目。那么答案就是

$$
\begin{aligned}
&\text{calc}(\textit{high}) - \text{calc}(\textit{low}-1)\\
=\ &\text{calc}(\textit{high}) - \text{calc}(\textit{low}) + \text{valid}(\textit{low})
\end{aligned}
$$

由于 $\textit{low}$ 是个很大的数字，不方便减一（Python 用户可以无视），所以用 $\text{valid}(\textit{low})$ 表示：如果 $\textit{low}$ 是步进数字，那么多减了 $1$，再加 $1$ 补回来。

如何计算 $\text{calc}(n)$ 呢？（下文用 $s$ 表示 $n$ 的字符串形式）

一种通用套路是，定义 $f(i,\textit{pre}, \textit{isLimit},\textit{isNum})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{pre}$ 表示上一个数位的值。如果 $\textit{isNum}$ 为 `false`，可以忽略 $\textit{pre}$。
- $\textit{isLimit}$ 表示当前是否受到了 $n$ 的约束（注意要构造的数字不能超过 $n$）。若为真，则第 $i$ 位填入的数字至多为 $s[i]$，否则可以是 $9$。如果在受到约束的情况下填了 $s[i]$，那么后续填入的数字仍会受到 $n$ 的约束。例如 $n=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{isNum}$ 表示 $i$ 前面的数位是否填了数字。若为假，则当前位可以跳过（不填数字），或者要填入的数字至少为 $1$；若为真，则要填入的数字可以从 $0$ 开始。例如 $n=123$，在 $i=0$ 时跳过的话，相当于后面要构造的是一个 $99$ 以内的数字了，如果 $i=1$ 不跳过，那么相当于构造一个 $10$ 到 $99$ 的两位数，如果 $i=1$ 也跳过，相当于构造的是一个 $9$ 以内的数字。

### 实现细节

递归入口：`f(0, 0, true, false)`，表示：

- 从 $s[0]$ 开始枚举；
- $\textit{pre}$ 初始化成什么都可以，因为填第一个数的时候是忽略 $\textit{pre}$ 的。
- 一开始要受到 $n$ 的约束（否则就可以随意填了，这肯定不行）；
- 一开始没有填数字。

递归中：

- 如果 $\textit{isNum}$ 为假，说明前面没有填数字，那么当前也可以不填数字。一旦从这里递归下去，$\textit{isLimit}$ 就可以置为 `false` 了，这是因为 $s[0]$ 必然是大于 $0$ 的，后面就不受到 $n$ 的约束了。或者说，最高位不填数字，后面无论怎么填都比 $n$ 小。
- 如果 $\textit{isNum}$ 为真，那么当前必须填一个数字。枚举填入的数字，根据 $\textit{isNum}$ 和 $\textit{isLimit}$ 来决定填入数字的范围。

递归终点：当 $i$ 等于 $s$ 长度时，如果 $\textit{isNum}$ 为真，则表示得到了一个合法数字（因为不合法的不会继续递归下去），返回 $1$，否则返回 $0$。

关于取模的细节，见文末的讲解。

### 答疑

**问**：记忆化四个状态有点麻烦，能不能只记忆化 $i$ 和 $\textit{pre}$ 这两个状态？

**答**：可以的！比如 $n=234$，第一位填 $2$，第二位填 $3$，后面无论怎么递归，都不会再次递归到第一位填 $2$，第二位填 $3$ 的情况，所以不需要记录。又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。

根据这个例子，我们可以只记录不受到约束时的状态 $(i,\textit{mask},\text{false},\text{true})$。比如 $n=456$，第一位（最高位）填的 $3$，那么继续递归，后面就可以随便填，所以状态 $(1,3,\text{false},\text{true})$ 就表示 $i=0$ 填 $3$，从 $i=1$ 往后随便填的方案数。

由于后面两个参数恒为 $\text{false}$ 和 $\text{true}$，所以可以不用记忆化，只记忆化 $i$ 和 $\textit{pre}$。

> 注：Python 有 `@cache`，可以无视上面说的。

**问**：能不能只记忆化 $i$？

**答**：这是不行的。想一想，我们为什么要用记忆化？如果递归到同一个状态时，计算出的结果是一样的，那么第二次递归到同一个状态，就可以直接返回第一次计算的结果了。通过保存第一次计算的结果，来优化时间复杂度。

由于前面选的数字会影响后面选的数字，两次递归到相同的 $i$，如果前面选的数字不一样，计算出的结果就可能是不一样的。如果只记忆化 $i$，就可能会算出错误的结果。

也可以这样理解：记忆化搜索要求递归函数无副作用（除了修改 `memo` 数组），从而保证递归到同一个状态时，计算出的结果是一样的。

```py [sol1-Python3]
class Solution:
    def countSteppingNumbers(self, low: str, high: str) -> int:
        MOD = 10 ** 9 + 7
        def calc(s: str) -> int:
            @cache  # 记忆化搜索
            def f(i: int, pre: int, is_limit: bool, is_num: bool) -> int:
                if i == len(s):
                    return int(is_num)  # is_num 为 True 表示得到了一个合法数字
                res = 0
                if not is_num:  # 可以跳过当前数位
                    res = f(i + 1, pre, False, False)
                low = 0 if is_num else 1  # 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
                up = int(s[i]) if is_limit else 9  # 如果前面填的数字都和 s 的一样，那么这一位至多填 s[i]（否则就超过 s 啦）
                for d in range(low, up + 1):  # 枚举要填入的数字 d
                    if not is_num or abs(d - pre) == 1:  # 第一位数字随便填，其余必须相差 1
                        res += f(i + 1, d, is_limit and d == up, True)
                return res % MOD
            return f(0, 0, True, False)
        return (calc(high) - calc(str(int(low) - 1))) % MOD
```

```java [sol1-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;

    public int countSteppingNumbers(String low, String high) {
        return (calc(high) - calc(low) + MOD + (valid(low) ? 1 : 0)) % MOD; // +MOD 防止算出负数
    }

    private char s[];
    private int memo[][];

    private int calc(String s) {
        this.s = s.toCharArray();
        int m = s.length();
        memo = new int[m][10];
        for (int i = 0; i < m; i++)
            Arrays.fill(memo[i], -1); // -1 表示没有计算过
        return f(0, 0, true, false);
    }

    private int f(int i, int pre, boolean isLimit, boolean isNum) {
        if (i == s.length)
            return isNum ? 1 : 0; // isNum 为 true 表示得到了一个合法数字
        if (!isLimit && isNum && memo[i][pre] != -1)
            return memo[i][pre];
        int res = 0;
        if (!isNum) // 可以跳过当前数位
            res = f(i + 1, pre, false, false);
        int up = isLimit ? s[i] - '0' : 9; // 如果前面填的数字都和 s 的一样，那么这一位至多填数字 s[i]（否则就超过 s 啦）
        for (int d = isNum ? 0 : 1; d <= up; d++) // 枚举要填入的数字 d
            if (!isNum || Math.abs(d - pre) == 1) // 第一位数字随便填，其余必须相差 1
                res = (res + f(i + 1, d, isLimit && d == up, true)) % MOD;
        if (!isLimit && isNum)
            memo[i][pre] = res;
        return res;
    }

    private boolean valid(String s) {
        for (int i = 1; i < s.length(); i++)
            if (Math.abs((int) s.charAt(i) - (int) s.charAt(i - 1)) != 1)
                return false;
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;

    int calc(string &s) {
        int m = s.length(), memo[m][10];
        memset(memo, -1, sizeof(memo)); // -1 表示没有计算过
        function<int(int, int, bool, bool)> f = [&](int i, int pre, bool is_limit, bool is_num) -> int {
            if (i == m)
                return is_num; // is_num 为 true 表示得到了一个合法数字
            if (!is_limit && is_num && memo[i][pre] != -1)
                return memo[i][pre];
            int res = 0;
            if (!is_num) // 可以跳过当前数位
                res = f(i + 1, pre, false, false);
            int up = is_limit ? s[i] - '0' : 9; // 如果前面填的数字都和 s 的一样，那么这一位至多填数字 s[i]（否则就超过 s 啦）
            for (int d = 1 - is_num; d <= up; ++d) // 枚举要填入的数字 d
                if (!is_num || abs(d - pre) == 1) // 第一位数字随便填，其余必须相差 1
                    res = (res + f(i + 1, d, is_limit && d == up, true)) % MOD;
            if (!is_limit && is_num)
                memo[i][pre] = res;
            return res;
        };
        return f(0, 0, true, false);
    }

    bool valid(string &s) {
        for (int i = 1; i < s.length(); i++)
            if (abs(int(s[i]) - int(s[i - 1])) != 1)
                return false;
        return true;
    }

public:
    int countSteppingNumbers(string low, string high) {
        return (calc(high) - calc(low) + MOD + valid(low)) % MOD; // +MOD 防止算出负数
    }
};
```

```go [sol1-Go]
const mod = 1_000_000_007

func calc(s string) int {
	memo := make([][10]int, len(s))
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, bool, bool) int
	f = func(i, pre int, isLimit, isNum bool) (res int) {
		if i == len(s) {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return 0
		}
		if !isLimit && isNum {
			p := &memo[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, pre, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 s 的一样，那么这一位至多填数字 s[i]（否则就超过 s 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if !isNum || abs(d-pre) == 1 { // 第一位数字随便填，其余必须相差 1
				res += f(i+1, d, isLimit && d == up, true)
			}
		}
		return res % mod // 记得取模，注意这可能会导致 calc(high) < calc(low)
	}
	return f(0, 0, true, false)
}

func valid(s string) bool {
	for i := 1; i < len(s); i++ {
		if abs(int(s[i-1])-int(s[i])) != 1 {
			return false
		}
	}
	return true
}

func countSteppingNumbers(low, high string) int {
	ans := (calc(high) - calc(low) + mod) % mod // +mod 防止算出负数
	if valid(low) {
		ans = (ans + 1) % mod
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD^2)$，其中 $n$ 为 $\textit{high}$ 的长度，$D=10$。由于每个状态只会计算一次，因此动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数为 $\mathcal{O}(nD)$，单个状态的计算时间为 $\mathcal{O}(D)$，因此时间复杂度为 $\mathcal{O}(nD^2)$。
- 空间复杂度：$\mathcal{O}(nD)$。

### 强化训练（数位 DP）

- [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/)（[题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/)）
- [面试题 17.06. 2出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci/)（[题解](https://leetcode.cn/problems/number-of-2s-in-range-lcci/solution/by-endlesscheng-x4mf/)）
- [600. 不含连续1的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/)（[题解](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/solution/by-endlesscheng-1egu/)）
- [902. 最大为 N 的数字组合](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/)（[数位 DP 通用模板](https://www.bilibili.com/video/BV1rS4y1s721/?t=33m22s) 33:22）
- [1012. 至少有 1 位重复的数字](https://leetcode.cn/problems/numbers-with-repeated-digits/)（[题解](https://leetcode.cn/problems/numbers-with-repeated-digits/solution/by-endlesscheng-c5vg/)）
- [1067. 范围内的数字计数](https://leetcode.cn/problems/digit-count-in-range/)
- [1397. 找到所有好字符串](https://leetcode.cn/problems/find-all-good-strings/)（有难度，需要结合一个经典字符串算法）

更多题目见我模板库中的 [dp.go](https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/dp.go#L1924)（搜索 `数位`）。

### 附：模运算

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

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。
