## v1.0：两次记忆化搜索

[v1.0 模板视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s)

把问题拆分成：

- 计算 $\le \textit{num}_2$ 的好整数个数，记作 $a$。
- 计算 $\le \textit{num}_1-1$ 的好整数个数，记作 $b$。

那么答案就是 $a-b$。

考虑到 $\textit{num}_1$ 是个字符串，不方便减一，可以改为计算 $\le \textit{num}_1$ 的合法数字个数，再单独判断 $\textit{num}_1$ 这个数是否合法。

把 [v1.0 模板](https://leetcode.cn/problems/numbers-with-repeated-digits/solution/by-endlesscheng-c5vg/) 中的参数 $\textit{mask}$ 去掉，加上参数 $\textit{sum}$，表示数位和。在递归中，如果 $\textit{sum}>\textit{maxSum}$ 则直接返回 $0$（因为 $\textit{sum}$ 不可能变小）。递归到终点时，如果 $\textit{sum}\ge \textit{minSum}$，说明我们成功构造出一个好整数，返回 $1$，否则返回 $0$。

此外，由于前导零对数位和无影响（$\textit{sum}+0=\textit{sum}$），模板中的 $\textit{isNum}$ 可以省略。

```py [sol-Python3]
class Solution:
    def count(self, num1: str, num2: str, min_sum: int, max_sum: int) -> int:
        def calc(high: str) -> int:
            @cache
            def dfs(i: int, s: int, is_limit: bool) -> int:
                if s > max_sum:  # 非法
                    return 0
                if i == len(high):
                    return s >= min_sum
                res = 0
                up = int(high[i]) if is_limit else 9
                for d in range(up + 1):  # 枚举当前数位填 d
                    res += dfs(i + 1, s + d, is_limit and d == up)
                return res
            return dfs(0, 0, True)

        is_num1_good = min_sum <= sum(map(int, num1)) <= max_sum
        return (calc(num2) - calc(num1) + is_num1_good) % 1_000_000_007
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int count(String num1, String num2, int minSum, int maxSum) {
        int ans = calc(num2, minSum, maxSum) - calc(num1, minSum, maxSum) + MOD; // 避免负数
        int sum = 0;
        for (char c : num1.toCharArray()) {
            sum += c - '0';
        }
        if (minSum <= sum && sum <= maxSum) {
            ans++; // num1 是合法的，补回来
        }
        return ans % MOD;
    }

    private int calc(String s, int minSum, int maxSum) {
        int n = s.length();
        int[][] memo = new int[n][Math.min(9 * n, maxSum) + 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1);
        }
        return dfs(0, 0, true, s.toCharArray(), minSum, maxSum, memo);
    }

    private int dfs(int i, int sum, boolean isLimit, char[] s, int minSum, int maxSum, int[][] memo) {
        if (sum > maxSum) { // 非法
            return 0;
        }
        if (i == s.length) {
            return sum >= minSum ? 1 : 0;
        }
        if (!isLimit && memo[i][sum] != -1) {
            return memo[i][sum];
        }

        int up = isLimit ? s[i] - '0' : 9;
        int res = 0;
        for (int d = 0; d <= up; d++) { // 枚举当前数位填 d
            res = (res + dfs(i + 1, sum + d, isLimit && (d == up), s, minSum, maxSum, memo)) % MOD;
        }

        if (!isLimit) {
            memo[i][sum] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    int calc(string &s, int min_sum, int max_sum) {
        int n = s.length();
        vector<vector<int>> memo(n, vector<int>(min(9 * n, max_sum) + 1, -1));
        function<int(int, int, bool)> dfs = [&](int i, int sum, bool is_limit) -> int {
            if (sum > max_sum) { // 非法
                return 0;
            }
            if (i == n) {
                return sum >= min_sum ? 1 : 0;
            }
            if (!is_limit && memo[i][sum] != -1) {
                return memo[i][sum];
            }

            int up = is_limit ? s[i] - '0' : 9;
            int res = 0;
            for (int d = 0; d <= up; d++) { // 枚举当前数位填 d
                res = (res + dfs(i + 1, sum + d, is_limit && d == up)) % MOD;
            }

            if (!is_limit) {
                memo[i][sum] = res;
            }
            return res;
        };
        return dfs(0, 0, true);
    }

public:
    int count(string num1, string num2, int min_sum, int max_sum) {
        int ans = calc(num2, min_sum, max_sum) - calc(num1, min_sum, max_sum) + MOD; // 避免负数
        int sum = 0;
        for (char c : num1) {
            sum += c - '0';
        }
        ans += min_sum <= sum && sum <= max_sum; // num1 是合法的，补回来
        return ans % MOD;
    }
};
```

```go [sol-Go]
func count(num1, num2 string, minSum, maxSum int) int {
	const mod = 1_000_000_007
	calc := func(s string) int {
		memo := make([][]int, len(s))
		for i := range memo {
			memo[i] = make([]int, min(9*len(s), maxSum)+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int, bool) int
		dfs = func(i, sum int, isLimit bool) (res int) {
			if sum > maxSum { // 非法
				return
			}
			if i == len(s) {
				if sum >= minSum { // 合法
					return 1
				}
				return
			}
			if !isLimit {
				p := &memo[i][sum]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}
			up := 9
			if isLimit {
				up = int(s[i] - '0')
			}
			for d := 0; d <= up; d++ { // 枚举当前数位填 d
				res = (res + dfs(i+1, sum+d, isLimit && d == up)) % mod
			}
			return
		}
		return dfs(0, 0, true)
	}
	ans := calc(num2) - calc(num1) + mod // 避免负数
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}
	if minSum <= sum && sum <= maxSum { // num1 是合法的，补回来
		ans++
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nmD)$，其中 $n$ 为 $\textit{nums}_2$ 的长度，$m=\min(9n, \textit{maxSum})$，$D=10$。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(D)$，因此时间复杂度为 $\mathcal{O}(nmD)$。
- 空间复杂度：$\mathcal{O}(nm)$。

## v2.0：一次记忆化搜索

v2.0 版本在 v1.0 的基础上做了改进，只需要一次记忆化搜索就能算出答案，效率更高。

[v2.0 模板视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s)

仿照 $\textit{isLimit}$，再添加一个参数 $\textit{limitLow}$，表示是否受到下界 $\textit{num}_1$ 的约束。为了让代码更清晰，原来的参数名 $\textit{isLimit}$ 改名为 $\textit{limitHigh}$。此外，$d$ 的最大值 $\textit{up}$ 改名为 $\textit{hi}$，即 $\textit{high}$ 的前两个字母。

为了方便计算，在 $\textit{num}_1$ 前面添加前导零，使其长度和 $\textit{num}_2$ 相等。

$\textit{limitLow}$ 的用法类似 $\textit{limitHigh}$，如果为 $\textit{limitLow}=\texttt{true}$，那么 $d$ 只能从 $\textit{num}_1[i]$ 开始枚举，否则可以从 $0$ 开始，这个值记作 $\textit{lo}$，即 $\textit{low}$ 的前两个字母。如果 $\textit{limitLow}=\texttt{true}$ 并且 $d=\textit{lo}$，那么往下递归时，传入的 $\textit{limitLow}$ 仍然为 $\texttt{true}$，否则为 $\texttt{false}$。

```py [sol-Python3]
class Solution:
    def count(self, num1: str, num2: str, min_sum: int, max_sum: int) -> int:
        high = list(map(int, num2))  # 避免在 dfs 中频繁调用 int()
        n = len(high)
        low = list(map(int, num1.zfill(n)))  # 补前导零（本题前导零不影响答案，可以补前导零，简化代码逻辑）

        @cache
        def dfs(i: int, s: int, limit_low: bool, limit_high: bool) -> int:
            if s > max_sum:  # 非法
                return 0
            if i == n:
                return 1 if s >= min_sum else 0

            lo = low[i] if limit_low else 0
            hi = high[i] if limit_high else 9

            res = 0
            for d in range(lo, hi + 1):  # 枚举当前数位填 d
                res += dfs(i + 1, s + d, limit_low and d == lo, limit_high and d == hi)
            return res

        return dfs(0, 0, True, True) % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int count(String num1, String num2, int minSum, int maxSum) {
        int n = num2.length();
        num1 = "0".repeat(n - num1.length()) + num1; // 补前导零（本题前导零不影响答案，可以补前导零，简化代码逻辑）

        int[][] memo = new int[n][Math.min(9 * n, maxSum) + 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1);
        }

        return dfs(0, 0, true, true, num1.toCharArray(), num2.toCharArray(), minSum, maxSum, memo);
    }

    private int dfs(int i, int sum, boolean limitLow, boolean limitHigh, char[] num1, char[] num2, int minSum, int maxSum, int[][] memo) {
        if (sum > maxSum) { // 非法
            return 0;
        }
        if (i == num2.length) {
            return sum >= minSum ? 1 : 0;
        }
        if (!limitLow && !limitHigh && memo[i][sum] != -1) {
            return memo[i][sum];
        }

        int lo = limitLow ? num1[i] - '0' : 0;
        int hi = limitHigh ? num2[i] - '0' : 9;

        int res = 0;
        for (int d = lo; d <= hi; d++) { // 枚举当前数位填 d
            res = (res + dfs(i + 1, sum + d, limitLow && d == lo, limitHigh && d == hi,
                             num1, num2, minSum, maxSum, memo)) % 1_000_000_007;
        }

        if (!limitLow && !limitHigh) {
            memo[i][sum] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int count(string num1, string num2, int min_sum, int max_sum) {
        int n = num2.size();
        num1 = string(n - num1.size(), '0') + num1; // 补前导零（本题前导零不影响答案，可以补前导零，简化代码逻辑）

        vector memo(n, vector<int>(min(9 * n, max_sum) + 1, -1));
        auto dfs = [&](this auto&& dfs, int i, int sum, bool limit_low, bool limit_high) -> int {
            if (sum > max_sum) { // 非法
                return 0;
            }
            if (i == n) {
                return sum >= min_sum;
            }
            if (!limit_low && !limit_high && memo[i][sum] != -1) {
                return memo[i][sum];
            }

            int lo = limit_low ? num1[i] - '0' : 0;
            int hi = limit_high ? num2[i] - '0' : 9;

            int res = 0;
            for (int d = lo; d <= hi; d++) { // 枚举当前数位填 d
                res = (res + dfs(i + 1, sum + d, limit_low && d == lo, limit_high && d == hi)) % 1'000'000'007;
            }

            if (!limit_low && !limit_high) {
                memo[i][sum] = res;
            }
            return res;
        };

        return dfs(0, 0, true, true);
    }
};
```

```go [sol-Go]
func count(num1, num2 string, minSum, maxSum int) int {
	const mod = 1_000_000_007
	n := len(num2)
	num1 = strings.Repeat("0", n-len(num1)) + num1 // 补前导零（本题前导零不影响答案，可以补前导零，简化代码逻辑）

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, min(9*n, maxSum)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, sum int, limitLow, limitHigh bool) (res int) {
		if sum > maxSum { // 非法
			return
		}
		if i == n {
			if sum >= minSum { // 合法
				return 1
			}
			return
		}

		if !limitLow && !limitHigh {
			p := &memo[i][sum]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = int(num1[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(num2[i] - '0')
		}

		for d := lo; d <= hi; d++ { // 枚举当前数位填 d
			res = (res + dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)) % mod
		}
		return
	}
	return dfs(0, 0, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nmD)$，其中 $n$ 为 $\textit{nums}_2$ 的长度，$m=\min(9n, \textit{maxSum})$，$D=10$。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(D)$，因此时间复杂度为 $\mathcal{O}(nmD)$。
- 空间复杂度：$\mathcal{O}(nm)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
