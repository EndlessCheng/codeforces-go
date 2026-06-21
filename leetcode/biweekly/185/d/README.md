[数位 DP v1.0 模板讲解](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s)

[数位 DP v2.0 模板讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s)（上下界数位 DP）

本题需要记录上一个填的数字 $\textit{pre}$。

枚举当前填的数字 $d$。如果 $d$ 是最高位（之前没有填过数字），或者 $|d - \textit{pre}|\le k$，那么可以填 $d$，继续递归。

递归到终点时，找到了一个好数，返回 $1$。

下面是数位 DP v2.1 模板。相比 v2.0，不需要写 $\textit{isNum}$ 参数（之前是否填过数字）。

## 答疑

**问**：下面的代码，为什么只在 `!limitLow && !limitHigh` 成立时才记忆化？（Python 选手可以跳过这个问题）

**答**：记忆化的原理是，当我们再次遇到相同状态时，可以直接返回 $\textit{memo}$ 中保存的结果。数位 DP 本质是暴力枚举，枚举每个数位填什么。$\textit{low}$ 是我们枚举的第一个数，$\textit{high}$ 是我们枚举的最后一个数。所以「填入的数字组成了 $\textit{low}$」以及「填入的数字组成了 $\textit{high}$」在整个递归过程中只会枚举一次。状态 $(\ldots, \textit{limitLow}, \textit{limitHigh})$ 中的 $\textit{limitLow}$ 和 $\textit{limitHigh}$ 如果其中一个是 $\texttt{true}$，说明我们正在填 $\textit{low}$ 或者正在填 $\textit{high}$，**这样的状态只会出现一次，不会再次遇到**，所以不需要记忆化这种状态。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def goodIntegers(self, l: int, r: int, k: int) -> int:
        low_s = list(map(int, str(l)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(r)))
        n = len(high_s)
        diff_lh = n - len(low_s)

        @cache
        def dfs(i: int, pre: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1  # 找到一个好数

            lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high_s[i] if limit_high else 9

            res = 0
            start = lo

            if limit_low and i < diff_lh:
                # 不填数字，上界不受约束
                res = dfs(i + 1, 0, True, False)
                start = 1  # 下面填数字，从 1 开始填

            # 如果在 diff_lh 之前填过数字，那么 limit_low 一定是 False
            is_first_num = limit_low and i <= diff_lh
            for d in range(start, hi + 1):
                if is_first_num or abs(d - pre) <= k:
                    res += dfs(i + 1, d, limit_low and d == lo, limit_high and d == hi)

            return res

        # pre 的初始值随意
        return dfs(0, 0, True, True)
```

```java [sol-Java]
class Solution {
    public long goodIntegers(long l, long r, int k) {
        char[] lowS = String.valueOf(l).toCharArray();
        char[] highS = String.valueOf(r).toCharArray();

        int n = highS.length;
        long[][] memo = new long[n][10];
        for (long[] row : memo) {
            Arrays.fill(row, -1);
        }

        // pre 的初始值随意
        return dfs(0, 0, true, true, lowS, highS, k, memo);
    }

    private long dfs(int i, int pre, boolean limitLow, boolean limitHigh, char[] lowS, char[] highS, int k, long[][] memo) {
        if (i == highS.length) {
            return 1; // 找到一个好数
        }

        if (!limitLow && !limitHigh && memo[i][pre] >= 0) {
            return memo[i][pre];
        }

        int diff = highS.length - lowS.length;
        int lo = limitLow && i >= diff ? lowS[i - diff] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long res = 0;
        int d = lo;

        if (limitLow && i < diff) {
            // 不填数字，上界不受约束
            res = dfs(i + 1, 0, true, false, lowS, highS, k, memo);
            d = 1; // 下面填数字，从 1 开始填
        }

        // 如果在 diff 之前填过数字，那么 limitLow 一定是 false
        boolean isFirst = limitLow && i <= diff;
        for (; d <= hi; d++) {
            if (isFirst || Math.abs(d - pre) <= k) {
                res += dfs(i + 1, d, limitLow && d == lo, limitHigh && d == hi, lowS, highS, k, memo);
            }
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
    long long goodIntegers(long long l, long long r, int k) {
        string low_s = to_string(l);
        string high_s = to_string(r);
        int n = high_s.size();
        int diff_lh = n - low_s.size();
        vector<array<long long, 10>> memo(n); // 保存 res+1，这样 memo[i][j] 只需初始化成 0

        auto dfs = [&](this auto&& dfs, int i, int pre, bool limit_low, bool limit_high) -> long long {
            if (i == n) {
                return 1; // 找到一个好数
            }

            if (!limit_low && !limit_high && memo[i][pre] > 0) {
                return memo[i][pre] - 1;
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long res = 0;
            int d = lo;

            if (limit_low && i < diff_lh) {
                // 不填数字，上界不受约束
                res = dfs(i + 1, 0, true, false);
                d = 1; // 下面填数字，从 1 开始填
            }

            // 如果在 diff_lh 之前填过数字，那么 limit_low 一定是 false
            bool is_first = limit_low && i <= diff_lh;
            for (; d <= hi; d++) {
                if (is_first || abs(d - pre) <= k) {
                    res += dfs(i + 1, d, limit_low && d == lo, limit_high && d == hi);
                }
            }

            if (!limit_low && !limit_high) {
                memo[i][pre] = res + 1;
            }
            return res;
        };

        // pre 的初始值随意
        return dfs(0, 0, true, true);
    }
};
```

```go [sol-Go]
func goodIntegers(l, r int64, k int) int64 {
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][10]int64, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int64
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			return 1 // 找到一个好数
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
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		// 如果在 diffLH 之前填过数字，那么 limitLow 一定是 false
		isFirst := limitLow && i <= diffLH
		for ; d <= hi; d++ {
			if isFirst || abs(d-pre) <= k {
				res += dfs(i+1, d, limitLow && d == lo, limitHigh && d == hi)
			}
		}
		return
	}

	// pre 的初始值随意
	return dfs(0, 0, true, true)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^2\log r)$，其中 $D=10$。由于每个状态只会计算一次，记忆化搜索的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(D\log r)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^2\log r)$。
- 空间复杂度：$\mathcal{O}(D\log r)$。保存多少状态，就需要多少空间。

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
