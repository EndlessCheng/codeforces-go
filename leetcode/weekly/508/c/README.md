**前置题目**：

1. 本题的无修改版本 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，[我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。
2. 本题做法类似 [1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/)，但 1186 题更简单，推荐先完成 1186 题，并阅读 [我的题解](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/solutions/2321829/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-hzz6/)。

子数组中的元素，从左到右，至多会经过三个阶段：

1. 不修改。
2. 修改。
3. 不修改。

按照状态机 DP 的套路，定义：

- $f[i+1][0]$ 表示右端点为 $i$ 的最大子数组和，且不修改任何元素。
- $f[i+1][1]$ 表示右端点为 $i$ 的最大子数组和，且修改了 $\textit{nums}[i]$。
- $f[i+1][2]$ 表示右端点为 $i$ 的最大子数组和，且在 $\textit{nums}[i]$ 的左边发生了修改（没有修改 $\textit{nums}[i]$）。

设 $x = \textit{nums}[i]$，设 $x$ 修改后的值为 $y$，分阶段讨论：

- **阶段 1**：和 53 题一样，考虑是否把 $x$ 和以 $i-1$ 结尾的子数组（处于阶段 1）拼起来，即 $f[i+1][0] = \max(f[i][0], 0) + x$。
- **阶段 2**：考虑是否把 $x$ 和以 $i-1$ 结尾的子数组（处于阶段 1 或者 2）拼起来，即 $f[i+1][1] = \max(f[i][0], f[i][1], 0) + y$。
- **阶段 3**：考虑是否把 $x$ 和以 $i-1$ 结尾的子数组（处于阶段 2 或者 3）拼起来，即 $f[i+1][2] = \max(f[i][1], f[i][2]) + x$。

初始值：$f[0][j] = 0$。

子数组的末尾可以停在任意阶段，所以答案为 $f[i][j]\ (i\ge 1)$ 的最大值。实际上，在不修改的情况下，无论子数组和是正还是负，修改肯定是好的（正用乘法，负用除法），所以只需考虑 $f[i][1]$ 和 $f[i][2]$ 中的最大值。

> **注**：本题不能只用最大子数组和计算。一个反例：$\textit{nums}=[2,-1,-1,2,-1,-1,2],\ k=2$。正确答案为 $5$，计算方法是把 $-1,-1,2,-1,-1$ 中的每个数都除以 $2$，得到 $\textit{nums}=[2,0,0,1,0,0,2]$。

[本题视频讲解](https://www.bilibili.com/video/BV1wwTu6EEcG/?t=12m20s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: list[int], k: int) -> int:
        def solve(is_mul: bool) -> int:
            # f[i+1][0] 表示右端点为 i 的最大子数组和，且不修改任何元素
            # f[i+1][1] 表示右端点为 i 的最大子数组和，且修改了 nums[i]
            # f[i+1][2] 表示右端点为 i 的最大子数组和，且在 nums[i] 的左边发生了修改（没有修改 nums[i]）
            f = [[0] * 3 for _ in range(len(nums) + 1)]
            res = -inf

            for i, x in enumerate(nums):
                y = x * k if is_mul else int(x / k)
                # 不修改 x，和 f[i][0] 拼起来，或者 x 是子数组的第一个数
                f[i + 1][0] = max(f[i][0], 0) + x
                # 修改 x，和 f[i][0] 或者 f[i][1] 拼起来，或者 y 是子数组的第一个数
                f[i + 1][1] = max(f[i][0], f[i][1], 0) + y
                # 不修改 x，和 f[i][1] 或者 f[i][2] 拼起来
                f[i + 1][2] = max(f[i][1], f[i][2]) + x
                # 枚举子数组的右端点为 i
                res = max(res, f[i + 1][1], f[i + 1][2])

            return res

        return max(solve(True), solve(False))
```

```java [sol-Java]
class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        return Math.max(solve(nums, k, true), solve(nums, k, false));
    }

    private long solve(int[] nums, int k, boolean isMul) {
        int n = nums.length;
        // f[i+1][0] 表示右端点为 i 的最大子数组和，且不修改任何元素
        // f[i+1][1] 表示右端点为 i 的最大子数组和，且修改了 nums[i]
        // f[i+1][2] 表示右端点为 i 的最大子数组和，且在 nums[i] 的左边发生了修改（没有修改 nums[i]）
        long[][] f = new long[n + 1][3];
        long res = Long.MIN_VALUE;

        for (int i = 0; i < n; i++) {
            long x = nums[i];
            long y = isMul ? x * k : x / k;
            // 不修改 x，和 f[i][0] 拼起来，或者 x 是子数组的第一个数
            f[i + 1][0] = Math.max(f[i][0], 0) + x;
            // 修改 x，和 f[i][0] 或者 f[i][1] 拼起来，或者 y 是子数组的第一个数
            f[i + 1][1] = Math.max(Math.max(f[i][0], f[i][1]), 0) + y;
            // 不修改 x，和 f[i][1] 或者 f[i][2] 拼起来
            f[i + 1][2] = Math.max(f[i][1], f[i][2]) + x;
            // 枚举子数组的右端点为 i
            res = Math.max(res, Math.max(f[i + 1][1], f[i + 1][2]));
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums, int k) {
        auto solve = [&](bool is_mul) -> long long {
            long long res = LLONG_MIN;
            int n = nums.size();
            // f[i+1][0] 表示右端点为 i 的最大子数组和，且不修改任何元素
            // f[i+1][1] 表示右端点为 i 的最大子数组和，且修改了 nums[i]
            // f[i+1][2] 表示右端点为 i 的最大子数组和，且在 nums[i] 的左边发生了修改（没有修改 nums[i]）
            vector<array<long long, 3>> f(n + 1);

            for (int i = 0; i < n; i++) {
                long long x = nums[i];
                long long y = is_mul ? x * k : x / k;
                // 不修改 x，和 f[i][0] 拼起来，或者 x 是子数组的第一个数
                f[i + 1][0] = max(f[i][0], 0LL) + x;
                // 修改 x，和 f[i][0] 或者 f[i][1] 拼起来，或者 y 是子数组的第一个数
                f[i + 1][1] = max(max(f[i][0], f[i][1]), 0LL) + y;
                // 不修改 x，和 f[i][1] 或者 f[i][2] 拼起来
                f[i + 1][2] = max(f[i][1], f[i][2]) + x;
                // 枚举子数组的右端点为 i
                res = max(res, max(f[i + 1][1], f[i + 1][2]));
            }

            return res;
        };

        return max(solve(true), solve(false));
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int, k int) int64 {
	solve := func(isMul bool) int64 {
		res := int64(math.MinInt)
		n := len(nums)
		// f[i+1][0] 表示右端点为 i 的最大子数组和，且不修改任何元素
		// f[i+1][1] 表示右端点为 i 的最大子数组和，且修改了 nums[i]
		// f[i+1][2] 表示右端点为 i 的最大子数组和，且在 nums[i] 的左边发生了修改（没有修改 nums[i]）
		f := make([][3]int64, n+1)
		for i, x := range nums {
			x := int64(x)
			y := x
			if isMul {
				y *= int64(k)
			} else {
				y /= int64(k)
			}
			// 不修改 x，和 f[i][0] 拼起来，或者 x 是子数组的第一个数
			f[i+1][0] = max(f[i][0], 0) + x
			// 修改 x，和 f[i][0] 或者 f[i][1] 拼起来，或者 y 是子数组的第一个数
			f[i+1][1] = max(f[i][0], f[i][1], 0) + y
			// 不修改 x，和 f[i][1] 或者 f[i][2] 拼起来
			f[i+1][2] = max(f[i][1], f[i][2]) + x
			// 枚举子数组的右端点为 i
			res = max(res, f[i+1][1], f[i+1][2])
		}
		return res
	}
	return max(solve(true), solve(false))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 空间优化

注意：为避免状态被覆盖，要按照 $j=2,1,0$ 的顺序计算 $f[i+1][j]$。

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: list[int], k: int) -> int:
        def solve(is_mul: bool) -> int:
            f0 = f1 = f2 = 0
            res = -inf

            for x in nums:
                y = x * k if is_mul else int(x / k)
                f2 = max(f1, f2) + x
                f1 = max(f0, f1, 0) + y
                f0 = max(f0, 0) + x
                res = max(res, f1, f2)

            return res

        return max(solve(True), solve(False))
```

```java [sol-Java]
class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        return Math.max(solve(nums, k, true), solve(nums, k, false));
    }

    private long solve(int[] nums, int k, boolean isMul) {
        long f0 = 0, f1 = 0, f2 = 0;
        long res = Long.MIN_VALUE;

        for (int x : nums) {
            long y = isMul ? (long) x * k : x / k;
            f2 = Math.max(f1, f2) + x;
            f1 = Math.max(Math.max(f0, f1), 0) + y;
            f0 = Math.max(f0, 0) + x;
            res = Math.max(res, Math.max(f1, f2));
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums, int k) {
        auto solve = [&](bool is_mul) -> long long {
            long long f0 = 0, f1 = 0, f2 = 0;
            long long res = LLONG_MIN;

            for (int x : nums) {
                long long y = is_mul ? 1LL * x * k : x / k;
                f2 = max(f1, f2) + x;
                f1 = max(max(f0, f1), 0LL) + y;
                f0 = max(f0, 0LL) + x;
                res = max(res, max(f1, f2));
            }

            return res;
        };

        return max(solve(true), solve(false));
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int, k int) int64 {
	solve := func(isMul bool) int64 {
		res := int64(math.MinInt)
		var f0, f1, f2 int64
		for _, x := range nums {
			x := int64(x)
			y := x
			if isMul {
				y *= int64(k)
			} else {
				y /= int64(k)
			}
			f2 = max(f1, f2) + x
			f1 = max(f0, f1, 0) + y
			f0 = max(f0, 0) + x
			res = max(res, f1, f2)
		}
		return res
	}
	return max(solve(true), solve(false))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**六、状态机 DP**」。

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
