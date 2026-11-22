枚举峰顶的下标为 $i$，枚举峰顶最终的值为 $j$，我们需要知道：

- 把前缀 $[0,i]$ 变成严格递增，且 $\textit{nums}[i] = j$，所需的最少操作次数是多少？
- 把后缀 $[i,n-1]$ 变成严格递减，且 $\textit{nums}[i] = j$，所需的最少操作次数是多少？

这是两个独立的问题，可以分别解决（前后缀分解）。

以前缀为例。枚举把 $\textit{nums}[i-1]$ 变成 $k\ (k < j)$，问题变成：

- 把前缀 $[0,i-1]$ 变成严格递增，且 $\textit{nums}[i-1] = k$，所需的最少操作次数是多少？

这是一个规模更小的子问题，所以可以用 DP 解决。

定义 $f[i][j]$ 表示把前缀 $[0,i]$ 变成严格递增，且 $\textit{nums}[i] = j$，所需的最少操作次数。

枚举把 $\textit{nums}[i-1]$ 变成 $k$，问题变成把前缀 $[0,i-1]$ 变成严格递增，且 $\textit{nums}[i-1] = k$，所需的最少操作次数，即 $f[i-1][k]$。加上把 $\textit{nums}[i]$ 变成 $j$ 的操作次数 $|\textit{nums}[i] - j|$，取最小值，得

$$
f[i][j] = |\textit{nums}[i] - j| + \min_{k=i}^{j-1} f[i-1][k]
$$

> **注**：由于前缀是严格递增的，所以 $\textit{nums}[i-1]$ 至少为 $i$，这是 $k$ 的最小值。

上式包含一个关于 $f[i-1][k]$ 的**前缀最小值**，所以可以一边枚举 $j$，一边用一个变量 $\textit{preMin}$ 维护 $f[i-1][k]$ 的最小值，把上式优化成

$$
f[i][j] = |\textit{nums}[i] - j| + \textit{preMin}
$$

初始值：$f[0][j] = |\textit{nums}[0] - j|$。

设前缀的 $f$ 数组为 $\textit{pre}$，后缀的 $f$ 数组为 $\textit{suf}$。枚举峰顶的下标为 $i$，枚举峰顶最终的值为 $j$，最少操作次数为

$$
\textit{pre}[i][j] + \textit{suf}[i][j] - |\textit{nums}[i] - j|
$$

注意 $|\textit{nums}[i] - j|$ 多算了一次，要减掉。

最后还剩下一个问题：峰顶的值最大枚举到多少？

设 $U=\max(\textit{nums})$。考虑一个极端情况，我们把 $\textit{nums}[0]$ 变成 $U$，那么从 $\textit{nums}[1]$ 到 $\textit{nums}[n-2]$ 要变成 $U+1,U+2,\ldots,U+n-2$。变成更大的数只会操作更多。所以峰顶的值最大枚举到 $U+n-2$。

如果 $n<3$，无解，返回 $-1$。

[本题视频讲解](https://www.bilibili.com/video/BV1qRyGzFExD/?t=3m25s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: list[int]) -> int:
        n = len(nums)
        if n < 3:
            return -1
        mx = max(nums) + n - 1

        def dp(a: list[int]) -> list[list[int]]:
            f = [[0] * mx for _ in range(n)]
            for j in range(1, mx):
                f[0][j] = abs(a[0] - j)
            for i in range(1, n - 1):
                pre_min = inf
                for j in range(i + 1, mx):
                    pre_min = min(pre_min, f[i - 1][j - 1])
                    f[i][j] = pre_min + abs(a[i] - j)
            return f

        pre = dp(nums)
        suf = dp(nums[::-1])[::-1]

        ans = inf
        # 枚举峰顶的位置和值
        for i in range(1, n - 1):
            # 从左到右严格递增，j 至少是 i+1
            # 从右到左严格递增，j 至少是 n-i
            for j in range(max(i + 1, n - i), mx):
                ans = min(ans, pre[i][j] + suf[i][j] - abs(nums[i] - j))  # 减去重复算的
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        if (n < 3) {
            return -1;
        }

        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        mx += n - 1;

        int[][] pre = dp(nums, mx);
        reverse(nums);
        int[][] suf = dp(nums, mx);
        reverse(suf);
        reverse(nums);

        int ans = Integer.MAX_VALUE;
        // 枚举峰顶的位置和值
        for (int i = 1; i < n - 1; i++) {
            // 从左到右严格递增，j 至少是 i+1
            // 从右到左严格递增，j 至少是 n-i
            for (int j = Math.max(i + 1, n - i); j < mx; j++) {
                ans = Math.min(ans, pre[i][j] + suf[i][j] - Math.abs(nums[i] - j)); // 减去重复算的
            }
        }
        return ans;
    }

    private int[][] dp(int[] nums, int mx) {
        int n = nums.length;
        int[][] f = new int[n][mx];
        for (int j = 1; j < mx; j++) {
            f[0][j] = Math.abs(nums[0] - j);
        }
        for (int i = 1; i < n - 1; i++) {
            int preMin = Integer.MAX_VALUE;
            for (int j = i + 1; j < mx; j++) {
                preMin = Math.min(preMin, f[i - 1][j - 1]);
                f[i][j] = preMin + Math.abs(nums[i] - j);
            }
        }
        return f;
    }

    private void reverse(int[] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }

    private void reverse(int[][] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int[] tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        if (n < 3) {
            return -1;
        }
        int mx = ranges::max(nums) + n - 1;

        auto dp = [&]() {
            vector f(n, vector<int>(mx));
            for (int j = 1; j < mx; j++) {
                f[0][j] = abs(nums[0] - j);
            }
            for (int i = 1; i < n - 1; i++) {
                int pre_min = INT_MAX;
                for (int j = i + 1; j < mx; j++) {
                    pre_min = min(pre_min, f[i - 1][j - 1]);
                    f[i][j] = pre_min + abs(nums[i] - j);
                }
            }
            return f;
        };

        auto pre = dp();
        ranges::reverse(nums);
        auto suf = dp();
        ranges::reverse(suf);
        ranges::reverse(nums);

        int ans = INT_MAX;
        // 枚举峰顶的位置和值
        for (int i = 1; i < n - 1; i++) {
            // 从左到右严格递增，j 至少是 i+1
            // 从右到左严格递增，j 至少是 n-i
            for (int j = max(i + 1, n - i); j < mx; j++) {
                ans = min(ans, pre[i][j] + suf[i][j] - abs(nums[i] - j)); // 减去重复算的
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}
	mx := slices.Max(nums) + n - 1

	dp := func() [][]int {
		f := make([][]int, n)
		for i := range f {
			f[i] = make([]int, mx)
		}

		for j := 1; j < mx; j++ {
			f[0][j] = abs(nums[0] - j)
		}
		for i := 1; i < n-1; i++ {
			preMin := math.MaxInt
			for j := i + 1; j < mx; j++ {
				preMin = min(preMin, f[i-1][j-1])
				f[i][j] = preMin + abs(nums[i]-j)
			}
		}

		return f
	}

	pre := dp()
	slices.Reverse(nums)
	suf := dp()
	slices.Reverse(suf)
	slices.Reverse(nums)

	ans := math.MaxInt
	// 枚举峰顶的位置和值
	for i := 1; i < n-1; i++ {
		// 从左到右严格递增，j 至少是 i+1
		// 从右到左严格递增，j 至少是 n-i
		for j := max(i+1, n-i); j < mx; j++ {
			ans = min(ans, pre[i][j]+suf[i][j]-abs(nums[i]-j)) // 减去重复算的
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+U))$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n(n+U))$。

**注**：用 **Slope Trick** 可以做到 $\mathcal{O}(n\log n)$ 的时间复杂度。相关题目为 [2263. 数组变为有序的最小操作次数](https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing/)（会员题）和 [LCP 59. 搭桥过河](https://leetcode.cn/problems/NfY1m5/)。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」和「**§7.6 多维 DP**」。

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
