如果没有长度约束，那么本题就是 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)，[我的题解](https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/)。

这引出了本题的两种做法：前缀和、动态规划。

## 方法一：前缀和

计算 $\textit{nums}$ 的前缀和数组 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

子数组 $[i,j)$ 的元素和为 $s[j]-s[i]$，长度为 $j-i$。

问题相当于：

- 计算最大的 $s[j]-s[i]$，满足 $i < j$ 且 $j-i$ 是 $k$ 的倍数。

> 注：限制 $i<j$ 是为了让子数组非空，符合题目要求。

枚举 $j$，要使 $s[j]-s[i]$ 尽量大，$s[i]$ 要尽量小。

要枚举 $i$ 吗？那样太慢了。

比如 $k=2$：

- 当 $j$ 是偶数时，比如 $j=6$，要使长度是 $k=2$ 的倍数，那么 $i$ 也必须是偶数 $0,2,4$。所以只需维护偶数下标的 $s[i]$ 的最小值，而不是遍历所有 $s[i]$。
- 当 $j$ 是奇数时，比如 $j=7$，要使长度是 $k=2$ 的倍数，那么 $i$ 也必须是奇数 $1,3,5$。所以只需维护奇数下标的 $s[i]$ 的最小值，而不是遍历所有 $s[i]$。

一般地，在遍历前缀和数组 $s$ 的同时，维护：

- 满足 $i < j$ 且 $i$ 与 $j$ 关于模 $k$ **同余**的 $s[i]$ 的最小值。

关于同余的概念，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1YeqHYSEhK/?t=7m56s)，欢迎点赞关注~

### 优化前

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: List[int], k: int) -> int:
        pre = list(accumulate(nums, initial=0))
        min_s = [inf] * k
        ans = -inf
        for j, s in enumerate(pre):
            i = j % k
            ans = max(ans, s - min_s[i])
            min_s[i] = min(min_s[i], s)
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        int n = nums.length;
        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        long[] minS = new long[k];
        Arrays.fill(minS, Long.MAX_VALUE / 2); // 防止下面减法溢出

        long ans = Long.MIN_VALUE;
        for (int j = 0; j < sum.length; j++) {
            int i = j % k;
            ans = Math.max(ans, sum[j] - minS[i]);
            minS[i] = Math.min(minS[i], sum[j]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums, int k) {
        int n = nums.size();
        vector<long long> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        vector<long long> min_s(k, LLONG_MAX / 2); // 防止下面减法溢出
        long long ans = LLONG_MIN;
        for (int j = 0; j < sum.size(); j++) {
            int i = j % k;
            ans = max(ans, sum[j] - min_s[i]);
            min_s[i] = min(min_s[i], sum[j]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int, k int) int64 {
	sum := make([]int, len(nums)+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	minS := make([]int, k)
	for i := range minS {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	ans := math.MinInt
	for j, s := range sum {
		i := j % k
		ans = max(ans, s-minS[i])
		minS[i] = min(minS[i], s)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 优化

一边计算前缀和，一边维护 $\textit{minS}$。

这里前缀和的下标从 $-1$ 开始，也就是定义 $s[-1] = 0$。由于 $-1$ 与 $k-1$ 模 $k$ 同余，所以初始化 $\textit{minS}[k-1] = 0$。

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: List[int], k: int) -> int:
        min_s = [inf] * k
        min_s[-1] = s = 0
        ans = -inf
        for j, x in enumerate(nums):
            s += x
            i = j % k
            ans = max(ans, s - min_s[i])
            min_s[i] = min(min_s[i], s)
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        long[] minS = new long[k];
        Arrays.fill(minS, 0, k - 1, Long.MAX_VALUE / 2); // 防止下面减法溢出

        long ans = Long.MIN_VALUE;
        long s = 0;
        for (int j = 0; j < nums.length; j++) {
            s += nums[j];
            int i = j % k;
            ans = Math.max(ans, s - minS[i]);
            minS[i] = Math.min(minS[i], s);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums, int k) {
        vector<long long> min_s(k, LLONG_MAX / 2); // 防止下面减法溢出
        min_s.back() = 0;

        long long ans = LLONG_MIN, s = 0;
        for (int j = 0; j < nums.size(); j++) {
            s += nums[j];
            int i = j % k;
            ans = max(ans, s - min_s[i]);
            min_s[i] = min(min_s[i], s);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int, k int) int64 {
	minS := make([]int, k)
	for i := range k - 1 {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	ans := math.MinInt
	s := 0
	for j, x := range nums {
		s += x
		i := j % k
		ans = max(ans, s-minS[i])
		minS[i] = min(minS[i], s)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 方法二：动态规划

把连续的 $k$ 个数绑在一起，合成一个数，就变成没有长度约束的 [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/) 了。

定义 $f[i+1]$ 表示以 $i$ 为右端点的、长度是 $k$ 的倍数的非空子数组的最大和。

设 $S$ 为子数组 $[i-k+1, i]$ 的元素和。分类讨论：

- 如果以 $i-k$ 为右端点的、长度是 $k$ 的倍数的非空子数组的最大和大于 $0$，那么可以与当前子数组 $[i-k+1, i]$ 拼接，得到一个更大的和。即 $f[i+1] = f[i-k+1] + S$。
- 否则，不拼接，$f[i+1] = S$。

所以状态转移方程为

$$
f[i+1] = \max(f[i-k+1], 0) + S\ \ \ (i+1\ge k)
$$

初始值：$f[i] = -\infty\ (i < k)$。

答案：$\max(f)$。

```py [sol-Python3]
class Solution:
    def maxSubarraySum(self, nums: list[int], k: int) -> int:
        f = [-inf] * (len(nums) + 1)
        s = 0  # 滑动窗口维护长为 k 的子数组的元素和
        for i, x in enumerate(nums):
            s += x
            left = i - k + 1
            if left < 0:
                continue
            f[i + 1] = max(f[left], 0) + s
            s -= nums[left]
        return max(f)
```

```java [sol-Java]
class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        int n = nums.length;
        long[] f = new long[n + 1];
        Arrays.fill(f, Long.MIN_VALUE);

        long ans = Long.MIN_VALUE;
        long sum = 0; // 滑动窗口维护长为 k 的子数组的元素和
        for (int i = 0; i < n; i++) {
            sum += nums[i];
            int left = i - k + 1;
            if (left < 0) {
                continue;
            }
            f[i + 1] = Math.max(f[left], 0L) + sum;
            ans = Math.max(ans, f[i + 1]);
            sum -= nums[left];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSubarraySum(vector<int>& nums, int k) {
        int n = nums.size();
        vector<long long> f(n + 1, LLONG_MIN);

        long long sum = 0; // 滑动窗口维护长为 k 的子数组的元素和
        for (int i = 0; i < n; i++) {
            sum += nums[i];
            int left = i - k + 1;
            if (left < 0) {
                continue;
            }
            f[i + 1] = max(f[left], 0LL) + sum;
            sum -= nums[left];
        }
        return ranges::max(f);
    }
};
```

```go [sol-Go]
func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MinInt
	}

	sum := 0 // 滑动窗口维护长为 k 的子数组的元素和
	for i, x := range nums {
		sum += x
		left := i - k + 1
		if left < 0 {
			continue
		}
		f[i+1] = max(f[left], 0) + sum
		sum -= nums[left]
	}
	return int64(slices.Max(f))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 数据结构题单的「**§1.2 前缀和与哈希表**」。
2. 动态规划题单的「**§1.3 最大子数组和**」。

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
