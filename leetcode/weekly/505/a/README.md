## 方法一：暴力枚举

枚举 $[\max(n-k, 1), n+k]$ 中的整数 $x$，如果 `n & x` 等于 $0$，把答案增加 $x$。

```py [sol-Python3]
class Solution:
    def sumOfGoodIntegers(self, n: int, k: int) -> int:
        ans = 0
        for x in range(max(n - k, 1), n + k + 1):
            if n & x == 0:
                ans += x
        return ans
```

```java [sol-Java]
class Solution {
    public int sumOfGoodIntegers(int n, int k) {
        int ans = 0;
        for (int x = Math.max(n - k, 1); x <= n + k; x++) {
            if ((n & x) == 0) {
                ans += x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfGoodIntegers(int n, int k) {
        int ans = 0;
        for (int x = max(n - k, 1); x <= n + k; x++) {
            if ((n & x) == 0) {
                ans += x;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumOfGoodIntegers(n, k int) (ans int) {
	for x := max(n-k, 1); x <= n+k; x++ {
		if n&x == 0 {
			ans += x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：上下界数位 DP

[数位 DP v1.0 模板讲解](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s)

[数位 DP v2.0 模板讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s)（上下界数位 DP）

本题要求 `n & x` 等于 $0$。如果 $n$ 的第 $i$ 位是 $1$，那么 $x$ 的第 $i$ 位不能填 $1$。

### 答疑

**问**：下面的代码，为什么只在 `!limitLow && !limitHigh` 成立时才记忆化？（Python 选手可以跳过这个问题）

**答**：记忆化的原理是，当我们再次遇到相同状态时，可以直接返回 $\textit{memo}$ 中保存的结果。数位 DP 本质是暴力枚举，枚举每个数位填什么。$\textit{low}$ 是我们枚举的第一个数，$\textit{high}$ 是我们枚举的最后一个数。所以「填入的数字组成了 $\textit{low}$」以及「填入的数字组成了 $\textit{high}$」在整个递归过程中只会枚举一次。状态 $(\ldots, \textit{limitLow}, \textit{limitHigh})$ 中的 $\textit{limitLow}$ 和 $\textit{limitHigh}$ 如果其中一个是 $\texttt{true}$，说明我们正在填 $\textit{low}$ 或者正在填 $\textit{high}$，**这样的状态只会出现一次，不会再次遇到**，所以不需要记忆化这种状态。

```py [sol-Python3]
class Solution:
    def sumOfGoodIntegers(self, n: int, k: int) -> int:
        low = max(n - k, 1)
        high = n + k
        m = high.bit_length()

        # dfs 返回两个数：子树合法二进制数个数，子树和
        @cache
        def dfs(i: int, limit_low: bool, limit_high: bool) -> tuple[int, int]:
            if i < 0:
                return 1, 0  # 如果没有特殊约束，那么能递归到终点的都是合法二进制数

            lo = low >> i & 1 if limit_low else 0
            hi = high >> i & 1 if limit_high else 1
            cnt = res = 0

            for d in range(lo, hi + 1):
                bit = d << i
                if n & bit:
                    continue
                sub_cnt, sub_sum = dfs(i - 1, limit_low and d == lo, limit_high and d == hi)
                cnt += sub_cnt  # 累加子树的合法二进制数个数
                res += sub_sum  # 累加子树和
                res += bit * sub_cnt  # bit 会出现在 sub_cnt 个数中（贡献法）

            return cnt, res

        return dfs(m - 1, True, True)[1]
```

```java [sol-Java]
class Solution {
    public int sumOfGoodIntegers(int n, int k) {
        int low = Math.max(n - k, 1);
        int high = n + k;
        int m = 32 - Integer.numberOfLeadingZeros(high);
        int[][] memo = new int[m][];
        return dfs(m - 1, true, true, low, high, n, memo)[1];
    }

    // dfs 返回两个数：子树合法二进制数个数，子树和
    private int[] dfs(int i, boolean limitLow, boolean limitHigh, int low, int high, int n, int[][] memo) {
        if (i < 0) {
            return new int[]{1, 0}; // 如果没有特殊约束，那么能递归到终点的都是合法二进制数
        }

        if (!limitLow && !limitHigh && memo[i] != null) {
            return memo[i];
        }

        int lo = limitLow ? low >> i & 1 : 0;
        int hi = limitHigh ? high >> i & 1 : 1;

        int cnt = 0;
        int sum = 0;

        for (int d = lo; d <= hi; d++) {
            int bit = d << i;
            if ((n & bit) > 0) { // 不满足要求
                continue;
            }
            int[] sub = dfs(i - 1, limitLow && d == lo, limitHigh && d == hi, low, high, n, memo);
            cnt += sub[0]; // 累加子树的合法二进制数个数
            sum += sub[1]; // 累加子树和
            sum += bit * sub[0]; // bit 会出现在 sub[0] 个数中（贡献法）
        }

        int[] res = new int[]{cnt, sum};
        if (!limitLow && !limitHigh) {
            memo[i] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfGoodIntegers(int n, int k) {
        int low = max(n - k, 1);
        int high = n + k;
        int m = bit_width(1u * high);
        vector<pair<int, int>> memo(m, {-1, -1});

        // dfs 返回两个数：子树合法二进制数个数，子树和
        auto dfs = [&](this auto&& dfs, int i, bool limit_low, bool limit_high) -> pair<int, int> {
            if (i < 0) {
                return {1, 0}; // 如果没有特殊约束，那么能递归到终点的都是合法二进制数
            }

            if (!limit_low && !limit_high && memo[i].first >= 0) {
                return memo[i];
            }

            int lo = limit_low ? low >> i & 1 : 0;
            int hi = limit_high ? high >> i & 1 : 1;
            int cnt = 0, sum = 0;

            for (int d = lo; d <= hi; d++) {
                int bit = d << i;
                if (n & bit) { // 不满足要求
                    continue;
                }
                auto [sub_cnt, sub_sum] = dfs(i - 1, limit_low && d == lo, limit_high && d == hi);
                cnt += sub_cnt; // 累加子树的合法二进制数个数
                sum += sub_sum; // 累加子树和
                sum += bit * sub_cnt; // bit 会出现在 sub_cnt 个数中（贡献法）
            }

            pair<int, int> res = {cnt, sum};
            if (!limit_low && !limit_high) {
                memo[i] = res;
            }
            return res;
        };

        return dfs(m - 1, true, true).second;
    }
};
```

```go [sol-Go]
func sumOfGoodIntegers(n, k int) int {
	low := max(n-k, 1)
	high := n + k
	m := bits.Len(uint(high))
	type pair struct{ cnt, sum int }
	memo := make([]pair, m)
	for i := range memo {
		memo[i].cnt = -1
	}

	// dfs 返回两个数：子树合法二进制数个数，子树和
	var dfs func(int, bool, bool) pair
	dfs = func(i int, limitLow, limitHigh bool) (res pair) {
		if i < 0 {
			return pair{1, 0} // 如果没有特殊约束，能递归到终点的都是合法二进制数
		}
		if !limitLow && !limitHigh {
			p := &memo[i]
			if p.cnt >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = low >> i & 1
		}
		hi := 1
		if limitHigh {
			hi = high >> i & 1
		}

		for d := lo; d <= hi; d++ {
			bit := d << i
			if n&bit > 0 { // 不满足要求
				continue
			}
			sub := dfs(i-1, limitLow && d == lo, limitHigh && d == hi)
			res.cnt += sub.cnt       // 累加子树的合法二进制数个数
			res.sum += sub.sum       // 累加子树和
			res.sum += bit * sub.cnt // bit 会出现在 sub.cnt 个数中（贡献法）
		}
		return
	}

	return dfs(m-1, true, true).sum
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log (n + k))$。由于每个状态只会计算一次，记忆化搜索的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\log (n + k))$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(\log (n + k))$。
- 空间复杂度：$\mathcal{O}(\log (n + k))$。保存多少状态，就需要多少空间。

## 专题训练

见下面动态规划题单的「**§10.2 统计合法元素的价值总和**」。

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
