## 方法一：数位 DP

**前置知识**：

[数位 DP v1.0 视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/)

[数位 DP v2.0 视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)（上下界数位 DP）

**状态定义**：$\textit{dfs}(i, \textit{pre}, \textit{limitLow},\textit{limitHigh})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{pre}$ 表示前一个**在路径中的**数字，初始值为 $0$。
- $\textit{limitHigh}$ 表示当前是否受到了 $\textit{high}$ 的约束（当前要填的数字不能超过 $\textit{high}$）。若为真，则第 $i$ 位填入的数字至多为 $\textit{high}[i]$，否则至多为 $9$，这个数记作 $\textit{hi}$。如果在受到约束的情况下填了 $\textit{hi}$，那么后续填入的数字仍会受到 $\textit{high}$ 的约束。例如 $\textit{high}=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{limitLow}$ 表示当前是否受到了 $\textit{low}$ 的约束（当前要填的数字不能低于 $\textit{low}$）。若为真，则第 $i$ 位填入的数字至少为 $\textit{low}[i]$，否则至少为 $0$，这个数记作 $\textit{lo}$。如果在受到约束的情况下填了 $\textit{lo}$，那么后续填入的数字仍会受到 $\textit{low}$ 的约束。

**状态转移**：枚举第 $i$ 位填数字 $d=\textit{lo},\textit{lo}+1\ldots,\textit{hi}$。继续递归，把 $\textit{i}$ 加一，如果 $i$ 在路径中，那么把 $\textit{pre}$ 置为 $d$，否则 $\textit{pre}$ 不变。特别地，如果 $i$ 在路径中，那么 $d$ 从 $\max(\textit{lo},\textit{pre})$ 开始枚举，从而保证路径上的数字是非递减的。

**递归终点**：$i=n$ 时，找到了一个合法数字，返回 $1$。

**递归入口**：$\textit{dfs}(0, 0, \texttt{true}, \texttt{true})$，参数分别表示：

- 从最高位开始。
- 假设第一个数字的前面是 $0$，这样第一个数字不会受到 $\textit{pre}$ 的约束。
- 一开始要受到 $\textit{low}$ 和 $\textit{high}$ 的约束（否则就可以随意填了，这肯定不行）。

为了方便判断 $i$ 是否在路径中，我们可以预处理一个长为 $n$ 的布尔数组，表示每个位置是否在路径中。

[本题视频讲解](https://www.bilibili.com/video/BV1agddBJEnQ/?t=15m52s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countGoodIntegersOnPath(self, l: int, r: int, directions: str) -> int:
        high_s = list(map(int, str(r)))  # 避免在 dfs 中频繁调用 int()
        n = len(high_s)
        low_s = list(map(int, str(l).zfill(n)))  # 补前导零，让 len(low_s) == n

        in_path = [False] * n
        in_path[-1] = True  # 终点一定在路径中
        pos = n - 1
        for d in reversed(directions):
            pos -= 1 if d == 'R' else 4  # 倒着往上相当于往左数 4 个位置
            if pos < 0:  # 只需要对网格图的后 n 个格子做标记
                break
            in_path[pos] = True  # 标记在路径中的格子

        @cache
        def dfs(i: int, pre: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:  # 成功到达终点
                return 1  # 找到了一个好整数

            lo = low_s[i] if limit_low else 0
            hi = high_s[i] if limit_high else 9

            res = 0
            # 如果当前位置在路径中，那么当前位置填的数必须 >= pre
            start = max(lo, pre) if in_path[i] else lo
            for d in range(start, hi + 1):
                res += dfs(i + 1,
                           d if in_path[i] else pre,
                           limit_low and d == lo,
                           limit_high and d == hi)
            return res

        return dfs(0, 0, True, True)
```

```java [sol-Java]
class Solution {
    public long countGoodIntegersOnPath(long l, long r, String directions) {
        char[] lowS = String.valueOf(l).toCharArray();
        char[] highS = String.valueOf(r).toCharArray();
        int n = highS.length;

        boolean[] inPath = new boolean[n];
        int pos = n - 16; // 右下角是下标 n-1，那么左上角是下标 n-16
        for (char d : directions.toCharArray()) {
            if (pos >= 0) { // 只需要对网格图的后 n 个格子做标记
                inPath[pos] = true; // 标记在路径中的格子
            }
            pos += d == 'R' ? 1 : 4; // 往下相当于往右数 4 个位置
        }
        inPath[n - 1] = true; // 终点一定在路径中

        long[][] memo = new long[n][10];
        for (long[] row : memo) {
            Arrays.fill(row, -1);
        }

        return dfs(0, 0, true, true, lowS, highS, inPath, memo);
    }

    private long dfs(int i, int pre, boolean limitLow, boolean limitHigh, char[] lowS, char[] highS, boolean[] inPath, long[][] memo) {
        if (i == highS.length) { // 成功到达终点
            return 1; // 找到了一个好整数
        }

        if (!limitLow && !limitHigh && memo[i][pre] >= 0) {
            return memo[i][pre];
        }

        int diff = highS.length - lowS.length;
        int lo = limitLow && i >= diff ? lowS[i - diff] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long res = 0;
        // 如果当前位置在路径中，那么当前位置填的数必须 >= pre
        int start = inPath[i] ? Math.max(lo, pre) : lo;
        for (int d = start; d <= hi; d++) {
            res += dfs(i + 1,
                    inPath[i] ? d : pre,
                    limitLow && d == lo,
                    limitHigh && d == hi,
                    lowS, highS, inPath, memo);
        }

        if (!limitLow && !limitHigh) {
            memo[i][pre] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countGoodIntegersOnPath(long long l, long long r, string directions) {
        string low_s = to_string(l);
        string high_s = to_string(r);
        int n = high_s.size();
        int diff_lh = n - low_s.size();
        vector<array<long long, 10>> memo(n);

        vector<int8_t> in_path(n);
        int pos = n - 16; // 右下角是下标 n-1，那么左上角是下标 n-16
        for (char d : directions) {
            if (pos >= 0) { // 只需要对网格图的后 n 个格子做标记
                in_path[pos] = true; // 标记在路径中的格子
            }
            pos += d == 'R' ? 1 : 4; // 往下相当于往右数 4 个位置
        }
        in_path[n - 1] = true; // 终点一定在路径中

        auto dfs = [&](this auto&& dfs, int i, int pre, bool limit_low, bool limit_high) -> long long {
            if (i == n) { // 成功到达终点
                return 1; // 找到了一个好整数
            }

            if (!limit_low && !limit_high && memo[i][pre] > 0) {
                return memo[i][pre] - 1; // 见下面注释
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long res = 0;
            // 如果当前位置在路径中，那么当前位置填的数必须 >= pre
            int start = in_path[i] ? max(lo, pre) : lo;
            for (int d = start; d <= hi; d++) {
                res += dfs(i + 1, in_path[i] ? d : pre, limit_low && d == lo, limit_high && d == hi);
            }

            if (!limit_low && !limit_high) {
                memo[i][pre] = res + 1; // 记忆化的时候多加一，这样 memo[i][j] 就不需要初始化成 -1 了
            }
            return res;
        };

        return dfs(0, 0, true, true);
    }
};
```

```go [sol-Go]
func countGoodIntegersOnPath(l, r int64, directions string) int64 {
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)

	inPath := make([]bool, n)
	pos := n - 16 // 右下角是下标 n-1，那么左上角是下标 n-16
	for _, d := range directions {
		if pos >= 0 { // 只需要对网格图的后 n 个格子做标记
			inPath[pos] = true // 标记在路径中的格子
		}
		if d == 'R' { // 往右
			pos++
		} else { // 往下
			pos += 4 // 相当于往右数 4 个位置
		}
	}
	inPath[n-1] = true // 终点一定在路径中

	diffLH := n - len(lowS)
	memo := make([][10]int64, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int64
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int64) {
		if i == n { // 成功到达终点
			return 1 // 找到了一个好整数
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
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		if inPath[i] { // 当前位置在路径中
			d = max(d, pre) // 当前位置填的数必须 >= pre
		}
		for ; d <= hi; d++ {
			p := pre
			if inPath[i] {
				p = d
			}
			res += dfs(i+1, p, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	return dfs(0, 0, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^2\log r)$，其中 $\mathcal{O}(\log r)$ 是 $r$ 的十进制长度，$D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(D\log r)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^2\log r)$。
- 空间复杂度：$\mathcal{O}(D\log r)$。保存多少状态，就需要多少空间。

## 方法二：组合数学

**前置题目**：[3519. 统计逐位非递减的整数](https://leetcode.cn/problems/count-numbers-with-non-decreasing-digits/)，[我的题解](https://leetcode.cn/problems/count-numbers-with-non-decreasing-digits/solutions/3649556/mo-ban-shang-xia-jie-shu-wei-dp-by-endle-rhuw/)。

为方便计算，用小于 $r+1$ 的合法数字个数减去小于 $l$ 的合法数字个数。

设 $r+1$ 的十进制字符串为 $s$，设 $\textit{hi}=\texttt{int}(s[i])$，设上一个在路径中的数为 $\textit{pre}$。

分类讨论：

- 如果 $s[i]$ 不在路径上，那么当前位置（在不受约束的情况下）可以填 $0,1,\ldots,\textit{hi}-1$，一共 $\textit{hi}$ 种数字。后面有 $n-1-i$ 个位置，设其中有 $m$ 个位置在路径上，那么：
  - 对于在路径上的位置，根据 3519 题解的公式，有 $\dbinom {m+9-\textit{pre}} {m}$ 种填法。
  - 对于不在路径上的位置，每个位置 $0$ 到 $9$ 随便填，有 $10^{n-1-i-m}$ 种填法。
  - 一共有 $\textit{hi}\cdot \dbinom {m+9-\textit{pre}} {m}\cdot 10^{n-1-i-m}$ 种填法。
- 如果 $s[i]$ 在路径上，那么：
  - 对于在路径上的位置（包括 $i$），根据 3519 题解的公式，有 $\dbinom {m+10-\textit{pre}} {m+1} - \dbinom {m+10-\textit{hi}} {m+1}$ 种填法。
  - 对于不在路径上的位置，每个位置 $0$ 到 $9$ 随便填，有 $10^{n-1-i-m}$ 种填法。
  - 一共有 $\left(\dbinom {m+10-\textit{pre}} {m+1} - \dbinom {m+10-\textit{hi}} {m+1}\right)\cdot 10^{n-1-i-m}$ 种填法。

```py [sol-Python3]
class Solution:
    def countGoodIntegersOnPath(self, l: int, r: int, directions: str) -> int:
        r = str(r + 1)  # 注意这里加一了
        n = len(r)
        l = str(l).zfill(n)

        in_path = [False] * n
        in_path[-1] = True  # 终点一定在路径中
        pos = n - 1
        for d in reversed(directions):
            pos -= 1 if d == 'R' else 4  # 倒着往上相当于往左数 4 个位置
            if pos < 0:  # 只需要对网格图的后 n 个格子做标记
                break
            in_path[pos] = True  # 标记在路径中的格子

        # suf[i] 表示后缀 [i, n-1] 在路径中的下标个数
        suf = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            suf[i] = suf[i + 1] + in_path[i]

        # 计算小于 r 的合法整数个数
        def solve(r: str) -> int:
            res = pre = 0
            for i, ch in enumerate(r):
                hi = int(ch)
                m = suf[i + 1]
                if not in_path[i]:
                    res += hi * comb(m + 9 - pre, m) * 10 ** (n - 1 - i - m)
                    continue
                if hi < pre:
                    break
                res += (comb(m + 10 - pre, m + 1) - comb(m + 10 - hi, m + 1)) * 10 ** (n - 1 - i - m)
                pre = hi  # 这一位填 hi，继续计算剩余数位的方案数
            return res

        return solve(r) - solve(l)
```

```java [sol-Java]
class Solution {
    private static final int MAX_M = 7;
    private static final long[][] comb = new long[MAX_M + 10][MAX_M + 1];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理组合数
        for (int i = 0; i < comb.length; i++) {
            comb[i][0] = 1;
            for (int j = 1; j < Math.min(i + 1, comb[i].length); j++) {
                comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
            }
        }
    }

    public long countGoodIntegersOnPath(long l, long r, String directions) {
        char[] highS = String.valueOf(r + 1).toCharArray(); // 注意这里加一了
        int n = highS.length;
        String tmp = String.valueOf(l);
        char[] lowS = ("0".repeat(n - tmp.length()) + tmp).toCharArray();

        boolean[] inPath = new boolean[n];
        int pos = n - 16; // 右下角是下标 n-1，那么左上角是下标 n-16
        for (char d : directions.toCharArray()) {
            if (pos >= 0) { // 只需要对网格图中的后 n 个格子做标记
                inPath[pos] = true; // 标记在路径中的格子
            }
            pos += d == 'R' ? 1 : 4; // 往下相当于往右数 4 个位置
        }
        inPath[n - 1] = true; // 终点一定在路径中

        // suf[i] 表示后缀 [i, n-1] 在路径中的下标个数
        int[] suf = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            suf[i] = suf[i + 1] + (inPath[i] ? 1 : 0);
        }

        long[] pow10 = new long[n + 1];
        pow10[0] = 1;
        for (int i = 1; i <= n; i++) {
            pow10[i] = pow10[i - 1] * 10;
        }

        return solve(highS, inPath, suf, pow10) - solve(lowS, inPath, suf, pow10);
    }

    private long solve(char[] s, boolean[] inPath, int[] suf, long[] pow10) {
        int n = s.length;
        long res = 0;
        int pre = 0;
        for (int i = 0; i < n; i++) {
            int hi = s[i] - '0';
            int m = suf[i + 1];
            if (!inPath[i]) {
                res += hi * comb[m + 9 - pre][m] * pow10[n - 1 - i - m];
                continue;
            }
            if (hi < pre) {
                break;
            }
            res += (comb[m + 10 - pre][m + 1] - comb[m + 10 - hi][m + 1]) * pow10[n - 1 - i - m];
            pre = hi; // 这一位填 hi，继续计算剩余数位的方案数
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MAX_M = 7;
long long comb[MAX_M + 10][MAX_M + 1];

int init = [] {
    // 预处理组合数
    for (int i = 0; i < MAX_M + 10; i++) {
        comb[i][0] = 1;
        for (int j = 1; j <= min(i, MAX_M); j++) {
            comb[i][j] = comb[i - 1][j - 1] + comb[i - 1][j];
        }
    }
    return 0;
}();

class Solution {
public:
    long long countGoodIntegersOnPath(long long l, long long r, string directions) {
        string high_s = to_string(r + 1); // 注意这里加一了
        int n = high_s.size();
        string low_s = to_string(l);
        low_s = string(n - low_s.size(), '0') + low_s;

        vector<int8_t> in_path(n);
        int pos = n - 16; // 右下角是下标 n-1，那么左上角是下标 n-16
        for (char d : directions) {
            if (pos >= 0) { // 只需要对网格图中的后 n 个格子做标记
                in_path[pos] = true; // 标记在路径中的格子
            }
            pos += d == 'R' ? 1 : 4; // 往下相当于往右数 4 个位置
        }
        in_path[n - 1] = true; // 终点一定在路径中

        // suf[i] 表示后缀 [i, n-1] 在路径中的下标个数
        vector<int> suf(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            suf[i] = suf[i + 1] + in_path[i];
        }

        vector<long long> pow10(n + 1);
        pow10[0] = 1;
        for (int i = 1; i <= n; i++) {
            pow10[i] = pow10[i - 1] * 10;
        }

        auto solve = [&](string& s) -> long long {
            long long res = 0;
            int pre = 0;
            for (int i = 0; i < n; i++) {
                int hi = s[i] - '0';
                int m = suf[i + 1];
                if (!in_path[i]) {
                    res += hi * comb[m + 9 - pre][m] * pow10[n - 1 - i - m];
                    continue;
                }
                if (hi < pre) {
                    break;
                }
                res += (comb[m + 10 - pre][m + 1] - comb[m + 10 - hi][m + 1]) * pow10[n - 1 - i - m];
                pre = hi; // 这一位填 hi，继续计算剩余数位的方案数
            }
            return res;
        };

        return solve(high_s) - solve(low_s);
    }
};
```

```go [sol-Go]
const maxM = 7

var comb [maxM + 10][maxM + 1]int

func init() {
	// 预处理组合数
	for i := range comb {
		comb[i][0] = 1
		for j := 1; j < min(i+1, len(comb[i])); j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}
}

func countGoodIntegersOnPath(l, r int64, directions string) int64 {
	highS := strconv.FormatInt(r+1, 10) // 注意这里加一了
	n := len(highS)
	lowS := strconv.FormatInt(l, 10)
	lowS = strings.Repeat("0", n-len(lowS)) + lowS

	inPath := make([]bool, n)
	pos := n - 16 // 右下角是下标 n-1，那么左上角是下标 n-16
	for _, d := range directions {
		if pos >= 0 { // 只需要对网格图中的后 n 个格子做标记
			inPath[pos] = true // 标记在路径中的格子
		}
		if d == 'R' { // 往右
			pos++
		} else { // 往下
			pos += 4 // 相当于往右数 4 个位置
		}
	}
	inPath[n-1] = true // 终点一定在路径中

	// suf[i] 表示后缀 [i, n-1] 在路径中的下标个数
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1]
		if inPath[i] {
			suf[i]++
		}
	}

	// 计算小于 r 的合法整数个数
	solve := func(r string) (res int) {
		pre := 0
		for i, ch := range r {
			hi := int(ch - '0')
			m := suf[i+1]
			if !inPath[i] {
				res += hi * comb[m+9-pre][m] * int(math.Pow10(n-1-i-m))
				continue
			}
			if hi < pre {
				break
			}
			res += (comb[m+10-pre][m+1] - comb[m+10-hi][m+1]) * int(math.Pow10(n-1-i-m))
			pre = hi // 这一位填 hi，继续计算剩余数位的方案数
		}
		return res
	}

	return int64(solve(highS) - solve(lowS))
}
```

#### 复杂度分析

不计入预处理的时间和空间。（Python3 预处理组合数的代码见 3519 题）

- 时间复杂度：$\mathcal{O}(\log r)$，其中 $\mathcal{O}(\log r)$ 是 $r$ 的十进制长度。
- 空间复杂度：$\mathcal{O}(\log r)$。

## 专题训练

见下面动态规划题单的「**十、数位 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
