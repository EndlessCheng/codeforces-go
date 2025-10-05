## 分析

设 $n$ 的十进制字符串为 $s$。

设当前数位 $d = s[i]$。一般来说，我们要把 $d$ 拆分成两个 $[1,9]$ 的数字之和。拆分出 $0$ 的情况后面会讨论。

例如 $n=234$，最低位的 $4$ 有两类拆分方式：

1. 不借位。$4$ 可以拆分成 $1+3,2+2,3+1$。
2. 借位，把 $234$ 视作 $220 + 14$。$14$ 可以拆分成 $5+9,6+8,7+7,8+6,9+5$。对于十位数，原本是 $3$，被个位数借了个 $1$ 后，变成了 $2$。

一般地，借位会影响当前数位的值，从而影响拆分方案数。

- 设当前数位 $d = s[i]$。
- 如果低位（$i+1$）发生了借位，那么把 $d$ 减少 $1$，表示借给低位一个 $1$。
- $d$ 也可以向高位（$i-1$）借 $1$，把 $d$ 增加 $10$。
- 特别地，如果 $d=0$，那么把 $d$ 减少 $1$ 后 $d=-1$，此时必须向高位借 $1$。

把 $d$ 拆分成两数之和，有哪些情况？

注意**前导零**并不算在内，例如 $n=234 = 231+3$，其中 $3$ 相当于 $003$，有两个前导零。

## 情况一：无前导零

设 $d=a+b$，其中 $a$ 和 $b$ 都是 $[1,9]$ 中的正整数。

分类讨论：

- $d$ 不向高位借 $1$，我们计算的是 $[1,9]$ 中的两数之和等于 $d$ 的方案数，即 $(a,b) = (1,d-1),(2,d-2),\ldots,(d-1,1)$，一共有 $d-1$ 个。特别地，如果 $d < 2$，那么无解。
- $d$ 向高位借 $1$，我们计算的是 $[1,9]$ 中的两数之和等于 $d+10$ 的方案数，即 $(a,b) = (d+1,9),(d+2,8),\ldots,(9,d+1)$，一共有 $9-d$ 个，或者说 $19-(d+10)$ 个。

## 情况二：有前导零

设 $d=a+b$，其中 $a$ 和 $b$ 至少有一个是 $0$。

假设 $a=0$，那么对于更高位的 $d'=a'+b'$，必须满足 $a'=0$，即**前导零**。

由于 $a=0$，那么 $b$ 必须是 $d$。

分类讨论：

- 如果 $d\ne 0$：
  - $b=d$，只有一种填法。如果 $d<0$ 则必须向高位借 $1$。
  - 特别地，如果我们刚开始填前导零，那么 $(a,b) = (0,d)$ 或者 $(a,b) = (d,0)$ 是对称的。我们无需分别计算 $a=0$ 和 $b=0$ 的情况，只需计算其中一种情况，然后把方案数乘以 $2$。
- 如果 $d = 0$，那么 $b=d=0$。
   - 如果 $i>0$，由于我们没有向高位借 $1$，那么 $b$ 的高位至少有一个非零数字，$b$ 前面不能都是 $0$。此时无解。
   - 否则 $i=0$，两个数的最高位都是 $0$，这是合法的。例如 $100 = 49+51$，$49$ 和 $51$ 的百位都是前导零。

根据上面的讨论，定义 $\textit{dfs}(i,\textit{borrowed},\textit{isNum})$ 表示在 $[0,i]$ 中填数字，且满足如下约束时的方案数：

- $\textit{borrowed} = \textit{true}$ 表示被低位（$i+1$）借 $1$。
- $\textit{isNum} = \textit{true}$ 表示右边所填数位没有 $0$（无前导零）。

按照前文的分类讨论，计算情况一和情况二的方案数。

**递归边界**：如果 $i<0$，那么 $\textit{borrowed}$ 必须是 $\textit{false}$。是 $\textit{false}$ 则返回 $1$，不是则返回 $0$。

**递归入口**：为了获知低位的行为（是否借 $1$），要从最低位开始递归，即 $\textit{dfs}(|s|-1,\texttt{false},\texttt{true})$。

[本题视频讲解](https://www.bilibili.com/video/BV1ESxKzeEt5/?t=52m18s)，欢迎点赞关注~

```py [sol-Python3]
# 返回两个 1~9 的整数和为 target 的方案数
def two_sum_ways(target: int) -> int:
    return max(min(target - 1, 19 - target), 0)  # 保证结果非负

class Solution:
    def countNoZeroPairs(self, n: int) -> int:
        s = list(map(int, str(n)))
        m = len(s)

        # borrowed = True 表示被低位（i+1）借了个 1
        # is_num = True 表示之前填的数位，两个数都无前导零
        @cache
        def dfs(i: int, borrowed: bool, is_num: bool) -> int:
            if i < 0:
                # borrowed 必须为 False
                return 0 if borrowed else 1

            d = s[i] - borrowed
            res = 0

            # 情况一：两个数位都不为 0
            if is_num:
                res = dfs(i - 1, False, True) * two_sum_ways(d)  # 不向高位借 1
                res += dfs(i - 1, True, True) * two_sum_ways(d + 10) # 向高位借 1

            # 情况二：其中一个数位填前导零
            if i < m - 1:  # 不能是最低位
                if d:
                    # 如果 d < 0，必须向高位借 1
                    # 如果 is_num = True，根据对称性，方案数要乘以 2
                    res += dfs(i - 1, d < 0, False) * (is_num + 1)
                elif i == 0:  # 两个数位都填 0，只有当 i = 0 的时候才有解
                    res += 1

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
                Arrays.fill(row, -1); // -1 表示没有计算过
            }
        }
        return dfs(m - 1, false, true, s, memo);
    }

    // borrowed = true 表示被低位（i+1）借了个 1
    // isNum = true 表示之前填的数位，两个数都无前导零
    private long dfs(int i, boolean borrowed, boolean isNum, char[] s, long[][][] memo) {
        if (i < 0) {
            // borrowed 必须为 false
            return borrowed ? 0 : 1;
        }

        int ib = borrowed ? 1 : 0;
        int in = isNum ? 1 : 0;
        if (memo[i][ib][in] != -1) { // 之前计算过
            return memo[i][ib][in];
        }

        int d = s[i] - '0' - (borrowed ? 1 : 0);
        long res = 0;

        // 情况一：两个数位都不为 0
        if (isNum) {
            res = dfs(i - 1, false, true, s, memo) * twoSumWays(d); // 不向高位借 1
            res += dfs(i - 1, true, true, s, memo) * twoSumWays(d + 10); // 向高位借 1
        }

        // 情况二：其中一个数位填前导零
        if (i < s.length - 1) { // 不能是最低位
            if (d != 0) {
                // 如果 d < 0，必须向高位借 1
                // 如果 isNum = true，根据对称性，方案数要乘以 2
                res += dfs(i - 1, d < 0, false, s, memo) * (isNum ? 2 : 1);
            } else if (i == 0) { // 两个数位都填 0，只有当 i = 0 的时候才有解
                res++;
            }
        }

        return memo[i][ib][in] = res; // 记忆化
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
    int two_sum_ways(int target) {
        return max(min(target - 1, 19 - target), 0); // 保证结果非负
    }

public:
    long long countNoZeroPairs(long long n) {
        string s = to_string(n);
        int m = s.size();
        vector memo(m, array<array<long long, 2>, 2>({{-1, -1}, {-1, -1}})); // -1 表示没有计算过

        // borrowed = true 表示被低位（i+1）借了个 1
        // is_num = true 表示之前填的数位，两个数都无前导零
        auto dfs = [&](this auto&& dfs, int i, bool borrowed, bool is_num) -> long long {
            if (i < 0) {
                // borrowed 必须为 false
                return !borrowed;
            }

            long long& res = memo[i][borrowed][is_num]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }

            int d = s[i] - '0' - borrowed;
            res = 0;

            // 情况一：两个数位都不为 0
            if (is_num) {
                res = dfs(i - 1, false, true) * two_sum_ways(d); // 不向高位借 1
                res += dfs(i - 1, true, true) * two_sum_ways(d + 10); // 向高位借 1
            }

            // 情况二：其中一个数位填前导零
            if (i < m - 1) { // 不能是最低位
                if (d) {
                    // 如果 d < 0，必须向高位借 1
                    // 如果 is_num = true，根据对称性，方案数要乘以 2
                    res += dfs(i - 1, d < 0, false) * (is_num + 1);
                } else if (i == 0) { // 两个数位都填 0，只有当 i = 0 的时候才有解
                    res++;
                }
            }

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

	// borrowed = 1 表示被低位（i+1）借了个 1
	// isNum = 1 表示之前填的数位，两个数都无前导零
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

		// 情况一：两个数位都不为 0
		if isNum > 0 {
			res = dfs(i-1, 0, 1) * twoSumWays(d)     // 不向高位借 1
			res += dfs(i-1, 1, 1) * twoSumWays(d+10) // 向高位借 1
		}

		// 情况二：其中一个数位填前导零
		if i < m-1 { // 不能是最低位
			if d != 0 {
				// 如果 d < 0，必须向高位借 1
				// 如果 isNum = 1，根据对称性，方案数要乘以 2
				res += dfs(i-1, isNeg(d), 0) * (isNum + 1)
			} else if i == 0 { // 两个数位都填 0，只有当 i = 0 的时候才有解
				res++
			}
		}
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

- 时间复杂度：$\mathcal{O}(\log n)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $n$ 的十进制长度 $\mathcal{O}(\log n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(\log n)$。保存多少状态，就需要多少空间。

## 相关题目

[1317. 将整数转换为两个无零整数的和](https://leetcode.cn/problems/convert-integer-to-the-sum-of-two-no-zero-integers/)

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
