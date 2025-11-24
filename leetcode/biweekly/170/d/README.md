## 前置知识

[数位 DP v1.0 模板讲解](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s)

[数位 DP v2.0 模板讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s)（上下界数位 DP）

## 写法一：波动值作为参数

在递归的过程中，我们需要知道：

- 当前在填从高到低第 $i$ 个数位（$i$ 从 $0$ 开始）。
- 前面产生的波动值（峰和谷的个数）为 $\textit{waviness}$。
- 为了判断 $i-2,i-1,i$ 三个位置是否形成峰或谷，我们需要：
    - 记录 $i-1,i-2$ 所填数字的大小关系 $\textit{lastCmp}$（用 $-1,0,1$ 分别表示小于、等于和大于）。
    - 记录 $i-1$ 所填数字 $\textit{lastDigit}$。
- 上下界约束参数 $\textit{limitLow}$ 和 $\textit{limitHigh}$，见数位 DP v2.0 模板讲解。

于是，定义 $\textit{dfs}(i,\textit{waviness},\textit{lastCmp},\textit{lastDigit},\textit{limitLow},\textit{limitHigh})$ 表示在上述情况下能得到的波动值之和。

枚举当前填的数字为 $d$，分类讨论：

- 如果当前我们填是最高有效位（或者说，之前没有填过非零数字），那么不产生峰谷，继续递归，$\textit{lastCmp}=0$，$\textit{lastDigit}=d$。累加返回值。
- 如果之前填过非零数字，那么设 $c$ 为 $d$ 和 $\textit{lastDigit}$ 的大小关系，如果 $c\ne 0$ 且 $c = -\textit{lastCmp}$（或者 $c\cdot \textit{lastCmp} = -1$），那么形成了一个峰或谷，把波动值 $\textit{waviness}$ 加一，继续递归，$\textit{lastCmp}=c$，$\textit{lastDigit}=d$。累加返回值。
- 注意上面的逻辑兼容前导零，所以无需单独处理前导零。

递归边界：如果 $i=n$，那么成功构造出一个数字，这个数字的波动值（峰谷个数）为 $\textit{waviness}$，返回 $\textit{waviness}$。

递归入口：$\textit{dfs}(0,0,0,0,\texttt{true},\texttt{true})$。

⚠**小技巧**：有没有觉得初始化 $\textit{memo}$ 数组为 $-1$ 比较麻烦？我们可以把要记忆化的值加一，并在取值时减一，就可以把 $\textit{memo}$ 数组初始化成 $0$ 了。

[本题视频讲解](https://www.bilibili.com/video/BV1fbUKBqEa7/?t=17m20s)，欢迎点赞关注~

下面是数位 DP v2.1 模板。相比 v2.0，不需要写 $\textit{isNum}$ 参数。

```py [sol-Python3]
class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        low_s = list(map(int, str(num1)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(num2)))
        n = len(high_s)
        diff_lh = n - len(low_s)

        @cache
        def dfs(i: int, waviness: int, last_cmp: int, last_digit: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return waviness

            lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high_s[i] if limit_high else 9

            res = 0
            is_num = not limit_low or i > diff_lh  # 前面是否填过数字
            for d in range(lo, hi + 1):
                # 当前填的数不是最高位，c 才有意义
                c = (d > last_digit) - (d < last_digit) if is_num else 0
                w = waviness
                if c * last_cmp < 0:  # 形成了一个峰或谷
                    w += 1
                res += dfs(i + 1, w, c, d, limit_low and d == lo, limit_high and d == hi)
            return res

        return dfs(0, 0, 0, 0, True, True)
```

```java [sol-Java]
class Solution {
    public long totalWaviness(long num1, long num2) {
        char[] lowS = Long.toString(num1).toCharArray();
        char[] highS = Long.toString(num2).toCharArray();
        int n = highS.length;
        long[][][][] memo = new long[n][n - 1][3][10]; // 一个数至多包含 n-2 个峰或谷
        return dfs(0, 0, 0, 0, true, true, lowS, highS, memo);
    }

    private long dfs(int i, int waviness, int lastCmp, int lastDigit, boolean limitLow, boolean limitHigh, char[] lowS, char[] highS, long[][][][] memo) {
        if (i == highS.length) {
            return waviness;
        }
        if (!limitLow && !limitHigh && memo[i][waviness][lastCmp + 1][lastDigit] > 0) {
            return memo[i][waviness][lastCmp + 1][lastDigit] - 1;
        }

        int diffLh = highS.length - lowS.length;
        int lo = limitLow && i >= diffLh ? lowS[i - diffLh] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long res = 0;
        boolean isNum = !limitLow || i > diffLh; // 前面是否填过数字
        for (int d = lo; d <= hi; d++) {
            // 当前填的数不是最高位，cmp 才有意义
            int cmp = isNum ? Integer.compare(d, lastDigit) : 0;
            int w = waviness + (cmp * lastCmp < 0 ? 1 : 0);
            res += dfs(i + 1, w, cmp, d, limitLow && d == lo, limitHigh && d == hi, lowS, highS, memo);
        }

        if (!limitLow && !limitHigh) {
            memo[i][waviness][lastCmp + 1][lastDigit] = res + 1;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long totalWaviness(long long num1, long long num2) {
        string low_s = to_string(num1);
        string high_s = to_string(num2);
        int n = high_s.size();
        int diff_lh = n - low_s.size();
        vector memo(n, vector<array<array<long long, 10>, 3>>(n));

        auto dfs = [&](this auto&& dfs, int i, int waviness, int last_cmp, int last_digit, bool limit_low, bool limit_high) -> long long {
            if (i == n) {
                return waviness;
            }
            long long& ref = memo[i][waviness][last_cmp + 1][last_digit];
            if (!limit_low && !limit_high && ref) {
                return ref - 1;
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long res = 0;
            bool is_num = !limit_low || i > diff_lh; // 前面是否填过数字
            for (int d = lo; d <= hi; d++) {
                // 当前填的数不是最高位，cmp 才有意义
                int cmp = is_num ? (d > last_digit) - (d < last_digit) : 0;
                int w = waviness + (cmp * last_cmp < 0);
                res += dfs(i + 1, w, cmp, d, limit_low && d == lo, limit_high && d == hi);
            }

            if (!limit_low && !limit_high) {
                ref = res + 1;
            }
            return res;
        };

        return dfs(0, 0, 0, 0, true, true);
    }
};
```

```go [sol-Go]
func totalWaviness(num1, num2 int64) int64 {
	lowS := strconv.FormatInt(num1, 10)
	highS := strconv.FormatInt(num2, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][][3][10]int, n)
	for i := range memo {
		memo[i] = make([][3][10]int, n-1) // 一个数至多包含 n-2 个峰或谷
	}

	var dfs func(int, int, int, int, bool, bool) int
	dfs = func(i, waviness, lastCmp, lastDigit int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return waviness
		}
		if !limitLow && !limitHigh {
			p := &memo[i][waviness][lastCmp+1][lastDigit]
			if *p > 0 {
				return *p - 1
			}
			defer func() { *p = res + 1 }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		isNum := !limitLow || i > diffLH // 前面是否填过数字
		for d := lo; d <= hi; d++ {
			c := 0
			if isNum { // 当前填的数不是最高位
				c = cmp.Compare(d, lastDigit)
			}
			w := waviness
			if c*lastCmp < 0 { // 形成了一个峰或谷
				w++
			}
			res += dfs(i+1, w, c, d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return int64(dfs(0, 0, 0, 0, true, true))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^2n^2)$，其中 $n = \mathcal{O}(\log \textit{num}_2)$ 是 $\textit{num}_2$ 的十进制长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(Dn^2)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^2n^2)$。
- 空间复杂度：$\mathcal{O}(Dn^2)$。

## 写法二：贡献法

去掉 $\textit{dfs}$ 中的 $\textit{waviness}$ 参数，写在返回值中。也就是**自底向上**计算所有 $\textit{waviness}$ 的总和。

回想一下如何统计一棵二叉树的节点个数：

- 递归左子树，拿到左子树的节点个数 $\textit{leftSize}$。
- 递归右子树，拿到右子树的节点个数 $\textit{rightSize}$。
- 加上当前节点（$1$ 个），得到当前子树的节点个数 $\textit{leftSize} + \textit{rightSize} + 1$。

本题如果只考虑每个数位怎么填，可以视作一棵十叉树。

统计波动值的总和：

- 递归子树，累加子树中产生的波动值。
- 对于当前数位 $i$，如果 $i-2,i-1,i$ 产生了一个峰谷，那么这个峰谷会对后面所有剩余数位的填法都**贡献**一次。比如剩余数位有 $42$ 种填法，那么这 $42$ 个数都在 $i-2,i-1,i$ 处产生了一个峰谷，所以波动值增加 $42$，累加到 $\textit{dfs}$ 的返回值中。

从上面的讨论中，我们还需要知道剩余数位有多少种填法，**每种填法都对应着十叉树的一个叶子**。虽然本题数位可以随便填没有特殊的约束，但为了代码的可扩展性，我把剩余数位的合法填法也作为返回值。

也就是说，$\textit{dfs}$ 返回两个数，第一个数是波动值，第二个数是剩余数位的合法填法方案数。这个方案数的计算方式与统计十叉树的**叶子**个数的计算方式是一样的。

```py [sol-Python3]
class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        low_s = list(map(int, str(num1)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(num2)))
        n = len(high_s)
        diff_lh = n - len(low_s)

        # dfs 返回两个数：子树波动值总和，子树合法数字个数
        @cache
        def dfs(i: int, last_cmp: int, last_digit: int, limit_low: bool, limit_high: bool) -> Tuple[int, int]:
            if i == n:
                return 0, 1  # 本题无特殊约束，能递归到终点的都是合法数字

            lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high_s[i] if limit_high else 9

            waviness_sum = num_cnt = 0
            is_num = not limit_low or i > diff_lh  # 前面是否填过数字
            for d in range(lo, hi + 1):
                # 当前填的数不是最高位，c 才有意义
                c = (d > last_digit) - (d < last_digit) if is_num else 0
                sub_waviness_sum, sub_num_cnt = dfs(i + 1, c, d, limit_low and d == lo, limit_high and d == hi)
                waviness_sum += sub_waviness_sum  # 累加子树的波动值
                num_cnt += sub_num_cnt  # 累加子树的合法数字个数
                if c * last_cmp < 0:  # 形成了一个峰或谷
                    waviness_sum += sub_num_cnt  # 这个峰谷会出现在 sub_num_cnt 个数字中
            return waviness_sum, num_cnt

        return dfs(0, 0, 0, True, True)[0]
```

```java [sol-Java]
class Solution {
    public long totalWaviness(long num1, long num2) {
        char[] lowS = Long.toString(num1).toCharArray();
        char[] highS = Long.toString(num2).toCharArray();
        int n = highS.length;
        long[][][][] memo = new long[n][3][10][];
        return dfs(0, 0, 0, true, true, lowS, highS, memo)[0];
    }

    // dfs 返回两个数：子树波动值总和，子树合法数字个数
    private long[] dfs(int i, int lastCmp, int lastDigit, boolean limitLow, boolean limitHigh, char[] lowS, char[] highS, long[][][][] memo) {
        if (i == highS.length) {
            return new long[]{0, 1}; // 本题无特殊约束，能递归到终点的都是合法数字
        }
        if (!limitLow && !limitHigh && memo[i][lastCmp + 1][lastDigit] != null) {
            return memo[i][lastCmp + 1][lastDigit];
        }

        int diffLh = highS.length - lowS.length;
        int lo = limitLow && i >= diffLh ? lowS[i - diffLh] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long wavinessSum = 0;
        long numCnt = 0;
        boolean isNum = !limitLow || i > diffLh; // 前面是否填过数字
        for (int d = lo; d <= hi; d++) {
            // 当前填的数不是最高位，cmp 才有意义
            int cmp = isNum ? Integer.compare(d, lastDigit) : 0;
            long[] sub = dfs(i + 1, cmp, d, limitLow && d == lo, limitHigh && d == hi, lowS, highS, memo);
            wavinessSum += sub[0]; // 累加子树的波动值
            numCnt += sub[1]; // 累加子树的合法数字个数
            if (cmp * lastCmp < 0) { // 形成了一个峰或谷
                wavinessSum += sub[1]; // 这个峰谷会出现在 sub[1] 个数字中
            }
        }

        long[] res = new long[]{wavinessSum, numCnt};
        if (!limitLow && !limitHigh) {
            memo[i][lastCmp + 1][lastDigit] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long totalWaviness(long long num1, long long num2) {
        string low_s = to_string(num1);
        string high_s = to_string(num2);
        int n = high_s.size();
        int diff_lh = n - low_s.size();
        vector<array<array<pair<long long, long long>, 10>, 3>> memo(n);

        // dfs 返回两个数：子树波动值总和，子树合法数字个数
        auto dfs = [&](this auto&& dfs, int i, int last_cmp, int last_digit, bool limit_low, bool limit_high) -> pair<long long, long long> {
            if (i == n) {
                return {0, 1}; // 本题无特殊约束，能递归到终点的都是合法数字
            }
            auto& ref = memo[i][last_cmp + 1][last_digit];
            if (!limit_low && !limit_high && ref.second) {
                return ref;
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long waviness_sum = 0, num_cnt = 0;
            bool is_num = !limit_low || i > diff_lh; // 前面是否填过数字
            for (int d = lo; d <= hi; d++) {
                // 当前填的数不是最高位，cmp 才有意义
                int cmp = is_num ? (d > last_digit) - (d < last_digit) : 0;
                auto [sub_waviness_sum, sub_num_cnt] = dfs(i + 1, cmp, d, limit_low && d == lo, limit_high && d == hi);
                waviness_sum += sub_waviness_sum; // 累加子树的波动值
                num_cnt += sub_num_cnt; // 累加子树的合法数字个数
                if (cmp * last_cmp < 0) { // 形成了一个峰或谷
                    waviness_sum += sub_num_cnt; // 这个峰谷会出现在 sub_num_cnt 个数字中
                }
            }

            pair<long long, long long> res = {waviness_sum, num_cnt};
            if (!limit_low && !limit_high) {
                ref = res;
            }
            return res;
        };

        return dfs(0, 0, 0, true, true).first;
    }
};
```

```go [sol-Go]
func totalWaviness(num1, num2 int64) int64 {
	lowS := strconv.FormatInt(num1, 10)
	highS := strconv.FormatInt(num2, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	type pair struct{ wavinessSum, numCnt int }
	memo := make([][3][10]pair, n)

	var dfs func(int, int, int, bool, bool) pair
	dfs = func(i, lastCmp, lastDigit int, limitLow, limitHigh bool) (res pair) {
		if i == n {
			return pair{0, 1} // 本题无特殊约束，能递归到终点的都是合法数字
		}
		if !limitLow && !limitHigh {
			p := &memo[i][lastCmp+1][lastDigit]
			if p.numCnt > 0 {
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

		isNum := !limitLow || i > diffLH // 前面是否填过数字
		for d := lo; d <= hi; d++ {
			c := 0
			if isNum { // 当前填的数不是最高位
				c = cmp.Compare(d, lastDigit)
			}
			sub := dfs(i+1, c, d, limitLow && d == lo, limitHigh && d == hi)
			res.wavinessSum += sub.wavinessSum // 累加子树的波动值
			res.numCnt += sub.numCnt // 累加子树的合法数字个数
			if c*lastCmp < 0 { // 形成了一个峰或谷
				res.wavinessSum += sub.numCnt // 这个峰谷会出现在 sub.numCnt 个数字中
			}
		}
		return
	}
	return int64(dfs(0, 0, 0, true, true).wavinessSum)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^2n)$，其中 $n = \mathcal{O}(\log \textit{num}_2)$ 是 $\textit{num}_2$ 的十进制长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(Dn)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^2n)$。
- 空间复杂度：$\mathcal{O}(Dn)$。

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
