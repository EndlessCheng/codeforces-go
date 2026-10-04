做法同 [1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/)，[我的题解](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/solutions/2321829/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-hzz6/)。

在 1186 的基础上，只需增加一个布尔参数 $\textit{rev}$，表示是否要把当前数 $\textit{nums}[i]$ 取反。每选一个数，就把 $\textit{rev}$ 取反。一开始 $\textit{rev} = \texttt{false}$。

[本题视频讲解](https://www.bilibili.com/video/BV1FiHj6uESs/?t=9m8s)，欢迎点赞关注~

## 写法一：记忆化搜索

```py [sol-Python3]
class Solution:
    def maxAlternatingSum(self, nums: list[int]) -> int:
        # 只需在 1186 的基础上增加参数 rev
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int, rev: bool) -> int:
            if i == n:
                return -inf  # 子数组至少要有一个数，不合法
            x = -nums[i] if rev else nums[i]
            if j == 0:
                return max(dfs(i + 1, 0, not rev), 0) + x
            return max(dfs(i + 1, 1, not rev) + x, dfs(i + 1, 0, rev))

        n = len(nums)
        ans = max(max(dfs(i, 0, False), dfs(i, 1, False)) for i in range(n))
        dfs.cache_clear()
        return ans
```

```java [sol-Java]
class Solution {
    public long maxAlternatingSum(int[] nums) {
        int n = nums.length;
        long[][][] memo = new long[n][2][2];
        for (long[][] mat : memo) {
            mat[0][0] = mat[0][1] = mat[1][0] = mat[1][1] = Long.MIN_VALUE;
        }

        long ans = Long.MIN_VALUE;
        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, Math.max(dfs(i, 0, 0, nums, memo), dfs(i, 1, 0, nums, memo)));
        }
        return ans;
    }

    // 只需在 1186 的基础上增加参数 rev
    private long dfs(int i, int j, int rev, int[] nums, long[][][] memo) {
        if (i == nums.length) { // 子数组至少要有一个数，不合法
            return Long.MIN_VALUE / 2; // 除 2 防止负数相加溢出
        }
        if (memo[i][j][rev] != Long.MIN_VALUE) {
            return memo[i][j][rev]; // 之前计算过
        }

        int x = rev == 0 ? nums[i] : -nums[i];
        if (j == 0) {
            return memo[i][j][rev] = Math.max(dfs(i + 1, 0, rev ^ 1, nums, memo), 0) + x;
        }
        return memo[i][j][rev] = Math.max(dfs(i + 1, 1, rev ^ 1, nums, memo) + x, dfs(i + 1, 0, rev, nums, memo));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        int n = nums.size();
        vector<array<array<long long, 2>, 2>> memo(n, {{{LLONG_MIN, LLONG_MIN}, {LLONG_MIN, LLONG_MIN}}});

        // 只需在 1186 的基础上增加参数 rev
        auto dfs = [&](this auto&& dfs, int i, int j, bool rev) -> long long {
            if (i == n) { // 子数组至少要有一个数，不合法
                return LLONG_MIN / 2; // 除 2 防止负数相加溢出
            }
            auto& res = memo[i][j][rev]; // 注意这里是引用
            if (res != LLONG_MIN) {
                return res; // 之前计算过
            }

            int x = rev ? -nums[i] : nums[i];
            if (j == 0) {
                return res = max(dfs(i + 1, 0, !rev), 0LL) + x;
            }
            return res = max(dfs(i + 1, 1, !rev) + x, dfs(i + 1, 0, rev));
        };

        long long ans = LLONG_MIN;
        for (int i = 0; i < n; i++) {
            ans = max(ans, max(dfs(i, 0, false), dfs(i, 1, false)));
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	memo := make([][2][2]int, n)
	for i := range memo {
		memo[i] = [2][2]int{{math.MinInt, math.MinInt}, {math.MinInt, math.MinInt}}
	}

	// 只需在 1186 的基础上增加参数 rev
	var dfs func(int, int, int) int
	dfs = func(i, j, rev int) (res int) {
		if i == n { // 子数组至少要有一个数，不合法
			return math.MinInt / 2 // 除 2 防止负数相加溢出
		}
		p := &memo[i][j][rev]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化

		x := nums[i]
		if rev > 0 {
			x = -x
		}

		if j == 0 {
			return max(dfs(i+1, 0, rev^1), 0) + x
		}
		return max(dfs(i+1, 1, rev^1)+x, dfs(i+1, 0, rev))
	}

	ans := math.MinInt
	for i := range nums {
		ans = max(ans, dfs(i, 0, 0), dfs(i, 1, 0))
	}
	return int64(ans)
}
```

## 写法二：递推

```py [sol-Python3]
class Solution:
    def maxAlternatingSum(self, nums: list[int]) -> int:
        n = len(nums)
        f = [[[-inf] * 2 for _ in range(2)] for _ in range(n + 1)]
        ans = -inf

        for i in range(n - 1, -1, -1):
            x = nums[i]
            f[i][0][0] = max(f[i + 1][0][1], 0) + x
            f[i][0][1] = max(f[i + 1][0][0], 0) - x
            f[i][1][0] = max(f[i + 1][1][1] + x, f[i + 1][0][0])
            f[i][1][1] = max(f[i + 1][1][0] - x, f[i + 1][0][1])
            ans = max(ans, f[i][0][0], f[i][1][0])

        return ans
```

```java [sol-Java]
class Solution {
    public long maxAlternatingSum(int[] nums) {
        int n = nums.length;
        long[][][] f = new long[n + 1][2][2];
        for (long[][] mat : f) {
            mat[0][0] = mat[0][1] = mat[1][0] = mat[1][1] = Long.MIN_VALUE / 2;
        }

        long ans = Long.MIN_VALUE;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            f[i][0][0] = Math.max(f[i + 1][0][1], 0) + x;
            f[i][0][1] = Math.max(f[i + 1][0][0], 0) - x;
            f[i][1][0] = Math.max(f[i + 1][1][1] + x, f[i + 1][0][0]);
            f[i][1][1] = Math.max(f[i + 1][1][0] - x, f[i + 1][0][1]);
            ans = Math.max(ans, Math.max(f[i][0][0], f[i][1][0]));
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        constexpr long long NEG_INF = LLONG_MIN / 2;
        int n = nums.size();
        vector<array<array<long long, 2>, 2>> f(n + 1, {{{NEG_INF, NEG_INF}, {NEG_INF, NEG_INF}}});
        long long ans = NEG_INF;

        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            f[i][0][0] = max(f[i + 1][0][1], 0LL) + x;
            f[i][0][1] = max(f[i + 1][0][0], 0LL) - x;
            f[i][1][0] = max(f[i + 1][1][1] + x, f[i + 1][0][0]);
            f[i][1][1] = max(f[i + 1][1][0] - x, f[i + 1][0][1]);
            ans = max(ans, max(f[i][0][0], f[i][1][0]));
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxAlternatingSum(nums []int) int64 {
	const negInf = math.MinInt / 2
	n := len(nums)
	f := make([][2][2]int, n+1)
	f[n] = [2][2]int{{negInf, negInf}, {negInf, negInf}} // 除 2 防止负数相加溢出
	ans := negInf
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		f[i][0][0] = max(f[i+1][0][1], 0) + x
		f[i][0][1] = max(f[i+1][0][0], 0) - x
		f[i][1][0] = max(f[i+1][1][1]+x, f[i+1][0][0])
		f[i][1][1] = max(f[i+1][1][0]-x, f[i+1][0][1])
		ans = max(ans, f[i][0][0], f[i][1][0])
	}
	return int64(ans)
}
```

## 写法三：空间优化

```py [sol-Python3]
class Solution:
    def maxAlternatingSum(self, nums: list[int]) -> int:
        ans = f00 = f01 = f10 = f11 = -inf
        for x in reversed(nums):
            f10, f11 = max(f11 + x, f00), max(f10 - x, f01)
            f00, f01 = max(f01, 0) + x, max(f00, 0) - x
            ans = max(ans, f00, f10)
        return ans
```

```java [sol-Java]
class Solution {
    public long maxAlternatingSum(int[] nums) {
        final long NEG_INF = Long.MIN_VALUE / 2;
        long f00 = NEG_INF, f01 = NEG_INF, f10 = NEG_INF, f11 = NEG_INF;
        long ans = NEG_INF;
        for (int i = nums.length - 1; i >= 0; i--) {
            int x = nums[i];
            long newF10 = Math.max(f11 + x, f00);
            long newF11 = Math.max(f10 - x, f01);
            f10 = newF10;
            f11 = newF11;

            long newF00 = Math.max(f01, 0) + x;
            long newF01 = Math.max(f00, 0) - x;
            f00 = newF00;
            f01 = newF01;

            ans = Math.max(ans, Math.max(f00, f10));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums) {
        constexpr long long NEG_INF = LLONG_MIN / 2;
        long long f00 = NEG_INF, f01 = NEG_INF, f10 = NEG_INF, f11 = NEG_INF;
        long long ans = NEG_INF;
        for (int i = nums.size() - 1; i >= 0; i--) {
            int x = nums[i];
            tie(f10, f11) = pair(max(f11 + x, f00), max(f10 - x, f01));
            tie(f00, f01) = pair(max(f01, 0LL) + x, max(f00, 0LL) - x);
            ans = max(ans, max(f00, f10));
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxAlternatingSum(nums []int) int64 {
	const negInf = math.MinInt / 2
	f00, f01, f10, f11 := negInf, negInf, negInf, negInf
	ans := negInf
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		f10, f11 = max(f11+x, f00), max(f10-x, f01)
		f00, f01 = max(f01, 0)+x, max(f00, 0)-x
		ans = max(ans, f00, f10)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**六、状态机 DP**」「**§7.3 子数组 DP**」和「**专题：前后缀分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
