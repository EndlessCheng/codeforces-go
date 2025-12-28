视频讲解：

- [数位 DP v1.0 模板讲解](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s)
- [数位 DP v2.0 模板讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s)（上下界数位 DP）

对于本题，我们需要知道最终 $奇数位之和 = 偶数位之和$ 是否成立，这等价于 $奇数位之和 - 偶数位之和 = 0$。

所以只需维护 $\textit{diff} = 奇数位之和 - 偶数位之和$。

此外，还需要知道当前填的是奇数位还是偶数位，用参数 $\textit{parity}$ 表示。

注意最小满足要求的数是 $11$，如果 $\textit{high} < 11$，可以直接返回 $0$。

此外，更新 $\textit{low}$ 为 $\max(\textit{low},11)$，以保证数字至少是两位数。

```py [sol-Python3]
class Solution:
    def countBalanced(self, low: int, high: int) -> int:
        # 最小的满足要求的数是 11
        if high < 11:
            return 0

        low = max(low, 11)
        low_s = list(map(int, str(low)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(high)))
        n = len(high_s)
        diff_lh = n - len(low_s)

        @cache
        def dfs(i: int, diff: int, parity: bool, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1 if diff == 0 else 0

            lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high_s[i] if limit_high else 9

            res = 0
            start = lo

            # 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
            if limit_low and i < diff_lh:
                # 不填数字，上界不受约束
                res = dfs(i + 1, diff, parity, True, False)
                start = 1  # 下面填数字，至少从 1 开始填

            for d in range(start, hi + 1):
                res += dfs(i + 1,
                           diff + (d if parity else -d),
                           not parity,  # 下一个位置奇偶性翻转
                           limit_low and d == lo,
                           limit_high and d == hi)

            return res

        return dfs(0, 0, True, True, True)
```

```java [sol-Java]
class Solution {
    public long countBalanced(long low, long high) {
        // 最小的满足要求的数是 11
        if (high < 11) {
            return 0;
        }

        low = Math.max(low, 11);
        char[] lowS = String.valueOf(low).toCharArray();
        char[] highS = String.valueOf(high).toCharArray();

        int n = highS.length;
        // diff 至少 floor(n/2) * 9，至多 ceil(n/2) * 9，值域大小 n * 9
        long[][][] memo = new long[n][n * 9 + 1][2];

        return dfs(0, n / 2 * 9, 1, true, true, lowS, highS, memo);
    }

    private long dfs(int i, int diff, int parity, boolean limitLow, boolean limitHigh, char[] lowS, char[] highS, long[][][] memo) {
        int n = highS.length;
        if (i == n) {
            return diff == n / 2 * 9 ? 1 : 0;
        }

        if (!limitLow && !limitHigh && memo[i][diff][parity] > 0) {
            return memo[i][diff][parity] - 1; // 记忆化的时候 +1，这里减掉
        }

        int diffLH = n - lowS.length;
        int lo = limitLow && i >= diffLH ? lowS[i - diffLH] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long res = 0;
        int d = lo;

        // 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
        if (limitLow && i < diffLH) {
            // 不填数字，上界不受约束
            res = dfs(i + 1, diff, parity, true, false, lowS, highS, memo);
            d = 1; // 下面填数字，至少从 1 开始填
        }

        for (; d <= hi; d++) {
            res += dfs(i + 1,
                    diff + (parity > 0 ? d : -d),
                    parity ^ 1, // 下一个位置奇偶性翻转
                    limitLow && d == lo,
                    limitHigh && d == hi,
                    lowS, highS, memo);
        }

        if (!limitLow && !limitHigh) {
            memo[i][diff][parity] = res + 1; // 记忆化的时候加一，这样 memo 可以初始化成 0
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countBalanced(long long low, long long high) {
        // 最小的满足要求的数是 11
        if (high < 11) {
            return 0;
        }

        low = max(low, 11LL);
        string low_s = to_string(low);
        string high_s = to_string(high);
        int n = high_s.size();
        int diff_lh = n - low_s.size();

        // diff 至少 floor(n/2) * 9，至多 ceil(n/2) * 9，值域大小 n * 9
        vector memo(n, vector<array<long long, 2>>(n * 9 + 1, {-1, -1}));

        auto dfs = [&](this auto&& dfs, int i, int diff, bool parity, bool limit_low, bool limit_high) -> long long {
            if (i == n) {
                return diff == n / 2 * 9;
            }

            if (!limit_low && !limit_high && memo[i][diff][parity] >= 0) {
                return memo[i][diff][parity];
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long res = 0;
            int d = lo;

            // 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
            if (limit_low && i < diff_lh) {
                // 不填数字，上界不受约束
                res = dfs(i + 1, diff, parity, true, false);
                d = 1; // 下面填数字，至少从 1 开始填
            }

            for (; d <= hi; d++) {
                res += dfs(i + 1,
                           diff + (parity ? d : -d),
                           !parity, // 下一个位置奇偶性翻转
                           limit_low && d == lo,
                           limit_high && d == hi);
            }

            if (!limit_low && !limit_high) {
                memo[i][diff][parity] = res;
            }
            return res;
        };

        return dfs(0, n / 2 * 9, true, true, true);
    }
};
```

```go [sol-Go]
func countBalanced(low, high int64) int64 {
	// 最小的满足要求的数是 11
	if high < 11 {
		return 0
	}

	low = max(low, 11)
	lowS := strconv.FormatInt(low, 10)
	highS := strconv.FormatInt(high, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][][2]int64, n)
	for i := range memo {
		// diff 至少 floor(n/2) * 9，至多 ceil(n/2) * 9，值域大小 n * 9
		memo[i] = make([][2]int64, n*9+1)
	}

	var dfs func(int, int, int, bool, bool) int64
	dfs = func(i, diff, parity int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			if diff != 0 { // 不合法
				return 0
			}
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][diff+n/2*9][parity] // 保证下标非负
			if *p > 0 {
				return *p - 1
			}
			defer func() { *p = res + 1 }() // 记忆化的时候加一，这样 memo 可以初始化成 0
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
		// 通过 limit_low 和 i 可以判断能否不填数字，无需 isNum 参数
		if limitLow && i < diffLH { // 可以不填任何数
			res = dfs(i+1, diff, parity, true, false) // 上界无约束
			d = 1 // 下面填数字，至少从 1 开始填
		}

		for ; d <= hi; d++ {
			// 下一个位置奇偶性翻转
			res += dfs(i+1, diff+(parity*2-1)*d, parity^1, 
				limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return dfs(0, 0, 1, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2D^2)$，其中 $n = \mathcal{O}(\log \textit{high})$ 是 $\textit{high}$ 的十进制长度，$D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2D)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(n^2D^2)$。
- 空间复杂度：$\mathcal{O}(n^2D)$。保存多少状态，就需要多少空间。

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
