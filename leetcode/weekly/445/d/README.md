## 方法一：数位 DP

**前置知识**：

[数位 DP v1.0 视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/)

[数位 DP v2.0 视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)（上下界数位 DP）

把 $l$ 和 $r$ 转成 $b$ 进制，然后套数位 DP v2.0 模板。

**状态定义**：$\textit{dfs}(i, \textit{pre}, \textit{limitLow},\textit{limitHigh})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{pre}$ 表示前一个位置填的数字，初始值为 $0$。
- $\textit{limitHigh}$ 表示当前是否受到了 $\textit{high}$ 的约束（我们要构造的数字不能超过 $\textit{high}$）。若为真，则第 $i$ 位填入的数字至多为 $\textit{high}[i]$，否则至多为 $b-1$，这个数记作 $\textit{hi}$。如果在受到约束的情况下填了 $\textit{hi}$，那么后续填入的数字仍会受到 $\textit{high}$ 的约束。例如 $\textit{high}=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{limitLow}$ 表示当前是否受到了 $\textit{low}$ 的约束（我们要构造的数字不能低于 $\textit{low}$）。若为真，则第 $i$ 位填入的数字至少为 $\textit{low}[i]$，否则至少为 $0$，这个数记作 $\textit{lo}$。如果在受到约束的情况下填了 $\textit{lo}$，那么后续填入的数字仍会受到 $\textit{low}$ 的约束。

**状态转移**：枚举第 $i$ 位填数字 $d=\max(\textit{lo},\textit{pre}),\textit{lo}+1,\ldots,\textit{hi}$。继续递归，把 $\textit{i}$ 加一，把 $\textit{pre}$ 置为 $d$。

**递归终点**：$i=n$ 时，找到了一个合法数字，返回 $1$。

**递归入口**：$\textit{dfs}(0, 0, \texttt{true}, \texttt{true})$，表示：

- 从最高位开始。
- 假设第一个数字的前面是 $0$，这样第一个数字不会受到 $\textit{pre}$ 的约束。
- 一开始要受到 $\textit{low}$ 和 $\textit{high}$ 的约束（否则就可以随意填了，这肯定不行）。

> 注：本题答案较小，可以只在返回前取模。（理由见方法二的公式）

[本题视频讲解](https://www.bilibili.com/video/BV1e3dBYLEDz/?t=29m32s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countNumbers(self, l: str, r: str, b: int) -> int:
        # 把 s 转成 b 进制，库函数写法见【Numpy】代码
        def trans(s: str) -> List[int]:
            x = int(s)
            digits = []
            while x:
                x, r = divmod(x, b)
                digits.append(r)
            digits.reverse()
            return digits

        high = trans(r)
        n = len(high)
        low = trans(l)
        low = [0] * (n - len(low)) + low

        @cache
        def dfs(i: int, pre: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1

            lo = low[i] if limit_low else 0
            hi = high[i] if limit_high else b - 1

            res = 0
            for d in range(max(lo, pre), hi + 1):
                res += dfs(i + 1, d, limit_low and d == lo, limit_high and d == hi)
            return res

        return dfs(0, 0, True, True) % 1_000_000_007
```

```py [sol-NumPy]
import numpy as np

class Solution:
    def countNumbers(self, l: str, r: str, b: int) -> int:
        # 把 s 转成 b 进制
        def trans(s: str) -> List[int]:
            t = np.base_repr(int(s), base=b)
            return list(map(int, t))

        high = trans(r)
        n = len(high)
        low = trans(l)
        low = [0] * (n - len(low)) + low

        @cache
        def dfs(i: int, pre: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1

            lo = low[i] if limit_low else 0
            hi = high[i] if limit_high else b - 1

            res = 0
            for d in range(max(lo, pre), hi + 1):
                res += dfs(i + 1, d, limit_low and d == lo, limit_high and d == hi)
            return res

        return dfs(0, 0, True, True) % 1_000_000_007
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    private static final int MOD = 1_000_000_007;

    public int countNumbers(String l, String r, int b) {
        char[] low = trans(l, b);
        char[] high = trans(r, b);
        int[][] memo = new int[high.length][b];
        for (int[] row : memo) {
            Arrays.fill(row, -1);
        }
        return dfs(0, 0, true, true, b, low, high, memo);
    }

    // 把十进制字符串 s 转成 b 进制字符数组
    private char[] trans(String s, int b) {
        return new BigInteger(s).toString(b).toCharArray();
    }

    private int dfs(int i, int pre, boolean limitLow, boolean limitHigh, int b, char[] low, char[] high, int[][] memo) {
        if (i == high.length) {
            return 1;
        }
        if (!limitLow && !limitHigh && memo[i][pre] >= 0) {
            return memo[i][pre];
        }

        int diffLH = high.length - low.length;
        int lo = limitLow && i >= diffLH ? low[i - diffLH] - '0' : 0;
        int hi = limitHigh ? high[i] - '0' : b - 1;

        long res = 0;
        for (int d = Math.max(lo, pre); d <= hi; d++) {
            res += dfs(i + 1, d, limitLow && d == lo, limitHigh && d == hi, b, low, high, memo);
        }
        res %= MOD;

        if (!limitLow && !limitHigh) {
            memo[i][pre] = (int) res;
        }
        return (int) res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 把十进制字符串 s 转成 b 进制
    // 用小学学过的【竖式除法】计算，读者可以先用竖式除法算算 1234÷10，再对照下面的代码
    vector<int> trans(string& s, int b) {
        for (char& c : s) {
            c -= '0';
        }
        vector<int> digits;
        while (!s.empty()) {
            string nxt_s; // 用竖式除法计算 s / b 得到的商（十进制）
            int rem = 0; // s % b
            for (char c : s) {
                rem = rem * 10 + c;
                int q = rem / b; // 商
                if (q || !nxt_s.empty()) { // 忽略前导零
                    nxt_s.push_back(q);
                }
                rem = rem % b;
            }
            digits.push_back(rem);
            s = move(nxt_s);
        }
        ranges::reverse(digits);
        return digits;
    }

public:
    int countNumbers(string l, string r, int b) {
        vector<int> low = trans(l, b);
        vector<int> high = trans(r, b);
        int n = high.size();
        int diff_lh = n - low.size();

        vector memo(n, vector<int>(b, -1));
        auto dfs = [&](this auto&& dfs, int i, int pre, bool limit_low, bool limit_high) -> int {
            if (i == n) {
                return 1;
            }
            if (!limit_low && !limit_high && memo[i][pre] >= 0) {
                return memo[i][pre];
            }

            int lo = limit_low && i >= diff_lh ? low[i - diff_lh] : 0;
            int hi = limit_high ? high[i] : b - 1;

            long long res = 0;
            for (int d = max(lo, pre); d <= hi; d++) {
                res += dfs(i + 1, d, limit_low && d == lo, limit_high && d == hi);
            }
            res %= 1'000'000'007;

            if (!limit_low && !limit_high) {
                memo[i][pre] = res;
            }
            return res;
        };
        return dfs(0, 0, true, true);
    }
};
```

```go [sol-Go]
func trans(s string, b int) string {
	x := &big.Int{}
	fmt.Fscan(strings.NewReader(s), x)
	return x.Text(b) // 转成 b 进制
}

func countNumbers(l, r string, b int) int {
	lowS := trans(l, b)
	highS := trans(r, b)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, b)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := b - 1
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		for d := max(lo, pre); d <= hi; d++ {
			res += dfs(i+1, d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return dfs(0, 0, true, true) % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2 + nb^2)$，其中 $n\approx m\log_b 10$，表示进制转换后的长度，$m$ 是 $r$ 的长度。进制转换的时间复杂度为 $\mathcal{O}(n^2)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nb)$，单个状态的计算时间为 $\mathcal{O}(b)$，所以动态规划的时间复杂度为 $\mathcal{O}(nb^2)$。
- 空间复杂度：$\mathcal{O}(nb)$。保存多少状态，就需要多少空间。

## 方法二：组合数学

假设第 $i$ 位填了数字 $j$，并且剩余的 $m=n-1-i$ 个位置不受约束，那么问题变成：

- 构造长为 $m$ 的非递减序列，元素范围 $[j,b-1]$，能构造多少个不同的序列？

根据 [3251. 单调数组对的数目 II 我的题解](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-ii/solutions/2876190/qian-zhui-he-you-hua-dppythonjavacgo-by-3biek/) 的方法二，方案数为

$$
\binom {m+b-1-j} m
$$

为方便计算，考虑用小于 $r+1$ 的合法数字个数减去小于 $l$ 的合法数字个数。

设 $r+1$ 进制转换后的字符串为 $s$。

设上一个数位填的是 $\textit{pre} = \texttt{int}(s[i-1])$，枚举当前位置填数字 $j=\textit{pre},\textit{pre}+1,\ldots, \textit{hi}-1$，其中 $\textit{hi}=\texttt{int}(s[i])$。剩余数位不受到 $s$ 的约束，方案数之和为

$$
\sum_{j=b-\textit{hi}}^{b-1-\textit{pre}} \binom {m+j} m
$$

根据上指标求和恒等式

$$
\binom {m} m + \binom {m+1} m + \cdots + \binom {m+k} m = \binom {m+k+1} {m+1}
$$

方案数之和化简为

$$
\begin{aligned}
    & \sum_{j=b-\textit{hi}}^{b-1-\textit{pre}} \binom {m+j} m      \\
={} & \sum_{j=0}^{b-1-\textit{pre}} \binom {m+j} m - \sum_{j=0}^{b-1-\textit{hi}} \binom {m+j} m       \\
={} & \binom {m+b-\textit{pre}} {m+1} - \binom {m+b-\textit{hi}} {m+1}      \\
={} & \binom {m+b-\textit{pre}} {b-1-\textit{pre}} - \binom {m+b-\textit{hi}} {b-1-\textit{hi}}      \\
\end{aligned}
$$

由于 $b-1 < 10$ 很小，可以用 [递推式](https://leetcode.cn/problems/pascals-triangle-ii/solutions/3041965/yu-chu-li-pythonjavaccgojsrust-by-endles-9wtq/) 预处理组合数，无需求阶乘及其逆元。

在本题数据范围下，组合数在 64 位整数范围内，无需中途取模，只需在返回前取模。

```py [sol-Python3]
# 关于预处理组合数的写法，见【Python3 预处理】
import numpy as np

class Solution:
    def countNumbers(self, l: str, r: str, b: int) -> int:
        # 把 s 转成 b 进制
        def trans(s: str, inc: int) -> List[int]:
            x = int(s) + inc
            t = np.base_repr(x, base=b)
            return list(map(int, t))

        def calc(s: str, inc: int) -> int:
            s = trans(s, inc)
            # 计算小于 s 的合法数字个数
            # 为什么是小于？注意下面的代码，我们没有统计每个数位都填 s[i] 的情况
            res = pre = 0
            for i, hi in enumerate(s):
                if hi < pre:
                    break
                m = len(s) - 1 - i
                res += comb(m + b - pre, b - 1 - pre) - comb(m + b - hi, b - 1 - hi)  # 不受约束的方案数
                pre = hi  # 这一位填 hi，继续计算剩余数位的方案数
            return res

        # 小于 r+1 的合法数字个数 - 小于 l 的合法数字个数
        return (calc(r, 1) - calc(l, 0)) % 1_000_000_007
```

```py [sol-Python3 预处理]
import numpy as np

MAX_N = 333  # 进制转换后的最大长度
MAX_B = 10

# 预处理组合数
C = [[0] * MAX_B for _ in range(MAX_N + MAX_B)]
for i in range(len(C)):
    C[i][0] = 1
    for j in range(1, min(i + 1, MAX_B)):
        # 注意本题组合数较小，无需取模
        C[i][j] = C[i - 1][j - 1] + C[i - 1][j]

class Solution:
    def countNumbers(self, l: str, r: str, b: int) -> int:
        # 把 s 转成 b 进制
        def trans(s: str, inc: int) -> List[int]:
            x = int(s) + inc
            t = np.base_repr(x, base=b)
            return list(map(int, t))

        def calc(s: str, inc: int) -> int:
            s = trans(s, inc)
            # 计算小于 s 的合法数字个数
            # 为什么是小于？注意下面的代码，我们没有统计每个数位都填 s[i] 的情况
            res = pre = 0
            for i, hi in enumerate(s):
                if hi < pre:
                    break
                m = len(s) - 1 - i
                res += C[m + b - pre][b - 1 - pre] - C[m + b - hi][b - 1 - hi]  # 不受约束的方案数
                pre = hi  # 这一位填 hi，继续计算剩余数位的方案数
            return res

        # 小于 r+1 的合法数字个数 - 小于 l 的合法数字个数
        return (calc(r, 1) - calc(l, 0)) % 1_000_000_007
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MAX_N = 333; // 进制转换后的最大长度
    private static final int MAX_B = 10;
    private static final long[][] comb = new long[MAX_N + MAX_B][MAX_B];
    private static boolean done = false;

    // 这样写比 static block 更快
    private void init() {
        if (done) {
            return;
        }
        done = true;
        // 预处理组合数
        for (int i = 0; i < MAX_N + MAX_B; i++) {
            comb[i][0] = 1;
            for (int j = 1; j < Math.min(i + 1, MAX_B); j++) {
                // 注意本题组合数较小，无需取模
                comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
            }
        }
    }

    public int countNumbers(String l, String r, int b) {
        init();
        // 小于 r+1 的合法数字个数 - 小于 l 的合法数字个数
        return (int) ((calc(r, b, true) - calc(l, b, false)) % MOD);
    }

    // 把十进制字符串 s 转成 b 进制字符数组
    private char[] trans(String s, int b, boolean inc) {
        BigInteger x = new BigInteger(s);
        if (inc) {
            x = x.add(BigInteger.ONE);
        }
        return x.toString(b).toCharArray();
    }

    private long calc(String S, int b, boolean inc) {
        char[] s = trans(S, b, inc);
        // 计算小于 s 的合法数字个数
        // 为什么是小于？注意下面的代码，我们没有统计每个数位都填 s[i] 的情况
        long res = 0;
        int pre = 0;
        for (int i = 0; i < s.length && s[i] - '0' >= pre; i++) {
            int hi = s[i] - '0';
            int m = s.length - 1 - i;
            res += comb[m + b - pre][b - 1 - pre] - comb[m + b - hi][b - 1 - hi]; // 不受约束的方案数
            pre = hi; // 这一位填 hi，继续计算剩余数位的方案数
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MAX_N = 333; // 进制转换后的最大长度
const int MAX_B = 10;
long long comb[MAX_N + MAX_B][MAX_B];

int init = [] {
    // 预处理组合数
    for (int i = 0; i < MAX_N + MAX_B; i++) {
        comb[i][0] = 1;
        for (int j = 1; j < min(i + 1, MAX_B); j++) {
            // 注意本题组合数较小，无需取模
            comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
        }
    }
    return 0;
}();

class Solution {
    // 把十进制字符串 s 转成 b 进制
    // 用小学学过的【竖式除法】计算，读者可以先用竖式除法算算 1234÷10，再对照下面的代码
    vector<int> trans(string& s, int b) {
        for (char& c : s) {
            c -= '0';
        }
        vector<int> digits;
        while (!s.empty()) {
            string nxt_s; // 用竖式除法计算 s / b 得到的商（十进制）
            int rem = 0; // s % b
            for (char c : s) {
                rem = rem * 10 + c;
                int q = rem / b; // 商
                if (q || !nxt_s.empty()) { // 忽略前导零
                    nxt_s.push_back(q);
                }
                rem = rem % b;
            }
            digits.push_back(rem);
            s = move(nxt_s);
        }
        ranges::reverse(digits);
        return digits;
    }

    long long calc(string& s, int b, bool check_s) {
        vector<int> digits = trans(s, b);
        // 计算小于 s 的合法数字个数
	    // 为什么是小于？注意下面的代码，我们没有统计每个数位都填 digits[i] 的情况
        long long res = 0;
        int pre = 0;
        for (int i = 0; i < digits.size() && digits[i] >= pre; i++) {
            int hi = digits[i];
            int m = digits.size() - 1 - i;
            res += comb[m + b - pre][b - 1 - pre] - comb[m + b - hi][b - 1 - hi]; // 不受约束的方案数
            pre = hi; // 这一位填 hi，继续计算剩余数位的方案数
        }
        return res + (check_s && is_non_dec(digits)); // 单独判断 digits 是否合法
    }

    bool is_non_dec(vector<int>& digits) {
        for (int i = 1; i < digits.size(); i++) {
            if (digits[i - 1] > digits[i]) {
                return false;
            }
        }
        return true;
    }

public:
    int countNumbers(string l, string r, int b) {
        // 小于 r+1 的合法数字个数 - 小于 l 的合法数字个数
        return (calc(r, b, true) - calc(l, b, false)) % 1'000'000'007;
    }
};
```

```go [sol-Go]
const maxN = 333 // 进制转换后的最大长度
const maxB = 10

var comb [maxN + maxB][maxB]int

func init() {
	// 预处理组合数
	for i := 0; i < len(comb); i++ {
		comb[i][0] = 1
		for j := 1; j < min(i+1, maxB); j++ {
			// 注意本题组合数较小，无需取模
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}
}

func trans(s string, b int, inc bool) string {
	x := &big.Int{}
	fmt.Fscan(strings.NewReader(s), x)
	if inc {
		x.Add(x, big.NewInt(1))
	}
	return x.Text(b) // 转成 b 进制
}

func calc(s string, b int, inc bool) (res int) {
	s = trans(s, b, inc)
	// 计算小于 s 的合法数字个数
	// 为什么是小于？注意下面的代码，我们没有统计每个数位都填 s[i] 的情况
	pre := 0
	for i, d := range s {
		hi := int(d - '0')
		if hi < pre {
			break
		}
		m := len(s) - 1 - i
		res += comb[m+b-pre][b-1-pre] - comb[m+b-hi][b-1-hi] // 不受约束的方案数
		pre = hi // 这一位填 hi，继续计算剩余数位的方案数
	}
	return
}

func countNumbers(l, r string, b int) int {
	// 小于 r+1 的合法数字个数 - 小于 l 的合法数字个数
	return (calc(r, b, true) - calc(l, b, false)) % 1_000_000_007
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n\approx m\log_b 10$，表示进制转换后的长度，$m$ 是 $r$ 的长度。瓶颈在进制转换上。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单的「**十、数位 DP**」和数学题单的「**§2.2 组合计数**」。

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
