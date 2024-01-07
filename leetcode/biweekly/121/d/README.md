[v1.0 视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/)

[v2.0 视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)

定义 $\textit{dfs}(i,\textit{limitLow},\textit{limitHigh})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{limitHigh}$ 表示当前是否受到了 $\textit{finish}$ 的约束（我们要构造的数字不能超过 $\textit{finish}$）。若为真，则第 $i$ 位填入的数字至多为 $\textit{finish}[i]$，否则至多为 $9$，这个数记作 $\textit{hi}$。如果在受到约束的情况下填了 $\textit{finish}[i]$，那么后续填入的数字仍会受到 $\textit{finish}$ 的约束。例如 $\textit{finish}=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{limitLow}$ 表示当前是否受到了 $\textit{start}$ 的约束（我们要构造的数字不能低于 $\textit{start}$）。若为真，则第 $i$ 位填入的数字至少为 $\textit{start}[i]$，否则至少为 $0$，这个数记作 $\textit{lo}$。如果在受到约束的情况下填了 $\textit{start}[i]$，那么后续填入的数字仍会受到 $\textit{start}$ 的约束。

枚举第 $i$ 位填什么数字。

如果 $i< n - |s|$，那么可以填 $[\textit{lo}, \min(\textit{hi}, \textit{limit})]$ 内的数，否则只能填 $s[i-(n-|s|)]$。这里 $|s|$ 表示 $s$ 的长度。

为什么不能把 $\textit{hi}$ 置为 $\min(\textit{hi}, \textit{limit})$？请看 [视频](https://www.bilibili.com/video/BV1Fg4y1Q7wv/) 中举的反例。

递归终点：$\textit{dfs}(n,*,*)=1$，表示成功构造出一个合法数字。

递归入口：$\textit{dfs}(0, \texttt{true}, \texttt{true})$，表示：

- 从最高位开始枚举。
- 一开始要受到 $\textit{start}$ 和 $\textit{finish}$ 的约束（否则就可以随意填了，这肯定不行）。

### 答疑

**问**：记忆化三个状态有点麻烦，能不能只记忆化 $i$ 这个状态？

**答**：是可以的。比如 $\textit{finish}=234$，第一位填 $2$，第二位填 $3$，后面无论怎么递归，都不会再次递归到第一位填 $2$，第二位填 $3$ 的情况，所以不需要记录。对于 $\textit{start}$ 也同理。

根据这个例子，我们可以只记录不受到 $\textit{limitLow}$ 或 $\textit{limitHigh}$ 约束时的状态 $i$。相当于记忆化的是 $(i,\texttt{false},\texttt{false})$ 这个状态，因为其它状态只会递归访问一次。

```py [sol-Python3]
class Solution:
    def numberOfPowerfulInt(self, start: int, finish: int, limit: int, s: str) -> int:
        low = str(start)
        high = str(finish)
        n = len(high)
        low = '0' * (n - len(low)) + low  # 补前导零，和 high 对齐
        diff = n - len(s)

        @cache
        def dfs(i: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1

            # 第 i 个数位可以从 lo 枚举到 hi
            # 如果对数位还有其它约束，应当只在下面的 for 循环做限制，不应修改 lo 或 hi
            lo = int(low[i]) if limit_low else 0
            hi = int(high[i]) if limit_high else 9

            res = 0
            if i < diff:  # 枚举这个数位填什么
                for d in range(lo, min(hi, limit) + 1):
                    res += dfs(i + 1, limit_low and d == lo, limit_high and d == hi)
            else:  # 这个数位只能填 s[i-diff]
                x = int(s[i - diff])
                if lo <= x <= min(hi, limit):
                    res = dfs(i + 1, limit_low and x == lo, limit_high and x == hi)
            return res

        return dfs(0, True, True)
```

```java [sol-Java]
class Solution {
    public long numberOfPowerfulInt(long start, long finish, int limit, String s) {
        String low = Long.toString(start);
        String high = Long.toString(finish);
        int n = high.length();
        low = "0".repeat(n - low.length()) + low; // 补前导零，和 high 对齐
        long[] memo = new long[n];
        Arrays.fill(memo, -1);
        return dfs(0, true, true, low.toCharArray(), high.toCharArray(), limit, s.toCharArray(), memo);
    }

    private long dfs(int i, boolean limitLow, boolean limitHigh, char[] low, char[] high, int limit, char[] s, long[] memo) {
        if (i == high.length) {
            return 1;
        }

        if (!limitLow && !limitHigh && memo[i] != -1) {
            return memo[i]; // 之前计算过
        }

        // 第 i 个数位可以从 lo 枚举到 hi
        // 如果对数位还有其它约束，应当只在下面的 for 循环做限制，不应修改 lo 或 hi
        int lo = limitLow ? low[i] - '0' : 0;
        int hi = limitHigh ? high[i] - '0' : 9;

        long res = 0;
        if (i < high.length - s.length) { // 枚举这个数位填什么
            for (int d = lo; d <= Math.min(hi, limit); d++) {
                res += dfs(i + 1, limitLow && d == lo, limitHigh && d == hi, low, high, limit, s, memo);
            }
        } else { // 这个数位只能填 s[i-diff]
            int x = s[i - (high.length - s.length)] - '0';
            if (lo <= x && x <= Math.min(hi, limit)) {
                res = dfs(i + 1, limitLow && x == lo, limitHigh && x == hi, low, high, limit, s, memo);
            }
        }

        if (!limitLow && !limitHigh) {
            memo[i] = res; // 记忆化 (i,false,false)
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfPowerfulInt(long long start, long long finish, int limit, string s) {
        string low = to_string(start);
        string high = to_string(finish);
        int n = high.size();
        low = string(n - low.size(), '0') + low; // 补前导零，和 high 对齐
        int diff = n - s.size();

        vector<long long> memo(n, -1);
        function<long long(int, bool, bool)> dfs = [&](int i, bool limit_low, bool limit_high) -> long long {
            if (i == low.size()) {
                return 1;
            }

            if (!limit_low && !limit_high && memo[i] != -1) {
                return memo[i]; // 之前计算过
            }

            // 第 i 个数位可以从 lo 枚举到 hi
            // 如果对数位还有其它约束，应当只在下面的 for 循环做限制，不应修改 lo 或 hi
            int lo = limit_low ? low[i] - '0' : 0;
            int hi = limit_high ? high[i] - '0' : 9;

            long long res = 0;
            if (i < diff) { // 枚举这个数位填什么
                for (int d = lo; d <= min(hi, limit); d++) {
                    res += dfs(i + 1, limit_low && d == lo, limit_high && d == hi);
                }
            } else { // 这个数位只能填 s[i-diff]
                int x = s[i - diff] - '0';
                if (lo <= x && x <= min(hi, limit)) {
                    res = dfs(i + 1, limit_low && x == lo, limit_high && x == hi);
                }
            }

            if (!limit_low && !limit_high) {
                memo[i] = res; // 记忆化 (i,false,false)
            }
            return res;
        };
        return dfs(0, true, true);
    }
};
```

```go [sol-Go]
func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	low := strconv.FormatInt(start, 10)
	high := strconv.FormatInt(finish, 10)
	n := len(high)
	low = strings.Repeat("0", n-len(low)) + low // 补前导零，和 high 对齐
	diff := n - len(s)

	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int, bool, bool) int64
	dfs = func(i int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			return 1
		}
		
		if !limitLow && !limitHigh {
			p := &memo[i]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		// 第 i 个数位可以从 lo 枚举到 hi
		// 如果对数位还有其它约束，应当只在下面的 for 循环做限制，不应修改 lo 或 hi
		lo := 0
		if limitLow {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		if i < diff { // 枚举这个数位填什么
			for d := lo; d <= min(hi, limit); d++ {
				res += dfs(i+1, limitLow && d == lo, limitHigh && d == hi)
			}
		} else { // 这个数位只能填 s[i-diff]
			x := int(s[i-diff] - '0')
			if lo <= x && x <= min(hi, limit) {
				res += dfs(i+1, limitLow && x == lo, limitHigh && x == hi)
			}
		}
		return
	}
	return dfs(0, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD)$，其中 $n=\mathcal{O}(\log \textit{finish})$，$D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以动态规划的时间复杂度为 $\mathcal{O}(nD)$。
- 空间复杂度：$\mathcal{O}(n)$。即状态个数。

## 题单：数位 DP

- [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/)（[题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/)）
- [面试题 17.06. 2 出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci/)（[题解](https://leetcode.cn/problems/number-of-2s-in-range-lcci/solution/by-endlesscheng-x4mf/)）
- [357. 统计各位数字都不同的数字个数](https://leetcode.cn/problems/count-numbers-with-unique-digits/)
- [600. 不含连续 1 的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/)（[题解](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/solution/by-endlesscheng-1egu/)）
- [788. 旋转数字](https://leetcode.cn/problems/rotated-digits/)
- [902. 最大为 N 的数字组合](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/) 1990
- [2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/) 2120
- [1012. 至少有 1 位重复的数字](https://leetcode.cn/problems/numbers-with-repeated-digits/)（[题解](https://leetcode.cn/problems/numbers-with-repeated-digits/solution/by-endlesscheng-c5vg/)）2230
- [2827. 范围中美丽整数的数目](https://leetcode.cn/problems/number-of-beautiful-integers-in-the-range/) 2324
- [2719. 统计整数数目](https://leetcode.cn/problems/count-of-integers/) 2355
- [2801. 统计范围内的步进数字数目](https://leetcode.cn/problems/count-stepping-numbers-in-range/) 2367
- [1397. 找到所有好字符串](https://leetcode.cn/problems/find-all-good-strings/) 2667
- [1215. 步进数](https://leetcode.cn/problems/stepping-numbers/)（会员题）1675
- [248. 中心对称数 III](https://leetcode.cn/problems/strobogrammatic-number-iii/)（会员题）
- [1067. 范围内的数字计数](https://leetcode.cn/problems/digit-count-in-range/)（会员题）2025
- [1088. 易混淆数 II](https://leetcode.cn/problems/confusing-number-ii/)（会员题）2077

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
