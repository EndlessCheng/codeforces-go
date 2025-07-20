本题和 [3352. 统计小于 N 的 K 可约简整数](https://leetcode.cn/problems/count-k-reducible-numbers-less-than-n/) 几乎一样。

如果 $k=0$，那么只有 $x=1$ 满足要求，返回 $1$。

如果 $k=1$，那么只有 $x=2^1,2^2,\ldots,2^{m-1}$ 满足要求，其中 $m$ 是 $n$ 的二进制长度。这一共有 $m-1$ 个数。

如果 $k\ge 2$，计算方法同 3352 题，区别：

- 3352 题是小于 $n$，本题是小于等于 $n$。DFS 中无需判断是否严格小于 $n$。递归入口 $i$ 可以等于 $m$。
- 3352 题是小于等于 $k$，本题是恰好等于 $k$。递归入口的判断改成恰好等于 $k$。
- 本题无需取模。

```py [sol-Python3]
class Solution:
    def popcountDepth(self, n: int, k: int) -> int:
        if k == 0:
            return 1

        # 注：也可以不转成字符串，下面 dfs 用位运算取出 n 的第 i 位
        # 但转成字符串的通用性更好
        s = list(map(int, bin(n)[2:]))
        m = len(s)
        if k == 1:
            return m - 1

        @cache
        def dfs(i: int, left1: int, is_limit: bool) -> int:
            if i == m:
                return 0 if left1 else 1
            up = s[i] if is_limit else 1
            res = 0
            for d in range(min(up, left1) + 1):
                res += dfs(i + 1, left1 - d, is_limit and d == up)
            return res

        ans = 0
        f = [0] * (m + 1)
        for i in range(1, m + 1):
            f[i] = f[i.bit_count()] + 1
            if f[i] == k:
                # 计算有多少个二进制数恰好有 i 个 1
                ans += dfs(0, i, True)
        return ans
```

```java [sol-Java]
class Solution {
    public long popcountDepth(long n, int k) {
        if (k == 0) {
            return 1;
        }

        // 注：也可以不转成字符串，下面 dfs 用位运算取出 n 的第 i 位
        // 但转成字符串的通用性更好
        char[] s = Long.toBinaryString(n).toCharArray();
        int m = s.length;
        if (k == 1) {
            return m - 1;
        }

        long[][] memo = new long[m][m + 1];
        for (long[] row : memo) {
            Arrays.fill(row, -1);
        }

        long ans = 0;
        int[] f = new int[m + 1];
        for (int i = 1; i <= m; i++) {
            f[i] = f[Integer.bitCount(i)] + 1;
            if (f[i] == k) {
                // 计算有多少个二进制数恰好有 i 个 1
                ans += dfs(0, i, true, s, memo);
            }
        }
        return ans;
    }

    private long dfs(int i, int left1, boolean isLimit, char[] s, long[][] memo) {
        if (i == s.length) {
            return left1 == 0 ? 1 : 0;
        }
        if (!isLimit && memo[i][left1] != -1) {
            return memo[i][left1];
        }

        int up = isLimit ? s[i] - '0' : 1;
        long res = 0;
        for (int d = 0; d <= Math.min(up, left1); d++) {
            res += dfs(i + 1, left1 - d, isLimit && d == up, s, memo);
        }

        if (!isLimit) {
            memo[i][left1] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long popcountDepth(long long n, int k) {
        if (k == 0) {
            return 1;
        }

        int m = bit_width((uint64_t) n);
        if (k == 1) {
            return m - 1;
        }

        vector memo(m, vector<long long>(m + 1, -1));
        auto dfs = [&](this auto& dfs, int i, int left1, bool is_limit) -> long long {
            if (i < 0) {
                return left1 == 0;
            }
            if (!is_limit && memo[i][left1] != -1) {
                return memo[i][left1];
            }

            // 直接用位运算取出 n 的第 i 位
            int up = is_limit ? n >> i & 1 : 1;
            long long res = 0;
            for (int d = 0; d <= min(up, left1); d++) {
                res += dfs(i - 1, left1 - d, is_limit && d == up);
            }

            if (!is_limit) {
                memo[i][left1] = res;
            }
            return res;
        };

        long long ans = 0;
        vector<int> f(m + 1);
        for (uint32_t i = 1; i <= m; i++) {
            f[i] = f[popcount(i)] + 1;
            if (f[i] == k) {
                // 计算有多少个二进制数恰好有 i 个 1
                ans += dfs(m - 1, i, true);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func popcountDepth(n int64, k int) (ans int64) {
	if k == 0 {
		return 1
	}

	// 注：也可以不转成字符串，下面 dfs 用位运算取出 n 的第 i 位 
	// 但转成字符串的通用性更好
	s := strconv.FormatInt(n, 2)
	m := len(s)
	if k == 1 {
		return int64(m - 1)
	}

	memo := make([][]int64, m)
	for i := range memo {
		memo[i] = make([]int64, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool) int64
	dfs = func(i, left1 int, isLimit bool) (res int64) {
		if i == m {
			if left1 == 0 {
				return 1
			}
			return
		}
		if !isLimit {
			p := &memo[i][left1]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		up := 1
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= min(up, left1); d++ {
			res += dfs(i+1, left1-d, isLimit && d == up)
		}
		return
	}

	f := make([]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = f[bits.OnesCount(uint(i))] + 1
		if f[i] == k {
			// 计算有多少个二进制数恰好有 i 个 1
			ans += dfs(0, i, true)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log^2 n)$。
- 空间复杂度：$\mathcal{O}(\log^2 n)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
