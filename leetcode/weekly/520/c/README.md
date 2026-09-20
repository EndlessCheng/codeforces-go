设整个 $\textit{nums}$ 的交替和为 $S$。一次操作可以让 $S$ 最大增加多少？

分类讨论：

- 如果我们操作的是一个长为**偶数**的子数组。假设子数组为 $[a,b,c,d]$，操作后的顺序为 $[b,c,d,a]$。每个元素在交替和中的正负号都取反了（正号变负号，负号变正号）。
- 如果我们操作的是一个长为**奇数**的子数组。假设子数组为 $[a,b,c]$，操作后的顺序为 $[b,c,a]$。$a$ 的符号没变，其余元素在交替和中的正负号都取反了。所以该操作等价于去掉子数组的第一个数，我们操作的仍然是一个长为偶数的子数组。

所以操作相当于：

- 选择 $\textit{nums}$ 中的一个长为偶数的子数组，把每个元素都取反。

子数组中的 $\textit{nums}[i]$，操作前是 $(-1)^i \textit{nums}[i]$，操作后是 $- (-1)^i \textit{nums}[i]$，增加了 $b[i] = 2(-1)^{i+1} \textit{nums}[i]$。

问题转化成：

- 计算数组 $b$ 中的长为偶数的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)。
- 这也是 [3381. 长度可被 K 整除的子数组的最大元素和](https://leetcode.cn/problems/maximum-subarray-sum-with-length-divisible-by-k/) $k=2$ 的情况。

我们可以把 $b$ 中的元素两两一对，就变成没有偶数长度限制的 53 题了。

有两种情况：

- 按照下标 $(0,1), (2,3), (4,5), \ldots$ 两两一对，计算 53 题。此时 $b[i-1] + b[i] = 2(\textit{nums}[i] - \textit{nums}[i-1])$。
- 按照下标 $(1,2), (3,4), (5,6), \ldots$ 两两一对，计算 53 题。此时 $b[i-1] + b[i] = 2(\textit{nums}[i-1] - \textit{nums}[i])$。

定义 $f[i+1]$ 表示以 $b[i]$ 结尾的偶数长度最大子数组和。如果 $f[i-1] > 0$，那么我们可以与以 $b[i-2]$ 结尾的偶数长度最大子数组和拼起来，即

$$
f[i+1] = \max(f[i-1], 0) + b[i-1] + b[i]。
$$

初始值 $f[0] = f[1] = 0$。

最大增量为 $\max(f)$。

[本题视频讲解](https://www.bilibili.com/video/BV1MEeB65EjZ/?t=10m23s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxValue(self, nums: list[int]) -> int:
        # 先计算整个 nums 的交替和
        alter_sum = 0
        for i, x in enumerate(nums):
            alter_sum += -x if i % 2 else x

        n = len(nums)
        f = [0] * (n + 1)
        for i in range(1, n):
            d = nums[i] - nums[i - 1]  # * 2 提到了最后一行
            f[i + 1] = max(f[i - 1], 0) + (d if i % 2 else -d)

        return alter_sum + max(f) * 2
```

```java [sol-Java]
class Solution {
    public long maxValue(int[] nums) {
        // 先计算整个 nums 的交替和
        long alterSum = 0;
        for (int i = 0; i < nums.length; i++) {
            alterSum += i % 2 > 0 ? -nums[i] : nums[i];
        }

        int n = nums.length;
        long[] f = new long[n + 1];
        long mx = 0;
        for (int i = 1; i < n; i++) {
            int d = nums[i] - nums[i - 1]; // * 2 提到了最后一行
            f[i + 1] = Math.max(f[i - 1], 0) + (i % 2 > 0 ? d : -d);
            mx = Math.max(mx, f[i + 1]);
        }

        return alterSum + mx * 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxValue(vector<int>& nums) {
        // 先计算整个 nums 的交替和
        long long alter_sum = 0;
        for (int i = 0; i < nums.size(); i++) {
            alter_sum += i % 2 ? -nums[i] : nums[i];
        }

        int n = nums.size();
        vector<long long> f(n + 1);
        for (int i = 1; i < n; i++) {
            int d = nums[i] - nums[i - 1]; // * 2 提到了最后一行
            f[i + 1] = max(f[i - 1], 0LL) + (i % 2 ? d : -d);
        }

        return alter_sum + ranges::max(f) * 2;
    }
};
```

```go [sol-Go]
func maxValue(nums []int) int64 {
	// 先计算整个 nums 的交替和
	alterSum := 0
	for i, x := range nums {
		alterSum += x * (1 - i%2*2)
	}

	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i < n; i++ {
		d := (nums[i] - nums[i-1]) * (i%2*2 - 1)
		f[i+1] = max(f[i-1], 0) + d // * 2 提到了最后一行
	}

	return int64(alterSum + slices.Max(f)*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 空间优化 + 合并循环

```py [sol-Python3]
class Solution:
    def maxValue(self, nums: list[int]) -> int:
        alter_sum = nums[0]
        f0 = f1 = mx = 0
        for i in range(1, len(nums)):
            x = nums[i]
            alter_sum += -x if i % 2 else x
            d = x - nums[i - 1]
            f0, f1 = f1, max(f0, 0) + (d if i % 2 else -d)
            mx = max(mx, f1)
        return alter_sum + mx * 2
```

```java [sol-Java]
class Solution {
    public long maxValue(int[] nums) {
        long alterSum = nums[0];
        long f0 = 0;
        long f1 = 0;
        long mx = 0;
        for (int i = 1; i < nums.length; i++) {
            int x = nums[i];
            alterSum += i % 2 > 0 ? -x : x;
            int d = x - nums[i - 1];
            long newF = Math.max(f0, 0) + (i % 2 > 0 ? d : -d);
            f0 = f1;
            f1 = newF;
            mx = Math.max(mx, f1);
        }
        return alterSum + mx * 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxValue(vector<int>& nums) {
        long long alter_sum = nums[0];
        long long f0 = 0, f1 = 0, mx = 0;
        for (int i = 1; i < nums.size(); i++) {
            int x = nums[i];
            alter_sum += i % 2 ? -x : x;
            int d = x - nums[i - 1];
            long long new_f = max(f0, 0LL) + (i % 2 ? d : -d);
            f0 = f1;
            f1 = new_f;
            mx = max(mx, f1);
        }
        return alter_sum + mx * 2;
    }
};
```

```go [sol-Go]
func maxValue(nums []int) int64 {
	alterSum := nums[0]
	var f0, f1, mx int
	for i := 1; i < len(nums); i++ {
		alterSum += nums[i] * (1 - i%2*2)
		d := (nums[i] - nums[i-1]) * (i%2*2 - 1)
		f0, f1 = f1, max(f0, 0)+d
		mx = max(mx, f1)
	}
	return int64(alterSum + mx*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 动态规划题单的「**§1.3 最大子数组和**」。
2. 数据结构题单的「**§1.2 前缀和与哈希表**」。

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
