设 $n$ 的十进制字符串为 $s$。

设 $d = s[i]$。首先，我们明确一些**基本的运算规则**：

- 如果低位（$i+1$）有借位，那么把 $d$ 减一，表示借给低位一个 $1$。
- 如果减一后 $d<0$，那么必须向高位（$i-1$）借一位，把 $d$ 增加 $10$。

首先，讨论两个数位都不为 $0$ 的情况。

- $d$ 不向高位借位，我们计算的是 $[1,9]$ 中的两数之和等于 $d$ 的方案数，即 $(1,d-1]),(2,d-2),\ldots,(s[i-1],1)$，这一共有 $d-1$ 个。特别地，如果 $d < 2$，那么无解。
- $d$ 向高位借位，我们计算的是 $[1,9]$ 中的两数之和等于 $d+10$ 的方案数，这一共有 $9-d$ 个。特别地，如果 $d=9$，那么 $d+10=19$，无解。

然后，讨论其中一个数 $a$ 在 $s[i]$ 这一位上是 $0$ 的情况。

这意味着 $a$ 的更高位都必须是 $0$，即**前导零**。

分类讨论：

- 如果 $i>0$，那么另一个数 $b$ 在这一位上必须填 $d$。
   - 如果 $d = 0$，不合法。
   - 否则只有一种填法。
   - 根据对称性，$a$ 这一位填 $d$ 或者 $b$ 这一位填 $d$ 是对称的，方案数可以乘以 $2$。
- 如果 $i=0$，那么当 $d=0$ 时，另一个数位可以为 $0$，即两个数的最高位都是 $0$。例如 $100 = 49+51$。

根据上面的讨论，定义 $\textit{dfs}(i,\textit{borrowed},\textit{isNum})$ 表示在 $[0,i]$ 中填数字，且满足如下约束时的方案数：

- $\textit{borrowed}= \textit{true}$ 表示被低位（$i+1$）借位。
- $\textit{isNum}= \textit{true}$ 表示之前填的数位，两个数都不为 $0$（无前导零）。

**递归边界**：如果 $i<0$，那么 $\textit{borrowed}$ 必须是 $\textit{false}$，满足则返回 $1$，不满足则返回 $0$。

**递归入口**：从最低位开始，$\textit{dfs}(|s|-1,\texttt{false},\texttt{true})$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 返回两个 1~9 的整数和为 target 的方案数
def two_sum_ways(target: int) -> int:
    return max(min(target - 1, 19 - target), 0)  # 保证结果非负

class Solution:
    def countNoZeroPairs(self, n: int) -> int:
        s = list(map(int, str(n)))
        m = len(s)

        # borrow = True 表示被低位（i+1）借位
        # isNum = True 表示之前填的数位，两个数都不为 0（无前导零）
        @cache
        def dfs(i: int, borrowed: bool, is_num: bool) -> int:
            if i < 0:
                # borrowed 必须为 False
                return 0 if borrowed else 1

            d = s[i] - borrowed

            # 其中一个数必须填前导零
            if not is_num:
                # 在 i > 0 的情况下，另一个数必须不为 0（否则可以为 0，即两个数的最高位都是 0）
                if i > 0 and d == 0:
                    return 0
                # 如果 d < 0，必须向高位借位
                return dfs(i - 1, d < 0, False)

            # 令其中一个数从当前位置开始往左都是 0（前导零）
            res = 0
            if i < m - 1:
                if d != 0:  # 另一个数不为 0
                    res = dfs(i - 1, d < 0, False) * 2  # 根据对称性乘以 2
                elif i == 0:  # 最高位被借走
                    res = 1  # 两个数都是 0
                # else res = 0

            # 两个数位都不为 0
            res += dfs(i - 1, False, True) * two_sum_ways(d)  # 不向 i-1 借位
            res += dfs(i - 1, True, True) * two_sum_ways(d + 10)  # 向 i-1 借位
            return res

        return dfs(m - 1, False, True)
```

```java [sol-Java]
class Solution {
    public long countNoZeroPairs(long n) {
        char[] s = Long.toString(n).toCharArray();
        int m = s.length;
        long[][][] memo = new long[m][2][2];
        for (long[][] mat : memo) {
            for (long[] row : mat) {
                Arrays.fill(row, -1);
            }
        }
        return dfs(m - 1, false, true, s, memo);
    }

    // borrow = true 表示被低位（i+1）借位
    // isNum = true 表示之前填的数位，两个数都不为 0（无前导零）
    private long dfs(int i, boolean borrowed, boolean isNum, char[] s, long[][][] memo) {
        if (i < 0) {
            // borrowed 必须为 false
            return borrowed ? 0 : 1;
        }
        int ib = borrowed ? 1 : 0;
        int in = isNum ? 1 : 0;
        if (memo[i][ib][in] != -1) {
            return memo[i][ib][in];
        }

        int d = s[i] - '0' - (borrowed ? 1 : 0);

        // 其中一个数必须填前导零
        if (!isNum) {
            // 在 i > 0 的情况下，另一个数必须不为 0（否则可以为 0，即两个数的最高位都是 0）
            if (i > 0 && d == 0) {
                return memo[i][ib][in] = 0;
            }
            // 如果 d < 0，必须向高位借位
            return memo[i][ib][in] = dfs(i - 1, d < 0, false, s, memo);
        }

        // 令其中一个数从当前位置开始往左都是 0（前导零）
        long res = 0;
        if (i < s.length - 1) {
            if (d != 0) { // 另一个数不为 0
                res = dfs(i - 1, d < 0, false, s, memo) * 2; // 根据对称性乘以 2
            } else if (i == 0) { // 最高位被借走
                res = 1; // 两个数都是 0
            } // else res = 0
        }

        // 两个数位都不为 0
        res += dfs(i - 1, false, true, s, memo) * twoSumWays(d); // 不向 i-1 借位
        res += dfs(i - 1, true, true, s, memo) * twoSumWays(d + 10); // 向 i-1 借位
        return memo[i][ib][in] = res;
    }

    // 返回两个 1~9 的整数和为 target 的方案数
    private int twoSumWays(int target) {
        return Math.max(Math.min(target - 1, 19 - target), 0); // 保证结果非负
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回两个 1~9 的整数和为 target 的方案数
    int twoSumWays(int target) {
        return max(min(target - 1, 19 - target), 0); // 保证结果非负
    }

public:
    long long countNoZeroPairs(long long n) {
        string s = to_string(n);
        int m = s.size();
        vector memo(m, array<array<long long, 2>, 2>({{-1, -1}, {-1, -1}}));

        // borrow = true 表示被低位（i+1）借位
        // isNum = true 表示之前填的数位，两个数都不为 0（无前导零）
        auto dfs = [&](this auto&& dfs, int i, bool borrowed, bool is_num) -> long long {
            if (i < 0) {
                // borrowed 必须为 false
                return !borrowed;
            }
            long long& res = memo[i][borrowed][is_num];
            if (res != -1) {
                return res;
            }

            int d = s[i] - '0' - borrowed;

            // 其中一个数必须填前导零
            res = 0;
            if (!is_num) {
                // 在 i > 0 的情况下，另一个数必须不为 0（否则可以为 0，即两个数的最高位都是 0）
                if (i > 0 && d == 0) {
                    return res = 0;
                }
                // 如果 d < 0，必须向高位借位
                return res = dfs(i - 1, d < 0, false);
            }

            // 令其中一个数从当前位置开始往左都是 0（前导零）
            if (i < m - 1) {
                if (d != 0) { // 另一个数不为 0
                    res = dfs(i - 1, d < 0, false) * 2; // 根据对称性乘以 2
                } else if (i == 0) { // 最高位被借走
                    res = 1; // 两个数都是 0
                }
                // else res = 0
            }

            // 两个数位都不为 0
            res += dfs(i - 1, false, true) * twoSumWays(d); // 不向 i-1 借位
            res += dfs(i - 1, true, true) * twoSumWays(d + 10); // 向 i-1 借位
            return res;
        };

        return dfs(m - 1, false, true);
    }
};
```

```go [sol-Go]
// 返回两个 1~9 的整数和为 target 的方案数
func twoSumWays(target int) int {
	return max(min(target-1, 19-target), 0) // 保证结果非负
}

func countNoZeroPairs(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	m := len(s)
	memo := make([][2][2]int, m)
	for i := range memo {
		memo[i] = [2][2]int{{-1, -1}, {-1, -1}} // -1 表示没有计算过
	}

	// borrow = 1 表示被低位（i+1）借位
	// isNum = 1 表示之前填的数位，两个数都不为 0（无前导零）
	var dfs func(int, int, int) int
	dfs = func(i, borrowed, isNum int) (res int) {
		if i < 0 {
			// borrowed 必须为 0
			return borrowed ^ 1
		}

		p := &memo[i][borrowed][isNum]
		if *p >= 0 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化

		d := int(s[i]-'0') - borrowed
		// 其中一个数必须填前导零
		if isNum == 0 {
			// 在 i > 0 的情况下，另一个数必须不为 0（否则可以为 0，即两个数的最高位都是 0）
			if i > 0 && d == 0 {
				return 0
			}
			// 如果 d < 0，必须向高位借位
			return dfs(i-1, isNeg(d), 0)
		}

		// 令其中一个数从当前位置开始往左都是 0（前导零）
		if i < m-1 {
			if d != 0 { // 另一个数不为 0
				res = dfs(i-1, isNeg(d), 0) * 2 // 根据对称性乘以 2
			} else if i == 0 { // 最高位被借走
				res = 1 // 两个数都是 0
			} // else res = 0
		}

		// 两个数位都不为 0
		res += dfs(i-1, 0, 1) * twoSumWays(d)    // 不向 i-1 借位
		res += dfs(i-1, 1, 1) * twoSumWays(d+10) // 向 i-1 借位
		return
	}

	return int64(dfs(m-1, 0, 1))
}

// 返回 d 是否为负数
func isNeg(d int) int {
	if d < 0 {
		return 1
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(\log n)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
